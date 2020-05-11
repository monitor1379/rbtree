package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-11 14:41:40
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 19:23:33
 */

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/monitor1379/rbtree"
)

func TestBSTree(t *testing.T) {
	tree := rbtree.NewIntBSTree()

	tree.InsertOrReplace(10, 10)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(5, 5)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(3, 3)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(1, 1)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(15, 15)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(13, 13)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(4, 4)
	fmt.Println(tree.PrettyString())

	fmt.Println(tree.Search(1))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(3))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(5))
}

func BenchmarkBSTreeRandomInsert(b *testing.B) {
	b.ResetTimer()

	tree := rbtree.NewIntBSTree()
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(rand.Int(), i)
	}
}

func BenchmarkBSTreeWorstInsert(b *testing.B) {
	b.ResetTimer()

	tree := rbtree.NewIntBSTree()
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(i, i)
	}
}
