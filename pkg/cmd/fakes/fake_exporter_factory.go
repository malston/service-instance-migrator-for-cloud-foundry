// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/cmd"
	"github.com/vmware-tanzu/service-instance-migrator-for-cloud-foundry/pkg/migrate"
)

type FakeExporterFactory struct {
	NewOrgExporterStub        func(migrate.ServiceInstanceExporter) cmd.OrgExporter
	newOrgExporterMutex       sync.RWMutex
	newOrgExporterArgsForCall []struct {
		arg1 migrate.ServiceInstanceExporter
	}
	newOrgExporterReturns struct {
		result1 cmd.OrgExporter
	}
	newOrgExporterReturnsOnCall map[int]struct {
		result1 cmd.OrgExporter
	}
	NewSpaceExporterStub        func(migrate.ServiceInstanceExporter) cmd.SpaceExporter
	newSpaceExporterMutex       sync.RWMutex
	newSpaceExporterArgsForCall []struct {
		arg1 migrate.ServiceInstanceExporter
	}
	newSpaceExporterReturns struct {
		result1 cmd.SpaceExporter
	}
	newSpaceExporterReturnsOnCall map[int]struct {
		result1 cmd.SpaceExporter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExporterFactory) NewOrgExporter(arg1 migrate.ServiceInstanceExporter) cmd.OrgExporter {
	fake.newOrgExporterMutex.Lock()
	ret, specificReturn := fake.newOrgExporterReturnsOnCall[len(fake.newOrgExporterArgsForCall)]
	fake.newOrgExporterArgsForCall = append(fake.newOrgExporterArgsForCall, struct {
		arg1 migrate.ServiceInstanceExporter
	}{arg1})
	stub := fake.NewOrgExporterStub
	fakeReturns := fake.newOrgExporterReturns
	fake.recordInvocation("NewOrgExporter", []interface{}{arg1})
	fake.newOrgExporterMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeExporterFactory) NewOrgExporterCallCount() int {
	fake.newOrgExporterMutex.RLock()
	defer fake.newOrgExporterMutex.RUnlock()
	return len(fake.newOrgExporterArgsForCall)
}

func (fake *FakeExporterFactory) NewOrgExporterCalls(stub func(migrate.ServiceInstanceExporter) cmd.OrgExporter) {
	fake.newOrgExporterMutex.Lock()
	defer fake.newOrgExporterMutex.Unlock()
	fake.NewOrgExporterStub = stub
}

func (fake *FakeExporterFactory) NewOrgExporterArgsForCall(i int) migrate.ServiceInstanceExporter {
	fake.newOrgExporterMutex.RLock()
	defer fake.newOrgExporterMutex.RUnlock()
	argsForCall := fake.newOrgExporterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExporterFactory) NewOrgExporterReturns(result1 cmd.OrgExporter) {
	fake.newOrgExporterMutex.Lock()
	defer fake.newOrgExporterMutex.Unlock()
	fake.NewOrgExporterStub = nil
	fake.newOrgExporterReturns = struct {
		result1 cmd.OrgExporter
	}{result1}
}

func (fake *FakeExporterFactory) NewOrgExporterReturnsOnCall(i int, result1 cmd.OrgExporter) {
	fake.newOrgExporterMutex.Lock()
	defer fake.newOrgExporterMutex.Unlock()
	fake.NewOrgExporterStub = nil
	if fake.newOrgExporterReturnsOnCall == nil {
		fake.newOrgExporterReturnsOnCall = make(map[int]struct {
			result1 cmd.OrgExporter
		})
	}
	fake.newOrgExporterReturnsOnCall[i] = struct {
		result1 cmd.OrgExporter
	}{result1}
}

func (fake *FakeExporterFactory) NewSpaceExporter(arg1 migrate.ServiceInstanceExporter) cmd.SpaceExporter {
	fake.newSpaceExporterMutex.Lock()
	ret, specificReturn := fake.newSpaceExporterReturnsOnCall[len(fake.newSpaceExporterArgsForCall)]
	fake.newSpaceExporterArgsForCall = append(fake.newSpaceExporterArgsForCall, struct {
		arg1 migrate.ServiceInstanceExporter
	}{arg1})
	stub := fake.NewSpaceExporterStub
	fakeReturns := fake.newSpaceExporterReturns
	fake.recordInvocation("NewSpaceExporter", []interface{}{arg1})
	fake.newSpaceExporterMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeExporterFactory) NewSpaceExporterCallCount() int {
	fake.newSpaceExporterMutex.RLock()
	defer fake.newSpaceExporterMutex.RUnlock()
	return len(fake.newSpaceExporterArgsForCall)
}

func (fake *FakeExporterFactory) NewSpaceExporterCalls(stub func(migrate.ServiceInstanceExporter) cmd.SpaceExporter) {
	fake.newSpaceExporterMutex.Lock()
	defer fake.newSpaceExporterMutex.Unlock()
	fake.NewSpaceExporterStub = stub
}

func (fake *FakeExporterFactory) NewSpaceExporterArgsForCall(i int) migrate.ServiceInstanceExporter {
	fake.newSpaceExporterMutex.RLock()
	defer fake.newSpaceExporterMutex.RUnlock()
	argsForCall := fake.newSpaceExporterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExporterFactory) NewSpaceExporterReturns(result1 cmd.SpaceExporter) {
	fake.newSpaceExporterMutex.Lock()
	defer fake.newSpaceExporterMutex.Unlock()
	fake.NewSpaceExporterStub = nil
	fake.newSpaceExporterReturns = struct {
		result1 cmd.SpaceExporter
	}{result1}
}

func (fake *FakeExporterFactory) NewSpaceExporterReturnsOnCall(i int, result1 cmd.SpaceExporter) {
	fake.newSpaceExporterMutex.Lock()
	defer fake.newSpaceExporterMutex.Unlock()
	fake.NewSpaceExporterStub = nil
	if fake.newSpaceExporterReturnsOnCall == nil {
		fake.newSpaceExporterReturnsOnCall = make(map[int]struct {
			result1 cmd.SpaceExporter
		})
	}
	fake.newSpaceExporterReturnsOnCall[i] = struct {
		result1 cmd.SpaceExporter
	}{result1}
}

func (fake *FakeExporterFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newOrgExporterMutex.RLock()
	defer fake.newOrgExporterMutex.RUnlock()
	fake.newSpaceExporterMutex.RLock()
	defer fake.newSpaceExporterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExporterFactory) recordInvocation(key string, args []interface{}) {
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

var _ cmd.ExporterFactory = new(FakeExporterFactory)
