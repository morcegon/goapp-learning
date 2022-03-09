package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickers uint

		println("Enter your first name:")
		fmt.Scan(&firstName)

		println("Enter your last name:")
		fmt.Scan(&lastName)

		println("Enter your email address:")
		fmt.Scan(&email)

		println("Enter numeber of tickets:")
		fmt.Scan(&userTickers)

		if userTickers <= remainingTickets {
			remainingTickets = remainingTickets - userTickers
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v.\n", firstName, lastName, userTickers, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
		}
	}
}
