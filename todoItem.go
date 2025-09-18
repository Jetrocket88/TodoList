package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	id           int
	title        string
	desc         string
	importance   int
	optional     bool
	dateAssigned string
	dateDue      string
}

func NewItem(ID int, Title string, Desc string, DateAssigned string, DateDue string, Importance int, Optional bool) *Item {
	return &Item{id: ID, title: Title, desc: Desc, importance: Importance, optional: Optional, dateAssigned: DateAssigned, dateDue: DateDue}
}

var IDList []int

func parseTextToItem(itemList []Item, input string) []Item {
	dataArray := strings.Split(strings.TrimSpace(input), "\n")
	for i := range dataArray {
		dataArray[i] = strings.TrimSpace(dataArray[i])
		row := strings.Split(dataArray[i], ",")
		row5, err := strconv.Atoi(row[5])
		if err != nil {
			fmt.Println("Error: ", err)
			panic("Problem with strconvAtoi: ")
		}
		row0, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println("Error: ", err)
			panic("Problem with strconvAtoi: ")
		}
		IDList = append(IDList, row0)
		row6, err := strconv.ParseBool(row[6])
		if err != nil {
			fmt.Println("Error: ", err)
			panic("Problem with strconv ParseBool: ")
		}
		itemList = append(itemList, *NewItem(row0, row[1], row[2], row[3], row[4], row5, row6))
	}
	return itemList
}

func convertListToString(items []Item) string {
	outputString := ""

	for _, value := range items {
		outputString = outputString + strconv.Itoa(value.id) + "," + value.title + "," + value.desc + "," + value.dateAssigned + "," + value.dateDue + "," +
			strconv.Itoa(value.importance) + "," + strconv.FormatBool(value.optional) + "\n"
	}
	return outputString
}

func updateFile(fileName string, items []Item) {
	os.Remove(fileName)
	os.Create(fileName)
	data := []byte(convertListToString(items))
	os.WriteFile(fileName, data, 0644)
}

func readDataFromFile(itemList []Item, filePath string) []Item {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	return parseTextToItem(itemList, dataString)
}

func genValidID() int {
	return 1 + IDList[len(IDList)-1]
}

func addItem(items []Item) []Item {
	fmt.Println("Please add the following data separated by commmas")
	fmt.Println("(title),(description),(date assigned),(date due),(importance),(optional 1/0)")

	var userInput string
	fmt.Scanln(&userInput)
	if userInput == "" {
		fmt.Scanln(&userInput)
	}

	userInput = strconv.Itoa(genValidID()) + "," + userInput
	items = parseTextToItem(items, userInput)

	fmt.Println("Successfully added item, here are the values:")
	if len(items) > 0 {
		fmt.Println(items[len(items)-1])
	} else {
		fmt.Println("No items were added (this shouldn't happen)")
	}

	updateFile("items.txt", items)
	return items
}

func printTasks(items []Item) {
	for _, value := range items {
		fmt.Printf("%d, %s, %s, %s, %s, %d, %t\n", value.id, value.title, value.desc, value.dateAssigned, value.dateDue, value.importance, value.optional)
	}
}

func appendRemove(items []Item, removal int) []Item {
	var temp []Item
	for _, value := range items {
		if value.id != removal {
			temp = append(temp, value)
		}
	}
	return temp
}

func linearSearchRemoval(items []Item, index int) []Item {
	return appendRemove(items, index)
}

func linearSearch(items []Item, targetID int) int {
	for index, value := range items {
		if value.id == targetID {
			return index
		}
	}
	return -1
}

func removeItem(items []Item) []Item {
	var lower = 0
	var upper = 10
	if len(items) < upper {
		upper = len(items) - 1
	}

	var userInput = " "
	fmt.Println("Enter the ID of the task you would like to remove")
	fmt.Println("Showing tasks ", lower, ": ", upper)
	printTasks(items[lower:upper])
	fmt.Println("Enter + to see more tasks or enter the id of the task you want to remove")
	fmt.Scan(&userInput)
	for userInput == "+" {
		userInput = " "
		lower += 10
		upper += 10
		if len(items) < upper {
			upper = len(items) - 1
		}
		if upper > lower {
			fmt.Println("Showing tasks ", lower, ": ", upper)
			printTasks(items[lower:upper])
			fmt.Println("Enter + to see more tasks or enter the id of the task you want to remove")
		} else {
			fmt.Println("Thats all the items printed")
		}
		fmt.Scan(&userInput)
	}
	idToRemove, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println(err)
		panic("idk")
	}
	fmt.Println("Row to remove:")
	indexToRemove := linearSearch(items, idToRemove)
	fmt.Println(items[indexToRemove])
	items = linearSearchRemoval(items, idToRemove)
	fmt.Println("Row removed")
	updateFile("items.txt", items)
	return items
}
