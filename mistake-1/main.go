package main

import "fmt"

func main() {
	// Let's declare a variable x and assign it a value of 10
	x := 10

	// Let's declare a pointer p that points to x
	p := &x

	// Let's print the value of x using the pointer
	fmt.Println(*p) // Output: 10

	// Let's modify the value of x using the pointer
	*p = 20

	// Let's print the modified value of x
	fmt.Println(x) // Output: 20

	//p := new(int)  // We allocate memory for an int and return a pointer to it
	*p = 10         // We then assigned the value 10 to the memory location pointed to by p
	fmt.Println(*p) // Output: 10

	// We declare a new struct named book
	type Book struct {
		Name   string
		Author string
		Price  int
	}

	b := Book{"Half of a sun", "John doe", 8000}
	//p := &b.Price        // Let's get the memory address of the Price field of the Book struct
	*p = 7000            // Let's modify the Price field using the pointer
	fmt.Println(b.Price) // Output: 7000

	//x := 10
	//p := *x        // Incorrect use of the * operator
	fmt.Println(p) // invalid operation: cannot indirect x (variable of type int)

	var q *int      // Declare a pointer q
	fmt.Println(*q) // Dereference q, which causes an error: panic: runtime error: invalid memory address or nil pointer dereference

	//p := new(int)
	*p = 10

	//var q *float64 = p // cannot use p (variable of type *int) as type *float64 in variable declaration
	fmt.Println(*q)

	

}
