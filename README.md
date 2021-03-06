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


You can create a rbtree object by:
```go
rbtree.NewIntRBTree()
rbtree.NewStringRBTree()
rbtree.NewRBTree(rbtree.Comparator)
```


Note that `rbtree.Comparator` is a type of `interface{}`: 
```go
// Comparetor.Compare returns an integer comparing two object lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//
type Comparator interface {
	Compare(interface{}, interface{}) int
}
```


Such as:
```go
type IntComparator struct{}

func (c *IntComparator) Compare(i, j interface{}) int {
	ii, ok := i.(int)
	if !ok {
		panic(ErrInvalidInputTypeOfComparator)
	}

	jj, ok := j.(int)
	if !ok {
		panic(ErrInvalidInputTypeOfComparator)
	}

	if ii < jj {
		return -1
	} else if ii == jj {
		return 0
	}
	return 1
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
```
BenchmarkRBTreeRandomInsert-8                    1000000              1280 ns/op
BenchmarkRBTreeRandomInsertAndSearch-8           1000000              1092 ns/op
BenchmarkRBTreeRandomDelete-8                    1000000              1189 ns/op
```

Binary-Search Tree:
```
BenchmarkBSTreeRandomInsert-8                    1000000              1491 ns/op
BenchmarkBSTreeWorstInsert-8                       50000            126385 ns/op
```


GODS library:
```
BenchmarkGodsRBTreeRandomInsert-8                1000000              1260 ns/op
BenchmarkGodsRBTreeRandomInsertAndSearch-8       1000000              1118 ns/op
```

Golang builtin map:
```
BenchmarkMapInsert-8                             3000000               500 ns/op
BenchmarkMapRandomInsertAndSearch-8             10000000               168 ns/op
```
