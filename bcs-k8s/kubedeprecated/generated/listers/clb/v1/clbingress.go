/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/clb/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClbIngressLister helps list ClbIngresses.
type ClbIngressLister interface {
	// List lists all ClbIngresses in the indexer.
	List(selector labels.Selector) (ret []*v1.ClbIngress, err error)
	// ClbIngresses returns an object that can list and get ClbIngresses.
	ClbIngresses(namespace string) ClbIngressNamespaceLister
	ClbIngressListerExpansion
}

// clbIngressLister implements the ClbIngressLister interface.
type clbIngressLister struct {
	indexer cache.Indexer
}

// NewClbIngressLister returns a new ClbIngressLister.
func NewClbIngressLister(indexer cache.Indexer) ClbIngressLister {
	return &clbIngressLister{indexer: indexer}
}

// List lists all ClbIngresses in the indexer.
func (s *clbIngressLister) List(selector labels.Selector) (ret []*v1.ClbIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClbIngress))
	})
	return ret, err
}

// ClbIngresses returns an object that can list and get ClbIngresses.
func (s *clbIngressLister) ClbIngresses(namespace string) ClbIngressNamespaceLister {
	return clbIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClbIngressNamespaceLister helps list and get ClbIngresses.
type ClbIngressNamespaceLister interface {
	// List lists all ClbIngresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ClbIngress, err error)
	// Get retrieves the ClbIngress from the indexer for a given namespace and name.
	Get(name string) (*v1.ClbIngress, error)
	ClbIngressNamespaceListerExpansion
}

// clbIngressNamespaceLister implements the ClbIngressNamespaceLister
// interface.
type clbIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClbIngresses in the indexer for a given namespace.
func (s clbIngressNamespaceLister) List(selector labels.Selector) (ret []*v1.ClbIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClbIngress))
	})
	return ret, err
}

// Get retrieves the ClbIngress from the indexer for a given namespace and name.
func (s clbIngressNamespaceLister) Get(name string) (*v1.ClbIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clbingress"), name)
	}
	return obj.(*v1.ClbIngress), nil
}
