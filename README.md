Here's a Go script that reads through a text file containing the Name, Email, and Role columns, compares the list of email addresses with another list of leavers, and produces a result that includes only the leavers:
1. Package and Imports: The script starts by defining the package and importing necessary packages.
2. Struct Definition: The Person struct holds the Name, Email, and Role of each person.
3. readFile Function:
    - Opens the file and handles potential errors.
    - Reads the file line by line, splits each line into parts, and checks if there are exactly 3 parts.
    - Adds valid lines to the persons slice as Person structs.
4. readLeavers Function:
    - Similar to readFile, it opens the file and reads it line by line.
    - Splits each line into parts and checks if there are exactly 2 parts.
    - Adds the email addresses to a map for quick lookup, converting them to lowercase to ensure case-insensitive comparison.
5. main Function:
    - Specifies the filenames for the persons and leavers files.
    - Reads data from these files using the defined functions.
    - Iterates through the persons and checks if their email addresses are in the leavers map.
    - Prints the details of the leavers.
This script is now well-commented to provide a clear understanding of each part and its functionality.

Running the Script:
1. Save the script in a file, e.g., main.go.
2. Create/Edit the persons.txt and leavers.txt files with the respective data.
3. Run the script using the Go command:
  go run main.go
  
This script assumes the data in persons.txt and leavers.txt is well-formatted. You might need to adjust it for different file formats or error handling as needed.
