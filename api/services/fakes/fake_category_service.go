// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
)

type FakeCategoryService struct {
	GetCategoriesStub        func() ([]models.Category, error)
	getCategoriesMutex       sync.RWMutex
	getCategoriesArgsForCall []struct {
	}
	getCategoriesReturns struct {
		result1 []models.Category
		result2 error
	}
	getCategoriesReturnsOnCall map[int]struct {
		result1 []models.Category
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCategoryService) GetCategories() ([]models.Category, error) {
	fake.getCategoriesMutex.Lock()
	ret, specificReturn := fake.getCategoriesReturnsOnCall[len(fake.getCategoriesArgsForCall)]
	fake.getCategoriesArgsForCall = append(fake.getCategoriesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCategories", []interface{}{})
	fake.getCategoriesMutex.Unlock()
	if fake.GetCategoriesStub != nil {
		return fake.GetCategoriesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCategoriesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCategoryService) GetCategoriesCallCount() int {
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	return len(fake.getCategoriesArgsForCall)
}

func (fake *FakeCategoryService) GetCategoriesCalls(stub func() ([]models.Category, error)) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = stub
}

func (fake *FakeCategoryService) GetCategoriesReturns(result1 []models.Category, result2 error) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = nil
	fake.getCategoriesReturns = struct {
		result1 []models.Category
		result2 error
	}{result1, result2}
}

func (fake *FakeCategoryService) GetCategoriesReturnsOnCall(i int, result1 []models.Category, result2 error) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = nil
	if fake.getCategoriesReturnsOnCall == nil {
		fake.getCategoriesReturnsOnCall = make(map[int]struct {
			result1 []models.Category
			result2 error
		})
	}
	fake.getCategoriesReturnsOnCall[i] = struct {
		result1 []models.Category
		result2 error
	}{result1, result2}
}

func (fake *FakeCategoryService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCategoryService) recordInvocation(key string, args []interface{}) {
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

var _ services.CategoryService = new(FakeCategoryService)
