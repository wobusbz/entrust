package entrust

import (
	"testing"
	"time"
)

func TestEntrustSlice(t *testing.T) {
	var list = new(EntrustSlice)

	*list = append(*list,
		&Entrust{
			EntrustNo:      "ABC0001",
			EntrustedPrice: 1.3,
			EntrustedTime:  time.Now(),
		},
		&Entrust{
			EntrustNo:      "ABC0002",
			EntrustedPrice: 1.2,
			EntrustedTime:  time.Now(),
		},
		&Entrust{
			EntrustNo:      "ABC0003",
			EntrustedPrice: 1.1,
			EntrustedTime:  time.Now(),
		},
		&Entrust{
			EntrustNo:      "ABC0004",
			EntrustedPrice: 1.6,
			EntrustedTime:  time.Now(),
		},
		&Entrust{
			EntrustNo:      "ABC0005",
			EntrustedPrice: 1,
			EntrustedTime:  time.Now(),
		},
	)
	list.Sort(false)
	list.InstertAssign(&Entrust{
		EntrustNo:      "ABC0006",
		EntrustedPrice: 0.9,
		EntrustedTime:  time.Now(),
	})
	list.InstertAssign(&Entrust{
		EntrustNo:      "ABC0007",
		EntrustedPrice: 0.2,
		EntrustedTime:  time.Now(),
	})
	list.InstertAssign(&Entrust{
		EntrustNo:      "ABC0008",
		EntrustedPrice: 1.5,
		EntrustedTime:  time.Now(),
	})
	list.InstertAssign(&Entrust{
		EntrustNo:      "ABC0009",
		EntrustedPrice: 1.5,
		EntrustedTime:  time.Now(),
	})
	list.InstertAssign(&Entrust{
		EntrustNo:      "ABC0010",
		EntrustedPrice: 0.1,
		EntrustedTime:  time.Now(),
	})
	list.InstertAssign(&Entrust{
		EntrustNo:      "ABC0010",
		EntrustedPrice: 0.01,
		EntrustedTime:  time.Now(),
	})
	for _, val := range *list {
		t.Log(val)

	}
	t.Log(list.Len())
}
