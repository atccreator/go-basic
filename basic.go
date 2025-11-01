// Lecture 42: Basic of Go Part-1
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func fun(){
	fmt.Println("Inside Fun: Hello, Mr. Go!");
}

func variableDeclarationExample(){
	var productName string = "Laptop" // Explicit type declaration
	companyName := "TechCorp" // Implicit type declaration here type is inferred as string and := is used for declaration
	var price float64 // Explicit type declaration
	price = 999.99
	var inStock = true // here type is inferred as bool
	const pi = 3.14 // constant declaration cannot be changed later

	fmt.Println("Product Name:", productName)
	fmt.Println("Company Name:", companyName)
	fmt.Println("Price:", price)
	fmt.Println("In Stock:", inStock)

	fmt.Println("loop example output:")
	loopExample()
}

func loopExample(){

	// traditional for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// for range loop python style
	// Example
	for i := range []int{1, 2, 3} {
		fmt.Println("Number:", i)
	}

	for i := range 3 {
		fmt.Println("Number:", i) // This will execute 3 times with i values 0, 1, 2
	}

	//Example 1: Iterating over a slice
	for i:= range []string{"apple", "banana", "cherry"} {
		fmt.Println("Fruit Index:", i)
	}
	//Example 2: Iterating over a map
	for key, value := range map[string]int{"a": 1, "b": 2, "c": 3} {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	//Example 3: Iterating over a channel
	ch := make(chan string)
	go func() {
		ch <- "Hello"
		ch <- "World"
		close(ch)
	}()
	for msg := range ch {
		fmt.Println("Received from channel:", msg)
	}

	// writing while loop using for
	count := 10

	for count > 0 {
		if count == 5 {
			count--
			continue
		}
		fmt.Println("Count:", count)
		count--
	}

	// Infinite loop example
	for {
		fmt.Println("Infinite Loop Example - Breaking after one iteration")
		break
	}
}

func ArrayExample(){

	// Dynamic array are called slices in Go

	var arr []int // declaration of empty integer array
	arr = append(arr, 10) // appending value to array
	arr = append(arr, 20, 30, 40) // appending multiple values to array
	fmt.Println("Array Elements of arr:", arr)

	var strArr = []string{"Go", "Python", "Java"} // declaration and initialization of string array
	fmt.Println("String Array Elements of strArr:", strArr)

	var arr2 []int
	fmt.Println("Length of arr2 before appending:", len(arr2))
	fmt.Println("Capacity of arr2 before appending:", cap(arr2)) // cap method gives the capacity of the array
	fmt.Println("Empity Array arr2:", arr2)

	// Create a slice with length 2 and capacity 5
	s := make([]int, 2, 5)
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))

	s = append(s, 10, 20) // Append two elements
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))

	s = append(s, 30) // Append another element, exceeding original capacity
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))


	// Different way of defining array with fixed size
	var fixedArr [3]int = [3]int{1, 2, 3} // with explicit type declaration
	fmt.Println("Fixed Size Array Elements of fixedArr:", fixedArr)
	var fixedArr2 = [5]string{"A", "B", "C", "D", "E"} // type inference implicitly
	fmt.Println("Fixed Size Array Elements of fixedArr2:", fixedArr2)
	var fixedArr3 [4]float64 // declaration of fixed size array
	fixedArr3[0] = 1.1
	fixedArr3[1] = 2.2
	fixedArr3[2] = 3.3
	fixedArr3[3] = 4.4
	fmt.Println("Fixed Size Array Elements of fixedArr3:", fixedArr3)

	for i ,val := range fixedArr3 {
		fmt.Printf("Index: %d, Value: %.2f\n", i, val)
	}

}

func UserInputExample(){
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // reads until the first space or newline.
	fmt.Println("Hello,", name)

	// Scanln example
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age) // reads until the newline character.
	fmt.Println("You are", age, "years old.")

	// Scanf example
	var name1 string
	var age1 int
	fmt.Print("Enter your name and age: ")
	fmt.Scanf("%s %d", &name1, &age1) // reads formatted input
	fmt.Println("Name:", name1, "| Age:", age1)

	// Using bufio for reading a full line including spaces
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)
	fmt.Println("Full Name:", fullName)

}

