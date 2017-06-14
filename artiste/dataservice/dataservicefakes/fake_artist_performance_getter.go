// This file was generated by counterfeiter
package dataservicefakes

import (
	"sync"

	"github.com/benwaine/artistprof/artiste/dataservice"
	"github.com/benwaine/artistprof/artiste/dataservice/clients"
)

type FakeArtistPerformanceGetter struct {
	GetArtistPerformancesStub        func(artistId string) ([]clients.PerformanceEvent, error)
	getArtistPerformancesMutex       sync.RWMutex
	getArtistPerformancesArgsForCall []struct {
		artistId string
	}
	getArtistPerformancesReturns struct {
		result1 []clients.PerformanceEvent
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtistPerformanceGetter) GetArtistPerformances(artistId string) ([]clients.PerformanceEvent, error) {
	fake.getArtistPerformancesMutex.Lock()
	fake.getArtistPerformancesArgsForCall = append(fake.getArtistPerformancesArgsForCall, struct {
		artistId string
	}{artistId})
	fake.recordInvocation("GetArtistPerformances", []interface{}{artistId})
	fake.getArtistPerformancesMutex.Unlock()
	if fake.GetArtistPerformancesStub != nil {
		return fake.GetArtistPerformancesStub(artistId)
	} else {
		return fake.getArtistPerformancesReturns.result1, fake.getArtistPerformancesReturns.result2
	}
}

func (fake *FakeArtistPerformanceGetter) GetArtistPerformancesCallCount() int {
	fake.getArtistPerformancesMutex.RLock()
	defer fake.getArtistPerformancesMutex.RUnlock()
	return len(fake.getArtistPerformancesArgsForCall)
}

func (fake *FakeArtistPerformanceGetter) GetArtistPerformancesArgsForCall(i int) string {
	fake.getArtistPerformancesMutex.RLock()
	defer fake.getArtistPerformancesMutex.RUnlock()
	return fake.getArtistPerformancesArgsForCall[i].artistId
}

func (fake *FakeArtistPerformanceGetter) GetArtistPerformancesReturns(result1 []clients.PerformanceEvent, result2 error) {
	fake.GetArtistPerformancesStub = nil
	fake.getArtistPerformancesReturns = struct {
		result1 []clients.PerformanceEvent
		result2 error
	}{result1, result2}
}

func (fake *FakeArtistPerformanceGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getArtistPerformancesMutex.RLock()
	defer fake.getArtistPerformancesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeArtistPerformanceGetter) recordInvocation(key string, args []interface{}) {
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

var _ dataservice.ArtistPerformanceGetter = new(FakeArtistPerformanceGetter)