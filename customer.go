package sleepy

import (
	"sync"
)

type Customer struct {
	wg *sync.WaitGroup
}

// NewCustomer create new Customer instance and initialize it
func NewCustomer() Customer {
	c := Customer{
		wg: &sync.WaitGroup{},
	}
	c.wg.Add(1)
	return c
}

// WaitToDone wait the func until the barber do not cut the hair
func (c Customer) WaitToDone() {
	c.wg.Wait()
}

func (c Customer) cut() {
	c.wg.Done()
}
