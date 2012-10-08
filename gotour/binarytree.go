package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
// Is there a left node? if yes, follow it
// Is there a right node? if yes, follow it
// No nodes? Push current node value to channel
	if(t.Left != nil){
		Walk(t.Left, ch)
	}
	if(t.Right != nil){
		Walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	//result <- ch
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
		
}