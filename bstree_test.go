package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-11 14:41:40
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 18:10:37
 */

import (
	"fmt"
	"testing"

	"github.com/monitor1379/rbtree"
)

func TestBSTree(t *testing.T) {
	bstree := rbtree.NewIntBSTree()
	bstree.InsertOrReplace(10, 10)
	fmt.Println(bstree.PrettyString())

	bstree.InsertOrReplace(5, 5)
	fmt.Println(bstree.PrettyString())

	bstree.InsertOrReplace(3, 3)
	fmt.Println(bstree.PrettyString())

	bstree.InsertOrReplace(1, 1)
	fmt.Println(bstree.PrettyString())

	bstree.InsertOrReplace(15, 15)
	fmt.Println(bstree.PrettyString())

	bstree.InsertOrReplace(13, 13)
	fmt.Println(bstree.PrettyString())

	bstree.InsertOrReplace(4, 4)
	fmt.Println(bstree.PrettyString())

	fmt.Println(bstree.Search(1))
	fmt.Println(bstree.Search(2))
	fmt.Println(bstree.Search(3))
	fmt.Println(bstree.Search(4))
	fmt.Println(bstree.Search(5))
}
