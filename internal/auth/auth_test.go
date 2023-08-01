package auth

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/client"
	"go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("registrySecretAuthGetter_GetKeyChain", func() {
	const (
		secretName      = "pull-push-secret"
		secretNamespace = "default"
	)

	var (
		ctrl          *gomock.Controller
		ctx           context.Context
		factory       *registryAuthGetterFactory
		mockClient    *client.MockClient
		fakeClientSet kubernetes.Interface
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockClient = client.NewMockClient(ctrl)
		fakeClientSet = fake.NewSimpleClientset()
		factory = NewRegistryAuthGetterFactory(mockClient, fakeClientSet).(*registryAuthGetterFactory)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should fail if it cannot get the secret", func() {

		mockClient.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).Return(errors.New("some error"))

		namespacedNamespace := types.NamespacedName{
			Name:      secretName,
			Namespace: secretNamespace,
		}
		registryAuthGetter := factory.newRegistryAuthGetter(namespacedNamespace)

		_, err := registryAuthGetter.GetKeyChain(ctx)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("cannot find secret"))
	})

	It("should fail if the secret doesn't contains auth data", func() {

		mockClient.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).DoAndReturn(
			func(_ interface{}, _ interface{}, s *v1.Secret, _ ...ctrlclient.GetOption) error {
				s.Type = v1.SecretTypeDockerConfigJson
				s.Data = map[string][]byte{
					v1.DockerConfigJsonKey: []byte("some data"),
				}
				return nil
			},
		)

		namespacedNamespace := types.NamespacedName{
			Name:      secretName,
			Namespace: secretNamespace,
		}
		registryAuthGetter := factory.newRegistryAuthGetter(namespacedNamespace)

		_, err := registryAuthGetter.GetKeyChain(ctx)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("could not create a keycahin from secret"))
	})

	It("should work as expected", func() {

		mockClient.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).Return(nil)

		namespacedNamespace := types.NamespacedName{
			Name:      secretName,
			Namespace: secretNamespace,
		}
		registryAuthGetter := factory.newRegistryAuthGetter(namespacedNamespace)

		_, err := registryAuthGetter.GetKeyChain(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
})

var _ = Describe("NewClusterAuthGetter", func() {
	var (
		ctrl       *gomock.Controller
		ctx        context.Context
		factory    *registryAuthGetterFactory
		mockClient *client.MockClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockClient = client.NewMockClient(ctrl)
		factory = NewRegistryAuthGetterFactory(mockClient, nil).(*registryAuthGetterFactory)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("shoudl work as expected", func() {
		namespacedNamespace := types.NamespacedName{
			Name:      pullSecretName,
			Namespace: pullSecretNamespace,
		}
		mockClient.EXPECT().Get(ctx, namespacedNamespace, gomock.Any()).Return(nil)

		registryAuthGetter := factory.NewClusterAuthGetter()
		_, err := registryAuthGetter.GetKeyChain(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
})
