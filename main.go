package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {

				fmt.Println("Our conference is sold out. Thank you for your interest.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Sorry, your name must be at least 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("Sorry, your email address is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Sorry, you have requested an invalid number of tickets.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("This conference has %v and %v tickets remain available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ticktes to your favorite events here!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you would like to purchase:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v. You have purchased %v tickets to %v. Please check %v for your confirmation.\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
}
