package sleepy

type barberMore struct {
	seats     chan Customer
	closeChan chan struct{}
}

func newBarberMore(seats int) Barber {
	return barberMore{
		seats:     make(chan Customer, seats),
		closeChan: make(chan struct{}),
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

// Close stop working
func (b barberMore) Close() {
	select {
	case b.closeChan <- struct{}{}:
	default:
	}
}

func (b barberMore) work() {
	for {
		select {
		case c := <-b.seats:
			c.cut()
		case _ = <-b.closeChan:
			return
		}
	}
}
