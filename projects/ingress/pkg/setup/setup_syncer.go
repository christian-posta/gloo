package setup

import (
	"context"

	"github.com/solo-io/gloo/projects/ingress/pkg/translator"

	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/projects/ingress/pkg/api/ingress"
	"k8s.io/client-go/kubernetes"

	"github.com/gogo/protobuf/types"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	gloodefaults "github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/projects/ingress/pkg/api/v1"
	"k8s.io/client-go/rest"
)

func Setup(ctx context.Context, kubeCache *kube.KubeCache, inMemoryCache memory.InMemoryResourceCache, settings *gloov1.Settings) error {
	var (
		cfg       *rest.Config
		clientset kubernetes.Interface
	)
	proxyFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		inMemoryCache,
		kubeCache,
		gloov1.ProxyCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	upstreamFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		inMemoryCache,
		kubeCache,
		gloov1.UpstreamCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	secretFactory, err := bootstrap.SecretFactoryForSettings(
		settings,
		inMemoryCache,
		gloov1.SecretCrd.Plural,
		&cfg,
		&clientset,
	)
	if err != nil {
		return err
	}

	refreshRate, err := types.DurationFromProto(settings.RefreshRate)
	if err != nil {
		return err
	}

	writeNamespace := settings.DiscoveryNamespace
	if writeNamespace == "" {
		writeNamespace = gloodefaults.GlooSystem
	}
	watchNamespaces := settings.WatchNamespaces
	var writeNamespaceProvided bool
	for _, ns := range watchNamespaces {
		if ns == writeNamespace {
			writeNamespaceProvided = true
			break
		}
	}
	if !writeNamespaceProvided {
		watchNamespaces = append(watchNamespaces, writeNamespace)
	}
	opts := Opts{
		WriteNamespace:  writeNamespace,
		WatchNamespaces: watchNamespaces,
		Proxies:         proxyFactory,
		Upstreams:       upstreamFactory,
		Secrets:         secretFactory,
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: refreshRate,
		},
	}

	return RunIngress()(opts)
}

func RunIngress() func(opts Opts) error {
	return func(opts Opts) error {
		opts.WatchOpts = opts.WatchOpts.WithDefaults()
		opts.WatchOpts.Ctx = contextutils.WithLogger(opts.WatchOpts.Ctx, "ingress")

		cfg, err := kubeutils.GetConfig("", "")
		if err != nil {
			return errors.Wrapf(err, "getting kube config")
		}
		kube, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			return errors.Wrapf(err, "getting kube client")
		}

		baseIngressClient := ingress.NewResourceClient(kube, &v1.Ingress{})
		ingressClient := v1.NewIngressClientWithBase(baseIngressClient)

		upstreamClient, err := gloov1.NewUpstreamClient(opts.Upstreams)
		if err != nil {
			return err
		}
		if err := upstreamClient.Register(); err != nil {
			return err
		}

		secretClient, err := gloov1.NewSecretClient(opts.Secrets)
		if err != nil {
			return err
		}

		proxyClient, err := gloov1.NewProxyClient(opts.Proxies)
		if err != nil {
			return err
		}
		if err := proxyClient.Register(); err != nil {
			return err
		}

		emitter := v1.NewApiEmitter(secretClient, upstreamClient, ingressClient)

		rpt := reporter.NewReporter("ingress", ingressClient.BaseClient())
		writeErrs := make(chan error)

		sync := translator.NewSyncer(opts.WriteNamespace, proxyClient, ingressClient, rpt, writeErrs)

		eventLoop := v1.NewApiEventLoop(emitter, sync)
		eventLoopErrs, err := eventLoop.Run(opts.WatchNamespaces, opts.WatchOpts)
		if err != nil {
			return err
		}
		go errutils.AggregateErrs(opts.WatchOpts.Ctx, writeErrs, eventLoopErrs, "event_loop")

		logger := contextutils.LoggerFrom(opts.WatchOpts.Ctx)

		go func() {
			for {
				select {
				case err := <-writeErrs:
					logger.Errorf("error: %v", err)
				case <-opts.WatchOpts.Ctx.Done():
					close(writeErrs)
					return
				}
			}
		}()
		return nil
	}
}
