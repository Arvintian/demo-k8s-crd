/*
Copyright Arvin
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1stable

import (
	v1stable "github.com/Arvintian/demo-k8s-crd/pkg/apis/example.com/v1stable"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PetLister helps list Pets.
type PetLister interface {
	// List lists all Pets in the indexer.
	List(selector labels.Selector) (ret []*v1stable.Pet, err error)
	// Pets returns an object that can list and get Pets.
	Pets(namespace string) PetNamespaceLister
	PetListerExpansion
}

// petLister implements the PetLister interface.
type petLister struct {
	indexer cache.Indexer
}

// NewPetLister returns a new PetLister.
func NewPetLister(indexer cache.Indexer) PetLister {
	return &petLister{indexer: indexer}
}

// List lists all Pets in the indexer.
func (s *petLister) List(selector labels.Selector) (ret []*v1stable.Pet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1stable.Pet))
	})
	return ret, err
}

// Pets returns an object that can list and get Pets.
func (s *petLister) Pets(namespace string) PetNamespaceLister {
	return petNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PetNamespaceLister helps list and get Pets.
type PetNamespaceLister interface {
	// List lists all Pets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1stable.Pet, err error)
	// Get retrieves the Pet from the indexer for a given namespace and name.
	Get(name string) (*v1stable.Pet, error)
	PetNamespaceListerExpansion
}

// petNamespaceLister implements the PetNamespaceLister
// interface.
type petNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pets in the indexer for a given namespace.
func (s petNamespaceLister) List(selector labels.Selector) (ret []*v1stable.Pet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1stable.Pet))
	})
	return ret, err
}

// Get retrieves the Pet from the indexer for a given namespace and name.
func (s petNamespaceLister) Get(name string) (*v1stable.Pet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1stable.Resource("pet"), name)
	}
	return obj.(*v1stable.Pet), nil
}
