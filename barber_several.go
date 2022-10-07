package sleepy

// This implementation has been used in case have waiting room with several seats
type barberSeveral struct {
	seats     chan Customer
	closeChan chan struct{}
}

func newBarberSeveral(seats int) Barber {
	return barberSeveral{
		seats:     make(chan Customer, seats),
		closeChan: make(chan struct{}),
	}
}

// StartWork Run in background the haircut process. It should not call multiple times
func (b barberSeveral) StartWork() {
	go b.work()
}

// EnterCustomer return false if all seats is taken
func (b barberSeveral) EnterCustomer(c Customer) bool {
	select {
	case b.seats <- c:
		return true
	default:
		return false
	}
}

// Close stop working
func (b barberSeveral) Close() {
	select {
	case b.closeChan <- struct{}{}:
	default:
	}
}

func (b barberSeveral) work() {
	for {
		select {
		case c := <-b.seats:
			c.cut()
		case _ = <-b.closeChan:
			return
		}
	}
}
