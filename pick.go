/* 
Package luckydip selects a random selection of items from a larger set
*/
package luckydip

import(
	"math/rand"
	"time"
)

// RandomSubset takes a slice and returns a random subset of the requested length
func RandomSubset(set []string, number int) []string {		
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(set), func(i, j int) { set[i], set[j] = set[j], set[i]})
	return set[:number]	
}