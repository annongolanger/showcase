// This file was generated by counterfeiter
package dataservicefakes

import (
	"sync"

	"github.com/benwaine/artistprof/artiste/dataservice"
)

type FakeGetSupportedArtists struct {
	GetSupportedArtistsStub        func() ([]dataservice.Artist, error)
	getSupportedArtistsMutex       sync.RWMutex
	getSupportedArtistsArgsForCall []struct{}
	getSupportedArtistsReturns     struct {
		result1 []dataservice.Artist
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetSupportedArtists) GetSupportedArtists() ([]dataservice.Artist, error) {
	fake.getSupportedArtistsMutex.Lock()
	fake.getSupportedArtistsArgsForCall = append(fake.getSupportedArtistsArgsForCall, struct{}{})
	fake.recordInvocation("GetSupportedArtists", []interface{}{})
	fake.getSupportedArtistsMutex.Unlock()
	if fake.GetSupportedArtistsStub != nil {
		return fake.GetSupportedArtistsStub()
	} else {
		return fake.getSupportedArtistsReturns.result1, fake.getSupportedArtistsReturns.result2
	}
}

func (fake *FakeGetSupportedArtists) GetSupportedArtistsCallCount() int {
	fake.getSupportedArtistsMutex.RLock()
	defer fake.getSupportedArtistsMutex.RUnlock()
	return len(fake.getSupportedArtistsArgsForCall)
}

func (fake *FakeGetSupportedArtists) GetSupportedArtistsReturns(result1 []dataservice.Artist, result2 error) {
	fake.GetSupportedArtistsStub = nil
	fake.getSupportedArtistsReturns = struct {
		result1 []dataservice.Artist
		result2 error
	}{result1, result2}
}

func (fake *FakeGetSupportedArtists) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSupportedArtistsMutex.RLock()
	defer fake.getSupportedArtistsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGetSupportedArtists) recordInvocation(key string, args []interface{}) {
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

var _ dataservice.GetSupportedArtists = new(FakeGetSupportedArtists)
