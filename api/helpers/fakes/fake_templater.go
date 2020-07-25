// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/helpers"
)

type FakeTemplater struct {
	GenerateTemplateStub        func(string, interface{}) (string, error)
	generateTemplateMutex       sync.RWMutex
	generateTemplateArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	generateTemplateReturns struct {
		result1 string
		result2 error
	}
	generateTemplateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTemplater) GenerateTemplate(arg1 string, arg2 interface{}) (string, error) {
	fake.generateTemplateMutex.Lock()
	ret, specificReturn := fake.generateTemplateReturnsOnCall[len(fake.generateTemplateArgsForCall)]
	fake.generateTemplateArgsForCall = append(fake.generateTemplateArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("GenerateTemplate", []interface{}{arg1, arg2})
	fake.generateTemplateMutex.Unlock()
	if fake.GenerateTemplateStub != nil {
		return fake.GenerateTemplateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateTemplateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTemplater) GenerateTemplateCallCount() int {
	fake.generateTemplateMutex.RLock()
	defer fake.generateTemplateMutex.RUnlock()
	return len(fake.generateTemplateArgsForCall)
}

func (fake *FakeTemplater) GenerateTemplateCalls(stub func(string, interface{}) (string, error)) {
	fake.generateTemplateMutex.Lock()
	defer fake.generateTemplateMutex.Unlock()
	fake.GenerateTemplateStub = stub
}

func (fake *FakeTemplater) GenerateTemplateArgsForCall(i int) (string, interface{}) {
	fake.generateTemplateMutex.RLock()
	defer fake.generateTemplateMutex.RUnlock()
	argsForCall := fake.generateTemplateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTemplater) GenerateTemplateReturns(result1 string, result2 error) {
	fake.generateTemplateMutex.Lock()
	defer fake.generateTemplateMutex.Unlock()
	fake.GenerateTemplateStub = nil
	fake.generateTemplateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTemplater) GenerateTemplateReturnsOnCall(i int, result1 string, result2 error) {
	fake.generateTemplateMutex.Lock()
	defer fake.generateTemplateMutex.Unlock()
	fake.GenerateTemplateStub = nil
	if fake.generateTemplateReturnsOnCall == nil {
		fake.generateTemplateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateTemplateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTemplater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateTemplateMutex.RLock()
	defer fake.generateTemplateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTemplater) recordInvocation(key string, args []interface{}) {
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

var _ helpers.Templater = new(FakeTemplater)