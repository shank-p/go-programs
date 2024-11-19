package main

import "fmt"

func solveMeFirst(a uint32,b uint32) uint32{
  return uint32(a+b)
}

func main() {
    var a, b uint32
	fmt.Print("Enter value for `a` : ")
    fmt.Scanf("%v", &a)
	fmt.Print("Enter value for `b` : ")
    fmt.Scanf("%v", &b)
	fmt.Println("Result : ", solveMeFirst(a,b))
}