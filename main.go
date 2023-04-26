package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("This conference has %v and %v tickets remain available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ticktes to your favorite events here!")

	var userName string
	var userTickets int
	// asks user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
