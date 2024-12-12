package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// ListContacts lists all contacts from the database
func ListContacts(db *sql.DB) {
	// SQL query to select all contacts
	query := "SELECT * FROM contact"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Iterate over the results and display them
	fmt.Println("\nCONTACT LIST: ")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	for rows.Next() {
		// Instance of the contact model
		contact := models.Contact{}

		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Without an email"
		}

		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	}
}

// GetContactByID retrieves a contact from the database by its ID
func GetContactByID(db *sql.DB, contactID int) {
	// SQL query to select a contact by its ID
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instance of the contact model
	contact := models.Contact{}
	var valueEmail sql.NullString // To handle null values

	// Scan the result into the contact model
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No contact found with ID %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Without an email"
	}

	fmt.Println("\nCONTACT DETAILS: ")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
}

// CreateContact registers a new contact in the database
func CreateContact(db *sql.DB, contact models.Contact) {
	// SQL statement to insert a new contact
	query := "INSERT INTO contact (name, email, phone) VALUES(?, ?, ?)"

	// Execute the SQL statement
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New contact successfully registered")
}

// UpdateContact updates an existing contact in the database
func UpdateContact(db *sql.DB, contact models.Contact) {
	// SQL statement to update a contact
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	// Execute the SQL statement
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact successfully updated")
}

// DeleteContact deletes an existing contact from the database
func DeleteContact(db *sql.DB, contactID int) {
	// SQL statement to delete a contact
	query := "DELETE FROM contact WHERE id = ?"
	// Execute the SQL statement
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact successfully deleted")
}
