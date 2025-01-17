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

// PodTemplateClusterLister can list PodTemplates across all workspaces, or scope down to a PodTemplateLister for one workspace.
type PodTemplateClusterLister interface {
	List(selector labels.Selector) (ret []*corev1.PodTemplate, err error)
	Cluster(cluster logicalcluster.Name) corev1listers.PodTemplateLister
}

type podTemplateClusterLister struct {
	indexer cache.Indexer
}

// NewPodTemplateClusterLister returns a new PodTemplateClusterLister.
func NewPodTemplateClusterLister(indexer cache.Indexer) *podTemplateClusterLister {
	return &podTemplateClusterLister{indexer: indexer}
}

// List lists all PodTemplates in the indexer across all workspaces.
func (s *podTemplateClusterLister) List(selector labels.Selector) (ret []*corev1.PodTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.PodTemplate))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get PodTemplates.
func (s *podTemplateClusterLister) Cluster(cluster logicalcluster.Name) corev1listers.PodTemplateLister {
	return &podTemplateLister{indexer: s.indexer, cluster: cluster}
}

// podTemplateLister implements the corev1listers.PodTemplateLister interface.
type podTemplateLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all PodTemplates in the indexer for a workspace.
func (s *podTemplateLister) List(selector labels.Selector) (ret []*corev1.PodTemplate, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*corev1.PodTemplate)
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

// PodTemplates returns an object that can list and get PodTemplates in one namespace.
func (s *podTemplateLister) PodTemplates(namespace string) corev1listers.PodTemplateNamespaceLister {
	return &podTemplateNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// podTemplateNamespaceLister implements the corev1listers.PodTemplateNamespaceLister interface.
type podTemplateNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all PodTemplates in the indexer for a given workspace and namespace.
func (s *podTemplateNamespaceLister) List(selector labels.Selector) (ret []*corev1.PodTemplate, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*corev1.PodTemplate)
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

// Get retrieves the PodTemplate from the indexer for a given workspace, namespace and name.
func (s *podTemplateNamespaceLister) Get(name string) (*corev1.PodTemplate, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("PodTemplate"), name)
	}
	return obj.(*corev1.PodTemplate), nil
}
