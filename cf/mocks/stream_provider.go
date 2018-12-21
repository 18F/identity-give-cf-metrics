// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/alphagov/paas-prometheus-exporter/cf"
	sonde_events "github.com/cloudfoundry/sonde-go/events"
)

type FakeAppStreamProvider struct {
	StartStub        func() (<-chan *sonde_events.Envelope, <-chan error)
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 <-chan *sonde_events.Envelope
		result2 <-chan error
	}
	startReturnsOnCall map[int]struct {
		result1 <-chan *sonde_events.Envelope
		result2 <-chan error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAppStreamProvider) Start() (<-chan *sonde_events.Envelope, <-chan error) {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startReturns.result1, fake.startReturns.result2
}

func (fake *FakeAppStreamProvider) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeAppStreamProvider) StartReturns(result1 <-chan *sonde_events.Envelope, result2 <-chan error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 <-chan *sonde_events.Envelope
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeAppStreamProvider) StartReturnsOnCall(i int, result1 <-chan *sonde_events.Envelope, result2 <-chan error) {
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 <-chan *sonde_events.Envelope
			result2 <-chan error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 <-chan *sonde_events.Envelope
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeAppStreamProvider) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeAppStreamProvider) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeAppStreamProvider) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppStreamProvider) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppStreamProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAppStreamProvider) recordInvocation(key string, args []interface{}) {
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

var _ cf.AppStreamProvider = new(FakeAppStreamProvider)
