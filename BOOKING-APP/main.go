package main

import (
	com "booking-app/comunicationUser"
	"booking-app/helper"
	"booking-app/models"
	"fmt"
	"sync"
)

// Variables
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]models.UserData, 0)

var wg = sync.WaitGroup{}

func main() {

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	firstName, lastName, email, userTickets := helper.GetUseInput()

	isValidFullName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	var isValidFormUserInput = isValidFullName && isValidEmail && isValidUserTickets

	// check total values of remainong tickets
	if isValidFormUserInput {

		newBookings, newRemainingTickets := helper.BookTicket(userTickets, firstName, lastName, email, remainingTickets, bookings, conferenceName)

		// update values
		remainingTickets = newRemainingTickets
		bookings = newBookings

		// number of thread
		wg.Add(1)
		go com.SendTicket(&wg, userTickets, firstName, lastName, email)

		firstNames := helper.GetFirstNames(bookings)

		fmt.Printf("These are all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else {

		if !isValidFullName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doen't contain @ sing")
		}
		if !isValidUserTickets {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}

	wg.Wait()
}
