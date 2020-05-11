package rbtree

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-11 14:39:28
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 18:09:58
 */

// BSTree is Binary Search Tree
type BSTree struct {
	root       *Node
	comparator Comparator
}

func NewBSTree(comparator Comparator) *BSTree {
	t := new(BSTree)
	t.comparator = comparator
	return t
}

func NewIntBSTree() *BSTree {
	return NewBSTree(&IntComparator{})
}

func NewStringBSTree() *BSTree {
	return NewBSTree(&StringComparator{})
}

func (t *BSTree) InsertOrReplace(key interface{}, value interface{}) {
	c, p, comp := t.SearchInsertPosition(key)

	// replace
	if c != nil {
		c.Value = value
		return
	}

	c = new(Node)
	c.Key = key
	c.Value = value

	// if tree is empty
	if p == nil {
		t.root = c
		return
	}

	// if tree is not empty(p is not nil)
	c.Parent = p
	if comp == -1 {
		p.Left = c
	} else {
		p.Right = c
	}
}

func (t *BSTree) SearchInsertPosition(key interface{}) (*Node, *Node, int) {
	c := t.root
	var p *Node
	var comp int

	for c != nil {
		comp = t.comparator.Compare(key, c.Key)
		if comp == -1 {
			p = c
			c = c.Left
		} else if comp == 1 {
			p = c
			c = c.Right
		} else {
			// comp == 0
			return c, p, comp
		}
	}
	return c, p, comp
}

func (t *BSTree) Search(key interface{}) (interface{}, bool) {
	c, _, _ := t.SearchInsertPosition(key)
	if c == nil {
		return nil, false
	}
	return c.Value, true
}

func (t *BSTree) PrettyString() string {
	return prettyString(t.root)
}
