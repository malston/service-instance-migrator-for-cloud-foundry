/* 
 *  Copyright 2022 VMware, Inc.
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  http://www.apache.org/licenses/LICENSE-2.0
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/om"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/om/httpclient"
)

type FakeOpsManager struct {
	GetBOSHCredentialsStub        func() (om.BoshCredentials, error)
	getBOSHCredentialsMutex       sync.RWMutex
	getBOSHCredentialsArgsForCall []struct {
	}
	getBOSHCredentialsReturns struct {
		result1 om.BoshCredentials
		result2 error
	}
	getBOSHCredentialsReturnsOnCall map[int]struct {
		result1 om.BoshCredentials
		result2 error
	}
	GetCertificateAuthoritiesStub        func() ([]httpclient.CA, error)
	getCertificateAuthoritiesMutex       sync.RWMutex
	getCertificateAuthoritiesArgsForCall []struct {
	}
	getCertificateAuthoritiesReturns struct {
		result1 []httpclient.CA
		result2 error
	}
	getCertificateAuthoritiesReturnsOnCall map[int]struct {
		result1 []httpclient.CA
		result2 error
	}
	GetDeployedProductCredentialsStub        func(string, string) (httpclient.DeployedProductCredential, error)
	getDeployedProductCredentialsMutex       sync.RWMutex
	getDeployedProductCredentialsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getDeployedProductCredentialsReturns struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}
	getDeployedProductCredentialsReturnsOnCall map[int]struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}
	GetStagedProductPropertiesStub        func(string) (map[string]httpclient.ResponseProperty, error)
	getStagedProductPropertiesMutex       sync.RWMutex
	getStagedProductPropertiesArgsForCall []struct {
		arg1 string
	}
	getStagedProductPropertiesReturns struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}
	getStagedProductPropertiesReturnsOnCall map[int]struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}
	ListDeployedProductsStub        func() ([]httpclient.DeployedProduct, error)
	listDeployedProductsMutex       sync.RWMutex
	listDeployedProductsArgsForCall []struct {
	}
	listDeployedProductsReturns struct {
		result1 []httpclient.DeployedProduct
		result2 error
	}
	listDeployedProductsReturnsOnCall map[int]struct {
		result1 []httpclient.DeployedProduct
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOpsManager) GetBOSHCredentials() (om.BoshCredentials, error) {
	fake.getBOSHCredentialsMutex.Lock()
	ret, specificReturn := fake.getBOSHCredentialsReturnsOnCall[len(fake.getBOSHCredentialsArgsForCall)]
	fake.getBOSHCredentialsArgsForCall = append(fake.getBOSHCredentialsArgsForCall, struct {
	}{})
	stub := fake.GetBOSHCredentialsStub
	fakeReturns := fake.getBOSHCredentialsReturns
	fake.recordInvocation("GetBOSHCredentials", []interface{}{})
	fake.getBOSHCredentialsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) GetBOSHCredentialsCallCount() int {
	fake.getBOSHCredentialsMutex.RLock()
	defer fake.getBOSHCredentialsMutex.RUnlock()
	return len(fake.getBOSHCredentialsArgsForCall)
}

func (fake *FakeOpsManager) GetBOSHCredentialsCalls(stub func() (om.BoshCredentials, error)) {
	fake.getBOSHCredentialsMutex.Lock()
	defer fake.getBOSHCredentialsMutex.Unlock()
	fake.GetBOSHCredentialsStub = stub
}

func (fake *FakeOpsManager) GetBOSHCredentialsReturns(result1 om.BoshCredentials, result2 error) {
	fake.getBOSHCredentialsMutex.Lock()
	defer fake.getBOSHCredentialsMutex.Unlock()
	fake.GetBOSHCredentialsStub = nil
	fake.getBOSHCredentialsReturns = struct {
		result1 om.BoshCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetBOSHCredentialsReturnsOnCall(i int, result1 om.BoshCredentials, result2 error) {
	fake.getBOSHCredentialsMutex.Lock()
	defer fake.getBOSHCredentialsMutex.Unlock()
	fake.GetBOSHCredentialsStub = nil
	if fake.getBOSHCredentialsReturnsOnCall == nil {
		fake.getBOSHCredentialsReturnsOnCall = make(map[int]struct {
			result1 om.BoshCredentials
			result2 error
		})
	}
	fake.getBOSHCredentialsReturnsOnCall[i] = struct {
		result1 om.BoshCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetCertificateAuthorities() ([]httpclient.CA, error) {
	fake.getCertificateAuthoritiesMutex.Lock()
	ret, specificReturn := fake.getCertificateAuthoritiesReturnsOnCall[len(fake.getCertificateAuthoritiesArgsForCall)]
	fake.getCertificateAuthoritiesArgsForCall = append(fake.getCertificateAuthoritiesArgsForCall, struct {
	}{})
	stub := fake.GetCertificateAuthoritiesStub
	fakeReturns := fake.getCertificateAuthoritiesReturns
	fake.recordInvocation("GetCertificateAuthorities", []interface{}{})
	fake.getCertificateAuthoritiesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) GetCertificateAuthoritiesCallCount() int {
	fake.getCertificateAuthoritiesMutex.RLock()
	defer fake.getCertificateAuthoritiesMutex.RUnlock()
	return len(fake.getCertificateAuthoritiesArgsForCall)
}

func (fake *FakeOpsManager) GetCertificateAuthoritiesCalls(stub func() ([]httpclient.CA, error)) {
	fake.getCertificateAuthoritiesMutex.Lock()
	defer fake.getCertificateAuthoritiesMutex.Unlock()
	fake.GetCertificateAuthoritiesStub = stub
}

func (fake *FakeOpsManager) GetCertificateAuthoritiesReturns(result1 []httpclient.CA, result2 error) {
	fake.getCertificateAuthoritiesMutex.Lock()
	defer fake.getCertificateAuthoritiesMutex.Unlock()
	fake.GetCertificateAuthoritiesStub = nil
	fake.getCertificateAuthoritiesReturns = struct {
		result1 []httpclient.CA
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetCertificateAuthoritiesReturnsOnCall(i int, result1 []httpclient.CA, result2 error) {
	fake.getCertificateAuthoritiesMutex.Lock()
	defer fake.getCertificateAuthoritiesMutex.Unlock()
	fake.GetCertificateAuthoritiesStub = nil
	if fake.getCertificateAuthoritiesReturnsOnCall == nil {
		fake.getCertificateAuthoritiesReturnsOnCall = make(map[int]struct {
			result1 []httpclient.CA
			result2 error
		})
	}
	fake.getCertificateAuthoritiesReturnsOnCall[i] = struct {
		result1 []httpclient.CA
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetDeployedProductCredentials(arg1 string, arg2 string) (httpclient.DeployedProductCredential, error) {
	fake.getDeployedProductCredentialsMutex.Lock()
	ret, specificReturn := fake.getDeployedProductCredentialsReturnsOnCall[len(fake.getDeployedProductCredentialsArgsForCall)]
	fake.getDeployedProductCredentialsArgsForCall = append(fake.getDeployedProductCredentialsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetDeployedProductCredentialsStub
	fakeReturns := fake.getDeployedProductCredentialsReturns
	fake.recordInvocation("GetDeployedProductCredentials", []interface{}{arg1, arg2})
	fake.getDeployedProductCredentialsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) GetDeployedProductCredentialsCallCount() int {
	fake.getDeployedProductCredentialsMutex.RLock()
	defer fake.getDeployedProductCredentialsMutex.RUnlock()
	return len(fake.getDeployedProductCredentialsArgsForCall)
}

func (fake *FakeOpsManager) GetDeployedProductCredentialsCalls(stub func(string, string) (httpclient.DeployedProductCredential, error)) {
	fake.getDeployedProductCredentialsMutex.Lock()
	defer fake.getDeployedProductCredentialsMutex.Unlock()
	fake.GetDeployedProductCredentialsStub = stub
}

func (fake *FakeOpsManager) GetDeployedProductCredentialsArgsForCall(i int) (string, string) {
	fake.getDeployedProductCredentialsMutex.RLock()
	defer fake.getDeployedProductCredentialsMutex.RUnlock()
	argsForCall := fake.getDeployedProductCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOpsManager) GetDeployedProductCredentialsReturns(result1 httpclient.DeployedProductCredential, result2 error) {
	fake.getDeployedProductCredentialsMutex.Lock()
	defer fake.getDeployedProductCredentialsMutex.Unlock()
	fake.GetDeployedProductCredentialsStub = nil
	fake.getDeployedProductCredentialsReturns = struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetDeployedProductCredentialsReturnsOnCall(i int, result1 httpclient.DeployedProductCredential, result2 error) {
	fake.getDeployedProductCredentialsMutex.Lock()
	defer fake.getDeployedProductCredentialsMutex.Unlock()
	fake.GetDeployedProductCredentialsStub = nil
	if fake.getDeployedProductCredentialsReturnsOnCall == nil {
		fake.getDeployedProductCredentialsReturnsOnCall = make(map[int]struct {
			result1 httpclient.DeployedProductCredential
			result2 error
		})
	}
	fake.getDeployedProductCredentialsReturnsOnCall[i] = struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetStagedProductProperties(arg1 string) (map[string]httpclient.ResponseProperty, error) {
	fake.getStagedProductPropertiesMutex.Lock()
	ret, specificReturn := fake.getStagedProductPropertiesReturnsOnCall[len(fake.getStagedProductPropertiesArgsForCall)]
	fake.getStagedProductPropertiesArgsForCall = append(fake.getStagedProductPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStagedProductPropertiesStub
	fakeReturns := fake.getStagedProductPropertiesReturns
	fake.recordInvocation("GetStagedProductProperties", []interface{}{arg1})
	fake.getStagedProductPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) GetStagedProductPropertiesCallCount() int {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	return len(fake.getStagedProductPropertiesArgsForCall)
}

func (fake *FakeOpsManager) GetStagedProductPropertiesCalls(stub func(string) (map[string]httpclient.ResponseProperty, error)) {
	fake.getStagedProductPropertiesMutex.Lock()
	defer fake.getStagedProductPropertiesMutex.Unlock()
	fake.GetStagedProductPropertiesStub = stub
}

func (fake *FakeOpsManager) GetStagedProductPropertiesArgsForCall(i int) string {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	argsForCall := fake.getStagedProductPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsManager) GetStagedProductPropertiesReturns(result1 map[string]httpclient.ResponseProperty, result2 error) {
	fake.getStagedProductPropertiesMutex.Lock()
	defer fake.getStagedProductPropertiesMutex.Unlock()
	fake.GetStagedProductPropertiesStub = nil
	fake.getStagedProductPropertiesReturns = struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) GetStagedProductPropertiesReturnsOnCall(i int, result1 map[string]httpclient.ResponseProperty, result2 error) {
	fake.getStagedProductPropertiesMutex.Lock()
	defer fake.getStagedProductPropertiesMutex.Unlock()
	fake.GetStagedProductPropertiesStub = nil
	if fake.getStagedProductPropertiesReturnsOnCall == nil {
		fake.getStagedProductPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]httpclient.ResponseProperty
			result2 error
		})
	}
	fake.getStagedProductPropertiesReturnsOnCall[i] = struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) ListDeployedProducts() ([]httpclient.DeployedProduct, error) {
	fake.listDeployedProductsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductsReturnsOnCall[len(fake.listDeployedProductsArgsForCall)]
	fake.listDeployedProductsArgsForCall = append(fake.listDeployedProductsArgsForCall, struct {
	}{})
	stub := fake.ListDeployedProductsStub
	fakeReturns := fake.listDeployedProductsReturns
	fake.recordInvocation("ListDeployedProducts", []interface{}{})
	fake.listDeployedProductsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpsManager) ListDeployedProductsCallCount() int {
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	return len(fake.listDeployedProductsArgsForCall)
}

func (fake *FakeOpsManager) ListDeployedProductsCalls(stub func() ([]httpclient.DeployedProduct, error)) {
	fake.listDeployedProductsMutex.Lock()
	defer fake.listDeployedProductsMutex.Unlock()
	fake.ListDeployedProductsStub = stub
}

func (fake *FakeOpsManager) ListDeployedProductsReturns(result1 []httpclient.DeployedProduct, result2 error) {
	fake.listDeployedProductsMutex.Lock()
	defer fake.listDeployedProductsMutex.Unlock()
	fake.ListDeployedProductsStub = nil
	fake.listDeployedProductsReturns = struct {
		result1 []httpclient.DeployedProduct
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) ListDeployedProductsReturnsOnCall(i int, result1 []httpclient.DeployedProduct, result2 error) {
	fake.listDeployedProductsMutex.Lock()
	defer fake.listDeployedProductsMutex.Unlock()
	fake.ListDeployedProductsStub = nil
	if fake.listDeployedProductsReturnsOnCall == nil {
		fake.listDeployedProductsReturnsOnCall = make(map[int]struct {
			result1 []httpclient.DeployedProduct
			result2 error
		})
	}
	fake.listDeployedProductsReturnsOnCall[i] = struct {
		result1 []httpclient.DeployedProduct
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBOSHCredentialsMutex.RLock()
	defer fake.getBOSHCredentialsMutex.RUnlock()
	fake.getCertificateAuthoritiesMutex.RLock()
	defer fake.getCertificateAuthoritiesMutex.RUnlock()
	fake.getDeployedProductCredentialsMutex.RLock()
	defer fake.getDeployedProductCredentialsMutex.RUnlock()
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOpsManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ om.OpsManager = new(FakeOpsManager)