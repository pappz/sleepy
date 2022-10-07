package sleepy

type Barber interface {
	StartWork()
	EnterCustomer(c Customer) bool
	Close()
}

// NewBarber based on the available seats return the best implementation
func NewBarber(seats int) Barber {
	if seats >= 1 {
		return newBarberMore(seats)
	} else {
		return newBarber0()
	}
}
