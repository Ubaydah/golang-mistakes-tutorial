package main

import "fmt"

type Speakable interface {
	Speak() string
	Listen() string
}
type Animal struct {
	Name string
}

func (a *Animal) Speak() string {
	return "I am an animal and my name is " + a.Name
}

type Printer interface {
	Print()
}

type LaserPrinter struct{}

func (lp LaserPrinter) Print() {
	// Implement the Print method for LaserPrinter
	fmt.Println("Welcome")
}

func main() {
	a := Animal{
		Name: "goat",
	}
	fmt.Println(a.Speak())

	var p Printer
	p = LaserPrinter{}
	p.Print()

	var i Speakable
	i.Speak() // runtime error: invalid memory address or nil pointer dereference

	//var i Speakable
	//i = &Animal{Name: "john doe"}
	//i.Speak()

	var myInterface interface{}

	if myInterface != nil {
		fmt.Println("myInterface is not nil")
	} else {
		fmt.Println("myInterface is nil")
	}

	//x.(T)

	type People struct {
	}

	// var f Speakable
	// s := f.(*People) //impossible type assertion: i.(*People)

	// s.Speak()

}
