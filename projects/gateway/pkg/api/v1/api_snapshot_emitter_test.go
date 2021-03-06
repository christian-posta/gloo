// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1           string
		namespace2           string
		cfg                  *rest.Config
		emitter              ApiEmitter
		gatewayClient        GatewayClient
		virtualServiceClient VirtualServiceClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		err := setup.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace2)
		kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		Expect(err).NotTo(HaveOccurred())

		cache := kuberc.NewKubeCache()
		// Gateway Constructor
		gatewayClientFactory := &factory.KubeResourceClientFactory{
			Crd:         GatewayCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		gatewayClient, err = NewGatewayClient(gatewayClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// VirtualService Constructor
		virtualServiceClientFactory := &factory.KubeResourceClientFactory{
			Crd:         VirtualServiceCrd,
			Cfg:         cfg,
			SharedCache: cache,
		}
		virtualServiceClient, err = NewVirtualServiceClient(virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewApiEmitter(gatewayClient, virtualServiceClient)
	})
	AfterEach(func() {
		setup.TeardownKube(namespace1)
		setup.TeardownKube(namespace2)
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *ApiSnapshot

		/*
			Gateway
		*/

		assertSnapshotGateways := func(expectGateways GatewayList, unexpectGateways GatewayList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectGateways {
						if _, err := snap.Gateways.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectGateways {
						if _, err := snap.Gateways.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := gatewayClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := gatewayClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}

		gateway1a, err := gatewayClient.Write(NewGateway(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		gateway1b, err := gatewayClient.Write(NewGateway(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotGateways(GatewayList{gateway1a, gateway1b}, nil)

		gateway2a, err := gatewayClient.Write(NewGateway(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		gateway2b, err := gatewayClient.Write(NewGateway(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotGateways(GatewayList{gateway1a, gateway1b, gateway2a, gateway2b}, nil)

		err = gatewayClient.Delete(gateway2a.Metadata.Namespace, gateway2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = gatewayClient.Delete(gateway2b.Metadata.Namespace, gateway2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotGateways(GatewayList{gateway1a, gateway1b}, GatewayList{gateway2a, gateway2b})

		err = gatewayClient.Delete(gateway1a.Metadata.Namespace, gateway1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = gatewayClient.Delete(gateway1b.Metadata.Namespace, gateway1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotGateways(nil, GatewayList{gateway1a, gateway1b, gateway2a, gateway2b})

		/*
			VirtualService
		*/

		assertSnapshotVirtualServices := func(expectVirtualServices VirtualServiceList, unexpectVirtualServices VirtualServiceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectVirtualServices {
						if _, err := snap.VirtualServices.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectVirtualServices {
						if _, err := snap.VirtualServices.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := virtualServiceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := virtualServiceClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}

		virtualService1a, err := virtualServiceClient.Write(NewVirtualService(namespace1, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		virtualService1b, err := virtualServiceClient.Write(NewVirtualService(namespace2, "angela"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualServices(VirtualServiceList{virtualService1a, virtualService1b}, nil)

		virtualService2a, err := virtualServiceClient.Write(NewVirtualService(namespace1, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		virtualService2b, err := virtualServiceClient.Write(NewVirtualService(namespace2, "bob"), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualServices(VirtualServiceList{virtualService1a, virtualService1b, virtualService2a, virtualService2b}, nil)

		err = virtualServiceClient.Delete(virtualService2a.Metadata.Namespace, virtualService2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = virtualServiceClient.Delete(virtualService2b.Metadata.Namespace, virtualService2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualServices(VirtualServiceList{virtualService1a, virtualService1b}, VirtualServiceList{virtualService2a, virtualService2b})

		err = virtualServiceClient.Delete(virtualService1a.Metadata.Namespace, virtualService1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = virtualServiceClient.Delete(virtualService1b.Metadata.Namespace, virtualService1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualServices(nil, VirtualServiceList{virtualService1a, virtualService1b, virtualService2a, virtualService2b})
	})
})
