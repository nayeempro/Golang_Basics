package main

import (
	"fmt"
	// "math"
	"strings"
)

func main() {
	// var deg float64 = 45

	// radiant := deg * (math.Pi) / 180

	// sinValue := math.Sin(radiant)

	// fmt.Println("sin 45 value is", sinValue)

	text := "golang is awesome"
	upper := strings.ToUpper(text)
	replaced := strings.ReplaceAll(upper, "AWESOME", "POWERFUL")
	fmt.Println("Modified:", replaced)
}
