package globel

import (
	"com.cyy/sudoku/utils"
	"strings"
)

var (
	// DataMap 数据仓库
	DataMap map[string]*utils.Observable
)

func init() {
	DataMap = make(map[string]*utils.Observable)
}

// SetDataStorage 设置数据存储信息，如果没有就加上，有就变更
func SetDataStorage(k string, v any) bool {
	// 边界值判断
	if v == nil || utils.StrLen(strings.TrimSpace(k)) == 0 {
		return false
	}
	// 先判断是否已经存在对应的数据存储信息
	oldVal, exists := DataMap[k]
	if exists {
		// 更换为新值
		oldVal.Set(v)
	} else {
		// 进行值的设定
		//创建新的观察者
		ob := utils.NewObservable(v)
		DataMap[k] = ob
	}
	return true
}

// GetDataStorage 通过Key获取数据存储信息
func GetDataStorage(k string) any {
	val, exists := DataMap[k]
	if exists {
		return val.Get()
	} else {
		return nil
	}
}

func GetDataObservable(k string) *utils.Observable {
	val, exists := DataMap[k]
	if exists {
		return val
	} else {
		return nil
	}
}
