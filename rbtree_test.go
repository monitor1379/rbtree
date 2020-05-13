package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:38:53
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 23:39:19
 */

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/monitor1379/rbtree"
)

func TestRBTree(t *testing.T) {
	tree := rbtree.NewIntRBTree()

	keys := []int{
		10,
		51,
		21,
		37,
		20,
		58,
		48,
		16,
		49,
		84,
		87,
		74,
		36,
		15,
	}

	for _, key := range keys {
		fmt.Println(key)
		tree.InsertOrReplace(key, key)
		fmt.Println(tree.PrettyString())
	}
}

func TestRBTreeRandomInsert(t *testing.T) {
	tree := rbtree.NewIntRBTree()
	for i := 0; i < 10000; i++ {
		tree.InsertOrReplace(rand.Int(), i)
	}
}

func BenchmarkRBTreeRandomInsert(b *testing.B) {
	tree := rbtree.NewIntRBTree()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(rand.Int(), i)
	}
}

func BenchmarkRBTreeRandomInsertAndSearch(b *testing.B) {
	tree := rbtree.NewIntRBTree()

	keys := []int{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, rand.Int())
	}

	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(keys[i], i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(keys[i])
	}
}

func TestRBTreeDelete(t *testing.T) {
	tree := rbtree.NewIntRBTree()

	keys := []int{}
	for i := 0; i < 10000; i++ {
		keys = append(keys, rand.Int()%10000)
	}

	for i := 0; i < 10000; i++ {
		tree.InsertOrReplace(keys[i], i)
	}

	fmt.Println(tree.PrettyString())
	for i := 0; i < 10000; i++ {
		fmt.Println("deleting:", keys[i])
		tree.Delete(keys[i])
	}
	fmt.Println(tree.PrettyString())
}

func BenchmarkRBTreeRandomDelete(b *testing.B) {
	tree := rbtree.NewIntRBTree()

	keys := rand.Perm(b.N)
	for i := 0; i < b.N; i++ {
		tree.InsertOrReplace(keys[i], i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(keys[i])
	}
}
