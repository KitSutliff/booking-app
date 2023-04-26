package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("This conference has %v and %v tickets remain available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ticktes to your favorite events here!")

}
