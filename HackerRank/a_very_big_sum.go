package main

import "fmt"

func main() {
	arrLen := 0
	fmt.Print("Enter size of array : ")
	fmt.Scanf("%d", &arrLen)

	arr := []int64 {}
	for i:=0; i<arrLen; i++ {
		num, _ := fmt.Scanf("%d")
		arr = append(arr, int64(num))
	}

	fmt.Println("INPUT ARR :", arr)

}