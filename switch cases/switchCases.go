package main

import "fmt"

// a conditional statement that evaluates an expression and compares it against a list of possible matches and executes the corresponding block of code
func main() {
	month := "December"
	switch month {
	case "January":
		fmt.Println("It's January")
	case "February":
		fmt.Println("It's February")
	case "March":
		fmt.Println("It's March")
	case "April":
		fmt.Println("It's April")
	case "May":
		fmt.Println("It's May")
	case "June":
		fmt.Println("It's June")
	case "July":
		fmt.Println("It's July")
	case "August":
		fmt.Println("It's August")
	case "September":
		fmt.Println("It's September")
	case "October":
		fmt.Println("It's October")
	case "November":
		fmt.Println("It's November")
	case "December":
		fmt.Println("It's December")
	default:
		fmt.Println("Probably a month from the Ethiopian or Chinese calender")
	}
	// Multiple expressions in case
	day := "Tuesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	}
	// Expressionless switch
	num := 120
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}
}
