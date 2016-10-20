// This file was generated by counterfeiter
package fakes

import (
	"lib/rules"
	"sync"
)

type Planner struct {
	GetRulesStub        func() (rules.RulesWithChain, error)
	getRulesMutex       sync.RWMutex
	getRulesArgsForCall []struct{}
	getRulesReturns     struct {
		result1 rules.RulesWithChain
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Planner) GetRules() (rules.RulesWithChain, error) {
	fake.getRulesMutex.Lock()
	fake.getRulesArgsForCall = append(fake.getRulesArgsForCall, struct{}{})
	fake.recordInvocation("GetRules", []interface{}{})
	fake.getRulesMutex.Unlock()
	if fake.GetRulesStub != nil {
		return fake.GetRulesStub()
	} else {
		return fake.getRulesReturns.result1, fake.getRulesReturns.result2
	}
}

func (fake *Planner) GetRulesCallCount() int {
	fake.getRulesMutex.RLock()
	defer fake.getRulesMutex.RUnlock()
	return len(fake.getRulesArgsForCall)
}

func (fake *Planner) GetRulesReturns(result1 rules.RulesWithChain, result2 error) {
	fake.GetRulesStub = nil
	fake.getRulesReturns = struct {
		result1 rules.RulesWithChain
		result2 error
	}{result1, result2}
}

func (fake *Planner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRulesMutex.RLock()
	defer fake.getRulesMutex.RUnlock()
	return fake.invocations
}

func (fake *Planner) recordInvocation(key string, args []interface{}) {
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
