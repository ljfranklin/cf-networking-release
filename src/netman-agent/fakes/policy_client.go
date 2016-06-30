// This file was generated by counterfeiter
package fakes

import (
	"netman-agent/models"
	"sync"
)

type PolicyClient struct {
	GetPoliciesStub        func() ([]models.Policy, error)
	getPoliciesMutex       sync.RWMutex
	getPoliciesArgsForCall []struct{}
	getPoliciesReturns     struct {
		result1 []models.Policy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyClient) GetPolicies() ([]models.Policy, error) {
	fake.getPoliciesMutex.Lock()
	fake.getPoliciesArgsForCall = append(fake.getPoliciesArgsForCall, struct{}{})
	fake.recordInvocation("GetPolicies", []interface{}{})
	fake.getPoliciesMutex.Unlock()
	if fake.GetPoliciesStub != nil {
		return fake.GetPoliciesStub()
	} else {
		return fake.getPoliciesReturns.result1, fake.getPoliciesReturns.result2
	}
}

func (fake *PolicyClient) GetPoliciesCallCount() int {
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	return len(fake.getPoliciesArgsForCall)
}

func (fake *PolicyClient) GetPoliciesReturns(result1 []models.Policy, result2 error) {
	fake.GetPoliciesStub = nil
	fake.getPoliciesReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	return fake.invocations
}

func (fake *PolicyClient) recordInvocation(key string, args []interface{}) {
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
