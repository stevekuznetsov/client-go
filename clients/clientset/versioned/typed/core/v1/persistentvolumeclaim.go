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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

// PersistentVolumeClaimsClusterGetter has a method to return a PersistentVolumeClaimsClusterInterface.
// A group's cluster client should implement this interface.
type PersistentVolumeClaimsClusterGetter interface {
	PersistentVolumeClaims() PersistentVolumeClaimsClusterInterface
}

// PersistentVolumeClaimsClusterInterface can operate on PersistentVolumeClaims across all clusters,
// or scope down to one cluster and return a PersistentVolumeClaimsNamespacer.
type PersistentVolumeClaimsClusterInterface interface {
	Cluster(logicalcluster.Name) PersistentVolumeClaimsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type persistentVolumeClaimsClusterInterface struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *persistentVolumeClaimsClusterInterface) Cluster(name logicalcluster.Name) PersistentVolumeClaimsNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &persistentVolumeClaimsNamespacer{clientCache: c.clientCache, name: name}
}

// List returns the entire collection of all PersistentVolumeClaims across all clusters.
func (c *persistentVolumeClaimsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PersistentVolumeClaims(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all PersistentVolumeClaims across all clusters.
func (c *persistentVolumeClaimsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PersistentVolumeClaims(metav1.NamespaceAll).Watch(ctx, opts)
}

// PersistentVolumeClaimsNamespacer can scope to objects within a namespace, returning a corev1client.PersistentVolumeClaimInterface.
type PersistentVolumeClaimsNamespacer interface {
	Namespace(string) corev1client.PersistentVolumeClaimInterface
}

type persistentVolumeClaimsNamespacer struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
	name        logicalcluster.Name
}

func (n *persistentVolumeClaimsNamespacer) Namespace(namespace string) corev1client.PersistentVolumeClaimInterface {
	return n.clientCache.ClusterOrDie(n.name).PersistentVolumeClaims(namespace)
}