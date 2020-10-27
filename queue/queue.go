package queue

import (
	"engine/entrust"
	"sync"
)

type EntrustQueue struct {
	me              sync.Mutex
	entrustSliceMap map[string]*entrust.EntrustSlice
}

func NewEntrustQueue(size int) *EntrustQueue {
	return &EntrustQueue{
		entrustSliceMap: make(map[string]*entrust.EntrustSlice, size),
	}
}

func (queue *EntrustQueue) GetKey(k string) *entrust.EntrustSlice {
	queue.me.Lock()
	defer queue.me.Unlock()
	if val, ok := queue.entrustSliceMap[k]; ok {
		return val
	}
	return nil
}

func (queue *EntrustQueue) Push(entru *entrust.Entrust) {
	queue.me.Lock()
	queue.entrustSliceMap[entru.ProductNo].InstertAssign(entru)
	queue.me.Unlock()
}
