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

type FakeClient struct {
	BoshEnvironmentStub        func() (string, string, string, error)
	boshEnvironmentMutex       sync.RWMutex
	boshEnvironmentArgsForCall []struct {
	}
	boshEnvironmentReturns struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}
	boshEnvironmentReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}
	CertificateAuthoritiesStub        func() ([]httpclient.CA, error)
	certificateAuthoritiesMutex       sync.RWMutex
	certificateAuthoritiesArgsForCall []struct {
	}
	certificateAuthoritiesReturns struct {
		result1 []httpclient.CA
		result2 error
	}
	certificateAuthoritiesReturnsOnCall map[int]struct {
		result1 []httpclient.CA
		result2 error
	}
	DeployedProductStub        func(string) (string, error)
	deployedProductMutex       sync.RWMutex
	deployedProductArgsForCall []struct {
		arg1 string
	}
	deployedProductReturns struct {
		result1 string
		result2 error
	}
	deployedProductReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DeployedProductCredentialsStub        func(string, string) (httpclient.DeployedProductCredential, error)
	deployedProductCredentialsMutex       sync.RWMutex
	deployedProductCredentialsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deployedProductCredentialsReturns struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}
	deployedProductCredentialsReturnsOnCall map[int]struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}
	StagedProductPropertiesStub        func(string) (map[string]httpclient.ResponseProperty, error)
	stagedProductPropertiesMutex       sync.RWMutex
	stagedProductPropertiesArgsForCall []struct {
		arg1 string
	}
	stagedProductPropertiesReturns struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}
	stagedProductPropertiesReturnsOnCall map[int]struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) BoshEnvironment() (string, string, string, error) {
	fake.boshEnvironmentMutex.Lock()
	ret, specificReturn := fake.boshEnvironmentReturnsOnCall[len(fake.boshEnvironmentArgsForCall)]
	fake.boshEnvironmentArgsForCall = append(fake.boshEnvironmentArgsForCall, struct {
	}{})
	stub := fake.BoshEnvironmentStub
	fakeReturns := fake.boshEnvironmentReturns
	fake.recordInvocation("BoshEnvironment", []interface{}{})
	fake.boshEnvironmentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeClient) BoshEnvironmentCallCount() int {
	fake.boshEnvironmentMutex.RLock()
	defer fake.boshEnvironmentMutex.RUnlock()
	return len(fake.boshEnvironmentArgsForCall)
}

func (fake *FakeClient) BoshEnvironmentCalls(stub func() (string, string, string, error)) {
	fake.boshEnvironmentMutex.Lock()
	defer fake.boshEnvironmentMutex.Unlock()
	fake.BoshEnvironmentStub = stub
}

