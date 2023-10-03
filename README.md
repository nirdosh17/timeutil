# Utility functions for Time
Ruby on Rails inspired utility functions for **Go** to manipulate Time.

## Install
`go get github.com/nirdosh17/timeutil`

## Usage
```go
  newT := timeutil.BeginningOfDay(time.Now())
  fmt.Println(newT)
```

## Available Functions
- BeginningOfDay(Time) Time
- BeginningOfHour(Time) Time
- BeginningOfMinute(Time) Time

- EndOfDay(Time) Time
- EndOfHour(Time) Time
- EndOfMinute(Time) Time

- MonthsAgo(Time, days)
- MonthsAfter(Time, days)

- DaysAgo(Time, days)
- DaysAfter(Time, days)
