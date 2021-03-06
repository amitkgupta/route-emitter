// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	. "github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

type FakeStagerBBS struct {
	WatchForCompletedTaskStub        func() (<-chan models.Task, chan<- bool, <-chan error)
	watchForCompletedTaskMutex       sync.RWMutex
	watchForCompletedTaskArgsForCall []struct{}
	watchForCompletedTaskReturns struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}
	DesireTaskStub        func(models.Task) error
	desireTaskMutex       sync.RWMutex
	desireTaskArgsForCall []struct {
		arg1 models.Task
	}
	desireTaskReturns struct {
		result1 error
	}
	ResolvingTaskStub        func(taskGuid string) error
	resolvingTaskMutex       sync.RWMutex
	resolvingTaskArgsForCall []struct {
		arg1 string
	}
	resolvingTaskReturns struct {
		result1 error
	}
	ResolveTaskStub        func(taskGuid string) error
	resolveTaskMutex       sync.RWMutex
	resolveTaskArgsForCall []struct {
		arg1 string
	}
	resolveTaskReturns struct {
		result1 error
	}
	GetAvailableFileServerStub        func() (string, error)
	getAvailableFileServerMutex       sync.RWMutex
	getAvailableFileServerArgsForCall []struct{}
	getAvailableFileServerReturns struct {
		result1 string
		result2 error
	}
}

func (fake *FakeStagerBBS) WatchForCompletedTask() (<-chan models.Task, chan<- bool, <-chan error) {
	fake.watchForCompletedTaskMutex.Lock()
	defer fake.watchForCompletedTaskMutex.Unlock()
	fake.watchForCompletedTaskArgsForCall = append(fake.watchForCompletedTaskArgsForCall, struct{}{})
	if fake.WatchForCompletedTaskStub != nil {
		return fake.WatchForCompletedTaskStub()
	} else {
		return fake.watchForCompletedTaskReturns.result1, fake.watchForCompletedTaskReturns.result2, fake.watchForCompletedTaskReturns.result3
	}
}

func (fake *FakeStagerBBS) WatchForCompletedTaskCallCount() int {
	fake.watchForCompletedTaskMutex.RLock()
	defer fake.watchForCompletedTaskMutex.RUnlock()
	return len(fake.watchForCompletedTaskArgsForCall)
}

func (fake *FakeStagerBBS) WatchForCompletedTaskReturns(result1 <-chan models.Task, result2 chan<- bool, result3 <-chan error) {
	fake.watchForCompletedTaskReturns = struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeStagerBBS) DesireTask(arg1 models.Task) error {
	fake.desireTaskMutex.Lock()
	defer fake.desireTaskMutex.Unlock()
	fake.desireTaskArgsForCall = append(fake.desireTaskArgsForCall, struct {
		arg1 models.Task
	}{arg1})
	if fake.DesireTaskStub != nil {
		return fake.DesireTaskStub(arg1)
	} else {
		return fake.desireTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) DesireTaskCallCount() int {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return len(fake.desireTaskArgsForCall)
}

func (fake *FakeStagerBBS) DesireTaskArgsForCall(i int) models.Task {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return fake.desireTaskArgsForCall[i].arg1
}

func (fake *FakeStagerBBS) DesireTaskReturns(result1 error) {
	fake.desireTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) ResolvingTask(arg1 string) error {
	fake.resolvingTaskMutex.Lock()
	defer fake.resolvingTaskMutex.Unlock()
	fake.resolvingTaskArgsForCall = append(fake.resolvingTaskArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.ResolvingTaskStub != nil {
		return fake.ResolvingTaskStub(arg1)
	} else {
		return fake.resolvingTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) ResolvingTaskCallCount() int {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return len(fake.resolvingTaskArgsForCall)
}

func (fake *FakeStagerBBS) ResolvingTaskArgsForCall(i int) string {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return fake.resolvingTaskArgsForCall[i].arg1
}

func (fake *FakeStagerBBS) ResolvingTaskReturns(result1 error) {
	fake.resolvingTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) ResolveTask(arg1 string) error {
	fake.resolveTaskMutex.Lock()
	defer fake.resolveTaskMutex.Unlock()
	fake.resolveTaskArgsForCall = append(fake.resolveTaskArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.ResolveTaskStub != nil {
		return fake.ResolveTaskStub(arg1)
	} else {
		return fake.resolveTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) ResolveTaskCallCount() int {
	fake.resolveTaskMutex.RLock()
	defer fake.resolveTaskMutex.RUnlock()
	return len(fake.resolveTaskArgsForCall)
}

func (fake *FakeStagerBBS) ResolveTaskArgsForCall(i int) string {
	fake.resolveTaskMutex.RLock()
	defer fake.resolveTaskMutex.RUnlock()
	return fake.resolveTaskArgsForCall[i].arg1
}

func (fake *FakeStagerBBS) ResolveTaskReturns(result1 error) {
	fake.resolveTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) GetAvailableFileServer() (string, error) {
	fake.getAvailableFileServerMutex.Lock()
	defer fake.getAvailableFileServerMutex.Unlock()
	fake.getAvailableFileServerArgsForCall = append(fake.getAvailableFileServerArgsForCall, struct{}{})
	if fake.GetAvailableFileServerStub != nil {
		return fake.GetAvailableFileServerStub()
	} else {
		return fake.getAvailableFileServerReturns.result1, fake.getAvailableFileServerReturns.result2
	}
}

func (fake *FakeStagerBBS) GetAvailableFileServerCallCount() int {
	fake.getAvailableFileServerMutex.RLock()
	defer fake.getAvailableFileServerMutex.RUnlock()
	return len(fake.getAvailableFileServerArgsForCall)
}

func (fake *FakeStagerBBS) GetAvailableFileServerReturns(result1 string, result2 error) {
	fake.getAvailableFileServerReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ StagerBBS = new(FakeStagerBBS)
