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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	storagev1alpha1 "k8s.io/api/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	storagev1alpha1listers "k8s.io/client-go/listers/storage/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// VolumeAttachmentClusterLister can list VolumeAttachments across all workspaces, or scope down to a VolumeAttachmentLister for one workspace.
type VolumeAttachmentClusterLister interface {
	List(selector labels.Selector) (ret []*storagev1alpha1.VolumeAttachment, err error)
	Cluster(cluster logicalcluster.Name) storagev1alpha1listers.VolumeAttachmentLister
}

type volumeAttachmentClusterLister struct {
	indexer cache.Indexer
}

// NewVolumeAttachmentClusterLister returns a new VolumeAttachmentClusterLister.
func NewVolumeAttachmentClusterLister(indexer cache.Indexer) *volumeAttachmentClusterLister {
	return &volumeAttachmentClusterLister{indexer: indexer}
}

// List lists all VolumeAttachments in the indexer across all workspaces.
func (s *volumeAttachmentClusterLister) List(selector labels.Selector) (ret []*storagev1alpha1.VolumeAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storagev1alpha1.VolumeAttachment))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get VolumeAttachments.
func (s *volumeAttachmentClusterLister) Cluster(cluster logicalcluster.Name) storagev1alpha1listers.VolumeAttachmentLister {
	return &volumeAttachmentLister{indexer: s.indexer, cluster: cluster}
}

// volumeAttachmentLister implements the storagev1alpha1listers.VolumeAttachmentLister interface.
type volumeAttachmentLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all VolumeAttachments in the indexer for a workspace.
func (s *volumeAttachmentLister) List(selector labels.Selector) (ret []*storagev1alpha1.VolumeAttachment, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*storagev1alpha1.VolumeAttachment))
	})
	return ret, err
}

// Get retrieves the VolumeAttachment from the indexer for a given workspace and name.
func (s *volumeAttachmentLister) Get(name string) (*storagev1alpha1.VolumeAttachment, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storagev1alpha1.Resource("VolumeAttachment"), name)
	}
	return obj.(*storagev1alpha1.VolumeAttachment), nil
}