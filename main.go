package main
import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend.")

	var userName string
	// ask user for their name

	userName = "Tom"

}
