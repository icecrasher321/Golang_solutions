package main
// URL -> https://tour.golang.org/moretypes/23
import (
	"golang.org/x/tour/wc"
	"strings"
)
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	var counters []int
	x := 0
	end := len(words)
	for i1 := 0; i1 < end; i1++ {    // Current element -> i1
		for i2 := 0; i2 < end; i2++ {  // Scans through the array to count -> i2
			if words[i1] == words[i2] {
				x += 1
			}
		}
		counters = append(counters, x)   // Appends it to an array called counters (contains all the word counts)
		x = 0
	}
	final := make(map[string]int)
	for i := 0; i < len(counters); i++ {
		final[words[i]] = counters[i]
	}
	return final
}
func main() {
	   wc.Test(WordCount)
}
