// This file was generated by counterfeiter
package fetcherfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/fetcher"
	"code.cloudfoundry.org/lager"
)

type FakeCacheDriver struct {
	FetchBlobStub        func(logger lager.Logger, id string, remoteBlobFunc fetcher.RemoteBlobFunc) ([]byte, int64, error)
	fetchBlobMutex       sync.RWMutex
	fetchBlobArgsForCall []struct {
		logger         lager.Logger
		id             string
		remoteBlobFunc fetcher.RemoteBlobFunc
	}
	fetchBlobReturns struct {
		result1 []byte
		result2 int64
		result3 error
	}
	fetchBlobReturnsOnCall map[int]struct {
		result1 []byte
		result2 int64
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCacheDriver) FetchBlob(logger lager.Logger, id string, remoteBlobFunc fetcher.RemoteBlobFunc) ([]byte, int64, error) {
	fake.fetchBlobMutex.Lock()
	ret, specificReturn := fake.fetchBlobReturnsOnCall[len(fake.fetchBlobArgsForCall)]
	fake.fetchBlobArgsForCall = append(fake.fetchBlobArgsForCall, struct {
		logger         lager.Logger
		id             string
		remoteBlobFunc fetcher.RemoteBlobFunc
	}{logger, id, remoteBlobFunc})
	fake.recordInvocation("FetchBlob", []interface{}{logger, id, remoteBlobFunc})
	fake.fetchBlobMutex.Unlock()
	if fake.FetchBlobStub != nil {
		return fake.FetchBlobStub(logger, id, remoteBlobFunc)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.fetchBlobReturns.result1, fake.fetchBlobReturns.result2, fake.fetchBlobReturns.result3
}

func (fake *FakeCacheDriver) FetchBlobCallCount() int {
	fake.fetchBlobMutex.RLock()
	defer fake.fetchBlobMutex.RUnlock()
	return len(fake.fetchBlobArgsForCall)
}

func (fake *FakeCacheDriver) FetchBlobArgsForCall(i int) (lager.Logger, string, fetcher.RemoteBlobFunc) {
	fake.fetchBlobMutex.RLock()
	defer fake.fetchBlobMutex.RUnlock()
	return fake.fetchBlobArgsForCall[i].logger, fake.fetchBlobArgsForCall[i].id, fake.fetchBlobArgsForCall[i].remoteBlobFunc
}

func (fake *FakeCacheDriver) FetchBlobReturns(result1 []byte, result2 int64, result3 error) {
	fake.FetchBlobStub = nil
	fake.fetchBlobReturns = struct {
		result1 []byte
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCacheDriver) FetchBlobReturnsOnCall(i int, result1 []byte, result2 int64, result3 error) {
	fake.FetchBlobStub = nil
	if fake.fetchBlobReturnsOnCall == nil {
		fake.fetchBlobReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 int64
			result3 error
		})
	}
	fake.fetchBlobReturnsOnCall[i] = struct {
		result1 []byte
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCacheDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchBlobMutex.RLock()
	defer fake.fetchBlobMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCacheDriver) recordInvocation(key string, args []interface{}) {
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

var _ fetcher.CacheDriver = new(FakeCacheDriver)
