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

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	appsv1listers "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
)

// DaemonSetClusterLister can list DaemonSets across all workspaces, or scope down to a DaemonSetLister for one workspace.
type DaemonSetClusterLister interface {
	List(selector labels.Selector) (ret []*appsv1.DaemonSet, err error)
	Cluster(cluster logicalcluster.Name) appsv1listers.DaemonSetLister
}

type daemonSetClusterLister struct {
	indexer cache.Indexer
}

// NewDaemonSetClusterLister returns a new DaemonSetClusterLister.
func NewDaemonSetClusterLister(indexer cache.Indexer) *daemonSetClusterLister {
	return &daemonSetClusterLister{indexer: indexer}
}

// List lists all DaemonSets in the indexer across all workspaces.
func (s *daemonSetClusterLister) List(selector labels.Selector) (ret []*appsv1.DaemonSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*appsv1.DaemonSet))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get DaemonSets.
func (s *daemonSetClusterLister) Cluster(cluster logicalcluster.Name) appsv1listers.DaemonSetLister {
	return &daemonSetLister{indexer: s.indexer, cluster: cluster}
}

// daemonSetLister implements the appsv1listers.DaemonSetLister interface.
type daemonSetLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all DaemonSets in the indexer for a workspace.
func (s *daemonSetLister) List(selector labels.Selector) (ret []*appsv1.DaemonSet, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*appsv1.DaemonSet)
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

// DaemonSets returns an object that can list and get DaemonSets in one namespace.
func (s *daemonSetLister) DaemonSets(namespace string) appsv1listers.DaemonSetNamespaceLister {
	return &daemonSetNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// daemonSetNamespaceLister implements the appsv1listers.DaemonSetNamespaceLister interface.
type daemonSetNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all DaemonSets in the indexer for a given workspace and namespace.
func (s *daemonSetNamespaceLister) List(selector labels.Selector) (ret []*appsv1.DaemonSet, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*appsv1.DaemonSet)
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

// Get retrieves the DaemonSet from the indexer for a given workspace, namespace and name.
func (s *daemonSetNamespaceLister) Get(name string) (*appsv1.DaemonSet, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1.Resource("DaemonSet"), name)
	}
	return obj.(*appsv1.DaemonSet), nil
}
