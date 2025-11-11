package scanner

import (
	"fmt"
	"todo/todo"

	"github.com/k0kubun/pp"
)

func printPrompt() {
	fmt.Println("Enter the command: ")
}

func printExit() {
	fmt.Println("Exit prog ... ")
}

func printSuccessAdd(title string) {
	fmt.Println("task ", title, " was successfully added")
	fmt.Println()
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Tasks list: ", tasks)
	fmt.Println()
}

func printDone(title string) {
	fmt.Println("Task ", title, " is done")
	fmt.Println()
}

func printDelete(title string) {
	fmt.Println("Task ", title, " is deleted")
	fmt.Println()
}

func printHelp() {
	fmt.Println("Commands:\nadd {title} {text ...}\ndelete {title}\nlist\ndone {title}\nexit")
	fmt.Println()
}

func printEvents(events []Event) {
	pp.Println("Events: ", events)
	fmt.Println()
}

func printResult(result string) {
	fmt.Println("Result after enter command: ", result)
	fmt.Println()
}

