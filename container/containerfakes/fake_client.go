// Code generated by counterfeiter. DO NOT EDIT.
package containerfakes

import (
	"sync"

	"github.com/topherbullock/garden-k8s/container"
)

type FakeClient struct {
	PodsStub        func(string) container.Pods
	podsMutex       sync.RWMutex
	podsArgsForCall []struct {
		arg1 string
	}
	podsReturns struct {
		result1 container.Pods
	}
	podsReturnsOnCall map[int]struct {
		result1 container.Pods
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Pods(arg1 string) container.Pods {
	fake.podsMutex.Lock()
	ret, specificReturn := fake.podsReturnsOnCall[len(fake.podsArgsForCall)]
	fake.podsArgsForCall = append(fake.podsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Pods", []interface{}{arg1})
	fake.podsMutex.Unlock()
	if fake.PodsStub != nil {
		return fake.PodsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.podsReturns.result1
}

func (fake *FakeClient) PodsCallCount() int {
	fake.podsMutex.RLock()
	defer fake.podsMutex.RUnlock()
	return len(fake.podsArgsForCall)
}

func (fake *FakeClient) PodsArgsForCall(i int) string {
	fake.podsMutex.RLock()
	defer fake.podsMutex.RUnlock()
	return fake.podsArgsForCall[i].arg1
}

func (fake *FakeClient) PodsReturns(result1 container.Pods) {
	fake.PodsStub = nil
	fake.podsReturns = struct {
		result1 container.Pods
	}{result1}
}

func (fake *FakeClient) PodsReturnsOnCall(i int, result1 container.Pods) {
	fake.PodsStub = nil
	if fake.podsReturnsOnCall == nil {
		fake.podsReturnsOnCall = make(map[int]struct {
			result1 container.Pods
		})
	}
	fake.podsReturnsOnCall[i] = struct {
		result1 container.Pods
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.podsMutex.RLock()
	defer fake.podsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ container.Client = new(FakeClient)