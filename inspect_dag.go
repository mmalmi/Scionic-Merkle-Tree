package main

import (
	"fmt"
	"os"
	"github.com/HORNET-Storage/Scionic-Merkle-Tree/v2/dag"
)

func main() {
	d, _ := dag.CreateDag(os.Args[1], false)
	root := d.Leafs[d.Root]
	fmt.Printf("Root.DagSize: %d\n", root.DagSize)
	fmt.Printf("Root.ContentSize: %d\n", root.ContentSize)
	fmt.Printf("Root.LeafCount: %d\n", root.LeafCount)
}
