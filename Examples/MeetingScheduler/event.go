package MeetingScheduler

import (
	"errors"
	"fmt"
)

type UserId string
type EventType string
type InviteStatus string

type Event interface {
	GetDetails() string
	AddPeople([]string) // sends event invite notification to the people
	RemovePeople([]string)
	ModifyTiming(int, int) error // Set the start and end of meeting time, return an error if it overlaps with some existing event for the owner
	Delete() error
}

type EventObject struct {
	ID           string
	Title        string
	Participants []string
	StartTime    int // ideally would be in millis
	EndTime      int
	Description  string
	Type         EventType
	Owner        string
	Status       string
}

const (
	INVITE_ONLY  = EventType("INVITE_ONLY")
	OPEN_FOR_ALL = EventType("OPEN_FOR_ALL")

	ACCEPT  = InviteStatus("ACCEPT")
	REJECT  = InviteStatus("REJECT")
	PENDING = InviteStatus("PENDING")
)

type User struct {
	Id          UserId
	Name        string
	Designation string
}

type UserEvent struct {
	UserId  string
	EventId string
	Status  string
}

func (event *EventObject) GetDetails() string {
	return fmt.Sprintf("Event with title %v is scheduled from %v to %v. It is going to be a %v event with the following participants %v. The event is owned by %v", event.Title, event.StartTime, event.EndTime, event.Type, event.Participants, event.Owner)
}

func (event *EventObject) AddPeople(userIds []string) {
	if event.Type == INVITE_ONLY {
		for _, user := range userIds {
			present := false
			for _, existingPeople := range event.Participants {
				if user == existingPeople {
					present = true
				}
			}
			if !present {
				event.Participants = append(event.Participants, user)
			}
		}
	}
}

func (event *EventObject) RemovePeople(userIds []string) {

	newParticipantList := []string{}
	if event.Type == INVITE_ONLY {
		for _, existing := range event.Participants {
			remove := false
			for _, user := range userIds {
				if existing == user {
					remove = true
				}
			}
			if !remove {
				newParticipantList = append(newParticipantList, existing)
			}
		}
		event.Participants = newParticipantList
	}
}

func (event *EventObject) ModifyTiming(startTime, endTime int) error {
	// ideally would check for overlapping events and return error accordingly
	if startTime >= endTime {
		return errors.New("Cannot have greater startime than endtime")
	}
	event.StartTime = startTime
	event.EndTime = endTime
	return nil
}

func (event *EventObject) Delete() error {
	if event.Status == "DELETED" {
		return errors.New("Event doesn't exist")
	}
	event.Status = "DELETED"
	return nil
}

