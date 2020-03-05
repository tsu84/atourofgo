// https://tour.golang.org/concurrency/7

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func StartWalk(t *tree.Tree, ch chan int) {
	Walk(t,ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go StartWalk(t1,c1)
	go StartWalk(t2,c2)
	
	m := make(map[int] int)
	for i1 := range c1 {
		m[i1] = i1
	}
	
	for i2 := range c2 {
		if i2 != m[i2] {
			return false
		}
	}
	
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
