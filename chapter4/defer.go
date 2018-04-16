package main

import "fmt"

//func main() {
//	a := 1
//	defer fmt.Println(a)
//	a = 0
//	defer fmt.Printf("%d\n",4/a)
//	defer fmt.Println("789")
//}


func main() {
	f := func() { fmt.Println("Test") }
	defer f()
	f = nil // 这里将函数的“值”设为nil，此时再尝试执行这个函数的时候
	defer f()
	defer fmt.Println("Test1")
}
