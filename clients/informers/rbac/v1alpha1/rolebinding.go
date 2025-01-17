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
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"

	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	rbacv1alpha1listers "github.com/kcp-dev/client-go/clients/listers/rbac/v1alpha1"
)

// RoleBindingClusterInformer provides access to a shared informer and lister for
// RoleBindings.
type RoleBindingClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() rbacv1alpha1listers.RoleBindingClusterLister
}

type roleBindingClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRoleBindingClusterInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRoleBindingClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRoleBindingClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRoleBindingClusterInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRoleBindingClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().RoleBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().RoleBindings().Watch(context.TODO(), options)
			},
		},
		&rbacv1alpha1.RoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *roleBindingClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRoleBindingClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *roleBindingClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacv1alpha1.RoleBinding{}, f.defaultInformer)
}

func (f *roleBindingClusterInformer) Lister() rbacv1alpha1listers.RoleBindingClusterLister {
	return rbacv1alpha1listers.NewRoleBindingClusterLister(f.Informer().GetIndexer())
}
