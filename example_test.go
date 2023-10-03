package timeutil

import (
	"fmt"
	"time"
)

func ExampleBeginningOfMonth() {
	dt, _ := time.Parse(time.RFC3339, "2023-01-14T12:45:53Z")
	result := BeginningOfMonth(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2023-01-01T00:00:00Z
}

func ExampleBeginningOfDay() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := BeginningOfDay(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-14T00:00:00Z
}

func ExampleBeginningOfHour() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := BeginningOfHour(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-14T12:00:00Z
}

func ExampleBeginningOfMinute() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := BeginningOfMinute(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-14T12:45:00Z
}

func ExampleEndOfMonth() {
	dt, _ := time.Parse(time.RFC3339, "2022-12-14T12:45:53Z")
	result := EndOfMonth(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2022-12-31T23:59:59Z
}

func ExampleEndOfDay() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := EndOfDay(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-14T23:59:59Z
}

func ExampleEndOfHour() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := EndOfHour(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-14T12:59:59Z
}

func ExampleEndOfMinute() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := EndOfMinute(dt).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-14T12:45:59Z
}

func ExampleMonthsAgo_fallingInSameYear() {
	dt, _ := time.Parse(time.RFC3339, "2023-04-04T12:45:53Z")
	result := MonthsAgo(dt, 2).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2023-02-01T00:00:00Z
}

func ExampleMonthsAgo_fallingInPrevYear() {
	dt, _ := time.Parse(time.RFC3339, "2023-01-14T12:45:53Z")
	result := MonthsAgo(dt, 2).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2022-11-01T00:00:00Z
}

func ExampleMonthsAfter_spillingToNextYear() {
	dt, _ := time.Parse(time.RFC3339, "2022-11-14T12:45:53Z")
	result := MonthsAfter(dt, 2).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2023-01-01T00:00:00Z
}

func ExampleMonthsAfter_fallingInSameYear() {
	dt, _ := time.Parse(time.RFC3339, "2022-11-14T12:45:53Z")
	result := MonthsAfter(dt, 1).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2022-12-01T00:00:00Z
}

func ExampleDaysAgo() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := DaysAgo(dt, 15).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-03-30T12:45:53Z
}

func ExampleDaysAfter() {
	dt, _ := time.Parse(time.RFC3339, "2021-04-14T12:45:53Z")
	result := DaysAfter(dt, 1).Format(time.RFC3339)
	fmt.Println(result)
	// Output: 2021-04-15T12:45:53Z
}
