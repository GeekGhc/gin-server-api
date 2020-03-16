package hash

import (
	"container/list"
	"gin-server-api/pkg/app"
)

type Map struct {
	slot []*list.List
}

type KV struct {
	k, v int
}

//hashMap 初始化
func HashMapConstructor() Map {
	return Map{make([]*list.List, app.HashMapSlotNum)}
}

func (m *Map) Put(k int, v int) {
	p := k % app.HashMapSlotNum
	//slot节点新建map
	if m.slot[p] == nil {
		m.slot[p] = list.New()
		m.slot[p].PushBack(KV{k, v})
		return
	}
	//遍历是否有相同k
	for x := m.slot[p].Front(); x != nil; x = x.Next() {
		//替换节点
		if x.Value.(KV).k == k {
			x.Value = KV{k, v}
			return
		}
	}
	m.slot[p].PushBack(KV{k, v})
}

func (m *Map) Get(k int) int {
	p := k % app.HashMapSlotNum
	if m.slot[p] == nil {
		return -1
	}
	for x := m.slot[p].Front(); x != nil; x = x.Next() {
		if k == x.Value.(KV).k {
			return x.Value.(KV).v
		}
	}
	return -1
}

func (m *Map) Remove(k int) {
	p := k % app.HashMapSlotNum
	if m.slot[p] == nil {
		return
	}
	for x := m.slot[p].Front(); x != nil; x = x.Next() {
		if x.Value.(KV).k == k {
			m.slot[p].Remove(x)
			return
		}
	}
}
