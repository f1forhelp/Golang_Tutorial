package main

import "fmt"

func main() {
	//Boolean
	fmt.Println("Hello World")
	i := 2 == 3
	fmt.Println(i)
	var b bool = true
	fmt.Println(b)

	//Integer
	var uint64Val uint64 = 9223372036854775807
	fmt.Println(uint64Val)

	//Float
	var floatVal float64 = 3.14
	fmt.Println(floatVal)

	//String
	var stringVal string = "Hello World"
	fmt.Println("First Cahr of string is : ", stringVal[0], " Last Char of string is : ", string(stringVal[len(stringVal)-1])) //First Cahr of string is :  72  Last Char of string is :  d

	str1 := "Hello"
	str2 := "World"

	fmt.Println("Concatinated String is : ", str1+str2) //Concatinated String is : HelloWorld

}
