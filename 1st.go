package main

import "fmt"

func main() {
	var name string = "Nayeem"
	var age int = 24
	var married bool // (Zero value false)
	// By default false
	var say = "Alhamdulillah"

	var a, b int = 23, 24
	var c, d = "go", true
	var (
		e, f    = "me", 23
		newName = "Md"
		isAdmin bool
	)

	fmt.Println(name, age, married, say, a, b, c, d, e, f, newName, isAdmin)
}
