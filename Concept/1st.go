package main

import "fmt"

var x = 5

func main() {
	// var name string = "Nayeem"
	// var age int = 24
	// var married bool // (Zero value false)
	// // By default false
	// var say = "Alhamdulillah"

	// var a, b int = 23, 24
	// var c, d = "go", true
	// var (
	// 	e, f    = "me", 23
	// 	newName = "Md"
	// 	isAdmin bool
	// )

	// const (
	// 	A = iota + 1
	// 	B
	// 	C
	// 	D
	// 	E
	// )
	// m, n, z := "nayeem", 25, true //Initialize and assign both
	// fmt.Println(name, age, married, say, a, b, c, d, e, f, newName, isAdmin, m, n, z)
	// fmt.Println("const value: ", A, B, C, D, E)

	//=============================Variable shadowing=============================

	// x := 19 // variable shadow  --> global x=5 (print x=19)
	// {
	// 	x := 20 // shadows outer x
	// 	fmt.Println("Inner block x:", x)
	// }
	// fmt.Println("Outer block x:", x)

	// err := fmt.Errorf("outer error")
	// if true {
	// 	err := fmt.Errorf("inner error") // shadows outer err
	// 	fmt.Println("Inner:", err)
	// }
	// fmt.Println("Outer:", err) // still "outer error"

}
