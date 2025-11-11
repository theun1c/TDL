package main

import (
	"todo/scanner"
	"todo/todo"
)

func main() {

	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Init()

}
