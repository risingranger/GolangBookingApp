package main

import "fmt"

func main() {
	conferenceName := "Go Conference 2023"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{} //This is a slice

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Println("We have total", conferenceTickets, "We currently have", remainingTickets, "tickets remaining")
	fmt.Printf("We have total %v We currently have %v tickets remaining", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
	// fmt.Println("Conference Name: ", )

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

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName //This can be done only if we have already defined the size of the array
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The Whole array is %v\n", bookings)
	fmt.Printf("Hi %v, your booking is confirmed\n", bookings[0])
	fmt.Printf("Slice type %v\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Your name is %v %v and your email is %v and you have booked %v tickets", firstName, lastName, email, userTickets)
	fmt.Printf("We have %v tickets remaining\n", remainingTickets)

	fmt.Printf("The Total number of bookings is %v\n", bookings)
}
