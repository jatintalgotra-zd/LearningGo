package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Check
// handle error - prints the error message from argument
func Check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// CloseFile
// close file and handle error, if any occurred
func CloseFile(file *os.File) {
	err := file.Close()
	Check(err)
}
func main() {
	// Check correct CLI Input
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	}

	// Open file using os.Open()
	f, err2 := os.Open(os.Args[1])
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// Using defer to close file using CloseFile function
	defer CloseFile(f)

	// Variables for counting info, warning and error entries
	var info, warning, errors int

	// Create a buffered reader for opened file
	buf := bufio.NewReader(f)
	// Create a Scanner to read from the buffered reader
	scanner := bufio.NewScanner(buf)

	// Loop over the file until scanner cannot scan another token
	for scanner.Scan() {
		// Convert current token from scanner to string for processing
		currLine := scanner.Text()
		// Switch case to handle INFO, WARNING and ERROR log
		switch {
		// Comparing Prefix to check if a log is Info, warning or error
		case strings.HasPrefix(currLine, "[INFO]"):
			info++
		case strings.HasPrefix(currLine, "[WARNING]"):
			warning++
		case strings.HasPrefix(currLine, "[ERROR]"):
			errors++
		}
	}

	// Calculate total no of Log Lines
	total := info + warning + errors

	// Output the formated result in counts and percentages of each
	fmt.Println("Analysis of file: ", os.Args[1])
	fmt.Println("Info: ", info, " entries")
	fmt.Printf("Percentage: %.2f %% \n", float64(info)/float64(total)*100)
	fmt.Println("Warnings: ", warning, " entries")
	fmt.Printf("Percentage: %.2f %% \n", float64(warning)/float64(total)*100)
	fmt.Println("Errors: ", errors, " entries")
	fmt.Printf("Percentage: %.2f %% \n", float64(errors)/float64(total)*100)
	fmt.Println("Total: ", total, " entries")
	fmt.Println("Analysed at: ", time.Now().Format("2006-01-02 15:04:05"))

}
