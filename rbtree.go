/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:44:01
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 23:07:09
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
