// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/config"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/om"
)

type FakeClientFactory struct {
	NewStub        func(string, []byte, om.CertAppender, om.OpsManagerFactory, om.UAAFactory, config.Authentication) (om.Client, error)
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 om.CertAppender
		arg4 om.OpsManagerFactory
		arg5 om.UAAFactory
		arg6 config.Authentication
	}
	newReturns struct {
		result1 om.Client
		result2 error
	}
	newReturnsOnCall map[int]struct {
		result1 om.Client
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientFactory) New(arg1 string, arg2 []byte, arg3 om.CertAppender, arg4 om.OpsManagerFactory, arg5 om.UAAFactory, arg6 config.Authentication) (om.Client, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.newMutex.Lock()
	ret, specificReturn := fake.newReturnsOnCall[len(fake.newArgsForCall)]
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 om.CertAppender
		arg4 om.OpsManagerFactory
		arg5 om.UAAFactory
		arg6 config.Authentication
	}{arg1, arg2Copy, arg3, arg4, arg5, arg6})
	stub := fake.NewStub
	fakeReturns := fake.newReturns
	fake.recordInvocation("New", []interface{}{arg1, arg2Copy, arg3, arg4, arg5, arg6})
	fake.newMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientFactory) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeClientFactory) NewCalls(stub func(string, []byte, om.CertAppender, om.OpsManagerFactory, om.UAAFactory, config.Authentication) (om.Client, error)) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = stub
}

func (fake *FakeClientFactory) NewArgsForCall(i int) (string, []byte, om.CertAppender, om.OpsManagerFactory, om.UAAFactory, config.Authentication) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	argsForCall := fake.newArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeClientFactory) NewReturns(result1 om.Client, result2 error) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 om.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClientFactory) NewReturnsOnCall(i int, result1 om.Client, result2 error) {
	fake.newMutex.Lock()
	defer fake.newMutex.Unlock()
	fake.NewStub = nil
	if fake.newReturnsOnCall == nil {
		fake.newReturnsOnCall = make(map[int]struct {
			result1 om.Client
			result2 error
		})
	}
	fake.newReturnsOnCall[i] = struct {
		result1 om.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ om.ClientFactory = new(FakeClientFactory)
