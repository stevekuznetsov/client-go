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

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	policyv1beta1client "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
)

// PodSecurityPoliciesClusterGetter has a method to return a PodSecurityPoliciesClusterInterface.
// A group's cluster client should implement this interface.
type PodSecurityPoliciesClusterGetter interface {
	PodSecurityPolicies() PodSecurityPoliciesClusterInterface
}

// PodSecurityPoliciesClusterInterface can operate on PodSecurityPolicies across all clusters,
// or scope down to one cluster and return a policyv1beta1client.PodSecurityPolicyInterface.
type PodSecurityPoliciesClusterInterface interface {
	Cluster(logicalcluster.Name) policyv1beta1client.PodSecurityPolicyInterface
	List(ctx context.Context, opts metav1.ListOptions) (*policyv1beta1.PodSecurityPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type podSecurityPoliciesClusterInterface struct {
	clientCache kcpclient.Cache[*policyv1beta1client.PolicyV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *podSecurityPoliciesClusterInterface) Cluster(name logicalcluster.Name) policyv1beta1client.PodSecurityPolicyInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).PodSecurityPolicies()
}

// List returns the entire collection of all PodSecurityPolicies across all clusters.
func (c *podSecurityPoliciesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*policyv1beta1.PodSecurityPolicyList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PodSecurityPolicies().List(ctx, opts)
}

// Watch begins to watch all PodSecurityPolicies across all clusters.
func (c *podSecurityPoliciesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PodSecurityPolicies().Watch(ctx, opts)
}
