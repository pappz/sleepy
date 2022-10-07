package sleepy

import (
	"sync"
)

// This implementation has been used in case the number of seats is 0
type barberSingle struct {
	mutex     *sync.Mutex
	wakeUp    chan Customer
	closeChan chan struct{}
}

func newBarberSingle() barberSingle {
	return barberSingle{
		mutex:     &sync.Mutex{},
		wakeUp:    make(chan Customer, 1),
		closeChan: make(chan struct{}),
	}
}

// StartWork Run in background the haircut process. It should not call multiple times
func (b barberSingle) StartWork() {
	go b.work()
}

// EnterCustomer return true if the shop is empty and the barber can start the work
func (b barberSingle) EnterCustomer(c Customer) bool {
	ok := b.mutex.TryLock()
	if !ok {
		return false
	}
	b.wakeUp <- c
	return true
}

// Close stop working
func (b barberSingle) Close() {
	select {
	case b.closeChan <- struct{}{}:
	default:
	}
}

func (b barberSingle) work() {
	for {
		select {
		case c := <-b.wakeUp:
			c.cut()
			b.mutex.Unlock()
		case _ = <-b.closeChan:
			return
		}
	}
}
