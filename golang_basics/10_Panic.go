package main

// import "fmt"

// func main() {
// 	panicExample2()
// }

// func panicExample1() {
// 	//Go dont have exceptions, generally when other languages have like opening a file that doesn't exist they throw exception .
// 	//In Go we have panics that are generated when we do something really evil like divide by zero.

// 	println("1")
// 	panic("This is panic")
// 	print("Non Reachable Code")

// 	//1
// 	//panic: This is panic

// 	//goroutine 1 [running]:
// 	//main.main()
// 	//	/Users/utkarshsharma/go/src/Golang_Tutorial/basics/10_Panic.go:5 +0x48
// 	//exit status 2
// }

// func panicExample2() {
// 	//In this example we have defer at the end of the function this can be problamatic as we dont have all the resourecs closed before any panic.
// 	fmt.Println("1")
// 	a, b := 0, 1
// 	println(b / a)
// 	defer fmt.Println("2")
// }

// func panicExample3() {
// 	//In this example we have defer at the beginning of the function this helps to close all the resources at before any panic.
// 	fmt.Println("1")
// 	a, b := 0, 1
// 	println(b / a)
// 	defer fmt.Println("2")
// }
