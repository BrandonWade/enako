// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/repositories"
)

type FakeAuthRepository struct {
	CreateAccountStub        func(string, string) (int64, error)
	createAccountMutex       sync.RWMutex
	createAccountArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createAccountReturns struct {
		result1 int64
		result2 error
	}
	createAccountReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthRepository) CreateAccount(arg1 string, arg2 string) (int64, error) {
	fake.createAccountMutex.Lock()
	ret, specificReturn := fake.createAccountReturnsOnCall[len(fake.createAccountArgsForCall)]
	fake.createAccountArgsForCall = append(fake.createAccountArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateAccount", []interface{}{arg1, arg2})
	fake.createAccountMutex.Unlock()
	if fake.CreateAccountStub != nil {
		return fake.CreateAccountStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createAccountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthRepository) CreateAccountCallCount() int {
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	return len(fake.createAccountArgsForCall)
}

func (fake *FakeAuthRepository) CreateAccountCalls(stub func(string, string) (int64, error)) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = stub
}

func (fake *FakeAuthRepository) CreateAccountArgsForCall(i int) (string, string) {
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	argsForCall := fake.createAccountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthRepository) CreateAccountReturns(result1 int64, result2 error) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = nil
	fake.createAccountReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthRepository) CreateAccountReturnsOnCall(i int, result1 int64, result2 error) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = nil
	if fake.createAccountReturnsOnCall == nil {
		fake.createAccountReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createAccountReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthRepository) recordInvocation(key string, args []interface{}) {
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

var _ repositories.AuthRepository = new(FakeAuthRepository)
