// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/helpers"
)

type FakeEmailObfuscator struct {
	ObfuscateStub        func(string) (string, error)
	obfuscateMutex       sync.RWMutex
	obfuscateArgsForCall []struct {
		arg1 string
	}
	obfuscateReturns struct {
		result1 string
		result2 error
	}
	obfuscateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmailObfuscator) Obfuscate(arg1 string) (string, error) {
	fake.obfuscateMutex.Lock()
	ret, specificReturn := fake.obfuscateReturnsOnCall[len(fake.obfuscateArgsForCall)]
	fake.obfuscateArgsForCall = append(fake.obfuscateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Obfuscate", []interface{}{arg1})
	fake.obfuscateMutex.Unlock()
	if fake.ObfuscateStub != nil {
		return fake.ObfuscateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.obfuscateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEmailObfuscator) ObfuscateCallCount() int {
	fake.obfuscateMutex.RLock()
	defer fake.obfuscateMutex.RUnlock()
	return len(fake.obfuscateArgsForCall)
}

func (fake *FakeEmailObfuscator) ObfuscateCalls(stub func(string) (string, error)) {
	fake.obfuscateMutex.Lock()
	defer fake.obfuscateMutex.Unlock()
	fake.ObfuscateStub = stub
}

func (fake *FakeEmailObfuscator) ObfuscateArgsForCall(i int) string {
	fake.obfuscateMutex.RLock()
	defer fake.obfuscateMutex.RUnlock()
	argsForCall := fake.obfuscateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEmailObfuscator) ObfuscateReturns(result1 string, result2 error) {
	fake.obfuscateMutex.Lock()
	defer fake.obfuscateMutex.Unlock()
	fake.ObfuscateStub = nil
	fake.obfuscateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEmailObfuscator) ObfuscateReturnsOnCall(i int, result1 string, result2 error) {
	fake.obfuscateMutex.Lock()
	defer fake.obfuscateMutex.Unlock()
	fake.ObfuscateStub = nil
	if fake.obfuscateReturnsOnCall == nil {
		fake.obfuscateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.obfuscateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEmailObfuscator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.obfuscateMutex.RLock()
	defer fake.obfuscateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmailObfuscator) recordInvocation(key string, args []interface{}) {
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

var _ helpers.EmailObfuscator = new(FakeEmailObfuscator)
