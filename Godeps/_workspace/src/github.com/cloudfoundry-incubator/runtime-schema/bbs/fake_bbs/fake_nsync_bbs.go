// This file was generated by counterfeiter
package fake_bbs

import (
	"github.com/cloudfoundry-incubator/runtime-schema/bbs"

	"github.com/cloudfoundry-incubator/runtime-schema/models"

	"sync"
)

type FakeNsyncBBS struct {
	DesireLRPStub        func(models.DesiredLRP) error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		arg1 models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 error
	}
	RemoveDesiredLRPByProcessGuidStub        func(guid string) error
	removeDesiredLRPByProcessGuidMutex       sync.RWMutex
	removeDesiredLRPByProcessGuidArgsForCall []struct {
		guid string
	}
	removeDesiredLRPByProcessGuidReturns struct {
		result1 error
	}
	GetAllDesiredLRPsByDomainStub        func(domain string) ([]models.DesiredLRP, error)
	getAllDesiredLRPsByDomainMutex       sync.RWMutex
	getAllDesiredLRPsByDomainArgsForCall []struct {
		domain string
	}
	getAllDesiredLRPsByDomainReturns struct {
		result1 []models.DesiredLRP
		result2 error
	}
	ChangeDesiredLRPStub        func(change models.DesiredLRPChange) error
	changeDesiredLRPMutex       sync.RWMutex
	changeDesiredLRPArgsForCall []struct {
		change models.DesiredLRPChange
	}
	changeDesiredLRPReturns struct {
		result1 error
	}
}

func (fake *FakeNsyncBBS) DesireLRP(arg1 models.DesiredLRP) error {
	fake.desireLRPMutex.Lock()
	defer fake.desireLRPMutex.Unlock()
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		arg1 models.DesiredLRP
	}{arg1})
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(arg1)
	} else {
		return fake.desireLRPReturns.result1
	}
}

func (fake *FakeNsyncBBS) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeNsyncBBS) DesireLRPArgsForCall(i int) models.DesiredLRP {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return fake.desireLRPArgsForCall[i].arg1
}

func (fake *FakeNsyncBBS) DesireLRPReturns(result1 error) {
	fake.desireLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNsyncBBS) RemoveDesiredLRPByProcessGuid(guid string) error {
	fake.removeDesiredLRPByProcessGuidMutex.Lock()
	defer fake.removeDesiredLRPByProcessGuidMutex.Unlock()
	fake.removeDesiredLRPByProcessGuidArgsForCall = append(fake.removeDesiredLRPByProcessGuidArgsForCall, struct {
		guid string
	}{guid})
	if fake.RemoveDesiredLRPByProcessGuidStub != nil {
		return fake.RemoveDesiredLRPByProcessGuidStub(guid)
	} else {
		return fake.removeDesiredLRPByProcessGuidReturns.result1
	}
}

func (fake *FakeNsyncBBS) RemoveDesiredLRPByProcessGuidCallCount() int {
	fake.removeDesiredLRPByProcessGuidMutex.RLock()
	defer fake.removeDesiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.removeDesiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeNsyncBBS) RemoveDesiredLRPByProcessGuidArgsForCall(i int) string {
	fake.removeDesiredLRPByProcessGuidMutex.RLock()
	defer fake.removeDesiredLRPByProcessGuidMutex.RUnlock()
	return fake.removeDesiredLRPByProcessGuidArgsForCall[i].guid
}

func (fake *FakeNsyncBBS) RemoveDesiredLRPByProcessGuidReturns(result1 error) {
	fake.removeDesiredLRPByProcessGuidReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNsyncBBS) GetAllDesiredLRPsByDomain(domain string) ([]models.DesiredLRP, error) {
	fake.getAllDesiredLRPsByDomainMutex.Lock()
	defer fake.getAllDesiredLRPsByDomainMutex.Unlock()
	fake.getAllDesiredLRPsByDomainArgsForCall = append(fake.getAllDesiredLRPsByDomainArgsForCall, struct {
		domain string
	}{domain})
	if fake.GetAllDesiredLRPsByDomainStub != nil {
		return fake.GetAllDesiredLRPsByDomainStub(domain)
	} else {
		return fake.getAllDesiredLRPsByDomainReturns.result1, fake.getAllDesiredLRPsByDomainReturns.result2
	}
}

func (fake *FakeNsyncBBS) GetAllDesiredLRPsByDomainCallCount() int {
	fake.getAllDesiredLRPsByDomainMutex.RLock()
	defer fake.getAllDesiredLRPsByDomainMutex.RUnlock()
	return len(fake.getAllDesiredLRPsByDomainArgsForCall)
}

func (fake *FakeNsyncBBS) GetAllDesiredLRPsByDomainArgsForCall(i int) string {
	fake.getAllDesiredLRPsByDomainMutex.RLock()
	defer fake.getAllDesiredLRPsByDomainMutex.RUnlock()
	return fake.getAllDesiredLRPsByDomainArgsForCall[i].domain
}

func (fake *FakeNsyncBBS) GetAllDesiredLRPsByDomainReturns(result1 []models.DesiredLRP, result2 error) {
	fake.getAllDesiredLRPsByDomainReturns = struct {
		result1 []models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeNsyncBBS) ChangeDesiredLRP(change models.DesiredLRPChange) error {
	fake.changeDesiredLRPMutex.Lock()
	defer fake.changeDesiredLRPMutex.Unlock()
	fake.changeDesiredLRPArgsForCall = append(fake.changeDesiredLRPArgsForCall, struct {
		change models.DesiredLRPChange
	}{change})
	if fake.ChangeDesiredLRPStub != nil {
		return fake.ChangeDesiredLRPStub(change)
	} else {
		return fake.changeDesiredLRPReturns.result1
	}
}

func (fake *FakeNsyncBBS) ChangeDesiredLRPCallCount() int {
	fake.changeDesiredLRPMutex.RLock()
	defer fake.changeDesiredLRPMutex.RUnlock()
	return len(fake.changeDesiredLRPArgsForCall)
}

func (fake *FakeNsyncBBS) ChangeDesiredLRPArgsForCall(i int) models.DesiredLRPChange {
	fake.changeDesiredLRPMutex.RLock()
	defer fake.changeDesiredLRPMutex.RUnlock()
	return fake.changeDesiredLRPArgsForCall[i].change
}

func (fake *FakeNsyncBBS) ChangeDesiredLRPReturns(result1 error) {
	fake.changeDesiredLRPReturns = struct {
		result1 error
	}{result1}
}

var _ bbs.NsyncBBS = new(FakeNsyncBBS)
