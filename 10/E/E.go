package main

import "time"

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t *Task) IsOverdue() bool {
	return time.Now().After(t.deadline)
}

func (t *Task) IsTopPriority() bool {
	return t.priority >= 4
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (tdl *ToDoList) TasksCount() int {
	return len(tdl.tasks)
}

func (tdl *ToDoList) NotesCount() int {
	return len(tdl.notes)
}

func (tdl *ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, task := range tdl.tasks {
		if task.IsTopPriority() {
			count++
		}
	}
	return count
}

func (tdl *ToDoList) CountOverdueTasks() int {
	count := 0
	for _, task := range tdl.tasks {
		if task.IsOverdue() {
			count++
		}
	}
	return count
}
