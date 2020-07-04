// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/services"
)

type FakeEmailService struct {
	SendAccountActivationEmailStub        func(string, string) error
	sendAccountActivationEmailMutex       sync.RWMutex
	sendAccountActivationEmailArgsForCall []struct {
		arg1 string
		arg2 string
	}
	sendAccountActivationEmailReturns struct {
		result1 error
	}
	sendAccountActivationEmailReturnsOnCall map[int]struct {
		result1 error
	}
	SendPasswordResetEmailStub        func(string, string) error
	sendPasswordResetEmailMutex       sync.RWMutex
	sendPasswordResetEmailArgsForCall []struct {
		arg1 string
		arg2 string
	}
	sendPasswordResetEmailReturns struct {
		result1 error
	}
	sendPasswordResetEmailReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmailService) SendAccountActivationEmail(arg1 string, arg2 string) error {
	fake.sendAccountActivationEmailMutex.Lock()
	ret, specificReturn := fake.sendAccountActivationEmailReturnsOnCall[len(fake.sendAccountActivationEmailArgsForCall)]
	fake.sendAccountActivationEmailArgsForCall = append(fake.sendAccountActivationEmailArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SendAccountActivationEmail", []interface{}{arg1, arg2})
	fake.sendAccountActivationEmailMutex.Unlock()
	if fake.SendAccountActivationEmailStub != nil {
		return fake.SendAccountActivationEmailStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendAccountActivationEmailReturns
	return fakeReturns.result1
}

func (fake *FakeEmailService) SendAccountActivationEmailCallCount() int {
	fake.sendAccountActivationEmailMutex.RLock()
	defer fake.sendAccountActivationEmailMutex.RUnlock()
	return len(fake.sendAccountActivationEmailArgsForCall)
}

func (fake *FakeEmailService) SendAccountActivationEmailCalls(stub func(string, string) error) {
	fake.sendAccountActivationEmailMutex.Lock()
	defer fake.sendAccountActivationEmailMutex.Unlock()
	fake.SendAccountActivationEmailStub = stub
}

func (fake *FakeEmailService) SendAccountActivationEmailArgsForCall(i int) (string, string) {
	fake.sendAccountActivationEmailMutex.RLock()
	defer fake.sendAccountActivationEmailMutex.RUnlock()
	argsForCall := fake.sendAccountActivationEmailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEmailService) SendAccountActivationEmailReturns(result1 error) {
	fake.sendAccountActivationEmailMutex.Lock()
	defer fake.sendAccountActivationEmailMutex.Unlock()
	fake.SendAccountActivationEmailStub = nil
	fake.sendAccountActivationEmailReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEmailService) SendAccountActivationEmailReturnsOnCall(i int, result1 error) {
	fake.sendAccountActivationEmailMutex.Lock()
	defer fake.sendAccountActivationEmailMutex.Unlock()
	fake.SendAccountActivationEmailStub = nil
	if fake.sendAccountActivationEmailReturnsOnCall == nil {
		fake.sendAccountActivationEmailReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendAccountActivationEmailReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEmailService) SendPasswordResetEmail(arg1 string, arg2 string) error {
	fake.sendPasswordResetEmailMutex.Lock()
	ret, specificReturn := fake.sendPasswordResetEmailReturnsOnCall[len(fake.sendPasswordResetEmailArgsForCall)]
	fake.sendPasswordResetEmailArgsForCall = append(fake.sendPasswordResetEmailArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SendPasswordResetEmail", []interface{}{arg1, arg2})
	fake.sendPasswordResetEmailMutex.Unlock()
	if fake.SendPasswordResetEmailStub != nil {
		return fake.SendPasswordResetEmailStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendPasswordResetEmailReturns
	return fakeReturns.result1
}

func (fake *FakeEmailService) SendPasswordResetEmailCallCount() int {
	fake.sendPasswordResetEmailMutex.RLock()
	defer fake.sendPasswordResetEmailMutex.RUnlock()
	return len(fake.sendPasswordResetEmailArgsForCall)
}

func (fake *FakeEmailService) SendPasswordResetEmailCalls(stub func(string, string) error) {
	fake.sendPasswordResetEmailMutex.Lock()
	defer fake.sendPasswordResetEmailMutex.Unlock()
	fake.SendPasswordResetEmailStub = stub
}

func (fake *FakeEmailService) SendPasswordResetEmailArgsForCall(i int) (string, string) {
	fake.sendPasswordResetEmailMutex.RLock()
	defer fake.sendPasswordResetEmailMutex.RUnlock()
	argsForCall := fake.sendPasswordResetEmailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEmailService) SendPasswordResetEmailReturns(result1 error) {
	fake.sendPasswordResetEmailMutex.Lock()
	defer fake.sendPasswordResetEmailMutex.Unlock()
	fake.SendPasswordResetEmailStub = nil
	fake.sendPasswordResetEmailReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEmailService) SendPasswordResetEmailReturnsOnCall(i int, result1 error) {
	fake.sendPasswordResetEmailMutex.Lock()
	defer fake.sendPasswordResetEmailMutex.Unlock()
	fake.SendPasswordResetEmailStub = nil
	if fake.sendPasswordResetEmailReturnsOnCall == nil {
		fake.sendPasswordResetEmailReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendPasswordResetEmailReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEmailService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendAccountActivationEmailMutex.RLock()
	defer fake.sendAccountActivationEmailMutex.RUnlock()
	fake.sendPasswordResetEmailMutex.RLock()
	defer fake.sendPasswordResetEmailMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmailService) recordInvocation(key string, args []interface{}) {
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

var _ services.EmailService = new(FakeEmailService)
