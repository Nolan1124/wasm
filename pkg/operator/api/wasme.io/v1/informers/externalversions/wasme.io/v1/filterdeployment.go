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

package v1

import (
	time "time"

	wasmeiov1 "github.com/solo-io/wasme/pkg/operator/api/wasme.io/v1"
	versioned "github.com/solo-io/wasme/pkg/operator/api/wasme.io/v1/clientset/versioned"
	internalinterfaces "github.com/solo-io/wasme/pkg/operator/api/wasme.io/v1/informers/externalversions/internalinterfaces"
	v1 "github.com/solo-io/wasme/pkg/operator/api/wasme.io/v1/listers/wasme.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FilterDeploymentInformer provides access to a shared informer and lister for
// FilterDeployments.
type FilterDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.FilterDeploymentLister
}

type filterDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFilterDeploymentInformer constructs a new informer for FilterDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilterDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFilterDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFilterDeploymentInformer constructs a new informer for FilterDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFilterDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WasmeV1().FilterDeployments(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WasmeV1().FilterDeployments(namespace).Watch(options)
			},
		},
		&wasmeiov1.FilterDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *filterDeploymentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFilterDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *filterDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wasmeiov1.FilterDeployment{}, f.defaultInformer)
}

func (f *filterDeploymentInformer) Lister() v1.FilterDeploymentLister {
	return v1.NewFilterDeploymentLister(f.Informer().GetIndexer())
}