/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	tenancyv1alpha1 "github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/apis/tenancy/v1alpha1"
	versioned "github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/client/listers/tenancy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualclusterInformer provides access to a shared informer and lister for
// Virtualclusters.
type VirtualclusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VirtualclusterLister
}

type virtualclusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualclusterInformer constructs a new informer for Virtualcluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualclusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualclusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualclusterInformer constructs a new informer for Virtualcluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualclusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenancyV1alpha1().Virtualclusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenancyV1alpha1().Virtualclusters(namespace).Watch(options)
			},
		},
		&tenancyv1alpha1.Virtualcluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualclusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualclusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualclusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tenancyv1alpha1.Virtualcluster{}, f.defaultInformer)
}

func (f *virtualclusterInformer) Lister() v1alpha1.VirtualclusterLister {
	return v1alpha1.NewVirtualclusterLister(f.Informer().GetIndexer())
}