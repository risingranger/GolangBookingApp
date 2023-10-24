package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference 2023"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var bookings[50] string //array declaration
	bookings := []string{} //This is a slice

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Println("We have total", conferenceTickets, "We currently have", remainingTickets, "tickets remaining")
	fmt.Printf("We have total %v We currently have %v tickets remaining", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
	// fmt.Println("Conference Name: ", )

	for remainingTickets > 0 && remainingTickets <= 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName //This can be done only if we have already defined the size of the array
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("The Whole array is %v\n", bookings)
			fmt.Printf("Hi %v, your booking is confirmed\n", bookings[0])
			fmt.Printf("Slice type %v\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Your name is %v %v and your email is %v and you have booked %v tickets\n", firstName, lastName, email, userTickets)
			fmt.Printf("We have %v tickets remaining\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Println("The first names of bookings are: ", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry! We are sold out! Please come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short. Please enter again.")
			}
			if !isValidEmail {
				fmt.Println("Email is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number is not valid")
			}
			// fmt.Printf("Sorry! We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			fmt.Println("Your input is invalid. Please try again")
		}

	}

}
