package main

import (
	"fmt"
	"time"
)

func ConvertMonth(month int) string {
	if month == 1 {
		return "January"
	} else if month == 2 {
		return "February"
	} else if month == 3 {
		return "March"
	} else if month == 4 {
		return "April"
	} else if month == 5 {
		return "May"
	} else if month == 6 {
		return "June"
	} else if month == 7 {
		return "July"
	} else if month == 8 {
		return "August"
	} else if month == 9 {
		return "September"
	} else if month == 10 {
		return "October"
	} else if month == 11 {
		return "November"
	} else if month == 12 {
		return "December"
	} else {
		return "null"
	}
}

func actualMonth() {
	fmt.Println(ConvertMonth(int(time.Now().Month())))
}

func main() {
	actualMonth()
}
