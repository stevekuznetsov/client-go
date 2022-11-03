//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	storagev1alpha1 "k8s.io/api/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	storagev1alpha1listers "k8s.io/client-go/listers/storage/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// CSIStorageCapacityClusterLister can list CSIStorageCapacities across all workspaces, or scope down to a CSIStorageCapacityLister for one workspace.
type CSIStorageCapacityClusterLister interface {
	List(selector labels.Selector) (ret []*storagev1alpha1.CSIStorageCapacity, err error)
	Cluster(cluster logicalcluster.Name) storagev1alpha1listers.CSIStorageCapacityLister
}

type cSIStorageCapacityClusterLister struct {
	indexer cache.Indexer
}

// NewCSIStorageCapacityClusterLister returns a new CSIStorageCapacityClusterLister.
func NewCSIStorageCapacityClusterLister(indexer cache.Indexer) *cSIStorageCapacityClusterLister {
	return &cSIStorageCapacityClusterLister{indexer: indexer}
}

// List lists all CSIStorageCapacities in the indexer across all workspaces.
func (s *cSIStorageCapacityClusterLister) List(selector labels.Selector) (ret []*storagev1alpha1.CSIStorageCapacity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storagev1alpha1.CSIStorageCapacity))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get CSIStorageCapacities.
func (s *cSIStorageCapacityClusterLister) Cluster(cluster logicalcluster.Name) storagev1alpha1listers.CSIStorageCapacityLister {
	return &cSIStorageCapacityLister{indexer: s.indexer, cluster: cluster}
}

// cSIStorageCapacityLister implements the storagev1alpha1listers.CSIStorageCapacityLister interface.
type cSIStorageCapacityLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all CSIStorageCapacities in the indexer for a workspace.
func (s *cSIStorageCapacityLister) List(selector labels.Selector) (ret []*storagev1alpha1.CSIStorageCapacity, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*storagev1alpha1.CSIStorageCapacity))
	})
	return ret, err
}

// CSIStorageCapacities returns an object that can list and get CSIStorageCapacities in one namespace.
func (s *cSIStorageCapacityLister) CSIStorageCapacities(namespace string) storagev1alpha1listers.CSIStorageCapacityNamespaceLister {
	return &cSIStorageCapacityNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// cSIStorageCapacityNamespaceLister implements the storagev1alpha1listers.CSIStorageCapacityNamespaceLister interface.
type cSIStorageCapacityNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all CSIStorageCapacities in the indexer for a given workspace and namespace.
func (s *cSIStorageCapacityNamespaceLister) List(selector labels.Selector) (ret []*storagev1alpha1.CSIStorageCapacity, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*storagev1alpha1.CSIStorageCapacity))
	})
	return ret, err
}

// Get retrieves the CSIStorageCapacity from the indexer for a given workspace, namespace and name.
func (s *cSIStorageCapacityNamespaceLister) Get(name string) (*storagev1alpha1.CSIStorageCapacity, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storagev1alpha1.Resource("CSIStorageCapacity"), name)
	}
	return obj.(*storagev1alpha1.CSIStorageCapacity), nil
}