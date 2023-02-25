package main

import (
	"fmt"
	"time"
)

// func printNumbers() {
//     for i := 1; i <= 10; i++ {
//         fmt.Println(i)
//     }
// }
func printNumbers(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	//go printNumbers() // create a goroutine
	//go printNumbers() // create another goroutine
	time.Sleep(1 * time.Second)
	// main function continues to execute here

	c := make(chan int)
	go printNumbers(c)
	for n := range c {
		fmt.Println(n)
	}

}

// func printNumbers(c chan int) {
//     for i := 1; i <= 10; i++ {
//         c <- i
//     }
//     close(c)
// }

// func main() {
//     c := make(chan int)
//     go printNumbers(c)
//     for n := range c {
//         fmt.Println(n)
//     }
//     d := make(chan int)
//     go printNumbers(d)
//     for n := range d {
//         fmt.Println(n)
//     }
// }


// func infiniteLoop() {
//     for {
//         fmt.Println("Here we are")
//     }
// }

// func main() {
//     go infiniteLoop()
// }

// func infiniteLoop(done chan bool) {
//     for {
//         select {
//         case <-done:
//             return
//         default:
//             // Do something
//         }
//     }
// }

// func main() {
//     done := make(chan bool)
//     go infiniteLoop(done)
//     // Do something else
//     done <- true
// }

