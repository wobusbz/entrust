package entrust

import (
	"errors"
	"sort"
	"time"
)

const (
	ANNUL_ENTRUST  = -2 // 取消撮合
	FAILED_ENTRUST = -1 // 撮合失败
	STAY_ENTRUST   = 0  // 待撮合
	PART_ENTRUST   = 1  // 部分撮合
	FINISH_ENTRUST = 2  // 撮合完成
)

type Entrust struct {
	EntrustNo      string    // 委托订单号
	EntrustedPrice float64   // 委托价
	EntrustedNum   int       // 委托数量
	EntrustedTime  time.Time // 委托时间
	ProductNo      string    // 商品编码
	EntrustStatus  int       // 撮合状态
	IsSysEntrust   int       // 是否是系统单  1 是 0 否  系统单优先于所有规则撮合
}

type EntrustSlice []*Entrust

func (e EntrustSlice) Len() int           { return len(e) }
func (e EntrustSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e EntrustSlice) Less(i, j int) bool { return e[i].EntrustedPrice < e[j].EntrustedPrice }

func (e EntrustSlice) Sort(reverse bool) {
	sort.Sort(e)
	if reverse {
		n := e.Len()
		for i := 0; i < n/2; i++ {
			e[i], e[n-i-1] = e[n-i-1], e[i]
		}
	}
}

func (e *EntrustSlice) GetEntrustNo(entrustNo string) (*Entrust, int) {
	for k, _ := range *e {
		if entrustNo == (*e)[k].EntrustNo {
			return (*e)[k], k
		}
	}
	return nil, -1
}

func (e *EntrustSlice) Remove(entrustNo string) (bool, error) {
	if _, k := e.GetEntrustNo(entrustNo); k != -1 {
		(*e) = append((*e)[k:], (*e)[:k+1]...)
		return true, nil
	}
	// copy((*e)[k:], (*e)[k+1:])
	return false, errors.New("委托单不存在")
}

func (e *EntrustSlice) Append(entrust *Entrust) {
	(*e) = append((*e), entrust)
}

// 插入规则 优先根据价格低高其次时间先后
func (e *EntrustSlice) InstertAssign(entrust *Entrust) {
	if e == nil {
		e.Append(entrust)
	} else {
		var index int
		var tempArry EntrustSlice
		for index < e.Len() {
			if entrust.EntrustedPrice > (*e)[index].EntrustedPrice {
				index++
				continue
			}
			if ((*e)[index].EntrustedPrice == entrust.EntrustedPrice) && (entrust.EntrustedTime.Unix() > (*e)[index].EntrustedTime.Unix()) {
				index++
				continue
			}
			tempArry = append(append(tempArry, entrust, (*e)[index]), (*e)[index+1:]...)
			var temp = (*e)[:index]
			e.Clean()
			(*e) = append((*e), append(temp, tempArry...)...)
			break
		}
	}
}

func (e *EntrustSlice) Clean() {
	(*e) = make(EntrustSlice, 0)
}
