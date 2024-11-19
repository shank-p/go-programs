/*
	1. Two Sum
	- https://leetcode.com/problems/two-sum/description/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) [2]int {
	diffMap := map[int]int {}
	for i:=0; i<len(nums); i++ {
		numIdx, ok := diffMap[nums[i]]
		if ok {
			return [2]int {numIdx, i}
		}
		diffMap[target-nums[i]] = i
	}
	return [2]int {-1, -1}
}

func main() {
	var inputStr string
	reader := bufio.NewReader(os.Stdin)

	var nums []int
	fmt.Print("Enter nums slice : ")
	inputStr, _ = reader.ReadString('\n')
	for _, num := range strings.Split(inputStr, " ") {
		num = strings.TrimSpace(num)
		num, _ := strconv.Atoi(num)
		nums = append(nums, num)
	}

	var target int
	fmt.Print("Enter target : ")
	fmt.Scan(&target)

	fmt.Println("Result : ", twoSum(nums, target))
}