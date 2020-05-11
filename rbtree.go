/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:44:01
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 21:20:01
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
	// case 1: c is root
	if c == t.root {
		c.Color = ColorBlack
		return
	}

	// case 2: c.parent is black
	if c.Parent.Color == ColorBlack {
		return
	}

	g := c.GetGrandparent()
	p := c.Parent
	u := c.GetUncle()

	// case 3: c.uncle is red
	// note that c.parent is red now
	if u != nil && u.Color == ColorRed {
		p.Color = ColorBlack
		u.Color = ColorBlack
		g.Color = ColorRed
		t.runCheckAndFix(g)
		return
	}

	// case 4: c.uncle is black
	// note that c.parent is red too now
	gpcRelation := t.getGPCRelation(g, p, c)
	if gpcRelation == gpcRelationLeftLine {
		// right rotate g
		t.rightRotate(g)
		if g == t.root {
			t.root = p
		}
		// recolor p and g
		p.Color = ColorBlack
		g.Color = ColorRed

	} else if gpcRelation == gpcRelationRightLine {
		// left  rotate g
		t.leftRotate(g)
		if g == t.root {
			t.root = p
		}
		// recolor p and g
		p.Color = ColorBlack
		g.Color = ColorRed

	} else if gpcRelation == gpcRelationLeftTriangle {
		// left rorate p
	} else {
		// right rotate p
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
	if c.Parent != p || p.Parent != g {
		panic("invalid g p c relation")
	}

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
