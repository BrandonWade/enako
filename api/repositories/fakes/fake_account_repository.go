// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type FakeAccountRepository struct {
	ActivateAccountStub        func(string) (bool, error)
	activateAccountMutex       sync.RWMutex
	activateAccountArgsForCall []struct {
		arg1 string
	}
	activateAccountReturns struct {
		result1 bool
		result2 error
	}
	activateAccountReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
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
	CreateActivationTokenStub        func(int64, string) (int64, error)
	createActivationTokenMutex       sync.RWMutex
	createActivationTokenArgsForCall []struct {
		arg1 int64
		arg2 string
	}
	createActivationTokenReturns struct {
		result1 int64
		result2 error
	}
	createActivationTokenReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetAccountStub        func(string) (*models.Account, error)
	getAccountMutex       sync.RWMutex
	getAccountArgsForCall []struct {
		arg1 string
	}
	getAccountReturns struct {
		result1 *models.Account
		result2 error
	}
	getAccountReturnsOnCall map[int]struct {
		result1 *models.Account
		result2 error
	}
	GetAccountByEmailStub        func(string) (*models.Account, error)
	getAccountByEmailMutex       sync.RWMutex
	getAccountByEmailArgsForCall []struct {
		arg1 string
	}
	getAccountByEmailReturns struct {
		result1 *models.Account
		result2 error
	}
	getAccountByEmailReturnsOnCall map[int]struct {
		result1 *models.Account
		result2 error
	}
	GetAccountByPasswordResetTokenStub        func(string) (*models.Account, error)
	getAccountByPasswordResetTokenMutex       sync.RWMutex
	getAccountByPasswordResetTokenArgsForCall []struct {
		arg1 string
	}
	getAccountByPasswordResetTokenReturns struct {
		result1 *models.Account
		result2 error
	}
	getAccountByPasswordResetTokenReturnsOnCall map[int]struct {
		result1 *models.Account
		result2 error
	}
	GetActivationTokenByAccountIDStub        func(int64) (*models.ActivationToken, error)
	getActivationTokenByAccountIDMutex       sync.RWMutex
	getActivationTokenByAccountIDArgsForCall []struct {
		arg1 int64
	}
	getActivationTokenByAccountIDReturns struct {
		result1 *models.ActivationToken
		result2 error
	}
	getActivationTokenByAccountIDReturnsOnCall map[int]struct {
		result1 *models.ActivationToken
		result2 error
	}
	UpdateActivationTokenLastSentAtStub        func(int64) (int64, error)
	updateActivationTokenLastSentAtMutex       sync.RWMutex
	updateActivationTokenLastSentAtArgsForCall []struct {
		arg1 int64
	}
	updateActivationTokenLastSentAtReturns struct {
		result1 int64
		result2 error
	}
	updateActivationTokenLastSentAtReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccountRepository) ActivateAccount(arg1 string) (bool, error) {
	fake.activateAccountMutex.Lock()
	ret, specificReturn := fake.activateAccountReturnsOnCall[len(fake.activateAccountArgsForCall)]
	fake.activateAccountArgsForCall = append(fake.activateAccountArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ActivateAccount", []interface{}{arg1})
	fake.activateAccountMutex.Unlock()
	if fake.ActivateAccountStub != nil {
		return fake.ActivateAccountStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.activateAccountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) ActivateAccountCallCount() int {
	fake.activateAccountMutex.RLock()
	defer fake.activateAccountMutex.RUnlock()
	return len(fake.activateAccountArgsForCall)
}

func (fake *FakeAccountRepository) ActivateAccountCalls(stub func(string) (bool, error)) {
	fake.activateAccountMutex.Lock()
	defer fake.activateAccountMutex.Unlock()
	fake.ActivateAccountStub = stub
}

func (fake *FakeAccountRepository) ActivateAccountArgsForCall(i int) string {
	fake.activateAccountMutex.RLock()
	defer fake.activateAccountMutex.RUnlock()
	argsForCall := fake.activateAccountArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountRepository) ActivateAccountReturns(result1 bool, result2 error) {
	fake.activateAccountMutex.Lock()
	defer fake.activateAccountMutex.Unlock()
	fake.ActivateAccountStub = nil
	fake.activateAccountReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) ActivateAccountReturnsOnCall(i int, result1 bool, result2 error) {
	fake.activateAccountMutex.Lock()
	defer fake.activateAccountMutex.Unlock()
	fake.ActivateAccountStub = nil
	if fake.activateAccountReturnsOnCall == nil {
		fake.activateAccountReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.activateAccountReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) CreateAccount(arg1 string, arg2 string) (int64, error) {
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

func (fake *FakeAccountRepository) CreateAccountCallCount() int {
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	return len(fake.createAccountArgsForCall)
}

func (fake *FakeAccountRepository) CreateAccountCalls(stub func(string, string) (int64, error)) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = stub
}

func (fake *FakeAccountRepository) CreateAccountArgsForCall(i int) (string, string) {
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	argsForCall := fake.createAccountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAccountRepository) CreateAccountReturns(result1 int64, result2 error) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = nil
	fake.createAccountReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) CreateAccountReturnsOnCall(i int, result1 int64, result2 error) {
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

func (fake *FakeAccountRepository) CreateActivationToken(arg1 int64, arg2 string) (int64, error) {
	fake.createActivationTokenMutex.Lock()
	ret, specificReturn := fake.createActivationTokenReturnsOnCall[len(fake.createActivationTokenArgsForCall)]
	fake.createActivationTokenArgsForCall = append(fake.createActivationTokenArgsForCall, struct {
		arg1 int64
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateActivationToken", []interface{}{arg1, arg2})
	fake.createActivationTokenMutex.Unlock()
	if fake.CreateActivationTokenStub != nil {
		return fake.CreateActivationTokenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createActivationTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) CreateActivationTokenCallCount() int {
	fake.createActivationTokenMutex.RLock()
	defer fake.createActivationTokenMutex.RUnlock()
	return len(fake.createActivationTokenArgsForCall)
}

func (fake *FakeAccountRepository) CreateActivationTokenCalls(stub func(int64, string) (int64, error)) {
	fake.createActivationTokenMutex.Lock()
	defer fake.createActivationTokenMutex.Unlock()
	fake.CreateActivationTokenStub = stub
}

func (fake *FakeAccountRepository) CreateActivationTokenArgsForCall(i int) (int64, string) {
	fake.createActivationTokenMutex.RLock()
	defer fake.createActivationTokenMutex.RUnlock()
	argsForCall := fake.createActivationTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAccountRepository) CreateActivationTokenReturns(result1 int64, result2 error) {
	fake.createActivationTokenMutex.Lock()
	defer fake.createActivationTokenMutex.Unlock()
	fake.CreateActivationTokenStub = nil
	fake.createActivationTokenReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) CreateActivationTokenReturnsOnCall(i int, result1 int64, result2 error) {
	fake.createActivationTokenMutex.Lock()
	defer fake.createActivationTokenMutex.Unlock()
	fake.CreateActivationTokenStub = nil
	if fake.createActivationTokenReturnsOnCall == nil {
		fake.createActivationTokenReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createActivationTokenReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetAccount(arg1 string) (*models.Account, error) {
	fake.getAccountMutex.Lock()
	ret, specificReturn := fake.getAccountReturnsOnCall[len(fake.getAccountArgsForCall)]
	fake.getAccountArgsForCall = append(fake.getAccountArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetAccount", []interface{}{arg1})
	fake.getAccountMutex.Unlock()
	if fake.GetAccountStub != nil {
		return fake.GetAccountStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAccountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) GetAccountCallCount() int {
	fake.getAccountMutex.RLock()
	defer fake.getAccountMutex.RUnlock()
	return len(fake.getAccountArgsForCall)
}

func (fake *FakeAccountRepository) GetAccountCalls(stub func(string) (*models.Account, error)) {
	fake.getAccountMutex.Lock()
	defer fake.getAccountMutex.Unlock()
	fake.GetAccountStub = stub
}

func (fake *FakeAccountRepository) GetAccountArgsForCall(i int) string {
	fake.getAccountMutex.RLock()
	defer fake.getAccountMutex.RUnlock()
	argsForCall := fake.getAccountArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountRepository) GetAccountReturns(result1 *models.Account, result2 error) {
	fake.getAccountMutex.Lock()
	defer fake.getAccountMutex.Unlock()
	fake.GetAccountStub = nil
	fake.getAccountReturns = struct {
		result1 *models.Account
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetAccountReturnsOnCall(i int, result1 *models.Account, result2 error) {
	fake.getAccountMutex.Lock()
	defer fake.getAccountMutex.Unlock()
	fake.GetAccountStub = nil
	if fake.getAccountReturnsOnCall == nil {
		fake.getAccountReturnsOnCall = make(map[int]struct {
			result1 *models.Account
			result2 error
		})
	}
	fake.getAccountReturnsOnCall[i] = struct {
		result1 *models.Account
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetAccountByEmail(arg1 string) (*models.Account, error) {
	fake.getAccountByEmailMutex.Lock()
	ret, specificReturn := fake.getAccountByEmailReturnsOnCall[len(fake.getAccountByEmailArgsForCall)]
	fake.getAccountByEmailArgsForCall = append(fake.getAccountByEmailArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetAccountByEmail", []interface{}{arg1})
	fake.getAccountByEmailMutex.Unlock()
	if fake.GetAccountByEmailStub != nil {
		return fake.GetAccountByEmailStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAccountByEmailReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) GetAccountByEmailCallCount() int {
	fake.getAccountByEmailMutex.RLock()
	defer fake.getAccountByEmailMutex.RUnlock()
	return len(fake.getAccountByEmailArgsForCall)
}

func (fake *FakeAccountRepository) GetAccountByEmailCalls(stub func(string) (*models.Account, error)) {
	fake.getAccountByEmailMutex.Lock()
	defer fake.getAccountByEmailMutex.Unlock()
	fake.GetAccountByEmailStub = stub
}

func (fake *FakeAccountRepository) GetAccountByEmailArgsForCall(i int) string {
	fake.getAccountByEmailMutex.RLock()
	defer fake.getAccountByEmailMutex.RUnlock()
	argsForCall := fake.getAccountByEmailArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountRepository) GetAccountByEmailReturns(result1 *models.Account, result2 error) {
	fake.getAccountByEmailMutex.Lock()
	defer fake.getAccountByEmailMutex.Unlock()
	fake.GetAccountByEmailStub = nil
	fake.getAccountByEmailReturns = struct {
		result1 *models.Account
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetAccountByEmailReturnsOnCall(i int, result1 *models.Account, result2 error) {
	fake.getAccountByEmailMutex.Lock()
	defer fake.getAccountByEmailMutex.Unlock()
	fake.GetAccountByEmailStub = nil
	if fake.getAccountByEmailReturnsOnCall == nil {
		fake.getAccountByEmailReturnsOnCall = make(map[int]struct {
			result1 *models.Account
			result2 error
		})
	}
	fake.getAccountByEmailReturnsOnCall[i] = struct {
		result1 *models.Account
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetAccountByPasswordResetToken(arg1 string) (*models.Account, error) {
	fake.getAccountByPasswordResetTokenMutex.Lock()
	ret, specificReturn := fake.getAccountByPasswordResetTokenReturnsOnCall[len(fake.getAccountByPasswordResetTokenArgsForCall)]
	fake.getAccountByPasswordResetTokenArgsForCall = append(fake.getAccountByPasswordResetTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetAccountByPasswordResetToken", []interface{}{arg1})
	fake.getAccountByPasswordResetTokenMutex.Unlock()
	if fake.GetAccountByPasswordResetTokenStub != nil {
		return fake.GetAccountByPasswordResetTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAccountByPasswordResetTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) GetAccountByPasswordResetTokenCallCount() int {
	fake.getAccountByPasswordResetTokenMutex.RLock()
	defer fake.getAccountByPasswordResetTokenMutex.RUnlock()
	return len(fake.getAccountByPasswordResetTokenArgsForCall)
}

func (fake *FakeAccountRepository) GetAccountByPasswordResetTokenCalls(stub func(string) (*models.Account, error)) {
	fake.getAccountByPasswordResetTokenMutex.Lock()
	defer fake.getAccountByPasswordResetTokenMutex.Unlock()
	fake.GetAccountByPasswordResetTokenStub = stub
}

func (fake *FakeAccountRepository) GetAccountByPasswordResetTokenArgsForCall(i int) string {
	fake.getAccountByPasswordResetTokenMutex.RLock()
	defer fake.getAccountByPasswordResetTokenMutex.RUnlock()
	argsForCall := fake.getAccountByPasswordResetTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountRepository) GetAccountByPasswordResetTokenReturns(result1 *models.Account, result2 error) {
	fake.getAccountByPasswordResetTokenMutex.Lock()
	defer fake.getAccountByPasswordResetTokenMutex.Unlock()
	fake.GetAccountByPasswordResetTokenStub = nil
	fake.getAccountByPasswordResetTokenReturns = struct {
		result1 *models.Account
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetAccountByPasswordResetTokenReturnsOnCall(i int, result1 *models.Account, result2 error) {
	fake.getAccountByPasswordResetTokenMutex.Lock()
	defer fake.getAccountByPasswordResetTokenMutex.Unlock()
	fake.GetAccountByPasswordResetTokenStub = nil
	if fake.getAccountByPasswordResetTokenReturnsOnCall == nil {
		fake.getAccountByPasswordResetTokenReturnsOnCall = make(map[int]struct {
			result1 *models.Account
			result2 error
		})
	}
	fake.getAccountByPasswordResetTokenReturnsOnCall[i] = struct {
		result1 *models.Account
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetActivationTokenByAccountID(arg1 int64) (*models.ActivationToken, error) {
	fake.getActivationTokenByAccountIDMutex.Lock()
	ret, specificReturn := fake.getActivationTokenByAccountIDReturnsOnCall[len(fake.getActivationTokenByAccountIDArgsForCall)]
	fake.getActivationTokenByAccountIDArgsForCall = append(fake.getActivationTokenByAccountIDArgsForCall, struct {
		arg1 int64
	}{arg1})
	fake.recordInvocation("GetActivationTokenByAccountID", []interface{}{arg1})
	fake.getActivationTokenByAccountIDMutex.Unlock()
	if fake.GetActivationTokenByAccountIDStub != nil {
		return fake.GetActivationTokenByAccountIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getActivationTokenByAccountIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) GetActivationTokenByAccountIDCallCount() int {
	fake.getActivationTokenByAccountIDMutex.RLock()
	defer fake.getActivationTokenByAccountIDMutex.RUnlock()
	return len(fake.getActivationTokenByAccountIDArgsForCall)
}

func (fake *FakeAccountRepository) GetActivationTokenByAccountIDCalls(stub func(int64) (*models.ActivationToken, error)) {
	fake.getActivationTokenByAccountIDMutex.Lock()
	defer fake.getActivationTokenByAccountIDMutex.Unlock()
	fake.GetActivationTokenByAccountIDStub = stub
}

func (fake *FakeAccountRepository) GetActivationTokenByAccountIDArgsForCall(i int) int64 {
	fake.getActivationTokenByAccountIDMutex.RLock()
	defer fake.getActivationTokenByAccountIDMutex.RUnlock()
	argsForCall := fake.getActivationTokenByAccountIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountRepository) GetActivationTokenByAccountIDReturns(result1 *models.ActivationToken, result2 error) {
	fake.getActivationTokenByAccountIDMutex.Lock()
	defer fake.getActivationTokenByAccountIDMutex.Unlock()
	fake.GetActivationTokenByAccountIDStub = nil
	fake.getActivationTokenByAccountIDReturns = struct {
		result1 *models.ActivationToken
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) GetActivationTokenByAccountIDReturnsOnCall(i int, result1 *models.ActivationToken, result2 error) {
	fake.getActivationTokenByAccountIDMutex.Lock()
	defer fake.getActivationTokenByAccountIDMutex.Unlock()
	fake.GetActivationTokenByAccountIDStub = nil
	if fake.getActivationTokenByAccountIDReturnsOnCall == nil {
		fake.getActivationTokenByAccountIDReturnsOnCall = make(map[int]struct {
			result1 *models.ActivationToken
			result2 error
		})
	}
	fake.getActivationTokenByAccountIDReturnsOnCall[i] = struct {
		result1 *models.ActivationToken
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) UpdateActivationTokenLastSentAt(arg1 int64) (int64, error) {
	fake.updateActivationTokenLastSentAtMutex.Lock()
	ret, specificReturn := fake.updateActivationTokenLastSentAtReturnsOnCall[len(fake.updateActivationTokenLastSentAtArgsForCall)]
	fake.updateActivationTokenLastSentAtArgsForCall = append(fake.updateActivationTokenLastSentAtArgsForCall, struct {
		arg1 int64
	}{arg1})
	fake.recordInvocation("UpdateActivationTokenLastSentAt", []interface{}{arg1})
	fake.updateActivationTokenLastSentAtMutex.Unlock()
	if fake.UpdateActivationTokenLastSentAtStub != nil {
		return fake.UpdateActivationTokenLastSentAtStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateActivationTokenLastSentAtReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccountRepository) UpdateActivationTokenLastSentAtCallCount() int {
	fake.updateActivationTokenLastSentAtMutex.RLock()
	defer fake.updateActivationTokenLastSentAtMutex.RUnlock()
	return len(fake.updateActivationTokenLastSentAtArgsForCall)
}

func (fake *FakeAccountRepository) UpdateActivationTokenLastSentAtCalls(stub func(int64) (int64, error)) {
	fake.updateActivationTokenLastSentAtMutex.Lock()
	defer fake.updateActivationTokenLastSentAtMutex.Unlock()
	fake.UpdateActivationTokenLastSentAtStub = stub
}

func (fake *FakeAccountRepository) UpdateActivationTokenLastSentAtArgsForCall(i int) int64 {
	fake.updateActivationTokenLastSentAtMutex.RLock()
	defer fake.updateActivationTokenLastSentAtMutex.RUnlock()
	argsForCall := fake.updateActivationTokenLastSentAtArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountRepository) UpdateActivationTokenLastSentAtReturns(result1 int64, result2 error) {
	fake.updateActivationTokenLastSentAtMutex.Lock()
	defer fake.updateActivationTokenLastSentAtMutex.Unlock()
	fake.UpdateActivationTokenLastSentAtStub = nil
	fake.updateActivationTokenLastSentAtReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) UpdateActivationTokenLastSentAtReturnsOnCall(i int, result1 int64, result2 error) {
	fake.updateActivationTokenLastSentAtMutex.Lock()
	defer fake.updateActivationTokenLastSentAtMutex.Unlock()
	fake.UpdateActivationTokenLastSentAtStub = nil
	if fake.updateActivationTokenLastSentAtReturnsOnCall == nil {
		fake.updateActivationTokenLastSentAtReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.updateActivationTokenLastSentAtReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.activateAccountMutex.RLock()
	defer fake.activateAccountMutex.RUnlock()
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	fake.createActivationTokenMutex.RLock()
	defer fake.createActivationTokenMutex.RUnlock()
	fake.getAccountMutex.RLock()
	defer fake.getAccountMutex.RUnlock()
	fake.getAccountByEmailMutex.RLock()
	defer fake.getAccountByEmailMutex.RUnlock()
	fake.getAccountByPasswordResetTokenMutex.RLock()
	defer fake.getAccountByPasswordResetTokenMutex.RUnlock()
	fake.getActivationTokenByAccountIDMutex.RLock()
	defer fake.getActivationTokenByAccountIDMutex.RUnlock()
	fake.updateActivationTokenLastSentAtMutex.RLock()
	defer fake.updateActivationTokenLastSentAtMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAccountRepository) recordInvocation(key string, args []interface{}) {
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

var _ repositories.AccountRepository = new(FakeAccountRepository)
