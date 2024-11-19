package main

import "fmt"

func compareTriplets(a [3]int32, b [3]int32) [2]int32 {
	var aliceWins, bobWins int32
	for i:=0; i<3; i++ {
		if a[i] > b[i] {
			aliceWins += 1
		} else if a[i] < b[i] {
			bobWins += 1
		} else {}
	}
	return [2]int32 {aliceWins, bobWins}
}

func main() {
	aliceScores := [3]int32 {}
	bobScores := [3]int32 {}

	fmt.Print("Enter Alice's scores : ")
	fmt.Scanf("%d %d %d", &aliceScores[0], &aliceScores[1], &aliceScores[2])
	fmt.Print("Enter Bob's scores : ")
	fmt.Scanf("%d %d %d", &bobScores[0], &bobScores[1], &bobScores[2])
	fmt.Println("Alice's and Bob's final scores are : ", compareTriplets(aliceScores, bobScores))
}