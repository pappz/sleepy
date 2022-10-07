package sleepy

import (
	"sync"
)

type barber0 struct {
	mutex  *sync.Mutex
	wakeUp chan Customer
}

func newBarber0() barber0 {
	return barber0{
		mutex:  &sync.Mutex{},
		wakeUp: make(chan Customer, 1),
	}
}

// StartWork Run in background the haircut process. It should not call multiple times
func (b barber0) StartWork() {
	go b.work()
}

// EnterCustomer return true if the shop is empty and the barber can start the work
func (b barber0) EnterCustomer(c Customer) bool {
	ok := b.mutex.TryLock()
	if !ok {
		return false
	}
	b.wakeUp <- c
	return true
}

func (b barber0) work() {
	for {
		c := <-b.wakeUp
		c.cut()
		b.mutex.Unlock()
	}
}
