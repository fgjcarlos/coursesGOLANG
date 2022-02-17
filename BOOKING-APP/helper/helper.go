package helper

import (
	"booking-app/models"
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	var isValidFullName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidUserTickets = (userTickets > 0) && (userTickets <= remainingTickets)

	return isValidEmail, isValidFullName, isValidUserTickets
}

func GreetUsers(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Println("We hace total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}

func GetFirstNames(bookings []models.UserData) []string {

	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}

	return firstNames
}

func GetUseInput() (string, string, string, int) {
	// var fullName string
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their first name, last name, email address, nº tickets
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func BookTicket(userTickets int, firstName string, lastName string, email string,
	remainingTickets int, bookings []models.UserData, conferenceName string) ([]models.UserData, int) {
	newRemainingTickets := remainingTickets - userTickets

	var userData = models.UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	newBookings := append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. Yo will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", newRemainingTickets, conferenceName)

	return newBookings, newRemainingTickets
}
