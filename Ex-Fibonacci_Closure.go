package main

// URL -> https://tour.golang.org/moretypes/26
import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1 := 0
	n2 := 1
	count := 0
	fmt.Println(n1)
	fmt.Println(n2)
	return func() int {
			count = n2 + n1
			n1 = n2
			n2 = count
		return count
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
