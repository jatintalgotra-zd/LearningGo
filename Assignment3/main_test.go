package main

import (
	"slices"
	"testing"
)

// TestAddTask
// test function for AddTask - checks for length, Task description and Task status
func TestAddTask(t *testing.T) {
	tasks := make([]Task, 0)
	nextId := GenId()
	task := "new task"
	id := AddTask(task, nextId, &tasks)

	idx := id - 1
	if len(tasks) != 1 {
		t.Errorf("Expected: 1 task, got: %d tasks", len(tasks))
	}
	if tasks[idx].Desc != task {
		t.Errorf("Expected task: %s, got: %s", task, tasks[id].Desc)
	}
	if tasks[idx].Status != false {
		t.Errorf("Expected status: false, got: %t", tasks[id].Status)
	}
}

// TestCompleteTask
// test function for CompleteTask - checks for Task status
func TestCompleteTask(t *testing.T) {
	tasks := []Task{
		{Desc: "test task"},
	}
	idx := 0
	CompleteTask(idx+1, &tasks)

	if tasks[idx].Status != true {
		t.Errorf("Expected status: true, got: %t", tasks[idx].Status)
	}
}

// TestListTasks
// test function for ListTasks - checks using slice of strings, if the task list matches test case
func TestListTasks(t *testing.T) {
	tasks := []Task{
		{Desc: "test task 1"},
		{Desc: "test task 2", Status: true},
		{Desc: "test task 3"},
	}
	expected := []string{
		"1. test task 1\n",
		"3. test task 3\n",
	}

	output := ListTasks(&tasks)

	if slices.Equal(expected, output) {
		t.Errorf("Expected: %v, got: %v", expected, output)
	} else {
		t.Logf("Expected: %v, got: %v", expected, output)
	}

}
