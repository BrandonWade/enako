// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"github.com/BrandonWade/enako/api/controllers"
)

type FakeAuthController struct {
	ActivateAccountStub        func(http.ResponseWriter, *http.Request)
	activateAccountMutex       sync.RWMutex
	activateAccountArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	CSRFStub        func(http.ResponseWriter, *http.Request)
	cSRFMutex       sync.RWMutex
	cSRFArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	LoginStub        func(http.ResponseWriter, *http.Request)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	LogoutStub        func(http.ResponseWriter, *http.Request)
	logoutMutex       sync.RWMutex
	logoutArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	RegisterUserStub        func(http.ResponseWriter, *http.Request)
	registerUserMutex       sync.RWMutex
	registerUserArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthController) ActivateAccount(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.activateAccountMutex.Lock()
	fake.activateAccountArgsForCall = append(fake.activateAccountArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("ActivateAccount", []interface{}{arg1, arg2})
	fake.activateAccountMutex.Unlock()
	if fake.ActivateAccountStub != nil {
		fake.ActivateAccountStub(arg1, arg2)
	}
}

func (fake *FakeAuthController) ActivateAccountCallCount() int {
	fake.activateAccountMutex.RLock()
	defer fake.activateAccountMutex.RUnlock()
	return len(fake.activateAccountArgsForCall)
}

func (fake *FakeAuthController) ActivateAccountCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.activateAccountMutex.Lock()
	defer fake.activateAccountMutex.Unlock()
	fake.ActivateAccountStub = stub
}

func (fake *FakeAuthController) ActivateAccountArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.activateAccountMutex.RLock()
	defer fake.activateAccountMutex.RUnlock()
	argsForCall := fake.activateAccountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthController) CSRF(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.cSRFMutex.Lock()
	fake.cSRFArgsForCall = append(fake.cSRFArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("CSRF", []interface{}{arg1, arg2})
	fake.cSRFMutex.Unlock()
	if fake.CSRFStub != nil {
		fake.CSRFStub(arg1, arg2)
	}
}

func (fake *FakeAuthController) CSRFCallCount() int {
	fake.cSRFMutex.RLock()
	defer fake.cSRFMutex.RUnlock()
	return len(fake.cSRFArgsForCall)
}

func (fake *FakeAuthController) CSRFCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.cSRFMutex.Lock()
	defer fake.cSRFMutex.Unlock()
	fake.CSRFStub = stub
}

func (fake *FakeAuthController) CSRFArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.cSRFMutex.RLock()
	defer fake.cSRFMutex.RUnlock()
	argsForCall := fake.cSRFArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthController) Login(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.loginMutex.Lock()
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("Login", []interface{}{arg1, arg2})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		fake.LoginStub(arg1, arg2)
	}
}

func (fake *FakeAuthController) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeAuthController) LoginCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeAuthController) LoginArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthController) Logout(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.logoutMutex.Lock()
	fake.logoutArgsForCall = append(fake.logoutArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("Logout", []interface{}{arg1, arg2})
	fake.logoutMutex.Unlock()
	if fake.LogoutStub != nil {
		fake.LogoutStub(arg1, arg2)
	}
}

func (fake *FakeAuthController) LogoutCallCount() int {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	return len(fake.logoutArgsForCall)
}

func (fake *FakeAuthController) LogoutCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.logoutMutex.Lock()
	defer fake.logoutMutex.Unlock()
	fake.LogoutStub = stub
}

func (fake *FakeAuthController) LogoutArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	argsForCall := fake.logoutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthController) RegisterUser(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.registerUserMutex.Lock()
	fake.registerUserArgsForCall = append(fake.registerUserArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("RegisterUser", []interface{}{arg1, arg2})
	fake.registerUserMutex.Unlock()
	if fake.RegisterUserStub != nil {
		fake.RegisterUserStub(arg1, arg2)
	}
}

func (fake *FakeAuthController) RegisterUserCallCount() int {
	fake.registerUserMutex.RLock()
	defer fake.registerUserMutex.RUnlock()
	return len(fake.registerUserArgsForCall)
}

func (fake *FakeAuthController) RegisterUserCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.registerUserMutex.Lock()
	defer fake.registerUserMutex.Unlock()
	fake.RegisterUserStub = stub
}

func (fake *FakeAuthController) RegisterUserArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.registerUserMutex.RLock()
	defer fake.registerUserMutex.RUnlock()
	argsForCall := fake.registerUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.activateAccountMutex.RLock()
	defer fake.activateAccountMutex.RUnlock()
	fake.cSRFMutex.RLock()
	defer fake.cSRFMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	fake.registerUserMutex.RLock()
	defer fake.registerUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthController) recordInvocation(key string, args []interface{}) {
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

var _ controllers.AuthController = new(FakeAuthController)
