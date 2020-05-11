package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:38:53
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 20:33:50
 */

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/monitor1379/rbtree"
)

func TestRBTree(t *testing.T) {
	tree := rbtree.NewIntRBTree()

	tree.InsertOrReplace(15, 15)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(5, 5)
	fmt.Println(tree.PrettyString())

	tree.InsertOrReplace(1, 1)
	fmt.Println(tree.PrettyString())

	// tree.InsertOrReplace(1, 1)
	// fmt.Println(tree.PrettyString())

	// tree.InsertOrReplace(15, 15)
	// fmt.Println(tree.PrettyString())

	// tree.InsertOrReplace(13, 13)
	// fmt.Println(tree.PrettyString())

	// tree.InsertOrReplace(4, 4)
	// fmt.Println(tree.PrettyString())

}

func BenchmarkRBTreeRandomInsert(b *testing.B) {
	b.ResetTimer()

	tree := rbtree.NewIntRBTree()
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(rand.Int(), i)
	}
}

func BenchmarkRBTreeWorstInsert(b *testing.B) {
	b.ResetTimer()

	tree := rbtree.NewIntRBTree()
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(i, i)
	}
}
