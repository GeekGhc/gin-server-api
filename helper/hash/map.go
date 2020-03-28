package hash

import (
	"container/list"
	"gin-server-api/app"
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
	for e := m.slot[p].Front(); e != nil; e = e.Next() {
		//替换节点
		if e.Value.(KV).k == k {
			e.Value = KV{k, v}
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
	for e := m.slot[p].Front(); e != nil; e = e.Next() {
		if k == e.Value.(KV).k {
			return e.Value.(KV).v
		}
	}
	return -1
}

func (m *Map) Remove(k int) {
	p := k % app.HashMapSlotNum
	if m.slot[p] == nil {
		return
	}
	for e := m.slot[p].Front(); e != nil; e = e.Next() {
		if e.Value.(KV).k == k {
			m.slot[p].Remove(e)
			return
		}
	}
}
