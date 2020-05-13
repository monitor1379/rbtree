package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-12 10:06:23
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-12 10:10:06
 */

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func TestGodsRBTree(t *testing.T) {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(1, 1)
	tree.Put(2, 2)
	data, err := tree.ToJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func BenchmarkGodsRBTreeRandomInsert(b *testing.B) {
	tree := redblacktree.NewWithIntComparator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(rand.Int(), i)
	}
}

func BenchmarkGodsRBTreeRandomInsertAndSearch(b *testing.B) {
	tree := redblacktree.NewWithIntComparator()

	keys := []int{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, rand.Int())
	}

	for i := 0; i < b.N; i++ {
		tree.Put(keys[i], i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(keys[i])
	}
}
