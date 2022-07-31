package main

// func main() {
// 	deferExample1()
// 	deferExample2()
// 	deferExample3()
// }

// func deferExample1() {

// 	//if we execute multiple defer at the same time then defer will be called in a LIFO manner.
// 	//This is done so that we can close resources in the opposite order it may be possible that child resource is dependent on parent resource.
// 	//Important: when dealing with lot of request it might be possible that you may be using a loop to open lot of resources at that time defer will only close after the function ends so be careful while using defer in those kind of scenario.

// 	defer println("First")
// 	defer println("Second")
// 	defer println("Third")
// 	//OP: Third Second First (As you can see first we get third got printed)
// }

// func deferExample2() {
// 	//Even though defer takes this statement to end but the data will be stored in method stack so it will call start rather then finish.
// 	a := "Start"
// 	defer println(a) //OP: Start
// 	a = "Finish"
// }

// func deferExample3() {
// 	var a string = "Start"

// 	var b *string = &a
// 	defer println(*b)
// 	*b = "Stop"
// 	println("1")

// 	//OP: 1 Start
// }
