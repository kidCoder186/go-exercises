package meetup

import (
	"time"
)

// WeekSchedule type
type WeekSchedule int

// WeekSchedule enum
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Date structure
type Date struct {
	day   int
	month int
	year  int
}

// Weekday returns the day of a week for a date
func (d Date) Weekday() time.Weekday {
	m := d.month - 2
	D := d.year % 100
	if d.month < 3 {
		D = (D + 99) % 100
		m = m + 12
	}
	C := d.year / 100
	f := d.day + (13*m-1)/5 + D + D/4 + C/4 - 2*C
	return time.Weekday(f % 7)
}

//DaysOfMonth returns the number days of the month
func (d Date) DaysOfMonth() int {
	switch d.month {
	case 2:
		if isLeapYear(d.year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func isLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

// Day calculates the date of meetup.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	firstDay := Date{1, int(month), year}
	firstWeekday := int(weekday-firstDay.Weekday()+7)%7 + 1

	switch week {
	case First:
		return firstWeekday
	case Second:
		return firstWeekday + 7
	case Third:
		return firstWeekday + 14
	case Fourth:
		return firstWeekday + 21
	case Teenth:
		for firstWeekday <= 12 {
			firstWeekday += 7
		}
		return firstWeekday
	case Last:
		for firstWeekday+7 <= firstDay.DaysOfMonth() {
			firstWeekday += 7
		}
		return firstWeekday
	default:
		return 0
	}
}
