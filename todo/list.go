package todo

import (
	"fmt"
	"time"
)

type List struct {
	tasks map[string]Task
}

// с точки зрения принципа разделения ответственности
// лучше всего в методы листа передавать уже созданную из вне задачу
// вместо передачи параметров текста и заголовка внутри методов
// AddTask(task) <-- better when --> AddTask(title, text ...)

// ресивер должен быть по указателю
// иначе мы значение будем менять у копии
// а не у изначального объекта !
func (l *List) AddTast(task Task) {
	// стоит выбор сохранения таск
	// 1 - слайс
	// 2 - мапа
	//
	// выбор: мапа (ключ - значение)
	// обоснование выбора:
	// удобнее всего, тк удаление/получение будет реализовано
	// через заголовок задачи (ключ)
	// а создание задачи будет через 2 параметра
	// название (ключ) - текст (значение)
	//
	// слайсом мы бы прогоняли каждый раз по циклу - неэффективно
	//
	// также в структуре листа мы бы удобно хранили значение таскс
	// мапа где ключ - строка (тайтл) а значение - объект таск

	// если такая задача существует "по названию"
	// то создаем новую с именет "тайтл + дата сейчас"
	// иначе просто создаем задачу, добавляя таск в мапу
	_, ok := l.tasks[task.Title]

	if ok {
		new_key := fmt.Sprintf("%v %v", task.Title, time.Now())
		l.tasks[new_key] = task
	} else {
		l.tasks[task.Title] = task
	}
}

func (l *List) DeleteTask(title string) string {

	_, ok := l.tasks[title]
	if !ok {
		return taskNotFound
	}

	delete(l.tasks, title)

	return ""
}

func (l *List) DoneTask(title string) string {
	task, ok := l.tasks[title]
	if !ok {
		return taskNotFound
	}

	// можео сделать и так, но следует придерживаться
	// принципа единой ответственности и выделить логику отметки
	// в отдельную функцию структуры таск
	//
	// task.IsDone = true
	// now := time.Now()
	// task.CompletedAt = &now

	task.Complete()

	l.tasks[title] = task

	return ""
}

// нужно обратить внимание!
// возвращать так мапу чревато какими-либо последствиями
//
// нужно нагуглить проблему
func (l *List) GetTask() map[string]Task {
	return l.tasks
}

