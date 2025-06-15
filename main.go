package main

import (
	"bufio"
	"fmt"
	"os"
)

// Task
// added fields for description and status
type Task struct {
	Desc   string
	Status bool
}

// GenId () func() int
// returns function to get next id
func GenId() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// Check (e error)
// handle error - prints the error message from argument
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// AddTask (description string, nextID func() int, mp map[int]*Task)
// adds new Task by generating id
func AddTask(description string, nextID func() int, mp map[int]*Task) {
	id := nextID()
	t1 := Task{Desc: description}
	mp[id] = &t1
	fmt.Println("Task added:", id, "-", description)

}

// ListTasks
// prints all pending tasks
func ListTasks(mp map[int]*Task) {
	for k, v := range mp {
		if !v.Status {
			fmt.Println(k, ". ", v.Desc)
		}
	}
}

// CompleteTask
// marks Task complete by id
func CompleteTask(id int, mp map[int]*Task) {
	t1 := mp[id]
	t1.Status = true
}

func main() {

	// map used to store tasks
	mp := make(map[int]*Task)
	// function to generate next id
	nextId := GenId()

	// using for to loop over user inputs
	for {
		// menu style input formatting
		fmt.Println("Input 1 to add Task, 2 to list pending Task, 3 to complete Task, 0 to quit")
		var n int
		_, err := fmt.Scan(&n)
		Check(err)

		// close program
		if n == 0 {
			break
		}

		// switch case to handle
		switch n {

		// enter new Task
		case 1:
			fmt.Println("Enter Task: ")
			var t1 string
			bufReader := bufio.NewReader(os.Stdin)
			t1, err := bufReader.ReadString('\n')
			Check(err)
			AddTask(t1, nextId, mp)

		// list pending tasks
		case 2:
			fmt.Println("Pending Task: ")
			ListTasks(mp)

		// make Task complete by id
		case 3:
			fmt.Println("Enter Task id: ")
			var id int
			_, err := fmt.Scan(&id)
			Check(err)
			_, ok := mp[id]
			if ok {
				fmt.Println("Marking task", id, "as completed...")
				CompleteTask(id, mp)
			} else {
				fmt.Println("Invalid task ID")
			}
		}
	}

}
