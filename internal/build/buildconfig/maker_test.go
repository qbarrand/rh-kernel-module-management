package buildconfig

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/hashstructure"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	buildv1 "github.com/openshift/api/build/v1"
	kmmv1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/build"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/client"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/constants"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/syncronizedmap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"
)

var _ = Describe("Maker_MakeBuildTemplate", func() {
	const (
		containerImage = "container-image"
		dockerFile     = "FROM some-image"
		moduleName     = "some-name"
		namespace      = "some-namespace"
		targetKernel   = "target-kernel"
	)

	var (
		ctrl            *gomock.Controller
		clnt            *client.MockClient
		maker           Maker
		mockBuildHelper *build.MockHelper
		ctx             context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		clnt = client.NewMockClient(ctrl)
		mockBuildHelper = build.NewMockHelper(ctrl)
		maker = NewMaker(clnt, mockBuildHelper, scheme)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	dockerfileConfigMap := v1.LocalObjectReference{Name: "configMapName"}
	dockerfileCMData := map[string]string{constants.DockerfileCMKey: dockerFile}

	It("should work as expected", func() {
		nodeSelector := map[string]string{"label-key": "label-value"}

		buildArgs := []kmmv1beta1.BuildArg{
			{
				Name:  "arg-1",
				Value: "value-1",
			},
			{
				Name:  "arg-2",
				Value: "value-2",
			},
		}

		buildSecrets := []v1.LocalObjectReference{
			{Name: "secret-1"},
			{Name: "secret-2"},
		}

		irs := v1.LocalObjectReference{Name: "push-secret"}

		mapping := kmmv1beta1.KernelMapping{
			ContainerImage: containerImage,
			Build: &kmmv1beta1.Build{
				BuildArgs:           buildArgs,
				DockerfileConfigMap: &dockerfileConfigMap,
				Secrets:             buildSecrets,
			},
		}

		mod := kmmv1beta1.Module{
			ObjectMeta: metav1.ObjectMeta{
				Name:      moduleName,
				Namespace: namespace,
			},
			Spec: kmmv1beta1.ModuleSpec{
				ModuleLoader: kmmv1beta1.ModuleLoaderSpec{
					Container: kmmv1beta1.ModuleLoaderContainerSpec{
						KernelMappings: []kmmv1beta1.KernelMapping{mapping},
					},
				},
				ImageRepoSecret: &irs,
				Selector:        nodeSelector,
			},
		}

		overrides := []kmmv1beta1.BuildArg{
			{
				Name:  "KERNEL_VERSION",
				Value: targetKernel,
			},
		}

		expected := buildv1.Build{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: moduleName + "-",
				Namespace:    namespace,
				Labels: map[string]string{
					constants.ModuleNameLabel:    moduleName,
					constants.TargetKernelTarget: targetKernel,
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion:         "kmm.sigs.k8s.io/v1beta1",
						Kind:               "Module",
						Name:               moduleName,
						Controller:         pointer.Bool(true),
						BlockOwnerDeletion: pointer.Bool(true),
					},
				},
			},
			Spec: buildv1.BuildSpec{
				CommonSpec: buildv1.CommonSpec{
					ServiceAccount: "builder",
					Source: buildv1.BuildSource{
						Dockerfile: pointer.String(dockerFile),
						Type:       buildv1.BuildSourceDockerfile,
					},
					Strategy: buildv1.BuildStrategy{
						Type: buildv1.DockerBuildStrategyType,
						DockerStrategy: &buildv1.DockerBuildStrategy{
							BuildArgs: append(
								envVarsFromKMMBuildArgs(buildArgs),
								v1.EnvVar{Name: "KERNEL_VERSION", Value: targetKernel},
							),
							Volumes: buildVolumesFromBuildSecrets(buildSecrets),
						},
					},
					Output: buildv1.BuildOutput{
						To: &v1.ObjectReference{
							Kind: "DockerImage",
							Name: containerImage,
						},
						PushSecret: &irs,
					},
					NodeSelector:   nodeSelector,
					MountTrustedCA: pointer.Bool(true),
				},
			},
		}

		hash, err := hashstructure.Hash(expected.Spec.CommonSpec.Source, nil)
		Expect(err).NotTo(HaveOccurred())
		annotations := map[string]string{buildHashAnnotation: fmt.Sprintf("%d", hash)}
		expected.SetAnnotations(annotations)

		gomock.InOrder(
			mockBuildHelper.EXPECT().GetRelevantBuild(mod, mapping).Return(mapping.Build),
			clnt.EXPECT().Get(ctx, types.NamespacedName{Name: dockerfileConfigMap.Name, Namespace: mod.Namespace}, gomock.Any()).DoAndReturn(
				func(_ interface{}, _ interface{}, cm *v1.ConfigMap) error {
					cm.Data = dockerfileCMData
					return nil
				},
			),
			mockBuildHelper.EXPECT().ApplyBuildArgOverrides(buildArgs, overrides).Return(append(buildArgs, kmmv1beta1.BuildArg{Name: "KERNEL_VERSION", Value: targetKernel})),
		)

		bc, err := maker.MakeBuildTemplate(ctx, mod, mapping, targetKernel, containerImage, true, nil)
		Expect(err).NotTo(HaveOccurred())

		Expect(
			cmp.Diff(&expected, bc),
		).To(
			BeEmpty(),
		)
	})

	Context(fmt.Sprintf("using %s", dtkBuildArg), func() {

		var (
			mockKODM *syncronizedmap.MockKernelOsDtkMapping
		)

		BeforeEach(func() {
			mockKODM = syncronizedmap.NewMockKernelOsDtkMapping(ctrl)
		})

		It("should fail if we couldn't get the DTK image", func() {

			build := &kmmv1beta1.Build{
				DockerfileConfigMap: &dockerfileConfigMap,
			}

			gomock.InOrder(
				mockBuildHelper.EXPECT().GetRelevantBuild(gomock.Any(), gomock.Any()).Return(build),
				clnt.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ interface{}, _ interface{}, cm *v1.ConfigMap) error {
						dockerfileData := fmt.Sprintf("FROM %s", dtkBuildArg)
						cm.Data = map[string]string{constants.DockerfileCMKey: dockerfileData}
						return nil
					},
				),
				mockKODM.EXPECT().GetImage(gomock.Any()).Return("", errors.New("random error")),
			)

			_, err := maker.MakeBuildTemplate(ctx, kmmv1beta1.Module{}, kmmv1beta1.KernelMapping{}, "", "", false, mockKODM)
			Expect(err).To(HaveOccurred())
		})

		It(fmt.Sprintf("should add a build arg if %s is used in the Dockerfile", dtkBuildArg), func() {

			const dtkImage = "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:111"

			build := &kmmv1beta1.Build{
				DockerfileConfigMap: &dockerfileConfigMap,
			}

			buildArgs := []kmmv1beta1.BuildArg{
				{
					Name:  dtkBuildArg,
					Value: dtkImage,
				},
			}

			gomock.InOrder(
				mockBuildHelper.EXPECT().GetRelevantBuild(gomock.Any(), gomock.Any()).Return(build),
				clnt.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ interface{}, _ interface{}, cm *v1.ConfigMap) error {
						dockerfileData := fmt.Sprintf("FROM %s", dtkBuildArg)
						cm.Data = map[string]string{constants.DockerfileCMKey: dockerfileData}
						return nil
					},
				),
				mockKODM.EXPECT().GetImage(gomock.Any()).Return(dtkImage, nil),
				mockBuildHelper.EXPECT().ApplyBuildArgOverrides(gomock.Any(), gomock.Any()).Return(buildArgs),
			)

			bct, err := maker.MakeBuildTemplate(ctx, kmmv1beta1.Module{}, kmmv1beta1.KernelMapping{}, "", "", false, mockKODM)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(bct.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs)).To(Equal(1))
			Expect(bct.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs[0].Name).To(Equal(buildArgs[0].Name))
			Expect(bct.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs[0].Value).To(Equal(buildArgs[0].Value))
		})
	})
})

