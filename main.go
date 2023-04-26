package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("This conference has %v and %v tickets remain available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ticktes to your favorite events here!")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// asks user for their name
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you would like to purchase:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v. You have purchased %v tickets to %v. Please check %v for your confirmation.\n", firstName, lastName, userTickets, conferenceName, email)
}
