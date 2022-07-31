package main

// import "fmt"

// func main() {

// 	//Constants:

// 	// const signConstValue2 float64 = math.Sin(1.57)	// math.Sin(1.57) (value of type float64) is not constant.

// 	// IN the below example we haven't assigned type to the constant.
// 	//So compiler will inffer the type of the constant.
// 	const a = 42
// 	var b int16 = 27
// 	fmt.Printf("%v ,%T\n", a+b, a+b)

// 	// const c int = 23
// 	// var d int16 = 27
// 	// fmt.Printf("%v ,%T\n", c+d, c+d) //invalid operation: c + d (mismatched types int and int16)

// 	//Enumerated Constants:
// 	const (
// 		e = iota
// 		f = iota
// 		g = iota
// 	)
// 	fmt.Println(e, f, g) //0 1 2

// 	//Auto incrementing constants.
// 	const (
// 		h = iota
// 		i
// 		j
// 	)

// 	fmt.Println(h, i, j) //0 1 2

// 	const (
// 		_ = iota + 5
// 		k
// 		l
// 	)

// 	fmt.Println(k, l) //6 7
// }
