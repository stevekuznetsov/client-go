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

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

// ResourceQuotaClusterLister can list ResourceQuotas across all workspaces, or scope down to a ResourceQuotaLister for one workspace.
type ResourceQuotaClusterLister interface {
	List(selector labels.Selector) (ret []*corev1.ResourceQuota, err error)
	Cluster(cluster logicalcluster.Name) corev1listers.ResourceQuotaLister
}

type resourceQuotaClusterLister struct {
	indexer cache.Indexer
}

// NewResourceQuotaClusterLister returns a new ResourceQuotaClusterLister.
func NewResourceQuotaClusterLister(indexer cache.Indexer) *resourceQuotaClusterLister {
	return &resourceQuotaClusterLister{indexer: indexer}
}

// List lists all ResourceQuotas in the indexer across all workspaces.
func (s *resourceQuotaClusterLister) List(selector labels.Selector) (ret []*corev1.ResourceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.ResourceQuota))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ResourceQuotas.
func (s *resourceQuotaClusterLister) Cluster(cluster logicalcluster.Name) corev1listers.ResourceQuotaLister {
	return &resourceQuotaLister{indexer: s.indexer, cluster: cluster}
}

// resourceQuotaLister implements the corev1listers.ResourceQuotaLister interface.
type resourceQuotaLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ResourceQuotas in the indexer for a workspace.
func (s *resourceQuotaLister) List(selector labels.Selector) (ret []*corev1.ResourceQuota, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.ResourceQuota))
	})
	return ret, err
}

// ResourceQuotas returns an object that can list and get ResourceQuotas in one namespace.
func (s *resourceQuotaLister) ResourceQuotas(namespace string) corev1listers.ResourceQuotaNamespaceLister {
	return &resourceQuotaNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// resourceQuotaNamespaceLister implements the corev1listers.ResourceQuotaNamespaceLister interface.
type resourceQuotaNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all ResourceQuotas in the indexer for a given workspace and namespace.
func (s *resourceQuotaNamespaceLister) List(selector labels.Selector) (ret []*corev1.ResourceQuota, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.ResourceQuota))
	})
	return ret, err
}

// Get retrieves the ResourceQuota from the indexer for a given workspace, namespace and name.
func (s *resourceQuotaNamespaceLister) Get(name string) (*corev1.ResourceQuota, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("ResourceQuota"), name)
	}
	return obj.(*corev1.ResourceQuota), nil
}