/*
	136. Single Number
	https://leetcode.com/problems/single-number/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func singleNumberInitial(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := numMap[num]; ok == false {
			numMap[num] = 1
		} else {
			numMap[num]+=1
		}
	}
	fmt.Println("Map :", numMap)
	for key, val := range numMap {
		if val==1 {
			fmt.Println("REUTURN :", key)
			return key
		}
	}
	return 0
}

func singleNumber(nums []int) int {
	// same numbers in nums gets cancelled out by XOR o/p
	// --> n^n = 0
	// --> n^0 = n
    result := 0
    for _, num := range nums{
        result ^= num
    }

    return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter input string `nums` : ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	var nums []int
	for _, strNum := range strings.Split(s, " ") {
		num, _ := strconv.Atoi(strNum)
		nums = append(nums, num)
	}

	fmt.Println("Result : ", singleNumber(nums))
}