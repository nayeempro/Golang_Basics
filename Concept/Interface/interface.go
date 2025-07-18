// ========================================
// GO INTERFACES COMPLETE GUIDE
// ========================================

// -----------------------------
// INTERFACES - BASIC DEFINITION
// -----------------------------
/*
In Go, an **interface** is a type that specifies a method set. If a type provides definitions for all the methods in the interface, it is said to "implement" the interface.
Interfaces allow you to write flexible and modular code.
They’re similar to contracts: if a type promises to do certain things (i.e., have certain methods), it can be used wherever that interface is expected.
*/

// -----------------------------
// REAL-WORLD EXAMPLE
// -----------------------------
/*
Example: Think of interfaces like a device charger port.
If a device supports the USB-C interface, any USB-C charger (implementation) can plug into it.
This makes code more flexible and swappable.
*/

// -----------------------------
// INTERFACE TYPE (DEFINITION & EXAMPLE)
// -----------------------------
package main

import "fmt"

// Define an interface
// Shape is an interface that requires Area method
// Any type that has Area() float64 method implements Shape

// BASIC INTERFACE EXAMPLE

type Shape interface {
	Area() float64
}

// Define a struct type
// Circle implements Shape interface because it has Area() method

type Circle struct {
	Radius float64
}

// Implement Area() for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func basicInterfaceExample() {
	var s Shape                    // s is a variable of interface type Shape
	s = Circle{Radius: 5.0}        // Assign Circle instance to interface variable
	fmt.Println("Area:", s.Area()) // Call interface method
}

// Output:
// Area: 78.5

// Explanation:
// - Shape interface defines a contract that any type with an Area method returns float64.
// - Circle struct satisfies the contract.
// - We can assign Circle to Shape because it implements Area().

// -----------------------------
// INTERMEDIATE INTERFACE EXAMPLE
// -----------------------------

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func interfaceWithMultipleTypes() {
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
	}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}

// Output:
// Area: 28.26
// Area: 20

// -----------------------------
// ADVANCED INTERFACE EXAMPLE
// -----------------------------
// Interface composition: combining multiple interfaces

type Printer interface {
	Print()
}

type Writer interface {
	Write() string
}

type PrintableWriter interface {
	Printer
	Writer
}

type Document struct {
	Content string
}

func (d Document) Print() {
	fmt.Println("Printing:", d.Content)
}

func (d Document) Write() string {
	return "Written: " + d.Content
}

func advancedInterfaceExample(p PrintableWriter) {
	p.Print()
	fmt.Println(p.Write())
}

// Output:
// Printing: My report
// Written: My report

// -----------------------------
// EMPTY INTERFACE - interface{}
// -----------------------------
/*
The empty interface `interface{}` can hold **any value** because every type implements at least 0 methods.
Useful for generic containers or when you don’t know the type in advance.
*/

func emptyInterfaceExample() {
	var x interface{}
	x = 42
	fmt.Println(x)
	x = "hello"
	fmt.Println(x)
	x = true
	fmt.Println(x)
}

// Output:
// 42
// hello
// true

// -----------------------------
// TYPE ASSERTION
// -----------------------------
/*
Used to retrieve the underlying value from an interface.
You must assert the type to access specific fields or methods.
*/

func typeAssertionExample(val interface{}) {
	str, ok := val.(string) // Try to convert interface{} to string
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Not a string")
	}
}

// Output for typeAssertionExample("Hi"):
// String value: Hi

// -----------------------------
// TYPE SWITCH
// -----------------------------
/*
Used to switch on the dynamic type of an interface value.
*/

func typeSwitchExample(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int value:", v)
	case string:
		fmt.Println("string value:", v)
	case bool:
		fmt.Println("bool value:", v)
	default:
		fmt.Println("unknown type")
	}
}

// -----------------------------
// COMMON METHODS & USES
// -----------------------------
/*
- Stringer: Allows custom string output via fmt.Stringer interface
- error: A built-in interface for handling errors
*/

// fmt.Stringer Example

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

func stringerExample() {
	p := Person{"Nayeem", 25}
	fmt.Println(p) // Will use String() automatically
}

// -----------------------------
// MAIN FUNCTION TO RUN ALL
// -----------------------------
func main() {
	fmt.Println("Basic Interface Example:")
	basicInterfaceExample()

	fmt.Println("\nInterface with Multiple Types:")
	interfaceWithMultipleTypes()

	fmt.Println("\nAdvanced Interface Example:")
	doc := Document{"My report"}
	advancedInterfaceExample(doc)

	fmt.Println("\nEmpty Interface Example:")
	emptyInterfaceExample()

	fmt.Println("\nType Assertion Example:")
	typeAssertionExample("Hi")
	typeAssertionExample(123)

	fmt.Println("\nType Switch Example:")
	typeSwitchExample("hello")
	typeSwitchExample(42)
	typeSwitchExample(false)

	fmt.Println("\nStringer Interface Example:")
	stringerExample()
}
