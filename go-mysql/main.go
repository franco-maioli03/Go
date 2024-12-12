package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Establish connection to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. List Contacts")
		fmt.Println("2. Get Contact by ID")
		fmt.Println("3. Create New Contact")
		fmt.Println("4. Update Contact")
		fmt.Println("5. Delete Contact")
		fmt.Println("6. Exit")
		fmt.Print("Select an option: ")

		// Read the option selected by the user
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Enter the Contact ID: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Print("Enter the ID of the contact you want to delete: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}

// Function to input contact details
func inputContactDetails(option int) models.Contact {
	// Read user input using bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Print("Enter the ID of the contact you want to edit: ")
		var idContact int
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}

	fmt.Print("Enter the contact's name: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Enter the contact's email: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Enter the contact's phone: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
