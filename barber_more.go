package sleepy

type barberMore struct {
	seats chan Customer
}

func newBarberMore(seats int) Barber {
	return barberMore{
		seats: make(chan Customer, seats),
	}
}

// StartWork Run in background the haircut process. It should not call multiple times
func (b barberMore) StartWork() {
	go b.work()
}

// EnterCustomer return false if all seats is taken
func (b barberMore) EnterCustomer(c Customer) bool {
	select {
	case b.seats <- c:
		return true
	default:
		return false
	}
}

func (b barberMore) work() {
	for {
		c := <-b.seats
		c.cut()
	}
}
