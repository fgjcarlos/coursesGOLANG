package comunicationuser

import (
	"fmt"
	"sync"
	"time"
)

func SendTicket(wg *sync.WaitGroup, userTickets int, firstName string, lastName string, email string) {

	// simulate load
	time.Sleep(3 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n %v \nto email to address %v\n", ticket, email)
	fmt.Println("#######################")

	defer wg.Done()
}
