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
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/cf"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/migrate"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/om"
)

type FakeClientHolder struct {
	CFClientStub        func(bool) cf.Client
	cFClientMutex       sync.RWMutex
	cFClientArgsForCall []struct {
		arg1 bool
	}
	cFClientReturns struct {
		result1 cf.Client
	}
	cFClientReturnsOnCall map[int]struct {
		result1 cf.Client
	}
	SourceBoshClientStub        func() bosh.Client
	sourceBoshClientMutex       sync.RWMutex
	sourceBoshClientArgsForCall []struct {
	}
	sourceBoshClientReturns struct {
		result1 bosh.Client
	}
	sourceBoshClientReturnsOnCall map[int]struct {
		result1 bosh.Client
	}
	SourceCFClientStub        func() cf.Client
	sourceCFClientMutex       sync.RWMutex
	sourceCFClientArgsForCall []struct {
	}
	sourceCFClientReturns struct {
		result1 cf.Client
	}
	sourceCFClientReturnsOnCall map[int]struct {
		result1 cf.Client
	}
	SourceOpsManClientStub        func() om.Client
	sourceOpsManClientMutex       sync.RWMutex
	sourceOpsManClientArgsForCall []struct {
	}
	sourceOpsManClientReturns struct {
		result1 om.Client
	}
	sourceOpsManClientReturnsOnCall map[int]struct {
		result1 om.Client
	}
	TargetBoshClientStub        func() bosh.Client
	targetBoshClientMutex       sync.RWMutex
	targetBoshClientArgsForCall []struct {
	}
	targetBoshClientReturns struct {
		result1 bosh.Client
	}
	targetBoshClientReturnsOnCall map[int]struct {
		result1 bosh.Client
	}
	TargetCFClientStub        func() cf.Client
	targetCFClientMutex       sync.RWMutex
	targetCFClientArgsForCall []struct {
	}
	targetCFClientReturns struct {
		result1 cf.Client
	}
	targetCFClientReturnsOnCall map[int]struct {
		result1 cf.Client
	}
	TargetOpsManClientStub        func() om.Client
	targetOpsManClientMutex       sync.RWMutex
	targetOpsManClientArgsForCall []struct {
	}
	targetOpsManClientReturns struct {
		result1 om.Client
	}
	targetOpsManClientReturnsOnCall map[int]struct {
		result1 om.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientHolder) CFClient(arg1 bool) cf.Client {
	fake.cFClientMutex.Lock()
	ret, specificReturn := fake.cFClientReturnsOnCall[len(fake.cFClientArgsForCall)]
	fake.cFClientArgsForCall = append(fake.cFClientArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.CFClientStub
	fakeReturns := fake.cFClientReturns
	fake.recordInvocation("CFClient", []interface{}{arg1})
	fake.cFClientMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) CFClientCallCount() int {
	fake.cFClientMutex.RLock()
	defer fake.cFClientMutex.RUnlock()
	return len(fake.cFClientArgsForCall)
}

func (fake *FakeClientHolder) CFClientCalls(stub func(bool) cf.Client) {
	fake.cFClientMutex.Lock()
	defer fake.cFClientMutex.Unlock()
	fake.CFClientStub = stub
}

func (fake *FakeClientHolder) CFClientArgsForCall(i int) bool {
	fake.cFClientMutex.RLock()
	defer fake.cFClientMutex.RUnlock()
	argsForCall := fake.cFClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClientHolder) CFClientReturns(result1 cf.Client) {
	fake.cFClientMutex.Lock()
	defer fake.cFClientMutex.Unlock()
	fake.CFClientStub = nil
	fake.cFClientReturns = struct {
		result1 cf.Client
	}{result1}
}

func (fake *FakeClientHolder) CFClientReturnsOnCall(i int, result1 cf.Client) {
	fake.cFClientMutex.Lock()
	defer fake.cFClientMutex.Unlock()
	fake.CFClientStub = nil
	if fake.cFClientReturnsOnCall == nil {
		fake.cFClientReturnsOnCall = make(map[int]struct {
			result1 cf.Client
		})
	}
	fake.cFClientReturnsOnCall[i] = struct {
		result1 cf.Client
	}{result1}
}

func (fake *FakeClientHolder) SourceBoshClient() bosh.Client {
	fake.sourceBoshClientMutex.Lock()
	ret, specificReturn := fake.sourceBoshClientReturnsOnCall[len(fake.sourceBoshClientArgsForCall)]
	fake.sourceBoshClientArgsForCall = append(fake.sourceBoshClientArgsForCall, struct {
	}{})
	stub := fake.SourceBoshClientStub
	fakeReturns := fake.sourceBoshClientReturns
	fake.recordInvocation("SourceBoshClient", []interface{}{})
	fake.sourceBoshClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) SourceBoshClientCallCount() int {
	fake.sourceBoshClientMutex.RLock()
	defer fake.sourceBoshClientMutex.RUnlock()
	return len(fake.sourceBoshClientArgsForCall)
}

func (fake *FakeClientHolder) SourceBoshClientCalls(stub func() bosh.Client) {
	fake.sourceBoshClientMutex.Lock()
	defer fake.sourceBoshClientMutex.Unlock()
	fake.SourceBoshClientStub = stub
}

func (fake *FakeClientHolder) SourceBoshClientReturns(result1 bosh.Client) {
	fake.sourceBoshClientMutex.Lock()
	defer fake.sourceBoshClientMutex.Unlock()
	fake.SourceBoshClientStub = nil
	fake.sourceBoshClientReturns = struct {
		result1 bosh.Client
	}{result1}
}

func (fake *FakeClientHolder) SourceBoshClientReturnsOnCall(i int, result1 bosh.Client) {
	fake.sourceBoshClientMutex.Lock()
	defer fake.sourceBoshClientMutex.Unlock()
	fake.SourceBoshClientStub = nil
	if fake.sourceBoshClientReturnsOnCall == nil {
		fake.sourceBoshClientReturnsOnCall = make(map[int]struct {
			result1 bosh.Client
		})
	}
	fake.sourceBoshClientReturnsOnCall[i] = struct {
		result1 bosh.Client
	}{result1}
}

func (fake *FakeClientHolder) SourceCFClient() cf.Client {
	fake.sourceCFClientMutex.Lock()
	ret, specificReturn := fake.sourceCFClientReturnsOnCall[len(fake.sourceCFClientArgsForCall)]
	fake.sourceCFClientArgsForCall = append(fake.sourceCFClientArgsForCall, struct {
	}{})
	stub := fake.SourceCFClientStub
	fakeReturns := fake.sourceCFClientReturns
	fake.recordInvocation("SourceCFClient", []interface{}{})
	fake.sourceCFClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) SourceCFClientCallCount() int {
	fake.sourceCFClientMutex.RLock()
	defer fake.sourceCFClientMutex.RUnlock()
	return len(fake.sourceCFClientArgsForCall)
}

func (fake *FakeClientHolder) SourceCFClientCalls(stub func() cf.Client) {
	fake.sourceCFClientMutex.Lock()
	defer fake.sourceCFClientMutex.Unlock()
	fake.SourceCFClientStub = stub
}

func (fake *FakeClientHolder) SourceCFClientReturns(result1 cf.Client) {
	fake.sourceCFClientMutex.Lock()
	defer fake.sourceCFClientMutex.Unlock()
	fake.SourceCFClientStub = nil
	fake.sourceCFClientReturns = struct {
		result1 cf.Client
	}{result1}
}

func (fake *FakeClientHolder) SourceCFClientReturnsOnCall(i int, result1 cf.Client) {
	fake.sourceCFClientMutex.Lock()
	defer fake.sourceCFClientMutex.Unlock()
	fake.SourceCFClientStub = nil
	if fake.sourceCFClientReturnsOnCall == nil {
		fake.sourceCFClientReturnsOnCall = make(map[int]struct {
			result1 cf.Client
		})
	}
	fake.sourceCFClientReturnsOnCall[i] = struct {
		result1 cf.Client
	}{result1}
}

func (fake *FakeClientHolder) SourceOpsManClient() om.Client {
	fake.sourceOpsManClientMutex.Lock()
	ret, specificReturn := fake.sourceOpsManClientReturnsOnCall[len(fake.sourceOpsManClientArgsForCall)]
	fake.sourceOpsManClientArgsForCall = append(fake.sourceOpsManClientArgsForCall, struct {
	}{})
	stub := fake.SourceOpsManClientStub
	fakeReturns := fake.sourceOpsManClientReturns
	fake.recordInvocation("SourceOpsManClient", []interface{}{})
	fake.sourceOpsManClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) SourceOpsManClientCallCount() int {
	fake.sourceOpsManClientMutex.RLock()
	defer fake.sourceOpsManClientMutex.RUnlock()
	return len(fake.sourceOpsManClientArgsForCall)
}

func (fake *FakeClientHolder) SourceOpsManClientCalls(stub func() om.Client) {
	fake.sourceOpsManClientMutex.Lock()
	defer fake.sourceOpsManClientMutex.Unlock()
	fake.SourceOpsManClientStub = stub
}

func (fake *FakeClientHolder) SourceOpsManClientReturns(result1 om.Client) {
	fake.sourceOpsManClientMutex.Lock()
	defer fake.sourceOpsManClientMutex.Unlock()
	fake.SourceOpsManClientStub = nil
	fake.sourceOpsManClientReturns = struct {
		result1 om.Client
	}{result1}
}

func (fake *FakeClientHolder) SourceOpsManClientReturnsOnCall(i int, result1 om.Client) {
	fake.sourceOpsManClientMutex.Lock()
	defer fake.sourceOpsManClientMutex.Unlock()
	fake.SourceOpsManClientStub = nil
	if fake.sourceOpsManClientReturnsOnCall == nil {
		fake.sourceOpsManClientReturnsOnCall = make(map[int]struct {
			result1 om.Client
		})
	}
	fake.sourceOpsManClientReturnsOnCall[i] = struct {
		result1 om.Client
	}{result1}
}

func (fake *FakeClientHolder) TargetBoshClient() bosh.Client {
	fake.targetBoshClientMutex.Lock()
	ret, specificReturn := fake.targetBoshClientReturnsOnCall[len(fake.targetBoshClientArgsForCall)]
	fake.targetBoshClientArgsForCall = append(fake.targetBoshClientArgsForCall, struct {
	}{})
	stub := fake.TargetBoshClientStub
	fakeReturns := fake.targetBoshClientReturns
	fake.recordInvocation("TargetBoshClient", []interface{}{})
	fake.targetBoshClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) TargetBoshClientCallCount() int {
	fake.targetBoshClientMutex.RLock()
	defer fake.targetBoshClientMutex.RUnlock()
	return len(fake.targetBoshClientArgsForCall)
}

func (fake *FakeClientHolder) TargetBoshClientCalls(stub func() bosh.Client) {
	fake.targetBoshClientMutex.Lock()
	defer fake.targetBoshClientMutex.Unlock()
	fake.TargetBoshClientStub = stub
}

func (fake *FakeClientHolder) TargetBoshClientReturns(result1 bosh.Client) {
	fake.targetBoshClientMutex.Lock()
	defer fake.targetBoshClientMutex.Unlock()
	fake.TargetBoshClientStub = nil
	fake.targetBoshClientReturns = struct {
		result1 bosh.Client
	}{result1}
}

func (fake *FakeClientHolder) TargetBoshClientReturnsOnCall(i int, result1 bosh.Client) {
	fake.targetBoshClientMutex.Lock()
	defer fake.targetBoshClientMutex.Unlock()
	fake.TargetBoshClientStub = nil
	if fake.targetBoshClientReturnsOnCall == nil {
		fake.targetBoshClientReturnsOnCall = make(map[int]struct {
			result1 bosh.Client
		})
	}
	fake.targetBoshClientReturnsOnCall[i] = struct {
		result1 bosh.Client
	}{result1}
}

func (fake *FakeClientHolder) TargetCFClient() cf.Client {
	fake.targetCFClientMutex.Lock()
	ret, specificReturn := fake.targetCFClientReturnsOnCall[len(fake.targetCFClientArgsForCall)]
	fake.targetCFClientArgsForCall = append(fake.targetCFClientArgsForCall, struct {
	}{})
	stub := fake.TargetCFClientStub
	fakeReturns := fake.targetCFClientReturns
	fake.recordInvocation("TargetCFClient", []interface{}{})
	fake.targetCFClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) TargetCFClientCallCount() int {
	fake.targetCFClientMutex.RLock()
	defer fake.targetCFClientMutex.RUnlock()
	return len(fake.targetCFClientArgsForCall)
}

func (fake *FakeClientHolder) TargetCFClientCalls(stub func() cf.Client) {
	fake.targetCFClientMutex.Lock()
	defer fake.targetCFClientMutex.Unlock()
	fake.TargetCFClientStub = stub
}

func (fake *FakeClientHolder) TargetCFClientReturns(result1 cf.Client) {
	fake.targetCFClientMutex.Lock()
	defer fake.targetCFClientMutex.Unlock()
	fake.TargetCFClientStub = nil
	fake.targetCFClientReturns = struct {
		result1 cf.Client
	}{result1}
}

func (fake *FakeClientHolder) TargetCFClientReturnsOnCall(i int, result1 cf.Client) {
	fake.targetCFClientMutex.Lock()
	defer fake.targetCFClientMutex.Unlock()
	fake.TargetCFClientStub = nil
	if fake.targetCFClientReturnsOnCall == nil {
		fake.targetCFClientReturnsOnCall = make(map[int]struct {
			result1 cf.Client
		})
	}
	fake.targetCFClientReturnsOnCall[i] = struct {
		result1 cf.Client
	}{result1}
}

func (fake *FakeClientHolder) TargetOpsManClient() om.Client {
	fake.targetOpsManClientMutex.Lock()
	ret, specificReturn := fake.targetOpsManClientReturnsOnCall[len(fake.targetOpsManClientArgsForCall)]
	fake.targetOpsManClientArgsForCall = append(fake.targetOpsManClientArgsForCall, struct {
	}{})
	stub := fake.TargetOpsManClientStub
	fakeReturns := fake.targetOpsManClientReturns
	fake.recordInvocation("TargetOpsManClient", []interface{}{})
	fake.targetOpsManClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClientHolder) TargetOpsManClientCallCount() int {
	fake.targetOpsManClientMutex.RLock()
	defer fake.targetOpsManClientMutex.RUnlock()
	return len(fake.targetOpsManClientArgsForCall)
}

func (fake *FakeClientHolder) TargetOpsManClientCalls(stub func() om.Client) {
	fake.targetOpsManClientMutex.Lock()
	defer fake.targetOpsManClientMutex.Unlock()
	fake.TargetOpsManClientStub = stub
}

func (fake *FakeClientHolder) TargetOpsManClientReturns(result1 om.Client) {
	fake.targetOpsManClientMutex.Lock()
	defer fake.targetOpsManClientMutex.Unlock()
	fake.TargetOpsManClientStub = nil
	fake.targetOpsManClientReturns = struct {
		result1 om.Client
	}{result1}
}

func (fake *FakeClientHolder) TargetOpsManClientReturnsOnCall(i int, result1 om.Client) {
	fake.targetOpsManClientMutex.Lock()
	defer fake.targetOpsManClientMutex.Unlock()
	fake.TargetOpsManClientStub = nil
	if fake.targetOpsManClientReturnsOnCall == nil {
		fake.targetOpsManClientReturnsOnCall = make(map[int]struct {
			result1 om.Client
		})
	}
	fake.targetOpsManClientReturnsOnCall[i] = struct {
		result1 om.Client
	}{result1}
}

func (fake *FakeClientHolder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cFClientMutex.RLock()
	defer fake.cFClientMutex.RUnlock()
	fake.sourceBoshClientMutex.RLock()
	defer fake.sourceBoshClientMutex.RUnlock()
	fake.sourceCFClientMutex.RLock()
	defer fake.sourceCFClientMutex.RUnlock()
	fake.sourceOpsManClientMutex.RLock()
	defer fake.sourceOpsManClientMutex.RUnlock()
	fake.targetBoshClientMutex.RLock()
	defer fake.targetBoshClientMutex.RUnlock()
	fake.targetCFClientMutex.RLock()
	defer fake.targetCFClientMutex.RUnlock()
	fake.targetOpsManClientMutex.RLock()
	defer fake.targetOpsManClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClientHolder) recordInvocation(key string, args []interface{}) {
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

var _ migrate.ClientHolder = new(FakeClientHolder)
