package rbtree

import (
	"errors"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-10 14:49:53
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-10 15:01:32
 */

var (
	ErrInvalidInputTypeOfComparator = errors.New("invalid input type of comparator")
)

// Comparetor.Compare returns an integer comparing two object lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//
type Comparator interface {
	Compare(interface{}, interface{}) int
}

type IntComparator struct{}

func (c *IntComparator) Compare(i, j interface{}) int {
	ii, ok := i.(int)
	if !ok {
		panic(ErrInvalidInputTypeOfComparator)
	}

	jj, ok := j.(int)
	if !ok {
		panic(ErrInvalidInputTypeOfComparator)
	}

	if ii < jj {
		return -1
	} else if ii == jj {
		return 0
	}
	return 1
}

type StringComparator struct{}

func (c *StringComparator) Compare(i, j interface{}) int {
	ii, ok := i.(string)
	if !ok {
		panic(ErrInvalidInputTypeOfComparator)
	}

	jj, ok := j.(string)
	if !ok {
		panic(ErrInvalidInputTypeOfComparator)
	}

	if ii < jj {
		return -1
	} else if ii == jj {
		return 0
	}
	return 1
}
