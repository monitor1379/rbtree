package rbtree

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:30:28
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-10 14:37:44
 */

type Node struct {
	Key    int
	Value  interface{}
	Color  Color
	Parent *Node
	Left   *Node
	Right  *Node
}

func (c *Node) GetGrandparent() *Node {
	if c.Parent == nil {
		return nil
	}
	return c.Parent.Parent
}

func (c *Node) GetUncle() *Node {
	g := c.GetGrandparent()
	if g == nil {
		return nil
	}
	if c.Parent == g.Left {
		return g.Right
	}
	return g.Left
}
