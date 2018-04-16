package uniq_set

import (
	"testing"
)


/**
1、文件名必须以xx_test.go命名

2、方法必须是Test[^a-z]开头

3、方法参数必须 t *testing.T

 */
func Test_BuildAndSoert(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(1)
	s.Add(10)
	s.Add(5)
	s.Add(10)

	s.Add(1)
	s.Add(1)
	s.Add(0)
	s.Add(2)
	s.Add(4)
	s.Add(3)

	s.Clear()
	if s.IsEmpty() {
		t.Log("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(6)
	s.Add(5)
	s.Add(5)

	if s.Has(2) {
		t.Log("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	t.Log("无序的切片", s.List())
	t.Log("有序的切片", s.SortList())

}
