// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type ArtifactClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Artifact, error)
	Write(resource *Artifact, opts clients.WriteOpts) (*Artifact, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (ArtifactList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan ArtifactList, <-chan error, error)
}

type artifactClient struct {
	rc clients.ResourceClient
}

func NewArtifactClient(rcFactory factory.ResourceClientFactory) (ArtifactClient, error) {
	return NewArtifactClientWithToken(rcFactory, "")
}

func NewArtifactClientWithToken(rcFactory factory.ResourceClientFactory, token string) (ArtifactClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Artifact{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Artifact resource client")
	}
	return NewArtifactClientWithBase(rc), nil
}

func NewArtifactClientWithBase(rc clients.ResourceClient) ArtifactClient {
	return &artifactClient{
		rc: rc,
	}
}

func (client *artifactClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *artifactClient) Register() error {
	return client.rc.Register()
}

func (client *artifactClient) Read(namespace, name string, opts clients.ReadOpts) (*Artifact, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Artifact), nil
}

func (client *artifactClient) Write(artifact *Artifact, opts clients.WriteOpts) (*Artifact, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(artifact, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Artifact), nil
}

func (client *artifactClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *artifactClient) List(namespace string, opts clients.ListOpts) (ArtifactList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToArtifact(resourceList), nil
}

func (client *artifactClient) Watch(namespace string, opts clients.WatchOpts) (<-chan ArtifactList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	artifactsChan := make(chan ArtifactList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				artifactsChan <- convertToArtifact(resourceList)
			case <-opts.Ctx.Done():
				close(artifactsChan)
				return
			}
		}
	}()
	return artifactsChan, errs, nil
}

func convertToArtifact(resources resources.ResourceList) ArtifactList {
	var artifactList ArtifactList
	for _, resource := range resources {
		artifactList = append(artifactList, resource.(*Artifact))
	}
	return artifactList
}
