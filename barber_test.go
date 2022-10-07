package sleepy

import (
	"testing"
	"time"
)

func TestBarber_0seats(t *testing.T) {
	barber := NewBarber(0)
	barber.StartWork()

	cust1 := NewCustomer()
	if true != barber.EnterCustomer(cust1) {
		t.Fatal("could not received the customer")
	}
	cust1.WaitToDone()
	barber.Close()
}

func TestBarber_0seatsMultiCustomer(t *testing.T) {
	barber := NewBarber(0)
	barber.StartWork()

	cust1 := NewCustomer()
	if true != barber.EnterCustomer(cust1) {
		t.Fatal("could not received the customer")
	}
	cust1.WaitToDone()

	cust2 := NewCustomer()
	if true != barber.EnterCustomer(cust2) {
		t.Fatal("could not received the customer")
	}
	cust2.WaitToDone()
	barber.Close()
}

func TestBarber_0seatsMultiCustomerCongestion(t *testing.T) {
	barber := NewBarber(0)

	cust1 := NewCustomer()
	if true != barber.EnterCustomer(cust1) {
		t.Fatal("could not received the customer")
	}

	cust2 := NewCustomer()
	if false != barber.EnterCustomer(cust2) {
		t.Fatal("could not received the customer")
	}
	barber.StartWork()
	cust1.WaitToDone()
	barber.Close()
}

func TestBarber_oneCustomer(t *testing.T) {
	barber := NewBarber(5)
	barber.StartWork()

	cust1 := NewCustomer()
	if true != barber.EnterCustomer(cust1) {
		t.Fatal("could not received the customer")
	}
	cust1.WaitToDone()
	barber.Close()
}

func TestBarber_multipleCustomer(t *testing.T) {
	barber := NewBarber(5)
	barber.StartWork()

	cust1 := NewCustomer()
	in := barber.EnterCustomer(cust1)
	if true != in {
		t.Fatal("invalid seats handling")
	}

	cust2 := NewCustomer()
	in = barber.EnterCustomer(cust2)
	if true != in {
		t.Fatal("invalid seats handling")
	}

	cust1.WaitToDone()
	cust2.WaitToDone()
	barber.Close()
}

func TestBarber_overLoad(t *testing.T) {
	b := NewBarber(1)
	// skip the startWork step

	cust1 := NewCustomer()
	if true != b.EnterCustomer(cust1) {
		t.Fatal("invalid seats handling")
	}

	cust2 := NewCustomer()
	if false != b.EnterCustomer(cust2) {
		t.Fatal("invalid seats handling")
	}
}

func TestBarber_overLoadAndRetry(t *testing.T) {
	barber := NewBarber(1)
	cust1 := NewCustomer()
	cust2 := NewCustomer()
	barber.EnterCustomer(cust1)

	if false != barber.EnterCustomer(cust2) {
		t.Fatal("invalid seats handling")
	}

	barber.StartWork()
	cust1.WaitToDone()

	if true != barber.EnterCustomer(cust2) {
		t.Fatal("invalid seats handling")
	}
	cust2.WaitToDone()
	barber.Close()
}

func TestBarber_burst(t *testing.T) {
	barber := NewBarber(1)
	barber.StartWork()

	customers := make([]Customer, 50)
	for i := 0; i < len(customers); i++ {
		c := NewCustomer()
		customers[i] = c
		go func(c Customer) {
			for true != barber.EnterCustomer(c) {
				time.Sleep(100 * time.Microsecond)
			}
		}(c)
	}

	for i := 0; i < len(customers); i++ {
		customers[i].WaitToDone()
	}
	barber.Close()
}
