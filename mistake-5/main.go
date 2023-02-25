package main

import "fmt"

// func divide(a, b int) (int, error) {
//     if b == 0 {
//         return 0, fmt.Errorf("cannot divide by zero")
//     }
//     return a / b, nil
// }

// func main() {
//     result, err := divide(5, 0)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     fmt.Println(result)
// }
func divide(a, b int) int {
    if b == 0 {
        panic("Cannot divide by zero")
    }
    return a / b
}

func main() {
    result := divide(5, 0)
    fmt.Println("Result:", result)
}

