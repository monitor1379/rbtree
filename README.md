# rbtree

Red-Black Tree Golang Implementation


## Installation

```
go get -u -v github.com/monitor1379/rbtree
```


## Example

```go
package main

import (
	"fmt"

	"github.com/monitor1379/rbtree"
)

func main() {
	tree := rbtree.NewStringRBTree()

	tree.Put("name", "monitor1379")
	tree.Put("age", 25)

	name, ok := tree.Get("name")
	fmt.Println(name, ok)
	// print: monitor1379 true

	tree.Remove("age")

	age, ok := tree.Get("age")
	fmt.Println(age, ok)
	// print: <nil> false

}

```



## Benchmark

Platform: 
```
goos: linux
goarch: amd64
pkg: github.com/monitor1379/rbtree
```

Run benchmark testing:
```
cd $GOPATH/src/github.com/monitor1379/rbtree
go test -v -bench=. -run=none .
```


RBTree:
BenchmarkRBTreeRandomInsert-8                    1000000              1280 ns/op
BenchmarkRBTreeRandomInsertAndSearch-8           1000000              1092 ns/op
BenchmarkRBTreeRandomDelete-8                    1000000              1189 ns/op

Binary-Search Tree:
BenchmarkBSTreeRandomInsert-8                    1000000              1491 ns/op
BenchmarkBSTreeWorstInsert-8                       50000            126385 ns/op

GODS library:
BenchmarkGodsRBTreeRandomInsert-8                1000000              1260 ns/op
BenchmarkGodsRBTreeRandomInsertAndSearch-8       1000000              1118 ns/op

Golang builtin map:
BenchmarkMapInsert-8                             3000000               500 ns/op
BenchmarkMapRandomInsertAndSearch-8             10000000               168 ns/op

```