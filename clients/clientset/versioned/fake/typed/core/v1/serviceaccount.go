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
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"

	kcpcorev1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var serviceAccountsResource = schema.GroupVersionResource{Group: "core", Version: "V1", Resource: "serviceaccounts"}
var serviceAccountsKind = schema.GroupVersionKind{Group: "core", Version: "V1", Kind: "ServiceAccount"}

type serviceAccountsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *serviceAccountsClusterClient) Cluster(cluster logicalcluster.Name) kcpcorev1.ServiceAccountsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &serviceAccountsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ServiceAccounts that match those selectors across all clusters.
func (c *serviceAccountsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceAccountList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(serviceAccountsResource, serviceAccountsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.ServiceAccountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceAccountList{ListMeta: obj.(*corev1.ServiceAccountList).ListMeta}
	for _, item := range obj.(*corev1.ServiceAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ServiceAccounts across all clusters.
func (c *serviceAccountsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(serviceAccountsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type serviceAccountsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *serviceAccountsNamespacer) Namespace(namespace string) corev1client.ServiceAccountInterface {
	return &serviceAccountsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type serviceAccountsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *serviceAccountsClient) Create(ctx context.Context, serviceAccount *corev1.ServiceAccount, opts metav1.CreateOptions) (*corev1.ServiceAccount, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(serviceAccountsResource, c.Cluster, c.Namespace, serviceAccount), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

func (c *serviceAccountsClient) Update(ctx context.Context, serviceAccount *corev1.ServiceAccount, opts metav1.UpdateOptions) (*corev1.ServiceAccount, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(serviceAccountsResource, c.Cluster, c.Namespace, serviceAccount), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

func (c *serviceAccountsClient) UpdateStatus(ctx context.Context, serviceAccount *corev1.ServiceAccount, opts metav1.UpdateOptions) (*corev1.ServiceAccount, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(serviceAccountsResource, c.Cluster, "status", c.Namespace, serviceAccount), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

func (c *serviceAccountsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(serviceAccountsResource, c.Cluster, c.Namespace, name, opts), &corev1.ServiceAccount{})
	return err
}

func (c *serviceAccountsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(serviceAccountsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ServiceAccountList{})
	return err
}

func (c *serviceAccountsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.ServiceAccount, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(serviceAccountsResource, c.Cluster, c.Namespace, name), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

// List takes label and field selectors, and returns the list of ServiceAccounts that match those selectors.
func (c *serviceAccountsClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceAccountList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(serviceAccountsResource, serviceAccountsKind, c.Cluster, c.Namespace, opts), &corev1.ServiceAccountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceAccountList{ListMeta: obj.(*corev1.ServiceAccountList).ListMeta}
	for _, item := range obj.(*corev1.ServiceAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *serviceAccountsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(serviceAccountsResource, c.Cluster, c.Namespace, opts))
}

func (c *serviceAccountsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.ServiceAccount, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(serviceAccountsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

func (c *serviceAccountsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ServiceAccountApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ServiceAccount, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(serviceAccountsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

func (c *serviceAccountsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ServiceAccountApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ServiceAccount, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(serviceAccountsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.ServiceAccount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

func (c *serviceAccountsClient) CreateToken(ctx context.Context, serviceAccountName string, tokenRequest *authenticationv1.TokenRequest, opts metav1.CreateOptions) (*authenticationv1.TokenRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateSubresourceAction(serviceAccountsResource, c.Cluster, serviceAccountName, "token", c.Namespace, tokenRequest), &authenticationv1.TokenRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authenticationv1.TokenRequest), err
}