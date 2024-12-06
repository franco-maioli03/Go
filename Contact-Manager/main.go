package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Contact structure
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Save contacts to a JSON file
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Load contacts from a JSON file
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Slice of contacts
	var contacts []Contact

	// Load existing contacts from the JSON file
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error loading contacts: ", err)
	}
	// Create a bufio instance
	reader := bufio.NewReader(os.Stdin)

	for {
		// Show options to the user
		fmt.Print("==== CONTACT MANAGER ====\n",
			"1. Add a contact\n",
			"2. Show all contacts\n",
			"3. Exit\n",
			"Choose an option: ")

		// Read user's choice
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error reading option: ", err)
			return
		}

		// Handle user options
		switch option {
		case 1:
			// Enter and create contact
			var c Contact
			fmt.Print("Name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')

			// Add contact to the slice
			contacts = append(contacts, c)

			// Save to JSON file
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error saving contact: ", err)
			}
		case 2:
			// Show all contacts
			fmt.Println("=============================")
			for index, contact := range contacts {
				fmt.Printf("%d. Name: %s Email: %s Phone: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=============================")
		case 3:
			// Exit the program
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
