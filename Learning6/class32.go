package main

import (
	"fmt"
	"time"
)

func class32() {
	now := time.Now()
	fmt.Println("Current time: ", now)
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month:",now.Month())
	fmt.Println("Day",now.Day())
	fmt.Println("Hour: ", now.Hour())
	fmt.Println("Minute",now.Minute())

	fmt.Println("Formatted: ", now.Format("2006-01-02"))
	fmt.Println("Formatted:",now.Format("02-Jan-2006"))
	start:= time.Now()
	time.Sleep(2*time.Second)
	duration:=time.Since(start)
	fmt.Println("Duration",duration)
}