package scanner

import (
	"bufio"
	"os"
	"strings"
	"todo/todo"
)

type Scanner struct {
	todoList *todo.List
	events   []Event
}

func NewScanner(todoList *todo.List) *Scanner {
	return &Scanner{
		todoList: todoList,
	}
}

func (s *Scanner) Init() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()

		ok := scanner.Scan()

		if !ok {
			return
		}

		input := scanner.Text()

		result := s.stringHandler(input)
		if result != emptyString {
			if result == exitCall {
				printExit()
				return
			}

			printResult(result)
		}

		event := NewEvent(result, input)
		s.events = append(s.events, event)
	}
}

func (s *Scanner) stringHandler(str string) string {
	fields := strings.Fields(str)

	if len(fields) == 0 {
		return emptyString
	}

	cmd := fields[0]

	if cmd == "exit" {
		return exitCall
	}

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "done" {
		return s.cmdDone(fields)
	}

	if cmd == "delete" {
		return s.cmdDelete(fields)
	}

	if cmd == "help" {
		return s.cmdHelp(fields)
	}

	if cmd == "events" {
		return s.cmdEvents(fields)
	}

	return unknownCommand
}

func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return argumentsError
	}

	title := fields[1]
	text := ""
	space := ""
	for i := 2; i < len(fields); i++ {
		text += space + fields[i]
		space = " "
	}

	task := todo.NewTask(title, text)

	s.todoList.AddTast(task)

	printSuccessAdd(title)

	return ""
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return argumentsError
	}

	tasks := s.todoList.GetTasks()

	printTasks(tasks)

	return ""
}

func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return argumentsError
	}

	title := fields[1]

	doneTaskResult := s.todoList.DoneTask(title)

	if doneTaskResult != "" {
		return doneTaskResult
	}

	printDone(title)

	return ""

}

func (s *Scanner) cmdDelete(fields []string) string {
	if len(fields) != 2 {
		return argumentsError
	}

	title := fields[1]

	deleteTaskResult := s.todoList.DeleteTask(title)

	if deleteTaskResult != "" {
		return deleteTaskResult
	}

	printDelete(title)

	return ""

}

func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return argumentsError
	}

	printHelp()

	return ""
}

func (s *Scanner) cmdEvents(fields []string) string {
	if len(fields) != 1 {
		return argumentsError
	}

	if len(s.events) == 0 {
		return "empty events"
	}

	printEvents(s.events)

	return ""
}
