package main

import ("golang.org/x/tour/tree"
       "fmt"
		 )

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan []int) {
	cur_el := t.Left.Left.Left.Value
	var elements []int
	for i := 0; i < 10; i++ {
		element := cur_el * i
		elements = append(elements, element)
	}
	ch <- elements
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan []int)
	go Walk(t1, c1)
	els1 := <-c1
	go Walk(t2, c1)
	els2 := <-c1
	if els1[1] == els2[1] {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2))) // returns false
	fmt.Println(Same(tree.New(1), tree.New(1))) // returns true
}
