package main

// import "fmt"

// func main() {
// 	//Maps:
// 	//Maps are unordered key-value pairs.
// 	//They also points to the same memory location as array does.

// 	//Ex: By using make we can assign map later.
// 	population := make(map[string]int)
// 	population["India"] = 13000
// 	population["China"] = 13000
// 	population["USA"] = 13000

// 	fmt.Println(population) //map[China:13000 India:13000 USA:13000]

// 	//Ex: without using make
// 	population2 := map[string]int{
// 		"India": 13000,
// 		"China": 13000,
// 		"USA":   13000,
// 	}

// 	fmt.Println(population2) //map[China:13000 India:13000 USA:13000]

// 	//Ex: In the below examplle we are trying to print the value of a key which is not present in the map.
// 	population3 := map[string]int{
// 		"India": 13000,
// 		"China": 13000,
// 		"USA":   13000,
// 	}

// 	fmt.Println(population3["Dosent Exist"]) //0

// 	//We can use the below approach to check for weather the value exist or not.
// 	res, ok := population3["Dosent Exist"]

// 	fmt.Println(res, ok) //0 false

// 	delete(population3, "USA")

// 	fmt.Println(population3) //map[China:13000 India:13000]

// }
