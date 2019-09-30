package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(Pick(2))
}

func Pick(num int) []string {
	pool := []string{"cat","dog","bat","car"}
	return randomSubset(pool,num)
}

func randomSubset(set []string, number int) []string {	
	results := make([]string,0)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < number; i++ {	
		randomChoice := set[rand.Intn(len(set))]
		results = append(results, randomChoice)		
	}

	return results
}

func generateIndexSet(set []string) []int {
	indicies := make([]int,len(set))

	for i := range set {
		indicies[i] = i
	}

	return indicies
}