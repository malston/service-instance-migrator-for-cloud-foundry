// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/cmd"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/config"
)

type FakeOrgExporter struct {
	ExportStub        func(context.Context, config.OpsManager, string, ...string) error
	exportMutex       sync.RWMutex
	exportArgsForCall []struct {
		arg1 context.Context
		arg2 config.OpsManager
		arg3 string
		arg4 []string
	}
	exportReturns struct {
		result1 error
	}
	exportReturnsOnCall map[int]struct {
		result1 error
	}
	ExportAllStub        func(context.Context, config.OpsManager, string) error
	exportAllMutex       sync.RWMutex
	exportAllArgsForCall []struct {
		arg1 context.Context
		arg2 config.OpsManager
		arg3 string
	}
	exportAllReturns struct {
		result1 error
	}
	exportAllReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrgExporter) Export(arg1 context.Context, arg2 config.OpsManager, arg3 string, arg4 ...string) error {
	fake.exportMutex.Lock()
	ret, specificReturn := fake.exportReturnsOnCall[len(fake.exportArgsForCall)]
	fake.exportArgsForCall = append(fake.exportArgsForCall, struct {
		arg1 context.Context
		arg2 config.OpsManager
		arg3 string
		arg4 []string
	}{arg1, arg2, arg3, arg4})
	stub := fake.ExportStub
	fakeReturns := fake.exportReturns
	fake.recordInvocation("Export", []interface{}{arg1, arg2, arg3, arg4})
	fake.exportMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrgExporter) ExportCallCount() int {
	fake.exportMutex.RLock()
	defer fake.exportMutex.RUnlock()
	return len(fake.exportArgsForCall)
}

func (fake *FakeOrgExporter) ExportCalls(stub func(context.Context, config.OpsManager, string, ...string) error) {
	fake.exportMutex.Lock()
	defer fake.exportMutex.Unlock()
	fake.ExportStub = stub
}

func (fake *FakeOrgExporter) ExportArgsForCall(i int) (context.Context, config.OpsManager, string, []string) {
	fake.exportMutex.RLock()
	defer fake.exportMutex.RUnlock()
	argsForCall := fake.exportArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeOrgExporter) ExportReturns(result1 error) {
	fake.exportMutex.Lock()
	defer fake.exportMutex.Unlock()
	fake.ExportStub = nil
	fake.exportReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrgExporter) ExportReturnsOnCall(i int, result1 error) {
	fake.exportMutex.Lock()
	defer fake.exportMutex.Unlock()
	fake.ExportStub = nil
	if fake.exportReturnsOnCall == nil {
		fake.exportReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.exportReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrgExporter) ExportAll(arg1 context.Context, arg2 config.OpsManager, arg3 string) error {
	fake.exportAllMutex.Lock()
	ret, specificReturn := fake.exportAllReturnsOnCall[len(fake.exportAllArgsForCall)]
	fake.exportAllArgsForCall = append(fake.exportAllArgsForCall, struct {
		arg1 context.Context
		arg2 config.OpsManager
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ExportAllStub
	fakeReturns := fake.exportAllReturns
	fake.recordInvocation("ExportAll", []interface{}{arg1, arg2, arg3})
	fake.exportAllMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrgExporter) ExportAllCallCount() int {
	fake.exportAllMutex.RLock()
	defer fake.exportAllMutex.RUnlock()
	return len(fake.exportAllArgsForCall)
}

func (fake *FakeOrgExporter) ExportAllCalls(stub func(context.Context, config.OpsManager, string) error) {
	fake.exportAllMutex.Lock()
	defer fake.exportAllMutex.Unlock()
	fake.ExportAllStub = stub
}

func (fake *FakeOrgExporter) ExportAllArgsForCall(i int) (context.Context, config.OpsManager, string) {
	fake.exportAllMutex.RLock()
	defer fake.exportAllMutex.RUnlock()
	argsForCall := fake.exportAllArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOrgExporter) ExportAllReturns(result1 error) {
	fake.exportAllMutex.Lock()
	defer fake.exportAllMutex.Unlock()
	fake.ExportAllStub = nil
	fake.exportAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrgExporter) ExportAllReturnsOnCall(i int, result1 error) {
	fake.exportAllMutex.Lock()
	defer fake.exportAllMutex.Unlock()
	fake.ExportAllStub = nil
	if fake.exportAllReturnsOnCall == nil {
		fake.exportAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.exportAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrgExporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exportMutex.RLock()
	defer fake.exportMutex.RUnlock()
	fake.exportAllMutex.RLock()
	defer fake.exportAllMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrgExporter) recordInvocation(key string, args []interface{}) {
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

var _ cmd.OrgExporter = new(FakeOrgExporter)
