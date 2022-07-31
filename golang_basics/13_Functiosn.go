package main

// import "fmt"

// func main() {
// 	name := "Some one"
// 	example1(&name)
// 	fmt.Println(name) //This is cat
// 	sum("This is message ", 2, 3, 4)
// 	fmt.Println("Sum is ", *sum1(2, 3, 4))

// 	//Functions are basically first class citizens. and we can do anything that we do with variables.
// 	var divide func(float64, float64) (float64, error) //Function Signature Declaration

// 	divide = func(f1, f2 float64) (float64, error) { //Assigning function
// 		if f2 != 0 {
// 			return 0, fmt.Errorf("Divide By Zero")
// 		}
// 		return f1 / f2, nil
// 	}

// 	if res, err := divide(2, 0); err != nil {
// 		fmt.Println("Error is :", err)
// 	} else {
// 		println(res)
// 	}

// }

// func example1(stringValue *string) {
// 	*stringValue = "This is cat"
// }

// func sum(message string, values ...int) {
// 	sum := 0
// 	for _, v := range values {
// 		sum += v
// 	}
// 	fmt.Printf("Message %v , Sum is %v\n", message, sum)
// }

// func sum1(values ...int) *int {
// 	// 	In go we can return a local variable as a pointer. So we are assigning pointer rather then copying value to temporary variable then return.
// 	// This is very important to note that we are accessing the pointer of a local variable of a function after its execution, this may seem weird as we are accessing the variable after that functional stack is popped out of the stack memory.
// 	// But in go,local pointers are promoted to the heap when they are accessed outside of the stack this promotion is automatic.
// 	// Its helps to increase performance as instead of copying the whole value we are passing the address as return value.

// 	result := 0

// 	for _, v := range values {
// 		result += v
// 	}

// 	return &result
// }

// //Named return in golang.
// func namedReturnExample(values ...int) (result int) {
// 	for _, v := range values {
// 		result += v
// 	}
// 	return result
// }

// //Returning multiple values.
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0.0, fmt.Errorf("Cannot Divide By Zero")
// 	}
// 	return a / b, nil
// }
