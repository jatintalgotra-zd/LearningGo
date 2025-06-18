package main

import "fmt"

// Logger
// interface declares method Log() with string input
type Logger interface {
	Log(string)
}

// ConsoleLogger
// struct defining message string for console
type ConsoleLogger struct {
}

// Log
// method to simulate logging to console
func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

// FileLogger
// struct defining message string for file
type FileLogger struct {
	logs []string
}

// Log
// method to simulate logging to file
func (f *FileLogger) Log(message string) {
	f.logs = append(f.logs, message)
	fmt.Println("File:", message)
}

// RemoteLogger
// struct defining message string for remote server
type RemoteLogger struct {
}

// Log
// method to simulate logging to remote server
func (r *RemoteLogger) Log(message string) {
	fmt.Println("Remote:", message)
}

func main() {
	// console, file and remote instances to simulate respective logging
	console := ConsoleLogger{}
	file := FileLogger{}
	remote := RemoteLogger{}

	// slice of type Logger with values of console, file and remote logger
	all := []Logger{
		&console,
		&file,
		&remote,
	}

	// calling LogAll function
	LogAll(all, "Hello!")

}

// LogAll
// function to simulate logging to any type which implements Logger interface. Passes on value using Log method
func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		logger.Log(message)
	}
}
