// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/models"
)

type FakeExpenseController struct {
	CreateExpenseStub        func(http.ResponseWriter, *models.Expense)
	createExpenseMutex       sync.RWMutex
	createExpenseArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *models.Expense
	}
	DeleteExpenseStub        func(http.ResponseWriter, *http.Request)
	deleteExpenseMutex       sync.RWMutex
	deleteExpenseArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	GetExpensesStub        func(http.ResponseWriter, *http.Request)
	getExpensesMutex       sync.RWMutex
	getExpensesArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	UpdateExpenseStub        func(http.ResponseWriter, *http.Request)
	updateExpenseMutex       sync.RWMutex
	updateExpenseArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExpenseController) CreateExpense(arg1 http.ResponseWriter, arg2 *models.Expense) {
	fake.createExpenseMutex.Lock()
	fake.createExpenseArgsForCall = append(fake.createExpenseArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *models.Expense
	}{arg1, arg2})
	fake.recordInvocation("CreateExpense", []interface{}{arg1, arg2})
	fake.createExpenseMutex.Unlock()
	if fake.CreateExpenseStub != nil {
		fake.CreateExpenseStub(arg1, arg2)
	}
}

func (fake *FakeExpenseController) CreateExpenseCallCount() int {
	fake.createExpenseMutex.RLock()
	defer fake.createExpenseMutex.RUnlock()
	return len(fake.createExpenseArgsForCall)
}

func (fake *FakeExpenseController) CreateExpenseCalls(stub func(http.ResponseWriter, *models.Expense)) {
	fake.createExpenseMutex.Lock()
	defer fake.createExpenseMutex.Unlock()
	fake.CreateExpenseStub = stub
}

func (fake *FakeExpenseController) CreateExpenseArgsForCall(i int) (http.ResponseWriter, *models.Expense) {
	fake.createExpenseMutex.RLock()
	defer fake.createExpenseMutex.RUnlock()
	argsForCall := fake.createExpenseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExpenseController) DeleteExpense(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.deleteExpenseMutex.Lock()
	fake.deleteExpenseArgsForCall = append(fake.deleteExpenseArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("DeleteExpense", []interface{}{arg1, arg2})
	fake.deleteExpenseMutex.Unlock()
	if fake.DeleteExpenseStub != nil {
		fake.DeleteExpenseStub(arg1, arg2)
	}
}

func (fake *FakeExpenseController) DeleteExpenseCallCount() int {
	fake.deleteExpenseMutex.RLock()
	defer fake.deleteExpenseMutex.RUnlock()
	return len(fake.deleteExpenseArgsForCall)
}

func (fake *FakeExpenseController) DeleteExpenseCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.deleteExpenseMutex.Lock()
	defer fake.deleteExpenseMutex.Unlock()
	fake.DeleteExpenseStub = stub
}

func (fake *FakeExpenseController) DeleteExpenseArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.deleteExpenseMutex.RLock()
	defer fake.deleteExpenseMutex.RUnlock()
	argsForCall := fake.deleteExpenseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExpenseController) GetExpenses(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.getExpensesMutex.Lock()
	fake.getExpensesArgsForCall = append(fake.getExpensesArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("GetExpenses", []interface{}{arg1, arg2})
	fake.getExpensesMutex.Unlock()
	if fake.GetExpensesStub != nil {
		fake.GetExpensesStub(arg1, arg2)
	}
}

func (fake *FakeExpenseController) GetExpensesCallCount() int {
	fake.getExpensesMutex.RLock()
	defer fake.getExpensesMutex.RUnlock()
	return len(fake.getExpensesArgsForCall)
}

func (fake *FakeExpenseController) GetExpensesCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.getExpensesMutex.Lock()
	defer fake.getExpensesMutex.Unlock()
	fake.GetExpensesStub = stub
}

func (fake *FakeExpenseController) GetExpensesArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.getExpensesMutex.RLock()
	defer fake.getExpensesMutex.RUnlock()
	argsForCall := fake.getExpensesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExpenseController) UpdateExpense(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.updateExpenseMutex.Lock()
	fake.updateExpenseArgsForCall = append(fake.updateExpenseArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("UpdateExpense", []interface{}{arg1, arg2})
	fake.updateExpenseMutex.Unlock()
	if fake.UpdateExpenseStub != nil {
		fake.UpdateExpenseStub(arg1, arg2)
	}
}

func (fake *FakeExpenseController) UpdateExpenseCallCount() int {
	fake.updateExpenseMutex.RLock()
	defer fake.updateExpenseMutex.RUnlock()
	return len(fake.updateExpenseArgsForCall)
}

func (fake *FakeExpenseController) UpdateExpenseCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.updateExpenseMutex.Lock()
	defer fake.updateExpenseMutex.Unlock()
	fake.UpdateExpenseStub = stub
}

func (fake *FakeExpenseController) UpdateExpenseArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.updateExpenseMutex.RLock()
	defer fake.updateExpenseMutex.RUnlock()
	argsForCall := fake.updateExpenseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExpenseController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createExpenseMutex.RLock()
	defer fake.createExpenseMutex.RUnlock()
	fake.deleteExpenseMutex.RLock()
	defer fake.deleteExpenseMutex.RUnlock()
	fake.getExpensesMutex.RLock()
	defer fake.getExpensesMutex.RUnlock()
	fake.updateExpenseMutex.RLock()
	defer fake.updateExpenseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExpenseController) recordInvocation(key string, args []interface{}) {
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

var _ controllers.ExpenseController = new(FakeExpenseController)
