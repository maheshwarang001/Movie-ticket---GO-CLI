package main

import "fmt"

func main() {

	const movieName string = "Classical Drama"
	const movieTicket uint16 = 200
	var remainingTicket uint16 = 200

	fmt.Println("Welcome to", movieName, "booking Application")
	fmt.Println("We have total of", movieTicket, "tickets and", remainingTicket, "tickets are still available")
	fmt.Println("Get your tickets here to attend")

	fmt.Println()

	var ticketList []string

	var firstName string
	var lastName string
	var email string
	var purchaseTicket uint16

	fmt.Println("Enter you're First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter you're Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter you're Email Id:")
	fmt.Scan(&email)

	fmt.Println("Enter Ticket Qty:")
	fmt.Scan(&purchaseTicket)

	if purchaseTicket <= remainingTicket {
		remainingTicket -= purchaseTicket

		ticketList = append(ticketList, firstName+" "+lastName)
		fmt.Printf("User %v %v booked %v tickets successfully.\n", firstName, lastName, purchaseTicket)
		fmt.Printf("You'll get tickets at email %v.\n", email)
		fmt.Printf("Remaining tickets: %v\n", remainingTicket)

	} else {
		fmt.Println("Sorry, we don't have enough tickets available.")
	}

	//for i := 0; i < len(ticketList); i++ {
	//	fmt.Println(ticketList[i])
	//}

	for index, value := range ticketList {
		fmt.Println(index, "=", value)
	}
}
