package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkAll(t *tree.Tree, ch chan int) {
	if t != nil {
		if t.Left != nil {
			WalkAll(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			WalkAll(t.Right, ch)
		}
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkAll(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 && ok2 {
			if v1 != v2 {
				return false
			}
		} else if !ok1 && !ok2 {
			return true
		} else {
			return false
		}
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t2))
}
