package main

import (
	"careercup/serializetree"
	"fmt"
)

func main() {
	t := serializetree.NewTree()
	t.Insert(3)
	t.Insert(5)
	t.Insert(6)
	fmt.Printf(t.Marshal())
}
