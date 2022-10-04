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

package policy

import (
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	"github.com/kcp-dev/client-go/clients/informers/policy/v1"
	"github.com/kcp-dev/client-go/clients/informers/policy/v1beta1"
)

type ClusterInterface interface {
	// V1 provides access to the shared informers in V1.
	V1() v1.ClusterInterface
	// V1beta1 provides access to the shared informers in V1beta1.
	V1beta1() v1beta1.ClusterInterface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &group{factory: f, tweakListOptions: tweakListOptions}
}

// V1 returns a new v1.ClusterInterface.
func (g *group) V1() v1.ClusterInterface {
	return v1.New(g.factory, g.tweakListOptions)
}

// V1beta1 returns a new v1beta1.ClusterInterface.
func (g *group) V1beta1() v1beta1.ClusterInterface {
	return v1beta1.New(g.factory, g.tweakListOptions)
}