var _ = Describe("envVarsFromKMMBuildArgs", func() {
	It("should return nil if args is nil", func() {
		Expect(envVarsFromKMMBuildArgs(nil)).To(BeNil())
	})

	It("should work as expected", func() {
		args := []kmmv1beta1.BuildArg{
			{Name: "arg1", Value: "value1"},
			{Name: "arg2", Value: "value2"},
		}

		expected := []v1.EnvVar{
			{Name: "arg1", Value: "value1"},
			{Name: "arg2", Value: "value2"},
		}

		Expect(envVarsFromKMMBuildArgs(args)).To(Equal(expected))
	})
})

var _ = Describe("buildVolumesFromBuildSecrets", func() {
	It("should return nil if secrets is nil", func() {
		Expect(buildVolumesFromBuildSecrets(nil)).To(BeNil())
	})

	It("should work as expected", func() {
		secrets := []v1.LocalObjectReference{
			{Name: "secret-1"},
			{Name: "secret-2"},
		}

		expectedVolumes := []buildv1.BuildVolume{
			{
				Name: "secret-secret-1",
				Source: buildv1.BuildVolumeSource{
					Type: buildv1.BuildVolumeSourceTypeSecret,
					Secret: &v1.SecretVolumeSource{
						SecretName: "secret-1",
						Optional:   pointer.Bool(false),
					},
				},
				Mounts: []buildv1.BuildVolumeMount{
					{DestinationPath: "/run/secrets/secret-1"},
				},
			},
			{
				Name: "secret-secret-2",
				Source: buildv1.BuildVolumeSource{
					Type: buildv1.BuildVolumeSourceTypeSecret,
					Secret: &v1.SecretVolumeSource{
						SecretName: "secret-2",
						Optional:   pointer.Bool(false),
					},
				},
				Mounts: []buildv1.BuildVolumeMount{
					{DestinationPath: "/run/secrets/secret-2"},
				},
			},
		}

		Expect(buildVolumesFromBuildSecrets(secrets)).To(Equal(expectedVolumes))
	})
})