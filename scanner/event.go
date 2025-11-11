package scanner

import "time"

type Event struct {
	Descriprion string
	UserInput   string
	EventDate   time.Time
}

func NewEvent(descriprion, userInput string) Event {
	return Event{
		Descriprion: descriprion,
		UserInput:   userInput,
		EventDate:   time.Now(),
	}
}
