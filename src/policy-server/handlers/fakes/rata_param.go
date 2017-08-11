// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"
)

type RataParam struct {
	ParamStub        func(*http.Request, string) string
	paramMutex       sync.RWMutex
	paramArgsForCall []struct {
		arg1 *http.Request
		arg2 string
	}
	paramReturns struct {
		result1 string
	}
	paramReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RataParam) Param(arg1 *http.Request, arg2 string) string {
	fake.paramMutex.Lock()
	ret, specificReturn := fake.paramReturnsOnCall[len(fake.paramArgsForCall)]
	fake.paramArgsForCall = append(fake.paramArgsForCall, struct {
		arg1 *http.Request
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Param", []interface{}{arg1, arg2})
	fake.paramMutex.Unlock()
	if fake.ParamStub != nil {
		return fake.ParamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.paramReturns.result1
}

func (fake *RataParam) ParamCallCount() int {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return len(fake.paramArgsForCall)
}

func (fake *RataParam) ParamArgsForCall(i int) (*http.Request, string) {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return fake.paramArgsForCall[i].arg1, fake.paramArgsForCall[i].arg2
}

func (fake *RataParam) ParamReturns(result1 string) {
	fake.ParamStub = nil
	fake.paramReturns = struct {
		result1 string
	}{result1}
}

func (fake *RataParam) ParamReturnsOnCall(i int, result1 string) {
	fake.ParamStub = nil
	if fake.paramReturnsOnCall == nil {
		fake.paramReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.paramReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *RataParam) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RataParam) recordInvocation(key string, args []interface{}) {
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
