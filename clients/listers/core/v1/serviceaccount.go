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

// ServiceAccountClusterLister can list ServiceAccounts across all workspaces, or scope down to a ServiceAccountLister for one workspace.
type ServiceAccountClusterLister interface {
	List(selector labels.Selector) (ret []*corev1.ServiceAccount, err error)
	Cluster(cluster logicalcluster.Name) corev1listers.ServiceAccountLister
}

type serviceAccountClusterLister struct {
	indexer cache.Indexer
}

// NewServiceAccountClusterLister returns a new ServiceAccountClusterLister.
func NewServiceAccountClusterLister(indexer cache.Indexer) *serviceAccountClusterLister {
	return &serviceAccountClusterLister{indexer: indexer}
}

// List lists all ServiceAccounts in the indexer across all workspaces.
func (s *serviceAccountClusterLister) List(selector labels.Selector) (ret []*corev1.ServiceAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.ServiceAccount))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ServiceAccounts.
func (s *serviceAccountClusterLister) Cluster(cluster logicalcluster.Name) corev1listers.ServiceAccountLister {
	return &serviceAccountLister{indexer: s.indexer, cluster: cluster}
}

// serviceAccountLister implements the corev1listers.ServiceAccountLister interface.
type serviceAccountLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ServiceAccounts in the indexer for a workspace.
func (s *serviceAccountLister) List(selector labels.Selector) (ret []*corev1.ServiceAccount, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*corev1.ServiceAccount)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// ServiceAccounts returns an object that can list and get ServiceAccounts in one namespace.
func (s *serviceAccountLister) ServiceAccounts(namespace string) corev1listers.ServiceAccountNamespaceLister {
	return &serviceAccountNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// serviceAccountNamespaceLister implements the corev1listers.ServiceAccountNamespaceLister interface.
type serviceAccountNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all ServiceAccounts in the indexer for a given workspace and namespace.
func (s *serviceAccountNamespaceLister) List(selector labels.Selector) (ret []*corev1.ServiceAccount, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*corev1.ServiceAccount)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the ServiceAccount from the indexer for a given workspace, namespace and name.
func (s *serviceAccountNamespaceLister) Get(name string) (*corev1.ServiceAccount, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("ServiceAccount"), name)
	}
	return obj.(*corev1.ServiceAccount), nil
}
