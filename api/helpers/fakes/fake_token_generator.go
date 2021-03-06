// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/helpers"
)

type FakeTokenGenerator struct {
	CreateTokenStub        func(int) string
	createTokenMutex       sync.RWMutex
	createTokenArgsForCall []struct {
		arg1 int
	}
	createTokenReturns struct {
		result1 string
	}
	createTokenReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenGenerator) CreateToken(arg1 int) string {
	fake.createTokenMutex.Lock()
	ret, specificReturn := fake.createTokenReturnsOnCall[len(fake.createTokenArgsForCall)]
	fake.createTokenArgsForCall = append(fake.createTokenArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("CreateToken", []interface{}{arg1})
	fake.createTokenMutex.Unlock()
	if fake.CreateTokenStub != nil {
		return fake.CreateTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createTokenReturns
	return fakeReturns.result1
}

func (fake *FakeTokenGenerator) CreateTokenCallCount() int {
	fake.createTokenMutex.RLock()
	defer fake.createTokenMutex.RUnlock()
	return len(fake.createTokenArgsForCall)
}

func (fake *FakeTokenGenerator) CreateTokenCalls(stub func(int) string) {
	fake.createTokenMutex.Lock()
	defer fake.createTokenMutex.Unlock()
	fake.CreateTokenStub = stub
}

func (fake *FakeTokenGenerator) CreateTokenArgsForCall(i int) int {
	fake.createTokenMutex.RLock()
	defer fake.createTokenMutex.RUnlock()
	argsForCall := fake.createTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTokenGenerator) CreateTokenReturns(result1 string) {
	fake.createTokenMutex.Lock()
	defer fake.createTokenMutex.Unlock()
	fake.CreateTokenStub = nil
	fake.createTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenGenerator) CreateTokenReturnsOnCall(i int, result1 string) {
	fake.createTokenMutex.Lock()
	defer fake.createTokenMutex.Unlock()
	fake.CreateTokenStub = nil
	if fake.createTokenReturnsOnCall == nil {
		fake.createTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.createTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTokenMutex.RLock()
	defer fake.createTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTokenGenerator) recordInvocation(key string, args []interface{}) {
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

var _ helpers.TokenGenerator = new(FakeTokenGenerator)
