package services

import (
	"time"

	"go.mau.fi/whatsmeow/types"
)

type Reminder struct {
	Message string
	To      types.JID
	Time 	time.Time

}

var reminderList = make(map[time.Time][]Reminder)


func AddReminder(reminder Reminder){
	
	reminderList[reminder.Time] = append(reminderList[reminder.Time], reminder)
}

func GetReminders(time time.Time)( []Reminder){
	remindersAtTime:= reminderList[time]
	return remindersAtTime
}

// func createReminder(message string,To types.JID)(Reminder){
// 	// parse message string to get time message to set reminder
// }


