// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageTagMirrorSetLister helps list ImageTagMirrorSets.
// All objects returned here must be treated as read-only.
type ImageTagMirrorSetLister interface {
	// List lists all ImageTagMirrorSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ImageTagMirrorSet, err error)
	// Get retrieves the ImageTagMirrorSet from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ImageTagMirrorSet, error)
	ImageTagMirrorSetListerExpansion
}

// imageTagMirrorSetLister implements the ImageTagMirrorSetLister interface.
type imageTagMirrorSetLister struct {
	indexer cache.Indexer
}

// NewImageTagMirrorSetLister returns a new ImageTagMirrorSetLister.
func NewImageTagMirrorSetLister(indexer cache.Indexer) ImageTagMirrorSetLister {
	return &imageTagMirrorSetLister{indexer: indexer}
}

// List lists all ImageTagMirrorSets in the indexer.
func (s *imageTagMirrorSetLister) List(selector labels.Selector) (ret []*v1.ImageTagMirrorSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ImageTagMirrorSet))
	})
	return ret, err
}

// Get retrieves the ImageTagMirrorSet from the index for a given name.
func (s *imageTagMirrorSetLister) Get(name string) (*v1.ImageTagMirrorSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("imagetagmirrorset"), name)
	}
	return obj.(*v1.ImageTagMirrorSet), nil
}
