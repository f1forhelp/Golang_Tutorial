package main

// import "fmt"

// func main() {
// 	//If statement

// 	//Ex:
// 	if true {
// 		fmt.Println("This is true")
// 	} else {
// 		fmt.Println("This is false")
// 	}

// 	//Ex:
// 	if age, ok := tempMap["age"]; ok {
// 		fmt.Println(age)
// 	} else {
// 		fmt.Println("Not found")
// 	}

// 	//Switch Statement

// 	//Ex:

// 	roll := 1

// 	switch roll {
// 	case 1:
// 		fmt.Println("1")
// 	case 2:
// 		fmt.Println("2")
// 	default:
// 		{
// 			fmt.Println("Default")
// 		}

// 	}

// 	//Ex: We can use initializer in switch statement

// 	switch i := 2 + 3; i {
// 	case 1, 5, 10:
// 		{
// 			println("Value is 1 , 5 or 10")
// 		}
// 	case 2, 4, 6:
// 		{
// 			println("Two , four or six")
// 		}
// 	default:
// 		{
// 			println("Another NUmber")
// 		}
// 	}

// 	//In go lang we have fallthrough keyword which is used to fall through the case. the default behaviou is break statement.
// 	//Ex:
// 	switch i := 2 + 3; i {
// 	case 1, 5, 10:
// 		println("Value is 1 , 5 or 10")
// 		fallthrough

// 	case 2, 4, 6:
// 		println("Two , four or six")
// 		fallthrough

// 	default:
// 		{
// 			println("Another NUmber")
// 		}
// 	}

// 	//Ex: Type switching
// 	var i interface{} = 1
// 	switch i.(type) {
// 	case int:
// 		println("Integer")
// 	case string:
// 		println("String")
// 	default:
// 		println("Unknown")
// 	}

// }

// var tempMap map[string]int = map[string]int{
// 	"Name": 1,
// }
