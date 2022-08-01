// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/migrate/cc"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/migrate/cc/db"
)

type FakeCloudControllerRepositoryFactory struct {
	NewCCDBStub        func(bool) (db.Repository, error)
	newCCDBMutex       sync.RWMutex
	newCCDBArgsForCall []struct {
		arg1 bool
	}
	newCCDBReturns struct {
		result1 db.Repository
		result2 error
	}
	newCCDBReturnsOnCall map[int]struct {
		result1 db.Repository
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudControllerRepositoryFactory) NewCCDB(arg1 bool) (db.Repository, error) {
	fake.newCCDBMutex.Lock()
	ret, specificReturn := fake.newCCDBReturnsOnCall[len(fake.newCCDBArgsForCall)]
	fake.newCCDBArgsForCall = append(fake.newCCDBArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.NewCCDBStub
	fakeReturns := fake.newCCDBReturns
	fake.recordInvocation("NewCCDB", []interface{}{arg1})
	fake.newCCDBMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudControllerRepositoryFactory) NewCCDBCallCount() int {
	fake.newCCDBMutex.RLock()
	defer fake.newCCDBMutex.RUnlock()
	return len(fake.newCCDBArgsForCall)
}

func (fake *FakeCloudControllerRepositoryFactory) NewCCDBCalls(stub func(bool) (db.Repository, error)) {
	fake.newCCDBMutex.Lock()
	defer fake.newCCDBMutex.Unlock()
	fake.NewCCDBStub = stub
}

func (fake *FakeCloudControllerRepositoryFactory) NewCCDBArgsForCall(i int) bool {
	fake.newCCDBMutex.RLock()
	defer fake.newCCDBMutex.RUnlock()
	argsForCall := fake.newCCDBArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCloudControllerRepositoryFactory) NewCCDBReturns(result1 db.Repository, result2 error) {
	fake.newCCDBMutex.Lock()
	defer fake.newCCDBMutex.Unlock()
	fake.NewCCDBStub = nil
	fake.newCCDBReturns = struct {
		result1 db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudControllerRepositoryFactory) NewCCDBReturnsOnCall(i int, result1 db.Repository, result2 error) {
	fake.newCCDBMutex.Lock()
	defer fake.newCCDBMutex.Unlock()
	fake.NewCCDBStub = nil
	if fake.newCCDBReturnsOnCall == nil {
		fake.newCCDBReturnsOnCall = make(map[int]struct {
			result1 db.Repository
			result2 error
		})
	}
	fake.newCCDBReturnsOnCall[i] = struct {
		result1 db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudControllerRepositoryFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newCCDBMutex.RLock()
	defer fake.newCCDBMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCloudControllerRepositoryFactory) recordInvocation(key string, args []interface{}) {
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

var _ cc.CloudControllerRepositoryFactory = new(FakeCloudControllerRepositoryFactory)
