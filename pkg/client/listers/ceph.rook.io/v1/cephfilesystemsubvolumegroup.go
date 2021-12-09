/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CephFilesystemSubVolumeGroupLister helps list CephFilesystemSubVolumeGroups.
// All objects returned here must be treated as read-only.
type CephFilesystemSubVolumeGroupLister interface {
	// List lists all CephFilesystemSubVolumeGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CephFilesystemSubVolumeGroup, err error)
	// CephFilesystemSubVolumeGroups returns an object that can list and get CephFilesystemSubVolumeGroups.
	CephFilesystemSubVolumeGroups(namespace string) CephFilesystemSubVolumeGroupNamespaceLister
	CephFilesystemSubVolumeGroupListerExpansion
}

// cephFilesystemSubVolumeGroupLister implements the CephFilesystemSubVolumeGroupLister interface.
type cephFilesystemSubVolumeGroupLister struct {
	indexer cache.Indexer
}

// NewCephFilesystemSubVolumeGroupLister returns a new CephFilesystemSubVolumeGroupLister.
func NewCephFilesystemSubVolumeGroupLister(indexer cache.Indexer) CephFilesystemSubVolumeGroupLister {
	return &cephFilesystemSubVolumeGroupLister{indexer: indexer}
}

// List lists all CephFilesystemSubVolumeGroups in the indexer.
func (s *cephFilesystemSubVolumeGroupLister) List(selector labels.Selector) (ret []*v1.CephFilesystemSubVolumeGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CephFilesystemSubVolumeGroup))
	})
	return ret, err
}

// CephFilesystemSubVolumeGroups returns an object that can list and get CephFilesystemSubVolumeGroups.
func (s *cephFilesystemSubVolumeGroupLister) CephFilesystemSubVolumeGroups(namespace string) CephFilesystemSubVolumeGroupNamespaceLister {
	return cephFilesystemSubVolumeGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CephFilesystemSubVolumeGroupNamespaceLister helps list and get CephFilesystemSubVolumeGroups.
// All objects returned here must be treated as read-only.
type CephFilesystemSubVolumeGroupNamespaceLister interface {
	// List lists all CephFilesystemSubVolumeGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CephFilesystemSubVolumeGroup, err error)
	// Get retrieves the CephFilesystemSubVolumeGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CephFilesystemSubVolumeGroup, error)
	CephFilesystemSubVolumeGroupNamespaceListerExpansion
}

// cephFilesystemSubVolumeGroupNamespaceLister implements the CephFilesystemSubVolumeGroupNamespaceLister
// interface.
type cephFilesystemSubVolumeGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CephFilesystemSubVolumeGroups in the indexer for a given namespace.
func (s cephFilesystemSubVolumeGroupNamespaceLister) List(selector labels.Selector) (ret []*v1.CephFilesystemSubVolumeGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CephFilesystemSubVolumeGroup))
	})
	return ret, err
}

// Get retrieves the CephFilesystemSubVolumeGroup from the indexer for a given namespace and name.
func (s cephFilesystemSubVolumeGroupNamespaceLister) Get(name string) (*v1.CephFilesystemSubVolumeGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cephfilesystemsubvolumegroup"), name)
	}
	return obj.(*v1.CephFilesystemSubVolumeGroup), nil
}
