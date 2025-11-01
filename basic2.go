// Lecture 43: Basic of Go Part-2
// Go comes with automatic garbage collection, which helps manage memory by automatically reclaiming memory that is no longer in use. 
// This reduces the chances of memory leaks and other related issues.
package main

import (
	"fmt"
)
// -------------- Example 1: --------------
// understanding Structs and Methods in Go

type Product struct {
	Name  string //field or data member
	Price float64 //field or data member
}

// Method or member function to calculate discounted price, here p is the receiver
func (p Product) DiscountedPrice(discount float64) float64 {
	return p.Price * (1 - discount)
}

// -------------- Example 2: --------------
// Struct constructor functions in Go

type Product2 struct {
	name string
	price float64
	company string
}

// this function acts as a constructor for Product2 struct (good practice to use New prefix for constructor functions) 
// using pointer and returning struct by reference
func NewProduct2(name string, price float64, company string) *Product2 {
	p := Product2{
		name: name,
		price: price,
		company: company,
	}
	return &p // this return the pointer to the Product2 struct
}

// this function acts as a constructor for Product2 struct without using pointer and returning struct by value/copy 
// (not recommended for large structs) unnecessary copying of data 
func NewProduct3(name string, price float64, company string) Product2 {
	p := Product2{
		name: name,
		price: price,
		company: company,
	}
	return p // this return the pointer to the Product2 struct
}

func PassByValueExample(copyOfProduct Product2) {
	copyOfProduct.company = "ChangedCompany" // this will not affect the original struct
}

func PassByReferenceExample(refToProduct *Product2) {
	refToProduct.company = "ChangedCompany" // this will affect the original struct
}

func main(){
	// Example 1:
	// product := Product{Name: "Iphone", Price: 1000.00} // Creating an instance of Product struct without using new keyword
	// discount := 0.1 // 10% discount
	// fmt.Printf("Original price of the %s is $%.2f\n", product.Name, product.Price)
	// fmt.Printf("The discounted price of the %s is $%.2f\n", product.Name, product.DiscountedPrice(discount))

	// Example 2:
	newProduct := NewProduct2("Laptop", 1500.00, "TechCorp")
	fmt.Println("Product Name:", (*newProduct).name) // manually dereferencing pointer to access name field
	fmt.Println("Product Price:", newProduct.price) // Go automatically dereferences pointers to access fields
	fmt.Println("Product Company:", newProduct.company)

	newProduct2 := NewProduct3("Smartphone", 800.00, "MobileInc")
	fmt.Println("Product2 Name:", newProduct2.name)
	fmt.Println("Product2 Price:", newProduct2.price)
	fmt.Println("Product2 Company:", newProduct2.company)


	PassByValueExample(newProduct2) // passing struct by value (copy)
	fmt.Println("After PassByValueExample, Product Company:", newProduct2.company) // company should remain unchanged

	PassByValueExample(*newProduct) // passing struct by value (copy)
	fmt.Println("After PassByValueExample, Product Company:", newProduct.company) // company should remain unchanged

	PassByReferenceExample(newProduct) // passing struct by reference (pointer)
	fmt.Println("After PassByReferenceExample, Product Company:", newProduct.company) // company should be changed

	

}