package rbtree_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-11 23:04:40
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-11 23:15:24
 */

import (
	"math/rand"
	"testing"
)

func BenchmarkMapInsert(b *testing.B) {
	b.ResetTimer()

	m := make(map[interface{}]interface{})
	for i := 0; i < b.N; i++ {
		m[rand.Int()] = i
	}
}

func BenchmarkMapRandomInsertAndSearch(b *testing.B) {
	m := make(map[interface{}]interface{})

	keys := []int{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, rand.Int())
	}

	for i := 0; i < b.N; i++ {
		m[keys[i]] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = m[keys[i]]
	}
}
