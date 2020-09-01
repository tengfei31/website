/*
 * @Author: wtf
 * @Date: 2020-09-01 16:00:52
 * @LastEditors: wtf
 * @LastEditTime: 2020-09-01 16:05:07
 * @Description: plase write Description
 */
package data

import "testing"

const url = "http://github.com/tengfei31/website"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("test.Add error")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}