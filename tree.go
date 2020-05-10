/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:44:01
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-10 15:26:44
 */

package rbtree

type Tree struct {
	root       *Node
	comparator Comparator
}

func NewTree(comparator Comparator) *Tree {
	tree := new(Tree)
	tree.comparator = comparator
	return tree
}

func NewIntTree() *Tree {
	return NewTree(&IntComparator{})
}

func NewStringTree() *Tree {
	return NewTree(&StringComparator{})
}

func (t *Tree) InsertOrReplace(key interface{}, value interface{}) {
	c, p := t.searchInsertPosition(key)

}

func (t *Tree) Delete(key interface{}) {

}

func (t *Tree) Search(key interface{}) (interface{}, bool) {
	c, _ := t.searchInsertPosition(key)
	if c == nil {
		return nil, false
	}
	return c.Value, true
}

func (t *Tree) searchInsertPosition(key interface{}) (*Node, *Node) {
	if t.root == nil {
		return nil, nil
	}

	var p *Node
	c := t.root

	for {
		s := t.comparator.Compare(key, c.Key)
		if s == -1 {
			p = c
			c = c.Left
		} else if s == 1 {
			p = c
			c = c.Right
		} else {
			return c, c.Parent
		}
		if c == nil {
			break
		}
	}

	return c, p
}
