package all

// Problem: 1185
// Title: Day of the Week
// Category: all
// Tags: all


import "time"

func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}
