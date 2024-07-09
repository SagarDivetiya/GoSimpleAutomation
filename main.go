package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Person struct to hold the data for each person
type Person struct {
	Name  string
	Email string
	Role  string
}

// readFile function to read the persons.txt file and return a slice of Person structs
func readFile(filename string) ([]Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // Return an error if the file cannot be opened
	}
	defer file.Close() // Ensure the file is closed after reading

	var persons []Person
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()        // Read each line
		parts := strings.Fields(line) // Split the line into parts (space-separated)
		if len(parts) != 3 {          // Check if the line has exactly 3 parts
			continue // Skip invalid lines
		}
		persons = append(persons, Person{parts[0], parts[1], parts[2]}) // Append the parsed data to persons slice
	}

	if err := scanner.Err(); err != nil {
		return nil, err // Return an error if there was an issue during scanning
	}

	return persons, nil // Return the slice of Person structs
}

// readLeavers function to read the leavers.txt file and return a map of leaver email addresses
func readLeavers(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // Return an error if the file cannot be opened
	}
	defer file.Close() // Ensure the file is closed after reading

	leavers := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()        // Read each line
		parts := strings.Fields(line) // Split the line into parts (space-separated)
		if len(parts) != 2 {          // Check if the line has exactly 2 parts
			continue // Skip invalid lines
		}
		leavers[strings.ToLower(parts[0])] = true // Store the email address in the map
	}

	if err := scanner.Err(); err != nil {
		return nil, err // Return an error if there was an issue during scanning
	}

	return leavers, nil // Return the map of leaver email addresses
}

func main() {
	personsFile := "persons.txt"
	leaversFile := "leavers.txt"

	// Read persons data from the file
	persons, err := readFile(personsFile)
	if err != nil {
		fmt.Println("Error reading persons file:", err)
		return
	}

	// Read leavers data from the file
	leavers, err := readLeavers(leaversFile)
	if err != nil {
		fmt.Println("Error reading leavers file:", err)
		return
	}

	// Compare and print the details of leavers
	for _, person := range persons {
		if leavers[strings.ToLower(person.Email)] { // Check if the email is in the leavers map
			fmt.Printf("%s %s %s\n", person.Name, person.Email, person.Role) // Print the details of the leaver
		}
	}
}
