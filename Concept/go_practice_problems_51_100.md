# Go Practice Problems (51–100)

## Problem 51: Check if Number is Prime
**Description**: Write a Go function to check whether a number is a prime number.
**Test Data**: `29`, `10`, `1`, `2`
**Expected Output**: `true`, `false`, `false`, `true`

## Problem 52: Factorial (Recursive)
**Description**: Write a recursive function to compute the factorial of a number.
**Test Data**: `5`, `0`, `1`
**Expected Output**: `120`, `1`, `1`

## Problem 53: Armstrong Number
**Description**: Check whether a number is an Armstrong number (sum of cubes of digits equals the number itself).
**Test Data**: `153`, `370`, `123`
**Expected Output**: `true`, `true`, `false`

## Problem 54: Fibonacci Series (First N Terms)
**Description**: Print the first `N` numbers in the Fibonacci sequence.
**Test Data**: `5`, `10`
**Expected Output**: `[0 1 1 2 3]`, `[0 1 1 2 3 5 8 13 21 34]`

## Problem 55: Reverse an Integer
**Description**: Reverse a given integer.
**Test Data**: `12345`, `9870`
**Expected Output**: `54321`, `789`

## Problem 56: Count Digits of a Number
**Description**: Count the number of digits in a given number.
**Test Data**: `12345`, `7`, `1000`
**Expected Output**: `5`, `1`, `4`

## Problem 57: Sum of Digits
**Description**: Find the sum of digits of a given number.
**Test Data**: `123`, `1001`
**Expected Output**: `6`, `2`

## Problem 58: Palindrome Number
**Description**: Check if a number is a palindrome.
**Test Data**: `121`, `1331`, `123`
**Expected Output**: `true`, `true`, `false`

## Problem 59: GCD (Greatest Common Divisor)
**Description**: Compute the GCD of two numbers.
**Test Data**: `48, 18`, `101, 103`
**Expected Output**: `6`, `1`

## Problem 60: LCM (Least Common Multiple)
**Description**: Compute the LCM of two numbers.
**Test Data**: `4, 5`, `10, 15`
**Expected Output**: `20`, `30`

## Problem 61: Swap Two Numbers (Without Temp)
**Description**: Swap two numbers without using a temporary variable.
**Test Data**: `a=3, b=5`
**Expected Output**: `a=5, b=3`

## Problem 62: Sum of Even Numbers in Array
**Description**: Sum all even numbers from an array.
**Test Data**: `[1,2,3,4,5,6]`
**Expected Output**: `12`

## Problem 63: Frequency of Elements in Array
**Description**: Count frequency of each element in an array.
**Test Data**: `[1,2,2,3,3,3]`
**Expected Output**: `{1:1, 2:2, 3:3}`

## Problem 64: Max and Min in Array
**Description**: Find the maximum and minimum values in an array.
**Test Data**: `[2,4,1,7,9]`
**Expected Output**: `Max: 9, Min: 1`

## Problem 65: Sum of Rows and Columns in 2D Array
**Description**: Compute the sum of each row and column in a 2D array.
**Test Data**: `[[1,2],[3,4]]`
**Expected Output**: `Row sums: [3, 7], Col sums: [4, 6]`

## Problem 66: Matrix Transpose
**Description**: Transpose a matrix.
**Test Data**: `[[1,2],[3,4]]`
**Expected Output**: `[[1,3],[2,4]]`

## Problem 67: Dot Product of Two Vectors
**Description**: Find the dot product of two vectors.
**Test Data**: `[1,2,3], [4,5,6]`
**Expected Output**: `32`

## Problem 68: Search Element in Array
**Description**: Search for an element in an array.
**Test Data**: `[1,2,3,4,5], 3`
**Expected Output**: `Found at index 2`

## Problem 69: Check Pangram
**Description**: Determine if a string is a pangram (contains all letters of alphabet).
**Test Data**: `"The quick brown fox jumps over a lazy dog"`
**Expected Output**: `true`

## Problem 70: Count Vowels and Consonants
**Description**: Count the number of vowels and consonants in a string.
**Test Data**: `"hello world"`
**Expected Output**: `Vowels: 3, Consonants: 7`

## Problem 71: Replace Character in String
**Description**: Replace all occurrences of a character in a string.
**Test Data**: `"banana", 'a', 'o'`
**Expected Output**: `"bonono"`

## Problem 72: Toggle Case of String
**Description**: Toggle the case of each character in a string.
**Test Data**: `"GoLang"`
**Expected Output**: `"gOlANG"`

## Problem 73: Validate Email Format
**Description**: Check whether a given string is a valid email format using regex.
**Test Data**: `"user@example.com"`
**Expected Output**: `true`

## Problem 74: Find Duplicates in Array
**Description**: Identify duplicates in an array.
**Test Data**: `[1,2,3,2,4,3,5]`
**Expected Output**: `[2, 3]`

