package main

// import "fmt"

// func main() {
// 	recoverExmaple1()
// 	fmt.Println("This is main")

// 	//As you can see we have handled the error with recover and rest of our program will continue.

// 	//1
// 	//Error Is  (0x10446f700,0x10447fca0)
// 	//This is main
// }

// func recoverExmaple1() {
// 	println("1")
// 	defer func() {
// 		if err := recover(); err != nil {
// 			println("Error Is ", err)
// 		} else {
// 			println("No Error")
// 		}
// 	}()
// 	panic("[This is panic]")
// }
