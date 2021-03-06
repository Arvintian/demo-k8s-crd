/*
Copyright Arvin
*/

// Code generated by client-gen. DO NOT EDIT.

package v1stable

import (
	v1stable "github.com/Arvintian/demo-k8s-crd/pkg/apis/example.com/v1stable"
	"github.com/Arvintian/demo-k8s-crd/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ExampleV1stableInterface interface {
	RESTClient() rest.Interface
	PetsGetter
}

// ExampleV1stableClient is used to interact with features provided by the example.com group.
type ExampleV1stableClient struct {
	restClient rest.Interface
}

func (c *ExampleV1stableClient) Pets(namespace string) PetInterface {
	return newPets(c, namespace)
}

// NewForConfig creates a new ExampleV1stableClient for the given config.
func NewForConfig(c *rest.Config) (*ExampleV1stableClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ExampleV1stableClient{client}, nil
}

// NewForConfigOrDie creates a new ExampleV1stableClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ExampleV1stableClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ExampleV1stableClient for the given RESTClient.
func New(c rest.Interface) *ExampleV1stableClient {
	return &ExampleV1stableClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1stable.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ExampleV1stableClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
