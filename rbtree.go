/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:44:01
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-12 11:16:01
 */

package rbtree

var (
	gpcRelationLeftLine      = 0
	gpcRelationRightLine     = 1
	gpcRelationLeftTriangle  = 2
	gpcRelationRightTriangle = 3
)

type RBTree struct {
	*BSTree
}

func NewRBTree(comparator Comparator) *RBTree {
	t := new(RBTree)
	t.BSTree = NewBSTree(comparator)
	return t
}

func NewIntRBTree() *RBTree {
	return NewRBTree(&IntComparator{})
}

func NewStringRBTree() *RBTree {
	return NewRBTree(&StringComparator{})
}

func (t *RBTree) Put(key interface{}, value interface{}) {
	t.InsertOrReplace(key, value)
}

func (t *RBTree) Get(key interface{}) (interface{}, bool) {
	return t.Search(key)
}

func (t *RBTree) Remove(key interface{}) {
	t.Delete(key)
}

func (t *RBTree) InsertOrReplace(key interface{}, value interface{}) {
	c, p, comp := t.SearchInsertPosition(key)

	// replace
	if c != nil {
		c.Value = value
		return
	}

	// insert
	t.size++

	c = new(Node)
	c.Key = key
	c.Value = value

	if p == nil {
		t.root = c
	} else {
		c.Parent = p
		if comp == -1 {
			p.Left = c
		} else {
			p.Right = c
		}
	}

	t.runCheckAndFix(c)
}

func (t *RBTree) runCheckAndFix(c *Node) {
	t.runCheckAndFixCIsRoot(c)
}

func (t *RBTree) runCheckAndFixCIsRoot(c *Node) {
	// case 1: c is root
	if c == t.root {
		c.Color = ColorBlack
		return
	}

	t.runCheckAndFixPIsBlack(c)
}

func (t *RBTree) runCheckAndFixPIsBlack(c *Node) {
	// case 2: c.parent is black
	if c.Parent.Color == ColorBlack {
		return
	}

	t.runCheckAndFixDoubleRed(c)
}

func (t *RBTree) runCheckAndFixDoubleRed(c *Node) {
	// case 3: c.uncle is red
	// note that c.parent is red now
	g := c.GetGrandparent()
	p := c.Parent
	u := c.GetUncle()

	if g == nil {
		return
	}

	if u != nil && u.Color == ColorRed {
		p.Color = ColorBlack
		u.Color = ColorBlack
		g.Color = ColorRed

		// recursive
		t.runCheckAndFix(g)
		return
	}

	t.runCheckAndFixGPCRelationTriangle(c)
}

func (t *RBTree) runCheckAndFixGPCRelationTriangle(c *Node) {
	// case 4: c.uncle is black and g/p/c relation is triangle

	// note that c.parent is red too now
	g := c.GetGrandparent()
	p := c.Parent

	gpcRelation := t.getGPCRelation(g, p, c)

	if gpcRelation == gpcRelationLeftTriangle {
		t.leftRotate(p)
		c = c.Left
	} else if gpcRelation == gpcRelationRightTriangle {
		t.rightRotate(p)
		c = c.Right
	}

	t.runCheckAndFixGPCRelationLine(c)
}

func (t *RBTree) runCheckAndFixGPCRelationLine(c *Node) {
	// case 5: c.uncle is black and g/p/c relation is line

	// note that c.parent is red too now
	g := c.GetGrandparent()
	p := c.Parent

	gpcRelation := t.getGPCRelation(g, p, c)

	if gpcRelation == gpcRelationLeftLine {
		// right rotate g and recolor p & g
		t.rightRotate(g)
	} else if gpcRelation == gpcRelationRightLine {
		// left  rotate g and recolor p & g
		t.leftRotate(g)
	}

	p.Color = ColorBlack
	g.Color = ColorRed
	if g == t.root {
		t.root = p
		p.Parent = nil
	}
}

func (t *RBTree) rightRotate(c *Node) {
	var (
		p  *Node // c.Parent
		l  *Node // c.Left
		lr *Node // c.Left.Right
	)

	p = c.Parent
	l = c.Left
	if l != nil {
		lr = l.Right
	}

	// move lr as c's left
	c.Left = lr
	if lr != nil {
		lr.Parent = c
	}

	// move l to c's position
	if p != nil {
		if c == p.Left {
			p.Left = l
		} else {
			p.Right = l
		}
		l.Parent = p
	}

	// set c as l's right
	l.Right = c
	c.Parent = l
}

func (t *RBTree) leftRotate(c *Node) {
	var (
		p  *Node // c.Parent
		r  *Node // c.Right
		rl *Node // c.Right.Left
	)

	p = c.Parent
	r = c.Right
	if r != nil {
		rl = r.Left
	}

	// move rl as c's right
	c.Right = rl
	if rl != nil {
		rl.Parent = c
	}

	// move r to c's position
	if p != nil {
		if c == p.Left {
			p.Left = r
		} else {
			p.Right = r
		}
		r.Parent = p
	}

	// set c as r's left
	r.Left = c
	c.Parent = r
}

