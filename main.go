package main

import (
  "fmt"
)

type Task struct {
    name string
    completed bool
}

type TodoList struct {
    tasks []Task
}

func (t *TodoList) AddTask(name string) {
    t.tasks = append(t.tasks, Task{name: name, completed: false})
}

func (t *TodoList) CompleteTask(name string) {
    for i, task := range t.tasks {
        if task.name == name {
            t.tasks[i].completed = true
        }
    }
}

func (t *TodoList) RemoveTask(name string) {
    for i, task := range t.tasks {
        if task.name == name {
            t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
        }
    }
}

func (t *TodoList) ListTasks() {
    for _, task := range t.tasks {
        fmt.Printf("%s - %t\n", task.name, task.completed)
    }
}

func main() {
    todo := TodoList{}
    todo.AddTask("Take out trash")
    todo.AddTask("Wash dishes")
    todo.AddTask("Do laundry")
    todo.CompleteTask("Wash dishes")
    todo.RemoveTask("Take out trash")
    todo.ListTasks()
}
