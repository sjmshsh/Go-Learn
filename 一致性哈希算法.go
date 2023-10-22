package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type ConsistenceHash struct {
	nodesMap        map[uint32]string // hash slot和虚拟node的映射关系
	nodesSlots      slots             // 虚拟node所有hash slot组成的切片
	NumVirtualNodes int               // 为每台机器在hash圆环上创建多少个虚拟Node
	lock            sync.Mutex        // 加锁, 实现并发安全
}

// 使用sort.Sort函数, 传入的参数需要实现的接口
type slots []uint32

func (s slots) Len() int {
	return len(s)
}

// Less 从小到大
func (s slots) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s slots) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 通过crc32函数计算散列值
func (h *ConsistenceHash) hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// 所有虚拟Node映射到hash slot排序后保存到切片
func (h *ConsistenceHash) sortNodesSlots() {
	// 先整体赋值
	slots := h.nodesSlots[:]
	for slot := range h.nodesMap {
		slots = append(slots, slot)
	}
	sort.Sort(slots)
	h.nodesSlots = slots
}

// AddNode 集群中添加机器
func (h *ConsistenceHash) AddNode(addr string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	// 根据定义的数量生成虚拟Node
	// addr加上不同的后缀计算散列得到每个虚拟Node的hash slot
	// 同一个机器的所有hash slot最终都指向同一个ip/port
	for i := 0; i < h.NumVirtualNodes; i++ {
		slot := h.hash(fmt.Sprintf("%s%d", addr, i))
		h.nodesMap[slot] = addr
	}
	h.sortNodesSlots()
}

// DeleteNode 从集群中摘除机器
func (h *ConsistenceHash) DeleteNode(addr string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	// 删除所有的虚拟节点
	for i := 0; i < h.NumVirtualNodes; i++ {
		slot := h.hash(fmt.Sprintf("%s%d", addr, i))
		delete(h.nodesMap, slot)
	}
	h.sortNodesSlots()
}

// SearchNode 查找用于读写某个key的Node
func (h *ConsistenceHash) SearchNode(key string) (string, error) {
	slot := h.hash(key)
	// 使用sort包的二分查找
	// 使用二分查找找到大于等于h1的最小值来确定虚拟Node
	f := func(x int) bool {
		return h.nodesSlots[x] >= slot
	}
	index := sort.Search(len(h.nodesSlots), f)
	if index >= len(h.nodesSlots) {
		index = 0
	}
	if addr, ok := h.nodesMap[h.nodesSlots[index]]; ok {
		return addr, nil
	} else {
		return addr, errors.New("not found")
	}
}
