package main

import "fmt"

func main() {

	userAge := map[string]int{
		"nayeem": 25,
		"tamim":  26,
	}

	if age, ok := userAge["nayeem"]; ok {
		fmt.Println("Nayeem's age is ", age)
	} else {
		fmt.Println("Nayeem not found")
	}
}
