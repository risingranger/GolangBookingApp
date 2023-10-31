package main

import (
	"booking-app/helper" // path refered from the GOPATH ie the mod file.This is for self created packages.
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference 2023"
var remainingTickets uint = 50

// var bookings[50] string //array declaration
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	//fmt.Println("We have total", conferenceTickets, "We currently have", remainingTickets, "tickets remaining")

	// for remainingTickets > 0 && remainingTickets <= 50 {
	firstName, lastName, email, userTickets := getUserInput()

	//packagename.FunctionName to call a function from another package

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// bookings[0] = firstName + " " + lastName //This can be done only if we have already defined the size of the array

		/*fmt.Printf("Hi %v, your booking is confirmed\n", bookings[0])
		fmt.Printf("Slice type %v\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))*/
		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1) // Wait for 1 goroutine to finish

		go sendTickets(userTickets, firstName, lastName, email) // go keyword is used to run the function in a separate thread. This is called go routine.

		firstNames := getFirstNames()
		fmt.Printf("First names of all the bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry! We are sold out! Please come back next year")
			//break
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
	// }
	wg.Wait() // Wait for all goroutines to finish
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v We currently have %v tickets remaining", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	//map[keytype]valuetype -> map is a key value pair. keytype can be string, int, float, bool, struct, array, slice, etc. valuetype can be anything.

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of all the bookings: %v\n", bookings)
	fmt.Printf("Your name is %v %v and your email is %v and you have booked %v tickets\n", firstName, lastName, email, userTickets)
	fmt.Printf("We have %v tickets remaining\n", remainingTickets)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //converts uint to string and 10 is the base
	//userData["tickets"] = fmt.Sprint(userTickets) //converts uint to string
	//user := map[string]string{}
	//fmt.Printf("The Whole array is %v\n", bookings)
}
func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("Sending tickets...")
	fmt.Printf("\n Sent: %v  to email address %v \n", ticket, email)
	fmt.Println("Tickets sent!")
	wg.Done() // Decrement the counter
}
