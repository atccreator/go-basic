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


func main() {
	// ArrayExample()
	// fun()
	// variableDeclarationExample()
	// Sum(5, 10)
	// UserInputExample()
	arr := []int{12, 11, 13, 5, 6}
	sortedArr := Insertionsort(arr)
	fmt.Println("Sorted array:", sortedArr)

	// Demonstrating underflow behavior
	var a uint8 = 0
	var b int8 = 0
	b--
	a--
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
}

