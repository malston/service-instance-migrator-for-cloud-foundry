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

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/bosh"
)

type FakeCertAppender struct {
	AppendCertsFromPEMStub        func([]byte) bool
	appendCertsFromPEMMutex       sync.RWMutex
	appendCertsFromPEMArgsForCall []struct {
		arg1 []byte
	}
	appendCertsFromPEMReturns struct {
		result1 bool
	}
	appendCertsFromPEMReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCertAppender) AppendCertsFromPEM(arg1 []byte) bool {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.appendCertsFromPEMMutex.Lock()
	ret, specificReturn := fake.appendCertsFromPEMReturnsOnCall[len(fake.appendCertsFromPEMArgsForCall)]
	fake.appendCertsFromPEMArgsForCall = append(fake.appendCertsFromPEMArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.AppendCertsFromPEMStub
	fakeReturns := fake.appendCertsFromPEMReturns
	fake.recordInvocation("AppendCertsFromPEM", []interface{}{arg1Copy})
	fake.appendCertsFromPEMMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCertAppender) AppendCertsFromPEMCallCount() int {
	fake.appendCertsFromPEMMutex.RLock()
	defer fake.appendCertsFromPEMMutex.RUnlock()
	return len(fake.appendCertsFromPEMArgsForCall)
}

func (fake *FakeCertAppender) AppendCertsFromPEMCalls(stub func([]byte) bool) {
	fake.appendCertsFromPEMMutex.Lock()
	defer fake.appendCertsFromPEMMutex.Unlock()
	fake.AppendCertsFromPEMStub = stub
}

func (fake *FakeCertAppender) AppendCertsFromPEMArgsForCall(i int) []byte {
	fake.appendCertsFromPEMMutex.RLock()
	defer fake.appendCertsFromPEMMutex.RUnlock()
	argsForCall := fake.appendCertsFromPEMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCertAppender) AppendCertsFromPEMReturns(result1 bool) {
	fake.appendCertsFromPEMMutex.Lock()
	defer fake.appendCertsFromPEMMutex.Unlock()
	fake.AppendCertsFromPEMStub = nil
	fake.appendCertsFromPEMReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCertAppender) AppendCertsFromPEMReturnsOnCall(i int, result1 bool) {
	fake.appendCertsFromPEMMutex.Lock()
	defer fake.appendCertsFromPEMMutex.Unlock()
	fake.AppendCertsFromPEMStub = nil
	if fake.appendCertsFromPEMReturnsOnCall == nil {
		fake.appendCertsFromPEMReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.appendCertsFromPEMReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCertAppender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appendCertsFromPEMMutex.RLock()
	defer fake.appendCertsFromPEMMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCertAppender) recordInvocation(key string, args []interface{}) {
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

var _ bosh.CertAppender = new(FakeCertAppender)
