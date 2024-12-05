package main

import (
	"bufio"
	"fmt"
	"os"
)

// Task structure
type Task struct {
	name        string
	description string
	completed   bool
}

// Task list structure
type TaskList struct {
	tasks []Task
}

// Method to add tasks
func (l *TaskList) addTask(t Task) {
	l.tasks = append(l.tasks, t)
}

// Method to mark task as completed
func (l *TaskList) markCompleted(index int) {
	l.tasks[index].completed = true
}

// Method to edit a task
func (l *TaskList) editTask(index int, t Task) {
	l.tasks[index] = t
}

// Method to delete a task
func (l *TaskList) deleteTask(index int) {
	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
}

func main() {

	// Instance of task list
	list := TaskList{}

	// Infinite menu loop
	reader := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Select an option: \n",
			"1. Add task\n",
			"2. Mark task as completed\n",
			"3. Edit task\n",
			"4. Delete task\n",
			"5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t Task
			fmt.Print("Enter the task name: ")
			t.name, _ = reader.ReadString('\n')
			fmt.Print("Enter the task description: ")
			t.description, _ = reader.ReadString('\n')
			list.addTask(t)
			fmt.Println("Task added successfully")
		case 2:
			var index int
			fmt.Print("Enter the index of the task to mark as completed: ")
			fmt.Scanln(&index)
			list.markCompleted(index)
			fmt.Println("Task marked as completed successfully")
		case 3:
			var index int
			var t Task
			fmt.Print("Enter the index of the task to update: ")
			fmt.Scanln(&index)
			fmt.Print("Enter the task name: ")
			t.name, _ = reader.ReadString('\n')
			fmt.Print("Enter the task description: ")
			t.description, _ = reader.ReadString('\n')
			list.editTask(index, t)
			fmt.Println("Task updated successfully")
		case 4:
			var index int
			fmt.Print("Enter the index of the task to delete: ")
			fmt.Scanln(&index)
			list.deleteTask(index)
			fmt.Println("Task deleted successfully")
		case 5:
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid option.")
		}

		// List all tasks
		fmt.Println("Task list:")
		fmt.Println("============================================================")
		for i, t := range list.tasks {
			fmt.Printf("%d. %s - %s - Completed: %t \n", i, t.name, t.description, t.completed)
		}
		fmt.Println("============================================================")
	}
}