func Insertionsort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}


func MapExample(){
	// Map is a collection of key-value pairs just like dictionary in python and object in javascript and hashmap in java
	
	// declaration of map key is string and value is int
	studentGrades := map[string]int{
		"Alice": 85,
		"Bob": 90,
		"Charlie": 78,
	}

	fmt.Println("Student Grades:", studentGrades)

	// iterating over map
	for student, grade := range studentGrades {
		fmt.Printf("%s: %d\n", student, grade)
	}

	// Only declaration of map
	employeeSalaries := make(map[string]float64) // initializing map using make function, make is used to create slices, maps, and channels and it is built-in function in Go

	// adding key-value pairs to map
	employeeSalaries["John"] = 50000.50
	employeeSalaries["Jane"] = 60000.75

	fmt.Println("Employee Salaries:", employeeSalaries)

	// accessing value using key
	johnSalary := employeeSalaries["John"]
	fmt.Println("John's Salary:", johnSalary)

	// deleting key-value pair from map
	delete(employeeSalaries, "Jane")
	fmt.Println("Employee Salaries after deletion:", employeeSalaries)

	// checking if a key exists
	salary, exists := employeeSalaries["Jane"]
	if exists {
		fmt.Println("Jane's Salary:", salary)
	} else {
		fmt.Println("Jane's record not found.")
	}

	// declaring empty map
	customMAP := map[string]string{}
	fmt.Println("Empty Custom Map:", customMAP)

	customMAP["Language"] = "Go"
	customMAP["Version"] = "1.18"
	customMAP["Author"] = "Google"
	customMAP["License"] = "BSD"
	for key, value := range customMAP {
		fmt.Println(key, ":", value)
	}

	
}

// intersting odd and even example
func oddEvenExample(num int) (string, int) {
	if num%2 == 0 {
		return "Even", 0
	}else{
		return "Odd", 1
	}
}

// Demosnstrating pointer 
func pointerExample(){
	var x int = 10
	var p *int = &x // p is a pointer that holds the address of x
	var y **int = &p // y is a pointer to pointer that holds the address of p
	var z int = **y
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (Address of x):", p)
	fmt.Println("Value pointed by p (Value of x):", *p)
	fmt.Println("Value of y (Address of p):", y)
	fmt.Println("Value pointed by y (Value of p):", *y)
	fmt.Println("Value pointed by pointer to pointer y (Value of x):", **y)
	fmt.Println("Value of z (Value of x):", z)
}

// More complex pointer example using array of pointers and struct pointers
func complexPointerExample(){
	type Person struct {
		Name string
		Age  int
	}

	alice := Person{Name: "Alice", Age: 30}
	bob := Person{Name: "Bob", Age: 25}

	people := []Person{alice, bob}
	var p *Person = &people[0] // pointer to the first element
	var q **Person = &p        // pointer to pointer

	fmt.Println("Name:", (*q).Name) // output: Name: Alice
	fmt.Println("Age:", (*q).Age) // output: Age: 30
}


func main() {
	// ArrayExample()
	// fun()
	// variableDeclarationExample()
	// Sum(5, 10)
	// UserInputExample()
	// arr := []int{12, 11, 13, 5, 6}
	// sortedArr := Insertionsort(arr)
	// fmt.Println("Sorted array:", sortedArr)
	// MapExample()
	// val1, val2 := oddEvenExample(7)
	// fmt.Println("Odd/Even Check:", val1, "| Value:", val2)
	// pointerExample()
	complexPointerExample()

	// Demonstrating underflow behavior
	var a uint8 = 0
	var b int8 = 0
	b--
	a--
	// fmt.Println("Value of a:", a)
	// fmt.Println("Value of b:", b)
}

