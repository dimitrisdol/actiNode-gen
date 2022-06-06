/*
Copyright 2022.

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
	"context"
	time "time"

	testv1 "github.com/dimitrisdol/actiNode-gen/apis/test/v1"
	versioned "github.com/dimitrisdol/actiNode-gen/generated/clientset/versioned"
	internalinterfaces "github.com/dimitrisdol/actiNode-gen/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/dimitrisdol/actiNode-gen/generated/listers/test/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ActiNodeInformer provides access to a shared informer and lister for
// ActiNodes.
type ActiNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ActiNodeLister
}

type actiNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewActiNodeInformer constructs a new informer for ActiNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewActiNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredActiNodeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredActiNodeInformer constructs a new informer for ActiNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredActiNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TestV1().ActiNodes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TestV1().ActiNodes(namespace).Watch(context.TODO(), options)
			},
		},
		&testv1.ActiNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *actiNodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredActiNodeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *actiNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&testv1.ActiNode{}, f.defaultInformer)
}

func (f *actiNodeInformer) Lister() v1.ActiNodeLister {
	return v1.NewActiNodeLister(f.Informer().GetIndexer())
}
