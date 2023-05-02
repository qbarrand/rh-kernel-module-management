package ocpbuild

import (
	"context"
	"errors"
	"fmt"

	buildv1 "github.com/openshift/api/build/v1"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/build"
	buildutils "github.com/rh-ecosystem-edge/kernel-module-management/internal/utils/build"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/rh-ecosystem-edge/kernel-module-management/internal/api"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/auth"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/module"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/registry"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/utils"
)

type manager struct {
	client          client.Client
	maker           Maker
	ocpBuildsHelper buildutils.OpenShiftBuildsHelper
	authFactory     auth.RegistryAuthGetterFactory
	registry        registry.Registry
}

func NewManager(
	client client.Client,
	maker Maker,
	ocpBuildsHelper buildutils.OpenShiftBuildsHelper,
	authFactory auth.RegistryAuthGetterFactory,
	registry registry.Registry) build.Manager {
	return &manager{
		client:          client,
		maker:           maker,
		ocpBuildsHelper: ocpBuildsHelper,
		authFactory:     authFactory,
		registry:        registry,
	}
}

func (m *manager) GarbageCollect(ctx context.Context, modName, namespace string, owner metav1.Object) ([]string, error) {

	//Garbage Collection noti (yet) implemented for Build
	return nil, nil
}

func (m *manager) ShouldSync(ctx context.Context, mld *api.ModuleLoaderData) (bool, error) {
	// if there is no build specified skip
	if !module.ShouldBeBuilt(mld) {
		return false, nil
	}

	targetImage := mld.ContainerImage

	// if build AND sign are specified, then we will build an intermediate image
	// and let sign produce the one specified in targetImage
	if module.ShouldBeSigned(mld) {
		targetImage = module.IntermediateImageName(mld.Name, mld.Namespace, targetImage)
	}

	// build is specified and targetImage is either the final image or the intermediate image
	// tag, depending on whether sign is specified or not. Either way, if targetImage exists
	// we can skip building it
	exists, err := module.ImageExists(ctx, m.authFactory, m.registry, mld, targetImage)
	if err != nil {
		return false, fmt.Errorf("failed to check existence of image %s: %w", targetImage, err)
	}

	return !exists, nil
}

func (m *manager) Sync(
	ctx context.Context,
	mld *api.ModuleLoaderData,
	pushImage bool,
	owner metav1.Object,
) (buildutils.Status, error) {

	logger := log.FromContext(ctx)

	buildTemplate, err := m.maker.MakeBuildTemplate(ctx, mld, pushImage, owner)
	if err != nil {
		return "", fmt.Errorf("could not make Build template: %v", err)
	}

	build, err := m.ocpBuildsHelper.GetBuild(ctx, mld)
	if err != nil {
		if !errors.Is(err, buildutils.ErrNoMatchingBuild) {
			return "", fmt.Errorf("error getting the build: %v", err)
		}

		logger.Info("Creating Build")

		if err = m.client.Create(ctx, buildTemplate); err != nil {
			return "", fmt.Errorf("could not create Build: %v", err)
		}

		return buildutils.StatusCreated, nil
	}

	changed, err := buildutils.IsBuildChanged(build, buildTemplate)
	if err != nil {
		return "", fmt.Errorf("could not determine if Build has changed: %v", err)
	}

	if changed {
		logger.Info("The module's build spec has been changed, deleting the current Build so a new one can be created", "name", build.Name)
		opts := []client.DeleteOption{
			client.PropagationPolicy(metav1.DeletePropagationBackground),
		}
		err = m.client.Delete(ctx, build, opts...)
		if err != nil {
			logger.Info(utils.WarnString(fmt.Sprintf("failed to delete Build %s: %v", build.Name, err)))
		}
		return buildutils.StatusInProgress, nil
	}

	switch build.Status.Phase {
	case buildv1.BuildPhaseComplete:
		return buildutils.StatusCompleted, nil
	case buildv1.BuildPhaseNew, buildv1.BuildPhasePending, buildv1.BuildPhaseRunning:
		return buildutils.StatusInProgress, nil
	case buildv1.BuildPhaseFailed:
		return buildutils.StatusFailed, fmt.Errorf("build failed: %v", build.Status.LogSnippet)
	default:
		return "", fmt.Errorf("unknown status: %v", build.Status)
	}
}