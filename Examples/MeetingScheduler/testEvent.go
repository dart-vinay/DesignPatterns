package MeetingScheduler

import "fmt"

func TestEvent() {
	// Get details of the event
	// Get calendar for the user
	// Accept/Reject event
	// Create an event
	// For a given set of users identify a common free slot of time

	event := new(EventObject)
	event.Title = "Meeting"
	event.ID = "3"
	event.Type = INVITE_ONLY
	event.AddPeople([]string{"3", "4"})
	event.RemovePeople([]string{"1"})

	fmt.Print(event.GetDetails())
}
