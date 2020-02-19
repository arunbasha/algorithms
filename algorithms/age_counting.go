package Algorithm

import (
	"time"
)
//Age Counting
func ageCounting(bDate time.Time) int{
	curDate := time.Now()
	bDay := bDate.Day()
	curDay := curDate.Day()

	years := curDate.Year() - bDate.Year()
	if isLeapYear(bDate) && !isLeapYear(curDate) && bDay > 60 {
		bDay -= 1
	}
	if isLeapYear(curDate) && !isLeapYear(bDate) && curDay > 60 {
		bDay += 1
	}

	if curDay < bDay {
		years -= 1
	}
	return years
}

func isLeapYear( date time.Time) bool {
	year := date.Year()
	if year % 400 == 0 {
		return true
	} else if year % 100 == 0 {
		return false
	} else if  year % 4 == 0 {
		return true
	}
	return false
}
