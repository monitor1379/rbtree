package rbtree

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:30:28
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 22:27:15
 */

const (

	// pretty print tree
	pipeString   = "│"
	lineString   = "──"
	branchString = "├"
	cornerString = "└"
	spaceString  = " "
)

type Node struct {
	Key    interface{}
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

func (c *Node) GetBrother() *Node {
	if c.Parent == nil {
		return nil
	}
	if c == c.Parent.Left {
		return c.Parent.Right
	}
	return c.Parent.Left
}

func printTree(node *Node, n, i int, prefix string, output io.Writer) {
	var curPrefix string
	var subPrefix string

	if i < n-1 {
		curPrefix = prefix + branchString + lineString
		subPrefix = prefix + pipeString + strings.Repeat(spaceString, 4)
	} else {
		curPrefix = prefix + cornerString + lineString
		subPrefix = prefix + strings.Repeat(spaceString, 5)
	}

	if node == nil {
		output.Write([]byte(fmt.Sprintf("%s> <nil>\n", curPrefix)))
	} else {
		output.Write([]byte(fmt.Sprintf("%s> {Node Key=%v Value=%v Color=%v}\n", curPrefix, node.Key, node.Value, node.Color)))
	}

	if node != nil {
		printTree(node.Left, 2, 0, subPrefix, output)
		printTree(node.Right, 2, 1, subPrefix, output)
	}
}

func prettyString(node *Node) string {
	buf := bytes.Buffer{}

	if node == nil {
		buf.Write([]byte("<nil>\n"))
	} else {
		buf.Write([]byte(fmt.Sprintf("{Node Key=%v Value=%v Color=%v} \n", node.Key, node.Value, node.Color)))
	}

	printTree(node.Left, 2, 0, "", &buf)
	printTree(node.Right, 2, 1, "", &buf)
	return buf.String()
}
