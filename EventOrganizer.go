package main

import (
	"fmt"
	"time"
)

// CustomDate represents a custom date format structure
type CustomDate struct {
	time.Time // Embedding time.Time within CustomDate
}

type OrganisingTeam struct {
	teamMemebers   []string
	primaryContact string
}

type NSPEvent struct {
	eventName      string
	primaryContact string
	EventDate      CustomDate
	organisingTeam OrganisingTeam
}

// NewCustomDate creates a new CustomDate with the given year, month, and day
func NewCustomDate(year, month, day int) CustomDate {
	return CustomDate{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)}
}

// FormatDate formats the date according to the provided layout
func (cd CustomDate) FormatDate(layout string) string {
	return cd.Time.Format(layout)
}

func main() {

	eventDate := NewCustomDate(2024, 1, 31)

	nspEvent := NSPEvent{eventName: "Nokia Hackthon", primaryContact: "8010*******",
		EventDate:      eventDate,
		organisingTeam: OrganisingTeam{[]string{"A", "B"}, "C"}}

	fmt.Println("struct : ", nspEvent)
}
