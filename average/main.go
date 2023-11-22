// Package average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/ms77grz/hfg/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := int64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/float64(sampleCount))
}
