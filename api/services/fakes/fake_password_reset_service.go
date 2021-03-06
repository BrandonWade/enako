// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
)

type FakePasswordResetService struct {
	CheckPasswordResetTokenIsValidStub        func(*models.PasswordResetToken) (bool, error)
	checkPasswordResetTokenIsValidMutex       sync.RWMutex
	checkPasswordResetTokenIsValidArgsForCall []struct {
		arg1 *models.PasswordResetToken
	}
	checkPasswordResetTokenIsValidReturns struct {
		result1 bool
		result2 error
	}
	checkPasswordResetTokenIsValidReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RequestPasswordResetStub        func(string) (string, error)
	requestPasswordResetMutex       sync.RWMutex
	requestPasswordResetArgsForCall []struct {
		arg1 string
	}
	requestPasswordResetReturns struct {
		result1 string
		result2 error
	}
	requestPasswordResetReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ResetPasswordStub        func(string, string) (bool, error)
	resetPasswordMutex       sync.RWMutex
	resetPasswordArgsForCall []struct {
		arg1 string
		arg2 string
	}
	resetPasswordReturns struct {
		result1 bool
		result2 error
	}
	resetPasswordReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	VerifyPasswordResetTokenStub        func(string) (bool, error)
	verifyPasswordResetTokenMutex       sync.RWMutex
	verifyPasswordResetTokenArgsForCall []struct {
		arg1 string
	}
	verifyPasswordResetTokenReturns struct {
		result1 bool
		result2 error
	}
	verifyPasswordResetTokenReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePasswordResetService) CheckPasswordResetTokenIsValid(arg1 *models.PasswordResetToken) (bool, error) {
	fake.checkPasswordResetTokenIsValidMutex.Lock()
	ret, specificReturn := fake.checkPasswordResetTokenIsValidReturnsOnCall[len(fake.checkPasswordResetTokenIsValidArgsForCall)]
	fake.checkPasswordResetTokenIsValidArgsForCall = append(fake.checkPasswordResetTokenIsValidArgsForCall, struct {
		arg1 *models.PasswordResetToken
	}{arg1})
	fake.recordInvocation("CheckPasswordResetTokenIsValid", []interface{}{arg1})
	fake.checkPasswordResetTokenIsValidMutex.Unlock()
	if fake.CheckPasswordResetTokenIsValidStub != nil {
		return fake.CheckPasswordResetTokenIsValidStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkPasswordResetTokenIsValidReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePasswordResetService) CheckPasswordResetTokenIsValidCallCount() int {
	fake.checkPasswordResetTokenIsValidMutex.RLock()
	defer fake.checkPasswordResetTokenIsValidMutex.RUnlock()
	return len(fake.checkPasswordResetTokenIsValidArgsForCall)
}

func (fake *FakePasswordResetService) CheckPasswordResetTokenIsValidCalls(stub func(*models.PasswordResetToken) (bool, error)) {
	fake.checkPasswordResetTokenIsValidMutex.Lock()
	defer fake.checkPasswordResetTokenIsValidMutex.Unlock()
	fake.CheckPasswordResetTokenIsValidStub = stub
}

func (fake *FakePasswordResetService) CheckPasswordResetTokenIsValidArgsForCall(i int) *models.PasswordResetToken {
	fake.checkPasswordResetTokenIsValidMutex.RLock()
	defer fake.checkPasswordResetTokenIsValidMutex.RUnlock()
	argsForCall := fake.checkPasswordResetTokenIsValidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePasswordResetService) CheckPasswordResetTokenIsValidReturns(result1 bool, result2 error) {
	fake.checkPasswordResetTokenIsValidMutex.Lock()
	defer fake.checkPasswordResetTokenIsValidMutex.Unlock()
	fake.CheckPasswordResetTokenIsValidStub = nil
	fake.checkPasswordResetTokenIsValidReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) CheckPasswordResetTokenIsValidReturnsOnCall(i int, result1 bool, result2 error) {
	fake.checkPasswordResetTokenIsValidMutex.Lock()
	defer fake.checkPasswordResetTokenIsValidMutex.Unlock()
	fake.CheckPasswordResetTokenIsValidStub = nil
	if fake.checkPasswordResetTokenIsValidReturnsOnCall == nil {
		fake.checkPasswordResetTokenIsValidReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkPasswordResetTokenIsValidReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) RequestPasswordReset(arg1 string) (string, error) {
	fake.requestPasswordResetMutex.Lock()
	ret, specificReturn := fake.requestPasswordResetReturnsOnCall[len(fake.requestPasswordResetArgsForCall)]
	fake.requestPasswordResetArgsForCall = append(fake.requestPasswordResetArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RequestPasswordReset", []interface{}{arg1})
	fake.requestPasswordResetMutex.Unlock()
	if fake.RequestPasswordResetStub != nil {
		return fake.RequestPasswordResetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestPasswordResetReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePasswordResetService) RequestPasswordResetCallCount() int {
	fake.requestPasswordResetMutex.RLock()
	defer fake.requestPasswordResetMutex.RUnlock()
	return len(fake.requestPasswordResetArgsForCall)
}

func (fake *FakePasswordResetService) RequestPasswordResetCalls(stub func(string) (string, error)) {
	fake.requestPasswordResetMutex.Lock()
	defer fake.requestPasswordResetMutex.Unlock()
	fake.RequestPasswordResetStub = stub
}

func (fake *FakePasswordResetService) RequestPasswordResetArgsForCall(i int) string {
	fake.requestPasswordResetMutex.RLock()
	defer fake.requestPasswordResetMutex.RUnlock()
	argsForCall := fake.requestPasswordResetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePasswordResetService) RequestPasswordResetReturns(result1 string, result2 error) {
	fake.requestPasswordResetMutex.Lock()
	defer fake.requestPasswordResetMutex.Unlock()
	fake.RequestPasswordResetStub = nil
	fake.requestPasswordResetReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) RequestPasswordResetReturnsOnCall(i int, result1 string, result2 error) {
	fake.requestPasswordResetMutex.Lock()
	defer fake.requestPasswordResetMutex.Unlock()
	fake.RequestPasswordResetStub = nil
	if fake.requestPasswordResetReturnsOnCall == nil {
		fake.requestPasswordResetReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.requestPasswordResetReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) ResetPassword(arg1 string, arg2 string) (bool, error) {
	fake.resetPasswordMutex.Lock()
	ret, specificReturn := fake.resetPasswordReturnsOnCall[len(fake.resetPasswordArgsForCall)]
	fake.resetPasswordArgsForCall = append(fake.resetPasswordArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ResetPassword", []interface{}{arg1, arg2})
	fake.resetPasswordMutex.Unlock()
	if fake.ResetPasswordStub != nil {
		return fake.ResetPasswordStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resetPasswordReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePasswordResetService) ResetPasswordCallCount() int {
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	return len(fake.resetPasswordArgsForCall)
}

func (fake *FakePasswordResetService) ResetPasswordCalls(stub func(string, string) (bool, error)) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = stub
}

func (fake *FakePasswordResetService) ResetPasswordArgsForCall(i int) (string, string) {
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	argsForCall := fake.resetPasswordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePasswordResetService) ResetPasswordReturns(result1 bool, result2 error) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = nil
	fake.resetPasswordReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) ResetPasswordReturnsOnCall(i int, result1 bool, result2 error) {
	fake.resetPasswordMutex.Lock()
	defer fake.resetPasswordMutex.Unlock()
	fake.ResetPasswordStub = nil
	if fake.resetPasswordReturnsOnCall == nil {
		fake.resetPasswordReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.resetPasswordReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) VerifyPasswordResetToken(arg1 string) (bool, error) {
	fake.verifyPasswordResetTokenMutex.Lock()
	ret, specificReturn := fake.verifyPasswordResetTokenReturnsOnCall[len(fake.verifyPasswordResetTokenArgsForCall)]
	fake.verifyPasswordResetTokenArgsForCall = append(fake.verifyPasswordResetTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("VerifyPasswordResetToken", []interface{}{arg1})
	fake.verifyPasswordResetTokenMutex.Unlock()
	if fake.VerifyPasswordResetTokenStub != nil {
		return fake.VerifyPasswordResetTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.verifyPasswordResetTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePasswordResetService) VerifyPasswordResetTokenCallCount() int {
	fake.verifyPasswordResetTokenMutex.RLock()
	defer fake.verifyPasswordResetTokenMutex.RUnlock()
	return len(fake.verifyPasswordResetTokenArgsForCall)
}

func (fake *FakePasswordResetService) VerifyPasswordResetTokenCalls(stub func(string) (bool, error)) {
	fake.verifyPasswordResetTokenMutex.Lock()
	defer fake.verifyPasswordResetTokenMutex.Unlock()
	fake.VerifyPasswordResetTokenStub = stub
}

func (fake *FakePasswordResetService) VerifyPasswordResetTokenArgsForCall(i int) string {
	fake.verifyPasswordResetTokenMutex.RLock()
	defer fake.verifyPasswordResetTokenMutex.RUnlock()
	argsForCall := fake.verifyPasswordResetTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePasswordResetService) VerifyPasswordResetTokenReturns(result1 bool, result2 error) {
	fake.verifyPasswordResetTokenMutex.Lock()
	defer fake.verifyPasswordResetTokenMutex.Unlock()
	fake.VerifyPasswordResetTokenStub = nil
	fake.verifyPasswordResetTokenReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) VerifyPasswordResetTokenReturnsOnCall(i int, result1 bool, result2 error) {
	fake.verifyPasswordResetTokenMutex.Lock()
	defer fake.verifyPasswordResetTokenMutex.Unlock()
	fake.VerifyPasswordResetTokenStub = nil
	if fake.verifyPasswordResetTokenReturnsOnCall == nil {
		fake.verifyPasswordResetTokenReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.verifyPasswordResetTokenReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePasswordResetService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkPasswordResetTokenIsValidMutex.RLock()
	defer fake.checkPasswordResetTokenIsValidMutex.RUnlock()
	fake.requestPasswordResetMutex.RLock()
	defer fake.requestPasswordResetMutex.RUnlock()
	fake.resetPasswordMutex.RLock()
	defer fake.resetPasswordMutex.RUnlock()
	fake.verifyPasswordResetTokenMutex.RLock()
	defer fake.verifyPasswordResetTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePasswordResetService) recordInvocation(key string, args []interface{}) {
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

var _ services.PasswordResetService = new(FakePasswordResetService)
