package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var br bool
	done := make(chan bool)
	c1 := make(chan int)
	c2 := make(chan int)
	go Walker(t1, c1)
	go Walker(t2, c2)
	go func() {
		for i := range c1 {
			if i == <-c2 {
				br = true
			} else {
				br = false
				break
			}
		}
		done <- br
	}()
	<-done
	return br
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("Are they same? - ", Same(t1, t2))

	t3 := tree.New(1)
	t4 := tree.New(2)
	fmt.Println("Are they same? - ", Same(t3, t4))
}
