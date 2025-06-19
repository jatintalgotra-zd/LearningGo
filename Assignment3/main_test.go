package main

import (
	"slices"
	"testing"
)

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
