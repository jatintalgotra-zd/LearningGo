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
func AddTask(description string, nextID func() int, tasks *[]Task) int {
	id := nextID()
	t1 := Task{Desc: description}
	*tasks = append(*tasks, t1)
	fmt.Println("Task added:", id, "-", description)
	return id
}

// ListTasks
// prints all pending tasks
func ListTasks(tasks *[]Task) []string {
	taskList := make([]string, 0)
	for k, v := range *tasks {
		if !v.Status {
			t := fmt.Sprintln(k+1, ". ", v.Desc)
			taskList = append(taskList, t)
		}
	}
	return taskList
}

// CompleteTask
// marks Task complete by id
func CompleteTask(id int, tasks *[]Task) {
	idx := id - 1
	if idx >= 0 && idx < len(*tasks) {
		fmt.Println("Marking task", id, "as completed...")
		(*tasks)[idx].Status = true

	} else {
		fmt.Println("Invalid task ID")
	}

}

func main() {

	// map used to store tasks
	tasks := make([]Task, 0)
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

		// switch case to call functions
		switch n {

		// enter new Task
		case 1:
			fmt.Println("Enter Task: ")
			var t1 string
			bufReader := bufio.NewReader(os.Stdin)
			t1, err2 := bufReader.ReadString('\n')
			Check(err2)
			AddTask(t1, nextId, &tasks)

		// list pending tasks
		case 2:
			fmt.Println("Pending Task: ")
			pendingTasks := ListTasks(&tasks)
			for id, task := range pendingTasks {
				fmt.Println(id, task)
			}

		// make Task complete by id
		case 3:
			fmt.Println("Enter Task id: ")
			var id int
			_, err3 := fmt.Scan(&id)
			Check(err3)
			CompleteTask(id, &tasks)
		}
	}

}
