package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Println("This conference has", conferenceTickets, "and", remainingTickets, "tickets remain available.")
	fmt.Println("Get ticktes to your favorite events here!")

}
