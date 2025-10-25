package main

import "fmt"

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


func main() {
	fun()
	variableDeclarationExample()
	// Sum(5, 10)
	var a uint8 = 0
	var b int8 = 0
	b--
	a--
	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
}

