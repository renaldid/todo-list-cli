package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== To-Do List CLI ===")

	// Membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	// Buatlah slice untuk menyimpan tugas
	tasks := make([]string, 0)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add a task")
		fmt.Println("2. Show task list")
		fmt.Println("3. Mark the task as complete")
		fmt.Println("4. Delete task")
		fmt.Println("0. Exit")

		fmt.Print("Select a menu: ")
		option, _ := reader.ReadString('\n')

		switch option {
		case "1\n":
			fmt.Print("Enter a task: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, task)
			fmt.Println("Task added successfully.")

		case "2\n":
			if len(tasks) == 0 {
				fmt.Println("There are no tasks in the list.")
			} else {
				fmt.Println("Task List:")
				for i, task := range tasks {
					fmt.Printf("%d. %s", i+1, task)
				}
			}

		case "3\n":
			if len(tasks) == 0 {
				fmt.Println("There are no tasks in the list.")
			} else {
				fmt.Print("Select the completed task number: ")
				var num int
				fmt.Scanln(&num)
				if num >= 1 && num <= len(tasks) {
					tasks = append(tasks[:num-1], tasks[num:]...)
					fmt.Println("The task is successfully marked as complete.")
				} else {
					fmt.Println("Invalid assignment number.")
				}
			}

		case "4\n":
			if len(tasks) == 0 {
				fmt.Println("There are no tasks in the list.")
			} else {
				fmt.Print("Select the task number you want to delete: ")
				var num int
				fmt.Scanln(&num)
				if num >= 1 && num <= len(tasks) {
					tasks = append(tasks[:num-1], tasks[num:]...)
					fmt.Println("Task deleted successfully.")
				} else {
					fmt.Println("Invalid assignment number.")
				}
			}

		case "0\n":
			fmt.Println("Thank you for using To-Do List CLI.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