func (t *RBTree) getGPCRelation(g, p, c *Node) int {
	if g.Left == p {
		if p.Left == c {
			return gpcRelationLeftLine
		} else {
			return gpcRelationLeftTriangle
		}
	} else {
		if p.Left == c {
			return gpcRelationRightTriangle
		} else {
			return gpcRelationRightLine
		}
	}
}

func getNodeColor(c *Node) Color {
	if c == nil {
		return ColorBlack
	}
	return c.Color
}

func (t *RBTree) Delete(key interface{}) {
	c, _, _ := t.SearchInsertPosition(key)
	if c == nil {
		return
	}

	if c.Left != nil && c.Right != nil {
		minimumSubNode := t.getMinimumSubNode(c.Right)
		c.Key = minimumSubNode.Key
		c.Value = minimumSubNode.Value
		c = minimumSubNode
	}

	var child *Node
	if c.Left == nil || c.Right == nil {
		if c.Right == nil {
			child = c.Left
		} else {
			child = c.Right
		}
		if c.Color == ColorBlack {
			c.Color = getNodeColor(child)
			t.deleteCase1(c)
		}
		t.replaceNode(c, child)
		if c.Parent == nil && child != nil {
			child.Color = ColorBlack
		}
	}
	t.size--

	/*
		var child *Node
		if c.Left != nil {
			child = c.Left
		} else {
			child = c.Right
		}

		// if c.Color == Red && child.Color == Black
		if c.Color == ColorRed {
			if child == nil || child.Color == ColorBlack {
				// remove c from the tree t
				t.replaceNode(c, child)
				return
			}
		}

		// if c.Color == Black && child.Color == Red
		if c.Color == ColorBlack {
			if child != nil && child.Color == ColorRed {
				child.Color = ColorBlack
				// remove c from the tree t
				t.replaceNode(c, child)
				return
			}
		}

		// c has at most one child, and
		// c.Color == Black && child.Color == Black
		t.deleteCase1(c)

		// remove c from the tree t
		t.replaceNode(c, child)
	*/
}

func (t *RBTree) deleteCase1(c *Node) {
	if c.Parent == nil {
		return
	}
	t.deleteCase2(c)
}

func (t *RBTree) deleteCase2(c *Node) {
	s := c.GetSibling()
	if getNodeColor(s) == ColorRed {
		c.Parent.Color = ColorRed
		s.Color = ColorBlack
		if c == c.Parent.Left {
			t.leftRotate(c.Parent)
		} else {
			t.rightRotate(c.Parent)
		}
	}
	t.deleteCase3(c)
}

func (t *RBTree) deleteCase3(c *Node) {
	s := c.GetSibling()
	if getNodeColor(c.Parent) == ColorBlack &&
		getNodeColor(s) == ColorBlack &&
		s != nil &&
		getNodeColor(s.Left) == ColorBlack &&
		getNodeColor(s.Right) == ColorBlack {
		s.Color = ColorRed
		t.deleteCase1(c.Parent)
	} else {
		t.deleteCase4(c)
	}
}

func (t *RBTree) deleteCase4(c *Node) {
	s := c.GetSibling()
	if getNodeColor(c.Parent) == ColorRed &&
		getNodeColor(s) == ColorBlack &&
		s != nil &&
		getNodeColor(s.Left) == ColorBlack &&
		getNodeColor(s.Right) == ColorBlack {
		s.Color = ColorRed
		c.Parent.Color = ColorBlack
	} else {
		t.deleteCase5(c)
	}
}

func (t *RBTree) deleteCase5(c *Node) {
	s := c.GetSibling()
	if c == c.Parent.Left &&
		getNodeColor(s) == ColorBlack &&
		s != nil &&
		getNodeColor(s.Left) == ColorRed &&
		getNodeColor(s.Right) == ColorBlack {
		s.Color = ColorRed
		s.Left.Color = ColorBlack
		t.rightRotate(s)
	} else if c == c.Parent.Right &&
		getNodeColor(s) == ColorBlack &&
		s != nil &&
		getNodeColor(s.Right) == ColorRed &&
		getNodeColor(s.Left) == ColorBlack {
		s.Color = ColorRed
		s.Right.Color = ColorBlack
		t.leftRotate(s)
	}
	t.deleteCase6(c)
}

func (t *RBTree) deleteCase6(c *Node) {
	s := c.GetSibling()
	if s == nil {
		return
	}
	s.Color = getNodeColor(c.Parent)
	c.Parent.Color = ColorBlack
	if c == c.Parent.Left && getNodeColor(s.Right) == ColorRed {
		s.Right.Color = ColorBlack
		t.leftRotate(c.Parent)
	} else if getNodeColor(s.Left) == ColorRed {
		s.Left.Color = ColorBlack
		t.rightRotate(c.Parent)
	}
}

func (t *RBTree) getMinimumSubNode(c *Node) *Node {
	for {
		if c.Left != nil {
			c = c.Left
		} else {
			break
		}
	}
	return c
}

func (t *RBTree) replaceNode(oldNode, newNode *Node) {
	if t.root == oldNode {
		t.root = newNode
	} else {
		if oldNode == oldNode.Parent.Left {
			oldNode.Parent.Left = newNode
		} else {
			oldNode.Parent.Right = newNode
		}
	}

	if newNode != nil {
		newNode.Parent = oldNode.Parent
	}

}
