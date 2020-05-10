package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:38:53
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-10 15:04:00
 */

import (
	"testing"

	"github.com/monitor1379/rbtree"
)

func TestNode(t *testing.T) {
	tree := rbtree.NewIntTree()
	tree.InsertOrReplace(10, 10)
	tree.InsertOrReplace(5, 5)
	tree.InsertOrReplace(15, 15)
	tree.InsertOrReplace(2, 2)
	tree.InsertOrReplace(8, 8)
	tree.InsertOrReplace(17, 17)

}
