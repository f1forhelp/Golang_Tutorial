package main

// import "fmt"

// func main() {
// 	fmt.Println("This is main")

// 	// example1()
// 	// example2()
// 	// example3()
// 	// example4()
// 	// example5()
// 	example6()
// }

// func example1() {
// 	var b int = 42
// 	var a *int = &b
// 	println(b)  //OP: 42
// 	println(*a) //OP: 42

// 	//This difference in address is because of refrences stores more then only adress like type of value.
// 	println(a)  //OP: 0x14000068f40
// 	println(&a) //OP: 0x14000068f50
// }

// func example2() {
// 	var a *int
// 	*a = 32
// 	println(a)

// 	//Error: panic: runtime error: invalid memory address or nil pointer dereference
// 	//This happened because pointer is not yet been initialized. Thus throws an error.
// }

// func example3() {
// 	b := 0
// 	var a *int = &b
// 	*a = 32
// 	println(a) //0x14000068f50
// }

// func example4() {
// 	//Hear new creates a pointer.
// 	var a *int = new(int)
// 	println(*a) //0 //As new initialize pointers with default values. So this will not throw an error.
// 	*a = 32
// 	println(a) //0x14000068f50
// }

// func example5() {
// 	var ms *myStruct
// 	var ms1 *myStruct

// 	fmt.Println(ms) //<nil>
// 	println(ms1)    //0x0
// }

// func example6() {
// 	var ms *myStruct
// 	fmt.Println(ms) //<nil>
// 	ms = new(myStruct)
// 	fmt.Println(ms)  //&{}
// 	fmt.Println(&ms) //0x1400008c020

// 	(*ms).foo = 42 //They Both Are same
// 	ms.foo = 42    //But This one is clear its basicall syntactic sugar

// 	fmt.Println((*ms).foo) //OP:42
// 	fmt.Println(ms.foo)    //OP: 42
// }

// type myStruct struct {
// 	foo int
// }
