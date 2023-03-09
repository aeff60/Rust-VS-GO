package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("America/New_York")
	start := time.Now().In(location)

	var sum float32
	for i := 1; i <= 100000; i++ {
		sum += float32(i)
	}
	// end := time.Now()
	// fmt.Println(end.Sub(start).Microseconds())

	duration := time.Since(start)
	fmt.Println("Sum:", sum)
	fmt.Println("Time elapsed:", duration.Microseconds())

	// fmt.Println("Time elapsed:", duration.String())
}
