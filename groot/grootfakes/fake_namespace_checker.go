// This file was generated by counterfeiter
package grootfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/groot"
)

type FakeNamespaceChecker struct {
	CheckStub        func(uidMappings, gidMappings []groot.IDMappingSpec) (bool, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		uidMappings []groot.IDMappingSpec
		gidMappings []groot.IDMappingSpec
	}
	checkReturns struct {
		result1 bool
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNamespaceChecker) Check(uidMappings []groot.IDMappingSpec, gidMappings []groot.IDMappingSpec) (bool, error) {
	var uidMappingsCopy []groot.IDMappingSpec
	if uidMappings != nil {
		uidMappingsCopy = make([]groot.IDMappingSpec, len(uidMappings))
		copy(uidMappingsCopy, uidMappings)
	}
	var gidMappingsCopy []groot.IDMappingSpec
	if gidMappings != nil {
		gidMappingsCopy = make([]groot.IDMappingSpec, len(gidMappings))
		copy(gidMappingsCopy, gidMappings)
	}
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		uidMappings []groot.IDMappingSpec
		gidMappings []groot.IDMappingSpec
	}{uidMappingsCopy, gidMappingsCopy})
	fake.recordInvocation("Check", []interface{}{uidMappingsCopy, gidMappingsCopy})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(uidMappings, gidMappings)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkReturns.result1, fake.checkReturns.result2
}

func (fake *FakeNamespaceChecker) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeNamespaceChecker) CheckArgsForCall(i int) ([]groot.IDMappingSpec, []groot.IDMappingSpec) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return fake.checkArgsForCall[i].uidMappings, fake.checkArgsForCall[i].gidMappings
}

func (fake *FakeNamespaceChecker) CheckReturns(result1 bool, result2 error) {
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeNamespaceChecker) CheckReturnsOnCall(i int, result1 bool, result2 error) {
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeNamespaceChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNamespaceChecker) recordInvocation(key string, args []interface{}) {
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

var _ groot.NamespaceChecker = new(FakeNamespaceChecker)
