// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"sync"

	beerscli "github.com/CodelyTV/golang-introduction/08-automated_tests/internal"
)

var (
	lockBeerRepoMockGetBeers sync.RWMutex
)

// Ensure, that BeerRepoMock does implement BeerRepo.
// If this is not the case, regenerate this file with moq.
var _ beerscli.BeerRepo = &BeerRepoMock{}

// BeerRepoMock is a mock implementation of BeerRepo.
//
//     func TestSomethingThatUsesBeerRepo(t *testing.T) {
//
//         // make and configure a mocked BeerRepo
//         mockedBeerRepo := &BeerRepoMock{
//             GetBeersFunc: func() ([]beerscli.Beer, error) {
// 	               panic("mock out the GetBeers method")
//             },
//         }
//
//         // use mockedBeerRepo in code that requires BeerRepo
//         // and then make assertions.
//
//     }
type BeerRepoMock struct {
	// GetBeersFunc mocks the GetBeers method.
	GetBeersFunc func() ([]beerscli.Beer, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetBeers holds details about calls to the GetBeers method.
		GetBeers []struct {
		}
	}
}

// GetBeers calls GetBeersFunc.
func (mock *BeerRepoMock) GetBeers() ([]beerscli.Beer, error) {
	if mock.GetBeersFunc == nil {
		panic("BeerRepoMock.GetBeersFunc: method is nil but BeerRepo.GetBeers was just called")
	}
	callInfo := struct {
	}{}
	lockBeerRepoMockGetBeers.Lock()
	mock.calls.GetBeers = append(mock.calls.GetBeers, callInfo)
	lockBeerRepoMockGetBeers.Unlock()
	return mock.GetBeersFunc()
}

// GetBeersCalls gets all the calls that were made to GetBeers.
// Check the length with:
//     len(mockedBeerRepo.GetBeersCalls())
func (mock *BeerRepoMock) GetBeersCalls() []struct {
} {
	var calls []struct {
	}
	lockBeerRepoMockGetBeers.RLock()
	calls = mock.calls.GetBeers
	lockBeerRepoMockGetBeers.RUnlock()
	return calls
}