func (fake *FakeClient) BoshEnvironmentReturns(result1 string, result2 string, result3 string, result4 error) {
	fake.boshEnvironmentMutex.Lock()
	defer fake.boshEnvironmentMutex.Unlock()
	fake.BoshEnvironmentStub = nil
	fake.boshEnvironmentReturns = struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeClient) BoshEnvironmentReturnsOnCall(i int, result1 string, result2 string, result3 string, result4 error) {
	fake.boshEnvironmentMutex.Lock()
	defer fake.boshEnvironmentMutex.Unlock()
	fake.BoshEnvironmentStub = nil
	if fake.boshEnvironmentReturnsOnCall == nil {
		fake.boshEnvironmentReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 string
			result4 error
		})
	}
	fake.boshEnvironmentReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeClient) CertificateAuthorities() ([]httpclient.CA, error) {
	fake.certificateAuthoritiesMutex.Lock()
	ret, specificReturn := fake.certificateAuthoritiesReturnsOnCall[len(fake.certificateAuthoritiesArgsForCall)]
	fake.certificateAuthoritiesArgsForCall = append(fake.certificateAuthoritiesArgsForCall, struct {
	}{})
	stub := fake.CertificateAuthoritiesStub
	fakeReturns := fake.certificateAuthoritiesReturns
	fake.recordInvocation("CertificateAuthorities", []interface{}{})
	fake.certificateAuthoritiesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CertificateAuthoritiesCallCount() int {
	fake.certificateAuthoritiesMutex.RLock()
	defer fake.certificateAuthoritiesMutex.RUnlock()
	return len(fake.certificateAuthoritiesArgsForCall)
}

func (fake *FakeClient) CertificateAuthoritiesCalls(stub func() ([]httpclient.CA, error)) {
	fake.certificateAuthoritiesMutex.Lock()
	defer fake.certificateAuthoritiesMutex.Unlock()
	fake.CertificateAuthoritiesStub = stub
}

func (fake *FakeClient) CertificateAuthoritiesReturns(result1 []httpclient.CA, result2 error) {
	fake.certificateAuthoritiesMutex.Lock()
	defer fake.certificateAuthoritiesMutex.Unlock()
	fake.CertificateAuthoritiesStub = nil
	fake.certificateAuthoritiesReturns = struct {
		result1 []httpclient.CA
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CertificateAuthoritiesReturnsOnCall(i int, result1 []httpclient.CA, result2 error) {
	fake.certificateAuthoritiesMutex.Lock()
	defer fake.certificateAuthoritiesMutex.Unlock()
	fake.CertificateAuthoritiesStub = nil
	if fake.certificateAuthoritiesReturnsOnCall == nil {
		fake.certificateAuthoritiesReturnsOnCall = make(map[int]struct {
			result1 []httpclient.CA
			result2 error
		})
	}
	fake.certificateAuthoritiesReturnsOnCall[i] = struct {
		result1 []httpclient.CA
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeployedProduct(arg1 string) (string, error) {
	fake.deployedProductMutex.Lock()
	ret, specificReturn := fake.deployedProductReturnsOnCall[len(fake.deployedProductArgsForCall)]
	fake.deployedProductArgsForCall = append(fake.deployedProductArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeployedProductStub
	fakeReturns := fake.deployedProductReturns
	fake.recordInvocation("DeployedProduct", []interface{}{arg1})
	fake.deployedProductMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) DeployedProductCallCount() int {
	fake.deployedProductMutex.RLock()
	defer fake.deployedProductMutex.RUnlock()
	return len(fake.deployedProductArgsForCall)
}

func (fake *FakeClient) DeployedProductCalls(stub func(string) (string, error)) {
	fake.deployedProductMutex.Lock()
	defer fake.deployedProductMutex.Unlock()
	fake.DeployedProductStub = stub
}

func (fake *FakeClient) DeployedProductArgsForCall(i int) string {
	fake.deployedProductMutex.RLock()
	defer fake.deployedProductMutex.RUnlock()
	argsForCall := fake.deployedProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DeployedProductReturns(result1 string, result2 error) {
	fake.deployedProductMutex.Lock()
	defer fake.deployedProductMutex.Unlock()
	fake.DeployedProductStub = nil
	fake.deployedProductReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeployedProductReturnsOnCall(i int, result1 string, result2 error) {
	fake.deployedProductMutex.Lock()
	defer fake.deployedProductMutex.Unlock()
	fake.DeployedProductStub = nil
	if fake.deployedProductReturnsOnCall == nil {
		fake.deployedProductReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.deployedProductReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeployedProductCredentials(arg1 string, arg2 string) (httpclient.DeployedProductCredential, error) {
	fake.deployedProductCredentialsMutex.Lock()
	ret, specificReturn := fake.deployedProductCredentialsReturnsOnCall[len(fake.deployedProductCredentialsArgsForCall)]
	fake.deployedProductCredentialsArgsForCall = append(fake.deployedProductCredentialsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DeployedProductCredentialsStub
	fakeReturns := fake.deployedProductCredentialsReturns
	fake.recordInvocation("DeployedProductCredentials", []interface{}{arg1, arg2})
	fake.deployedProductCredentialsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) DeployedProductCredentialsCallCount() int {
	fake.deployedProductCredentialsMutex.RLock()
	defer fake.deployedProductCredentialsMutex.RUnlock()
	return len(fake.deployedProductCredentialsArgsForCall)
}

func (fake *FakeClient) DeployedProductCredentialsCalls(stub func(string, string) (httpclient.DeployedProductCredential, error)) {
	fake.deployedProductCredentialsMutex.Lock()
	defer fake.deployedProductCredentialsMutex.Unlock()
	fake.DeployedProductCredentialsStub = stub
}

func (fake *FakeClient) DeployedProductCredentialsArgsForCall(i int) (string, string) {
	fake.deployedProductCredentialsMutex.RLock()
	defer fake.deployedProductCredentialsMutex.RUnlock()
	argsForCall := fake.deployedProductCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DeployedProductCredentialsReturns(result1 httpclient.DeployedProductCredential, result2 error) {
	fake.deployedProductCredentialsMutex.Lock()
	defer fake.deployedProductCredentialsMutex.Unlock()
	fake.DeployedProductCredentialsStub = nil
	fake.deployedProductCredentialsReturns = struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeployedProductCredentialsReturnsOnCall(i int, result1 httpclient.DeployedProductCredential, result2 error) {
	fake.deployedProductCredentialsMutex.Lock()
	defer fake.deployedProductCredentialsMutex.Unlock()
	fake.DeployedProductCredentialsStub = nil
	if fake.deployedProductCredentialsReturnsOnCall == nil {
		fake.deployedProductCredentialsReturnsOnCall = make(map[int]struct {
			result1 httpclient.DeployedProductCredential
			result2 error
		})
	}
	fake.deployedProductCredentialsReturnsOnCall[i] = struct {
		result1 httpclient.DeployedProductCredential
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) StagedProductProperties(arg1 string) (map[string]httpclient.ResponseProperty, error) {
	fake.stagedProductPropertiesMutex.Lock()
	ret, specificReturn := fake.stagedProductPropertiesReturnsOnCall[len(fake.stagedProductPropertiesArgsForCall)]
	fake.stagedProductPropertiesArgsForCall = append(fake.stagedProductPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StagedProductPropertiesStub
	fakeReturns := fake.stagedProductPropertiesReturns
	fake.recordInvocation("StagedProductProperties", []interface{}{arg1})
	fake.stagedProductPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) StagedProductPropertiesCallCount() int {
	fake.stagedProductPropertiesMutex.RLock()
	defer fake.stagedProductPropertiesMutex.RUnlock()
	return len(fake.stagedProductPropertiesArgsForCall)
}

func (fake *FakeClient) StagedProductPropertiesCalls(stub func(string) (map[string]httpclient.ResponseProperty, error)) {
	fake.stagedProductPropertiesMutex.Lock()
	defer fake.stagedProductPropertiesMutex.Unlock()
	fake.StagedProductPropertiesStub = stub
}

func (fake *FakeClient) StagedProductPropertiesArgsForCall(i int) string {
	fake.stagedProductPropertiesMutex.RLock()
	defer fake.stagedProductPropertiesMutex.RUnlock()
	argsForCall := fake.stagedProductPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) StagedProductPropertiesReturns(result1 map[string]httpclient.ResponseProperty, result2 error) {
	fake.stagedProductPropertiesMutex.Lock()
	defer fake.stagedProductPropertiesMutex.Unlock()
	fake.StagedProductPropertiesStub = nil
	fake.stagedProductPropertiesReturns = struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) StagedProductPropertiesReturnsOnCall(i int, result1 map[string]httpclient.ResponseProperty, result2 error) {
	fake.stagedProductPropertiesMutex.Lock()
	defer fake.stagedProductPropertiesMutex.Unlock()
	fake.StagedProductPropertiesStub = nil
	if fake.stagedProductPropertiesReturnsOnCall == nil {
		fake.stagedProductPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]httpclient.ResponseProperty
			result2 error
		})
	}
	fake.stagedProductPropertiesReturnsOnCall[i] = struct {
		result1 map[string]httpclient.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.boshEnvironmentMutex.RLock()
	defer fake.boshEnvironmentMutex.RUnlock()
	fake.certificateAuthoritiesMutex.RLock()
	defer fake.certificateAuthoritiesMutex.RUnlock()
	fake.deployedProductMutex.RLock()
	defer fake.deployedProductMutex.RUnlock()
	fake.deployedProductCredentialsMutex.RLock()
	defer fake.deployedProductCredentialsMutex.RUnlock()
	fake.stagedProductPropertiesMutex.RLock()
	defer fake.stagedProductPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ om.Client = new(FakeClient)
