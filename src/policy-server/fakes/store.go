// This file was generated by counterfeiter
package fakes

import (
	"policy-server/models"
	"policy-server/store"
	"sync"
)

type Store struct {
	CreateStub        func([]models.Policy) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []models.Policy
	}
	createReturns struct {
		result1 error
	}
	AllStub        func() ([]models.Policy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []models.Policy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Store) Create(arg1 []models.Policy) error {
	var arg1Copy []models.Policy
	if arg1 != nil {
		arg1Copy = make([]models.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []models.Policy
	}{arg1Copy})
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *Store) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Store) CreateArgsForCall(i int) []models.Policy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *Store) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) All() ([]models.Policy, error) {
	fake.allMutex.Lock()
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	} else {
		return fake.allReturns.result1, fake.allReturns.result2
	}
}

func (fake *Store) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *Store) AllReturns(result1 []models.Policy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return fake.invocations
}

func (fake *Store) recordInvocation(key string, args []interface{}) {
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

var _ store.Store = new(Store)
