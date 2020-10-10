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
	v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/apis/bkbcs.tencent.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BKDataApiConfigLister helps list BKDataApiConfigs.
type BKDataApiConfigLister interface {
	// List lists all BKDataApiConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.BKDataApiConfig, err error)
	// BKDataApiConfigs returns an object that can list and get BKDataApiConfigs.
	BKDataApiConfigs(namespace string) BKDataApiConfigNamespaceLister
	BKDataApiConfigListerExpansion
}

// bKDataApiConfigLister implements the BKDataApiConfigLister interface.
type bKDataApiConfigLister struct {
	indexer cache.Indexer
}

// NewBKDataApiConfigLister returns a new BKDataApiConfigLister.
func NewBKDataApiConfigLister(indexer cache.Indexer) BKDataApiConfigLister {
	return &bKDataApiConfigLister{indexer: indexer}
}

// List lists all BKDataApiConfigs in the indexer.
func (s *bKDataApiConfigLister) List(selector labels.Selector) (ret []*v1.BKDataApiConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BKDataApiConfig))
	})
	return ret, err
}

// BKDataApiConfigs returns an object that can list and get BKDataApiConfigs.
func (s *bKDataApiConfigLister) BKDataApiConfigs(namespace string) BKDataApiConfigNamespaceLister {
	return bKDataApiConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BKDataApiConfigNamespaceLister helps list and get BKDataApiConfigs.
type BKDataApiConfigNamespaceLister interface {
	// List lists all BKDataApiConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.BKDataApiConfig, err error)
	// Get retrieves the BKDataApiConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.BKDataApiConfig, error)
	BKDataApiConfigNamespaceListerExpansion
}

// bKDataApiConfigNamespaceLister implements the BKDataApiConfigNamespaceLister
// interface.
type bKDataApiConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BKDataApiConfigs in the indexer for a given namespace.
func (s bKDataApiConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.BKDataApiConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BKDataApiConfig))
	})
	return ret, err
}

// Get retrieves the BKDataApiConfig from the indexer for a given namespace and name.
func (s bKDataApiConfigNamespaceLister) Get(name string) (*v1.BKDataApiConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("bkdataapiconfig"), name)
	}
	return obj.(*v1.BKDataApiConfig), nil
}
