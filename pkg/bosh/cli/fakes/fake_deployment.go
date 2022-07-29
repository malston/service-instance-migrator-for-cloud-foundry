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

	"github.com/cloudfoundry/bosh-cli/director"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/bosh/cli"
)

type FakeDeployment struct {
	VMInfosStub        func() ([]director.VMInfo, error)
	vMInfosMutex       sync.RWMutex
	vMInfosArgsForCall []struct {
	}
	vMInfosReturns struct {
		result1 []director.VMInfo
		result2 error
	}
	vMInfosReturnsOnCall map[int]struct {
		result1 []director.VMInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeployment) VMInfos() ([]director.VMInfo, error) {
	fake.vMInfosMutex.Lock()
	ret, specificReturn := fake.vMInfosReturnsOnCall[len(fake.vMInfosArgsForCall)]
	fake.vMInfosArgsForCall = append(fake.vMInfosArgsForCall, struct {
	}{})
	stub := fake.VMInfosStub
	fakeReturns := fake.vMInfosReturns
	fake.recordInvocation("VMInfos", []interface{}{})
	fake.vMInfosMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeployment) VMInfosCallCount() int {
	fake.vMInfosMutex.RLock()
	defer fake.vMInfosMutex.RUnlock()
	return len(fake.vMInfosArgsForCall)
}

func (fake *FakeDeployment) VMInfosCalls(stub func() ([]director.VMInfo, error)) {
	fake.vMInfosMutex.Lock()
	defer fake.vMInfosMutex.Unlock()
	fake.VMInfosStub = stub
}

func (fake *FakeDeployment) VMInfosReturns(result1 []director.VMInfo, result2 error) {
	fake.vMInfosMutex.Lock()
	defer fake.vMInfosMutex.Unlock()
	fake.VMInfosStub = nil
	fake.vMInfosReturns = struct {
		result1 []director.VMInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) VMInfosReturnsOnCall(i int, result1 []director.VMInfo, result2 error) {
	fake.vMInfosMutex.Lock()
	defer fake.vMInfosMutex.Unlock()
	fake.VMInfosStub = nil
	if fake.vMInfosReturnsOnCall == nil {
		fake.vMInfosReturnsOnCall = make(map[int]struct {
			result1 []director.VMInfo
			result2 error
		})
	}
	fake.vMInfosReturnsOnCall[i] = struct {
		result1 []director.VMInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.vMInfosMutex.RLock()
	defer fake.vMInfosMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeployment) recordInvocation(key string, args []interface{}) {
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

var _ cli.Deployment = new(FakeDeployment)