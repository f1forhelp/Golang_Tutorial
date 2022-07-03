package main

// import (
// 	"fmt"
// )

// func main() {
// 	//Ex
// 	doctor := doctor{
// 		name: "John",
// 		age:  50,
// 	}

// 	fmt.Println(doctor) //{John 50 []}
// 	doctor.age = 22
// 	doctor.patients = []string{"John", "Jane"}
// 	fmt.Println(doctor) //{John 22 [John Jane]}

// 	//Ex: Anonymous struct

// 	person := struct {
// 		name string
// 		age  int
// 	}{
// 		name: "John",
// 		age:  50,
// 	}

// 	fmt.Println(person) //{John 50}

// 	//In go we don't have inheritance but rather composition (HAS - A) relationship
// 	// go produces syntactic sugar to access property of struct which HAS-A another struct.
// 	//This is also called embedding though there is an alternate called Interface.

// 	//Ex: Embedding struct
// 	class1 := class{
// 		SchoolName: "Some Random School",
// 		TeacherProp: Teacher{
// 			teacherName: "John",
// 			teacherAge:  50,
// 		},
// 		students: []student{
// 			{
// 				studentName: "John",
// 				studentRoll: 50,
// 			},
// 			{
// 				studentName: "Some",
// 				studentRoll: 23,
// 			},
// 		},
// 	}

// 	fmt.Println(class1) //{Go {John 50} [{John 50} {Some 23}]}

// 	//Ex: In the below example we can acess the properties of the embedded struct using the dot operator (Student Strength).
// 	class2 := class{}
// 	class2.ClassName = "Go"
// 	class2.studentStrength = 23
// 	class2.TeacherProp.teacherName = "Teacher Name"

// }

// type doctor struct {
// 	name     string
// 	age      int
// 	patients []string
// }

// type class struct {
// 	SchoolName string
// 	ClassProp
// 	TeacherProp Teacher
// 	students    []student
// }

// type student struct {
// 	studentName string
// 	studentRoll int
// }

// type Teacher struct {
// 	teacherName string
// 	teacherAge  int
// }

// type ClassProp struct {
// 	ClassName       string
// 	studentStrength int
// }
