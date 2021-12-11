package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	const minDistance, maxDistance = 56_000_000, 401_000_000
	var randomDistance = rand.Intn(maxDistance - minDistance)
	fmt.Println(minDistance + randomDistance)
}
