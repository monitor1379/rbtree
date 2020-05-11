# rbtree
Red Black Tree Golang Implementation


## Benchmark

goos: linux
goarch: amd64
pkg: github.com/monitor1379/rbtree

BenchmarkBSTreeRandomInsert-8                    1000000              1501 ns/op
BenchmarkBSTreeWorstInsert-8                       50000            131459 ns/op
BenchmarkMapInsert-8                             2000000               658 ns/op
BenchmarkMapRandomInsertAndSearch-8             10000000               169 ns/op
BenchmarkRBTreeRandomInsert-8                    1000000              1303 ns/op
BenchmarkRBTreeRandomInsertAndSearch-8           1000000              1116 ns/op