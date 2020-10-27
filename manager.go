package entrust

import (
	"engine/entrust"
	"engine/queue"
	"errors"
)

const (
	BUY_STATE  = 1
	SELL_STATE = 2

	ENTRUSTMAP = map[int]string{
		1: "买",
		2: "卖",
	}
)

type EntrustManager struct {
	BuyQueue  *queue.EntrustQueue
	SellQueue *queue.EntrustQueue

	LastTransactionPrice float64 // 上次成交价
	curMinPrice          float64 // 最新最低价
}

func NewEntrustManager() *EntrustManager {
	return &EntrustManager{
		BuyQueue:  queue.NewEntrustQueue(100),
		SellQueue: queue.NewEntrustQueue(100),
	}
}

func (manager *EntrustManager) Buy(enrust *entrust.Entrust) {
	manager.BuyQueue.Push(enrust)
}

func (manager *EntrustManager) Sell(enrust *entrust.Entrust) {
	manager.SellQueue.Push(enrust)
}

func (manager *EntrustManager) MatchmakingTrading() {

}

func (manager *EntrustManager) RecovEntrust(entrustState int, entrustNo, productNO string) (bool, error) {
	entrustStates, ok := ENTRUSTMAP[entrustState]
	if !ok {
		return false, errors.New("invalid entrustState")
	}
	if entrustStates == BUY_STATE {
		return manager.recovEntrust(manager.BuyQueue, entrustNo, productNO)
	} else {
		return manager.recovEntrust(manager.SellQueue, entrustNo, productNO)
	}
}

func (manager *EntrustManager) recovEntrust(q *queue.EntrustQueue, entrustNo, productNO string) (bool, error) {
	entrust := q.GetKey(productNO)
	if entrust == nil {
		return false, errors.New("produc no not found")
	}
	return entrust.Remove(entrustNo)
}

func (manager *EntrustManager) ManualEntrust(entrustState int, productNO string, entrustNo []string) {

}
