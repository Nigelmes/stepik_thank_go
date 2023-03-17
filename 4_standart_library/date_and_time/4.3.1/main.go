package main

import (
	"fmt"
	"time"
)

func isLeapYear(year int) bool {
	t := time.Date(year, time.December, 31, 23, 0, 0, 0, time.UTC)
	return t.YearDay() == 366
}

func main() {
	fmt.Println(isLeapYear(2020))

}
