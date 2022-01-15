package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

func WalkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int, 1)
	t2Chan := make(chan int, 1)
	go WalkAndClose(t1, t1Chan)
	go WalkAndClose(t2, t2Chan)
	for {
		read1, ok1 := <-t1Chan
		read2, ok2 := <-t2Chan
		// fmt.Println("Comparing - T1: ", read1, " T2: ", read2)
		if !ok1 && !ok2 {
			break
		}
		if ok1 != ok2 || read1 != read2 {
			return false
		}
	}
	return true
}

func main() {
	tree1 := tree.New(10)
	tree2 := tree.New(20)
	fmt.Println("Expecting False: ", Same(tree1, tree2))
	fmt.Println("Expecting True: ", Same(tree1, tree1))

}
