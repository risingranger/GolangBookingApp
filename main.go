package main

import "fmt"

func main() {
	var conferenceName = "Go Conference 2023"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Println("We have total", conferenceTickets, "We currently have", remainingTickets, "tickets remaining")
	fmt.Printf("We have total %v We currently have %v tickets remaining", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
	// fmt.Println("Conference Name: ", )

	var userName string
	var userTicket int
	userName = "John Doe"
	userTicket = 5
	fmt.Printf("User %v has booked %v tickets\n", userName, userTicket)
}
