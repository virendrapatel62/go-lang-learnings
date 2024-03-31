package models

import "time"

type Event struct {
	Id          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (event Event) Save() {
	// add it to a database
	events = append(events, event)
}
