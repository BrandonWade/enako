// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"github.com/BrandonWade/enako/api/helpers"
)

type FakeCookieStorer struct {
	GetStub        func(*http.Request, string) (helpers.SessionStorer, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 *http.Request
		arg2 string
	}
	getReturns struct {
		result1 helpers.SessionStorer
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 helpers.SessionStorer
		result2 error
	}
	IsAuthenticatedStub        func(*http.Request) (bool, error)
	isAuthenticatedMutex       sync.RWMutex
	isAuthenticatedArgsForCall []struct {
		arg1 *http.Request
	}
	isAuthenticatedReturns struct {
		result1 bool
		result2 error
	}
	isAuthenticatedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCookieStorer) Get(arg1 *http.Request, arg2 string) (helpers.SessionStorer, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 *http.Request
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCookieStorer) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCookieStorer) GetCalls(stub func(*http.Request, string) (helpers.SessionStorer, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCookieStorer) GetArgsForCall(i int) (*http.Request, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCookieStorer) GetReturns(result1 helpers.SessionStorer, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 helpers.SessionStorer
		result2 error
	}{result1, result2}
}

func (fake *FakeCookieStorer) GetReturnsOnCall(i int, result1 helpers.SessionStorer, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 helpers.SessionStorer
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 helpers.SessionStorer
		result2 error
	}{result1, result2}
}

func (fake *FakeCookieStorer) IsAuthenticated(arg1 *http.Request) (bool, error) {
	fake.isAuthenticatedMutex.Lock()
	ret, specificReturn := fake.isAuthenticatedReturnsOnCall[len(fake.isAuthenticatedArgsForCall)]
	fake.isAuthenticatedArgsForCall = append(fake.isAuthenticatedArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("IsAuthenticated", []interface{}{arg1})
	fake.isAuthenticatedMutex.Unlock()
	if fake.IsAuthenticatedStub != nil {
		return fake.IsAuthenticatedStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.isAuthenticatedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCookieStorer) IsAuthenticatedCallCount() int {
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	return len(fake.isAuthenticatedArgsForCall)
}

func (fake *FakeCookieStorer) IsAuthenticatedCalls(stub func(*http.Request) (bool, error)) {
	fake.isAuthenticatedMutex.Lock()
	defer fake.isAuthenticatedMutex.Unlock()
	fake.IsAuthenticatedStub = stub
}

func (fake *FakeCookieStorer) IsAuthenticatedArgsForCall(i int) *http.Request {
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	argsForCall := fake.isAuthenticatedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCookieStorer) IsAuthenticatedReturns(result1 bool, result2 error) {
	fake.isAuthenticatedMutex.Lock()
	defer fake.isAuthenticatedMutex.Unlock()
	fake.IsAuthenticatedStub = nil
	fake.isAuthenticatedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCookieStorer) IsAuthenticatedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isAuthenticatedMutex.Lock()
	defer fake.isAuthenticatedMutex.Unlock()
	fake.IsAuthenticatedStub = nil
	if fake.isAuthenticatedReturnsOnCall == nil {
		fake.isAuthenticatedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isAuthenticatedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCookieStorer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCookieStorer) recordInvocation(key string, args []interface{}) {
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

var _ helpers.CookieStorer = new(FakeCookieStorer)
