package main

// import "fmt"

// func main() {
// 	//Arrays:
// 	grades := [3]int{97, 92, 81}

// 	var students [3]string

// 	fmt.Println(grades)   //[97 92 81]
// 	fmt.Println(students) //[ ]

// 	//In golang if we assign an array to a variable, we are basically coppying the content of first array to the new variable rather then passing the pointer.

// 	//Ex: 1
// 	var gradesCopy [3]int = grades

// 	grades[0] = 23

// 	fmt.Println(grades)     //[23 92 81]
// 	fmt.Println(gradesCopy) //[97 92 81]

// 	//Ex: 2
// 	//In the below example we used ... to pass size of array to constructor.
// 	fruits := [...]string{"apple", "banana", "orange"}

// 	newFruits := fruits
// 	newFruits[0] = "kiwi"

// 	fmt.Println(fruits)    //[apple banana orange]
// 	fmt.Println(newFruits) //[kiwi banana orange]

// 	//Slices:

// 	//Slices are a reference type. Slice are different from array as we don't have to pass size of array to constructor.

// 	slice1 := []int{1, 2, 3, 4, 5}

// 	var slice2 []int
// 	fmt.Println(slice2) //[]

// 	slice2 = slice1
// 	fmt.Println(slice2) //[1 2 3 4 5]

// 	slice2[0] = 10
// 	fmt.Println(slice1) //[10 2 3 4 5]
// 	fmt.Println(slice2) //[10 2 3 4 5]

// 	//Ex: 3
// 	//In the below example we are converting an array to slice.
// 	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
// 	var slice3 []int = array1[:]
// 	slice4 := array1[1:5]

// 	slice3[0] = 10
// 	slice4[1] = 20

// 	fmt.Println(array1) //[10 2 20 4 5]
// 	fmt.Println(slice3) //[10 2 20 4 5]
// 	fmt.Println(slice4) //[2 20 4 5]

// 	//Ex: 4
// 	//In the below example we are converting a slice to array.

// }
