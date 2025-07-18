package main

import (
	"fmt"
	// "math"
	"strings"
)

type CustomText string

func (ct CustomText) Shout() string {
	return strings.ToUpper(string(ct)) + "!"
}
func logMessage(prefix string, messages ...string) {
	for _, msg := range messages {
		fmt.Println(prefix + ": " + msg)
	}
}

func main() {
	// var deg float64 = 45

	// radiant := deg * (math.Pi) / 180

	// sinValue := math.Sin(radiant)

	// fmt.Println("sin 45 value is", sinValue)

	// text := "golang is awesome"
	// upper := strings.ToUpper(text)
	// replaced := strings.ReplaceAll(upper, "AWESOME", "POWERFUL")
	// fmt.Println("Modified:", replaced)

	// var msg CustomText = "golang is cool"
	// fmt.Println(msg.Shout())

	// Goto

	// 	i := 0

	// loop:
	// 	if i >= 5 {
	// 		print("5 or more")
	// 		if i > 8 {
	// 			return
	// 		}
	// 	}
	// 	fmt.Println("i =", i)
	// 	i++
	// 	goto loop // Jump to loop label

	logMessage("INFO", "Server started", "Listening on port 8080")

}
