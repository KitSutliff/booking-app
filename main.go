package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("This conference has %v and %v tickets remain available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ticktes to your favorite events here!")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v. You have purchased %v tickets to %v. Please check %v for your confirmation.\n", firstName, lastName, userTickets, conferenceName, email)
			fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
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
			if !isValidTickets {
				fmt.Println("Sorry, you have requested an invalid number of tickets.")
			}
		}
	}
}
