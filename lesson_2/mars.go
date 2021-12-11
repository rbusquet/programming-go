package main

import "fmt"

func main() {
	// Print("Something", 2, "etc") also works
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")

	// Printf
	fmt.Printf("My weight on the surface of Mars is %v lbs", 150*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
