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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	admissionregistrationv1beta1listers "k8s.io/client-go/listers/admissionregistration/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// ValidatingWebhookConfigurationClusterLister can list ValidatingWebhookConfigurations across all workspaces, or scope down to a ValidatingWebhookConfigurationLister for one workspace.
type ValidatingWebhookConfigurationClusterLister interface {
	List(selector labels.Selector) (ret []*admissionregistrationv1beta1.ValidatingWebhookConfiguration, err error)
	Cluster(cluster logicalcluster.Name) admissionregistrationv1beta1listers.ValidatingWebhookConfigurationLister
}

type validatingWebhookConfigurationClusterLister struct {
	indexer cache.Indexer
}

// NewValidatingWebhookConfigurationClusterLister returns a new ValidatingWebhookConfigurationClusterLister.
func NewValidatingWebhookConfigurationClusterLister(indexer cache.Indexer) *validatingWebhookConfigurationClusterLister {
	return &validatingWebhookConfigurationClusterLister{indexer: indexer}
}

// List lists all ValidatingWebhookConfigurations in the indexer across all workspaces.
func (s *validatingWebhookConfigurationClusterLister) List(selector labels.Selector) (ret []*admissionregistrationv1beta1.ValidatingWebhookConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ValidatingWebhookConfigurations.
func (s *validatingWebhookConfigurationClusterLister) Cluster(cluster logicalcluster.Name) admissionregistrationv1beta1listers.ValidatingWebhookConfigurationLister {
	return &validatingWebhookConfigurationLister{indexer: s.indexer, cluster: cluster}
}

// validatingWebhookConfigurationLister implements the admissionregistrationv1beta1listers.ValidatingWebhookConfigurationLister interface.
type validatingWebhookConfigurationLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ValidatingWebhookConfigurations in the indexer for a workspace.
func (s *validatingWebhookConfigurationLister) List(selector labels.Selector) (ret []*admissionregistrationv1beta1.ValidatingWebhookConfiguration, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*admissionregistrationv1beta1.ValidatingWebhookConfiguration)
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

// Get retrieves the ValidatingWebhookConfiguration from the indexer for a given workspace and name.
func (s *validatingWebhookConfigurationLister) Get(name string) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(admissionregistrationv1beta1.Resource("ValidatingWebhookConfiguration"), name)
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), nil
}