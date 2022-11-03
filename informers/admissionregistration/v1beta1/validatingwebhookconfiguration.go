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
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamadmissionregistrationv1beta1informers "k8s.io/client-go/informers/admissionregistration/v1beta1"
	upstreamadmissionregistrationv1beta1listers "k8s.io/client-go/listers/admissionregistration/v1beta1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	admissionregistrationv1beta1listers "github.com/kcp-dev/client-go/listers/admissionregistration/v1beta1"
)

// ValidatingWebhookConfigurationClusterInformer provides access to a shared informer and lister for
// ValidatingWebhookConfigurations.
type ValidatingWebhookConfigurationClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamadmissionregistrationv1beta1informers.ValidatingWebhookConfigurationInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() admissionregistrationv1beta1listers.ValidatingWebhookConfigurationClusterLister
}

type validatingWebhookConfigurationClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewValidatingWebhookConfigurationClusterInformer constructs a new informer for ValidatingWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewValidatingWebhookConfigurationClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredValidatingWebhookConfigurationClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredValidatingWebhookConfigurationClusterInformer constructs a new informer for ValidatingWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredValidatingWebhookConfigurationClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Watch(context.TODO(), options)
			},
		},
		&admissionregistrationv1beta1.ValidatingWebhookConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *validatingWebhookConfigurationClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredValidatingWebhookConfigurationClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *validatingWebhookConfigurationClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&admissionregistrationv1beta1.ValidatingWebhookConfiguration{}, f.defaultInformer)
}

func (f *validatingWebhookConfigurationClusterInformer) Lister() admissionregistrationv1beta1listers.ValidatingWebhookConfigurationClusterLister {
	return admissionregistrationv1beta1listers.NewValidatingWebhookConfigurationClusterLister(f.Informer().GetIndexer())
}

func (f *validatingWebhookConfigurationClusterInformer) Cluster(cluster logicalcluster.Name) upstreamadmissionregistrationv1beta1informers.ValidatingWebhookConfigurationInformer {
	return &validatingWebhookConfigurationInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type validatingWebhookConfigurationInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamadmissionregistrationv1beta1listers.ValidatingWebhookConfigurationLister
}

func (f *validatingWebhookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *validatingWebhookConfigurationInformer) Lister() upstreamadmissionregistrationv1beta1listers.ValidatingWebhookConfigurationLister {
	return f.lister
}