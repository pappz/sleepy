package main

import "github.com/pappz/sleepy"

func main() {
	barber := sleepy.NewBarber(5)
	customer := sleepy.NewCustomer()
	barber.StartWork()
	barber.EnterCustomer(customer)
	customer.WaitToDone()
}
