package uniq_set

import (
	"sync"
	"sort"
)


/**
   Golang 实现无重复元素的队列

 */

type UniqSet struct {
	mSet map[int]bool
	sync.RWMutex
}

func New() *UniqSet {

	return &UniqSet{
		mSet: map[int]bool{},
	}
}

func (this *UniqSet) Add(item int) {
	this.RLock()
	defer this.RUnlock()
	this.mSet[item] = true

}

func (this *UniqSet)Remove(item int) {
	this.RLock()
	defer this.RUnlock()
	delete(this.mSet, item)
}

func (this *UniqSet)Has(item int) bool {
	this.RLock()
	defer this.RUnlock()
	_, ok := this.mSet[item]
	return ok
}

func (this *UniqSet) Len() int {
	return len(this.List())
}

func (this *UniqSet) Clear() {
	this.RLock()
	this.RUnlock()
	this.mSet = map[int]bool{}
}

func (this *UniqSet) IsEmpty() bool {
	if this.Len() == 0 {
		return true
	}
	return false
}

func (this *UniqSet) List() []int {
	this.RLock()
	defer this.RUnlock()
	list := []int{}
	for item, _ := range this.mSet {
		list = append(list, item)
	}

	return list
}
func (this *UniqSet) SortList() []int {
	this.RLock()
	defer this.RUnlock()

	list := []int{}
	for item, _ := range this.mSet {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

