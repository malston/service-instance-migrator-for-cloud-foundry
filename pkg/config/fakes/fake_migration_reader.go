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

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/config"
)

type FakeMigrationReader struct {
	GetMigrationStub        func() (*config.Migration, error)
	getMigrationMutex       sync.RWMutex
	getMigrationArgsForCall []struct {
	}
	getMigrationReturns struct {
		result1 *config.Migration
		result2 error
	}
	getMigrationReturnsOnCall map[int]struct {
		result1 *config.Migration
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMigrationReader) GetMigration() (*config.Migration, error) {
	fake.getMigrationMutex.Lock()
	ret, specificReturn := fake.getMigrationReturnsOnCall[len(fake.getMigrationArgsForCall)]
	fake.getMigrationArgsForCall = append(fake.getMigrationArgsForCall, struct {
	}{})
	stub := fake.GetMigrationStub
	fakeReturns := fake.getMigrationReturns
	fake.recordInvocation("GetMigration", []interface{}{})
	fake.getMigrationMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMigrationReader) GetMigrationCallCount() int {
	fake.getMigrationMutex.RLock()
	defer fake.getMigrationMutex.RUnlock()
	return len(fake.getMigrationArgsForCall)
}

func (fake *FakeMigrationReader) GetMigrationCalls(stub func() (*config.Migration, error)) {
	fake.getMigrationMutex.Lock()
	defer fake.getMigrationMutex.Unlock()
	fake.GetMigrationStub = stub
}

func (fake *FakeMigrationReader) GetMigrationReturns(result1 *config.Migration, result2 error) {
	fake.getMigrationMutex.Lock()
	defer fake.getMigrationMutex.Unlock()
	fake.GetMigrationStub = nil
	fake.getMigrationReturns = struct {
		result1 *config.Migration
		result2 error
	}{result1, result2}
}

func (fake *FakeMigrationReader) GetMigrationReturnsOnCall(i int, result1 *config.Migration, result2 error) {
	fake.getMigrationMutex.Lock()
	defer fake.getMigrationMutex.Unlock()
	fake.GetMigrationStub = nil
	if fake.getMigrationReturnsOnCall == nil {
		fake.getMigrationReturnsOnCall = make(map[int]struct {
			result1 *config.Migration
			result2 error
		})
	}
	fake.getMigrationReturnsOnCall[i] = struct {
		result1 *config.Migration
		result2 error
	}{result1, result2}
}

func (fake *FakeMigrationReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMigrationMutex.RLock()
	defer fake.getMigrationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMigrationReader) recordInvocation(key string, args []interface{}) {
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

var _ config.MigrationReader = new(FakeMigrationReader)
