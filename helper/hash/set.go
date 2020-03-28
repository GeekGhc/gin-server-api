package hash

import (
	"container/list"
	"gin-server-api/app"
)

type Set struct {
	slot []*list.List
}

//HashSet 初始化
func HashSetConstructor() Set {
	return Set{make([]*list.List, app.HashMapSlotNum)}
}

func (s *Set) Add(k int) {
	p := k % app.HashSetSlotNum
	if s.slot[p] == nil {
		s.slot[p] = list.New()
		s.slot[p].PushBack(k)
		return
	}
	for e := s.slot[p].Front(); e != nil; e = e.Next() {
		if e.Value == k {
			return
		}
	}
	s.slot[p].PushBack(k)
}

func (s *Set) Remove(k int) {
	p := k % app.HashSetSlotNum
	if s.slot[p] == nil {
		return
	}
	for e := s.slot[p].Front(); e != nil; e = e.Next() {
		if e.Value == k {
			s.slot[p].Remove(e)
			return
		}
	}
}

func (s *Set) Contains(k int) bool {
	p := k % app.HashSetSlotNum
	if s.slot[p] == nil {
		return false
	}
	for e := s.slot[p].Front(); e != nil; e = e.Next() {
		if e.Value == k {
			return true
		}
	}
	return false
}
