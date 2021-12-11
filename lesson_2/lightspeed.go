package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	const sxtsSpeed = 100_800                  // km/h
	distance = 96_300_000                      // km
	const hoursPerDay, minutesPerHour = 24, 60 // h
	fmt.Println(distance/sxtsSpeed/hoursPerDay, " days")

	var weight = 100
	weight -= 2
}