## Problem 75: Count Words in a Sentence
**Description**: Count the number of words in a string.
**Test Data**: `"Go is simple and powerful"`
**Expected Output**: `5`

## Problem 76: Remove Whitespace from String
**Description**: Remove all whitespaces from a string.
**Test Data**: `" a b c "`
**Expected Output**: `"abc"`

## Problem 77: Compare Two Strings
**Description**: Compare two strings ignoring case.
**Test Data**: `"Go", "go"`
**Expected Output**: `true`

## Problem 78: Random Integer Generator
**Description**: Generate a random integer between 1 to 100.
**Test Data**: _none_
**Expected Output**: `e.g., 57`

## Problem 79: Slice Operations (Add, Remove)
**Description**: Demonstrate adding and removing elements in slices.
**Test Data**: `[1,2,3] → add 4 → remove index 1`
**Expected Output**: `[1,3,4]`

## Problem 80: Deep Copy of Slice
**Description**: Create a deep copy of a slice.
**Test Data**: `[1,2,3]`
**Expected Output**: `New slice with same data but independent`

## Problem 81: Nested Map Access
**Description**: Create and access values from a nested map.
**Test Data**: `map[string]map[string]int{"scores":{"math":90}}`
**Expected Output**: `90`

## Problem 82: Create Struct and Access Fields
**Description**: Define a struct and access its fields.
**Test Data**: `type Student struct {Name string; Age int}` → Student{"Ali", 20}`
**Expected Output**: `Ali 20`

## Problem 83: Struct with Method
**Description**: Create a method that works on a struct type.
**Test Data**: `type Circle struct {Radius float64}` → `Area()` method
**Expected Output**: `Area: 3.14 * r^2`

## Problem 84: Pointer Basics
**Description**: Demonstrate how pointers work in Go.
**Test Data**: `a := 5 → modify via pointer`
**Expected Output**: `a becomes modified`

## Problem 85: Pointer to Struct
**Description**: Use a pointer to a struct and modify its values.
**Test Data**: `*Student → change Age`
**Expected Output**: `Age updated`

## Problem 86: Array of Structs
**Description**: Store multiple structs in an array/slice.
**Test Data**: `[]Student{{"Ali", 20}, {"Sara", 22}}`
**Expected Output**: List of students

## Problem 87: Interface Implementation
**Description**: Define an interface and implement it with a struct.
**Test Data**: `type Shape interface {Area()}` → `Circle` implements
**Expected Output**: `Area calculated`

## Problem 88: Sorting Array
**Description**: Sort an array in ascending order.
**Test Data**: `[4,2,5,1]`
**Expected Output**: `[1,2,4,5]`

## Problem 89: Sorting Structs by Field
**Description**: Sort a slice of structs by one of the fields.
**Test Data**: `[]Student{{"Ali", 20}, {"Sara", 18}}`
**Expected Output**: `Sorted by age`

## Problem 90: Use of Defer Statement
**Description**: Show how defer works.
**Test Data**: `defer fmt.Println("A")` + `fmt.Println("B")`
**Expected Output**: `B A`

## Problem 91: Panic and Recover
**Description**: Handle panic using recover.
**Test Data**: Panic manually and recover
**Expected Output**: `"Recovered from panic"`

## Problem 92: File Read
**Description**: Read contents of a file.
**Test Data**: File `"data.txt"`
**Expected Output**: File contents

## Problem 93: File Write
**Description**: Write content to a file.
**Test Data**: `"output.txt", "Hello World"`
**Expected Output**: File created

## Problem 94: JSON Marshal/Unmarshal
**Description**: Convert struct to JSON and vice versa.
**Test Data**: Student struct → JSON → struct
**Expected Output**: Success

## Problem 95: Time Formatting
**Description**: Format current time in specific format.
**Test Data**: `"2006-01-02 15:04:05"`
**Expected Output**: Current date/time

## Problem 96: Goroutine Basics
**Description**: Run a function in a goroutine.
**Test Data**: `go print("Hello")`
**Expected Output**: Concurrent output

## Problem 97: Channel Communication
**Description**: Use channels to communicate between goroutines.
**Test Data**: `ch := make(chan int)`
**Expected Output**: Sent/received values

## Problem 98: Buffered Channel
**Description**: Use buffered channel to hold values.
**Test Data**: `make(chan int, 3)`
**Expected Output**: Stores 3 values without blocking

## Problem 99: Select Statement
**Description**: Use select to wait on multiple channels.
**Test Data**: `select { case <-ch1: ... case <-ch2: ... }`
**Expected Output**: One ready channel selected

## Problem 100: Mutex for Concurrency
**Description**: Use `sync.Mutex` to protect shared resource.
**Test Data**: Shared counter incremented by goroutines
**Expected Output**: Race-free counter