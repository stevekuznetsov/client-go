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
	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	authorizationv1beta1client "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
)

// LocalSubjectAccessReviewsClusterGetter has a method to return a LocalSubjectAccessReviewsClusterInterface.
// A group's cluster client should implement this interface.
type LocalSubjectAccessReviewsClusterGetter interface {
	LocalSubjectAccessReviews() LocalSubjectAccessReviewsClusterInterface
}

// LocalSubjectAccessReviewsClusterInterface can scope down to one cluster and return a LocalSubjectAccessReviewsNamespacer.
type LocalSubjectAccessReviewsClusterInterface interface {
	Cluster(logicalcluster.Name) LocalSubjectAccessReviewsNamespacer
}

type localSubjectAccessReviewsClusterInterface struct {
	clientCache kcpclient.Cache[*authorizationv1beta1client.AuthorizationV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *localSubjectAccessReviewsClusterInterface) Cluster(name logicalcluster.Name) LocalSubjectAccessReviewsNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &localSubjectAccessReviewsNamespacer{clientCache: c.clientCache, name: name}
}

// LocalSubjectAccessReviewsNamespacer can scope to objects within a namespace, returning a authorizationv1beta1client.LocalSubjectAccessReviewInterface.
type LocalSubjectAccessReviewsNamespacer interface {
	Namespace(string) authorizationv1beta1client.LocalSubjectAccessReviewInterface
}

type localSubjectAccessReviewsNamespacer struct {
	clientCache kcpclient.Cache[*authorizationv1beta1client.AuthorizationV1beta1Client]
	name        logicalcluster.Name
}

func (n *localSubjectAccessReviewsNamespacer) Namespace(namespace string) authorizationv1beta1client.LocalSubjectAccessReviewInterface {
	return n.clientCache.ClusterOrDie(n.name).LocalSubjectAccessReviews(namespace)
}
