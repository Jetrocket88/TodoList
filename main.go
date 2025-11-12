package main

import "fmt"

func main() {
	var input int
	fmt.Println("What would you like to do?")
	fmt.Println("All changes will be written to file")
	fmt.Println("(1) Add a task\n(2) Remove a task\n(3) View all tasks")
	fmt.Scan(&input)

	var items []Item
	items = readDataFromFile(items, "items.txt")

	switch input {
	case 1:
		items = addItem(items)
	case 2:
		items = removeItem(items)
	case 3:
		printTasks(items)
	}
}
