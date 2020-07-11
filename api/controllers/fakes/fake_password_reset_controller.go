// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"github.com/BrandonWade/enako/api/controllers"
)

type FakePasswordResetController struct {
	RequestPasswordResetStub        func(http.ResponseWriter, *http.Request)
	requestPasswordResetMutex       sync.RWMutex
	requestPasswordResetArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	ResetPasswordStub        func(http.ResponseWriter, *http.Request)
	resetPasswordMutex       sync.RWMutex
	resetPasswordArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	SetPasswordResetTokenStub        func(http.ResponseWriter, *http.Request)
	setPasswordResetTokenMutex       sync.RWMutex
	setPasswordResetTokenArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePasswordResetController) RequestPasswordReset(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.requestPasswordResetMutex.Lock()
	fake.requestPasswordResetArgsForCall = append(fake.requestPasswordResetArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("RequestPasswordReset", []interface{}{arg1, arg2})
	fake.requestPasswordResetMutex.Unlock()
	if fake.RequestPasswordResetStub != nil {
		fake.RequestPasswordResetStub(arg1, arg2)
	}
}

func (fake *FakePasswordResetController) RequestPasswordResetCallCount() int {
	fake.requestPasswordResetMutex.RLock()
	defer fake.requestPasswordResetMutex.RUnlock()
	return len(fake.requestPasswordResetArgsForCall)
}

func (fake *FakePasswordResetController) RequestPasswordResetCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.requestPasswordResetMutex.Lock()
	defer fake.requestPasswordResetMutex.Unlock()
	fake.RequestPasswordResetStub = stub
}

func (fake *FakePasswordResetController) RequestPasswordResetArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.requestPasswordResetMutex.RLock()
	defer fake.requestPasswordResetMutex.RUnlock()
	argsForCall := fake.requestPasswordResetArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePasswordResetController) ResetPassword(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.resetPasswordMutex.Lock()
	fake.resetPasswordArgsForCall = append(fake.resetPasswordArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("ResetPassword", []interface{}{arg1, arg2})
	fake.resetPasswordMutex.Unlock()
	if fake.ResetPasswordStub != nil {
		fake.ResetPasswordStub(arg1, arg2)
	}
}

func (fake *FakePasswordResetController) ResetPasswordCallCount() int {
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	return len(fake.resetPasswordArgsForCall)
}

func (fake *FakePasswordResetController) ResetPasswordCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = stub
}

func (fake *FakePasswordResetController) ResetPasswordArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	argsForCall := fake.resetPasswordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePasswordResetController) SetPasswordResetToken(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.setPasswordResetTokenMutex.Lock()
	fake.setPasswordResetTokenArgsForCall = append(fake.setPasswordResetTokenArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("SetPasswordResetToken", []interface{}{arg1, arg2})
	fake.setPasswordResetTokenMutex.Unlock()
	if fake.SetPasswordResetTokenStub != nil {
		fake.SetPasswordResetTokenStub(arg1, arg2)
	}
}

func (fake *FakePasswordResetController) SetPasswordResetTokenCallCount() int {
	fake.setPasswordResetTokenMutex.RLock()
	defer fake.setPasswordResetTokenMutex.RUnlock()
	return len(fake.setPasswordResetTokenArgsForCall)
}

func (fake *FakePasswordResetController) SetPasswordResetTokenCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.setPasswordResetTokenMutex.Lock()
	defer fake.setPasswordResetTokenMutex.Unlock()
	fake.SetPasswordResetTokenStub = stub
}

func (fake *FakePasswordResetController) SetPasswordResetTokenArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.setPasswordResetTokenMutex.RLock()
	defer fake.setPasswordResetTokenMutex.RUnlock()
	argsForCall := fake.setPasswordResetTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePasswordResetController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestPasswordResetMutex.RLock()
	defer fake.requestPasswordResetMutex.RUnlock()
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	fake.setPasswordResetTokenMutex.RLock()
	defer fake.setPasswordResetTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePasswordResetController) recordInvocation(key string, args []interface{}) {
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

var _ controllers.PasswordResetController = new(FakePasswordResetController)
