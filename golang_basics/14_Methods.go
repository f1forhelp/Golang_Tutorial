package main

// import "fmt"

// //  Method : a method is just a function that is executed in a known context.

// func main() {
// 	g := greeter{
// 		greetings: "Hello",
// 		name:      "Go",
// 	}
// 	g.greet()            //OP: Hello {Hello Go} Go
// 	g.greetWithPointer() //op: Hello &{Hello Go} Go
// }

// type greeter struct {
// 	greetings string
// 	name      string
// }

// // We can only have single receiver of context in function ie  (g greeter) we cant do like this  (g greeter,n name) this is wrong.
// func (g greeter) greet() { //We have passed greeter as context
// 	fmt.Println(g.greetings, g, g.name)
// }

// func (g *greeter) greetWithPointer() { //we have passed pointer of greeter as context
// 	fmt.Println(g.greetings, g, g.name)
// }
