package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func simpleArraySum(ar []int32) int32 {
    sum := int32(0)
	for _, num := range ar {
		sum += num
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter input array : ")
	inputStr, _ := reader.ReadString('\n')
	
	inputArr := []int32{}
	for _, token := range strings.Split(inputStr, " ") {
		token = strings.TrimSpace(token)
		num, _ := strconv.Atoi(token)
		inputArr = append(inputArr, int32(num))
	}
	fmt.Println("INPUT ARR :", inputArr)
	fmt.Println("Result : ", simpleArraySum(inputArr))
}
