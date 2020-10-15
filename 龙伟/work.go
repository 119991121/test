package main

import (
	"fmt"
	"time"
)

func main() {
	var year int
	var month int
	var day int
	now_year := int(time.Now().Year())
	now_month := int(time.Now().Month())
	now_day := int(time.Now().Day())
	array1 := [13]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	array2 := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	fmt.Println("What year were you born [YYYY]?")
	fmt.Println("What month were you born [MM]?")
	fmt.Println("What day were you born [DD]?")
	fmt.Scanln(&year, &month, &day)
	fmt.Printf("It looks like you were born on %d/%d/%d", day, month, year)

	if (now_year%4 == 0 && now_year%100 != 0) || now_year%400 == 0 {
		switch {
		case month == now_month && day == now_day:
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		case month == now_month && day < now_day:
			sum := now_day - day
			fmt.Println("It looks like your birthday has passed ", sum, "day.")
			fmt.Println("Hope you'are looking forward to the next one!")
		case month < now_month:
			middle := 0
			for i := month + 1; i < now_month; i++ {
				middle = middle + array1[i]
			}
			sum := array1[month] - day + middle + now_day
			fmt.Println("It looks like your birthday has passed ", sum, " day.")
			fmt.Println("Hope you'are looking forward to the next one!")
		case month == now_month && day > now_day:
			sum := day - now_day
			fmt.Println("It looks like there are", sum, " days from your birthday.")
			fmt.Println("Hope you're looking forward to it!")
		case month > now_month:
			middle := 0
			for i := now_month + 1; i < month; i++ {
				middle = middle + array1[i]
			}
			sum := array1[now_month] - now_day + middle + day
			fmt.Println("It looks like there are ", sum, "days from your birthday.")
			fmt.Println("Hope you're looking forward to it!")
		}
	} else {
		switch {
		case month == now_month && day == now_day:
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		case month == now_month && day < now_day:
			sum := now_day - day
			fmt.Println("It looks like your birthday has passed ", sum, "day.")
			fmt.Println("Hope you'are looking forward to the next one!")
		case month < now_month:
			middle := 0
			for i := month + 1; i < now_month; i++ {
				middle = middle + array2[i]
			}
			sum := array2[month] - day + middle + now_day
			fmt.Println("It looks like your birthday has passed ", sum, " day.")
			fmt.Println("Hope you'are looking forward to the next one!")
		case month == now_month && day > now_day:
			sum := day - now_day
			fmt.Println("It looks like there are", sum, " days from your birthday.")
			fmt.Println("Hope you're looking forward to it!")
		case month > now_month:
			middle := 0
			for i := now_month + 1; i < month; i++ {
				middle = middle + array2[i]
			}
			sum := array2[now_month] - now_day + middle + day
			fmt.Println("It looks like there are ", sum, "days from your birthday.")
			fmt.Println("Hope you're looking forward to it!")
		}

	}

}