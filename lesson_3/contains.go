package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	// Emerging from the cave, your eyes meet the blinding midday sun.
	// How do you declare a Boolean variable named wearShades?
	// var wearShades = true
	// There is a sign near the cave entrance.
	// How can you determine if the command contains the word “read”?
	command = "read sign"
	fmt.Println(strings.Contains(command, "read"))
}
