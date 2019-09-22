package main

import(
	"fmt"
)

func main() {
	fmt.Println(Pick(2))
}

func Pick(num int) []string {
	pool := []string{"cat","dog","bat","car"}
	return pool[0:num]
}
