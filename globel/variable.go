package globel

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/server"
	"com.cyy/sudoku/types"
	"strconv"
)

// 接口声明
var (
	sudoku     *_types.SudokuInfo
	saveSudoku *_types.SudokuInfo
)

// 变量声明
var (
	// 全局维护一个事件总线
	eventBus *event.Bus
)

// 初始化处理
func init() {
	CreateGameByLevel(LevelList[0])
	eventBus = event.NewEventBus()
}

// CreateGameByLevel 通过等级进行对应数独游戏的创建
func CreateGameByLevel(level _types.LevelEnum) {
	if sudoku == nil {
		sudoku = &_types.SudokuInfo{}
	}
	sudoku.LevelInfo = level
	sudoku.CurrGame = server.GenerateSudokuPuzzle(sudoku.LevelInfo.InitSudokuNum)
	// 步数信息留存
	sudoku.Step = _types.StepInfo{Num: 0, Info: make(map[int]*_types.GameDump)}

	// 数字重置处理
	for idx, _ := range sudoku.NumFillArr {
		sudoku.NumFillArr[idx] = 0
	}
	// 对局遍历，并获取其中每个数字的个数
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.NumFillArr[sudoku.CurrGame[i][j].Num]++
		}
	}

}

// ChangeGameDataVal 数独单元格点击
func ChangeGameDataVal(i int, j int, newVal int) bool {
	// 数字格式庞大暖
	if newVal < 1 || newVal > 9 {
		return false
	}
	// 判断对应单元格是否为初始化对局的单元格
	if !sudoku.CurrGame[i][j].IsHole {
		return false
	}
	oldVal := sudoku.CurrGame[i][j].Num
	// 如果对应的数字已满足9个，则不允许再次填入
	if sudoku.NumFillArr[newVal] == 9 && oldVal != newVal {
		return false
	}
	// 设置为false
	isValid := false

	// 如果选择的数字和要填的数字一致，则直接取消该数字
	if oldVal == newVal {
		newVal = 0
		isValid = true
	} else {
		// 进行判断
		isValid = server.IsValid(&sudoku.CurrGame, i, j, newVal)
	}
	if isValid {
		sudoku.CurrGame[i][j].Num = newVal //设值
		sudoku.Step.Num++
		//sudoku.Step.Info[sudoku.Step.Num] = sudoku.CurrGame // 步数加一,历史记录留存
		sudoku.Step.Info[sudoku.Step.Num] = &_types.GameDump{
			X:      i,
			Y:      j,
			OldVal: oldVal,
			NewVal: newVal,
		} // 步数加一,历史记录留存

		if oldVal == newVal { //填入相同的数字
			// 不做处理
		} else {
			// 对应数字回退处理，恢复原色
			if sudoku.NumFillArr[oldVal] == 9 && oldVal != 0 {
				eventBus.Publish(event.Event{Type: event.NumberFillRollback + strconv.Itoa(oldVal), Data: newVal})
			}
			sudoku.NumFillArr[oldVal]-- //对应原来的数字个数减少
			sudoku.NumFillArr[newVal]++ //对应数字个数增加
			//如果有9个数，则发布事件
			if sudoku.NumFillArr[newVal] == 9 && newVal != 0 {
				eventBus.Publish(event.Event{Type: event.NumberFillCompleted + strconv.Itoa(newVal), Data: newVal})
			}
		}

	}

	if 1 == sudoku.Step.Num { // 如果是走了第一步，则进行计时 TODO 考虑回滚操作
		eventBus.Publish(event.Event{Type: event.TimeReStart})
	}
	// 如果步数和当前等级数之和为数独总个数，则游戏胜利，发布胜利事件
	if 81 == sudoku.Step.Num+sudoku.LevelInfo.InitSudokuNum {
		eventBus.Publish(event.Event{Type: event.GameVictory})
	}
	return isValid
}

// GetGameData 获取当前游戏数据
func GetGameData() _types.Game {
	return sudoku.CurrGame
}

// GetGameLevel 获取当前等级信息
func GetGameLevel() _types.LevelEnum {
	return sudoku.LevelInfo
}

func GetGameDataVal(i int, j int) _types.GameCell {
	return sudoku.CurrGame[i][j]
}

// EventBus 获取事件总线
func EventBus() *event.Bus {
	return eventBus
}

// UndoStep 操作回滚
func UndoStep() {
	if sudoku.Step.Num == 0 { //初始盘不能撤销
		return
	}
	// 当前步骤的信息
	stepChange := sudoku.Step.Info[sudoku.Step.Num]
	sudoku.CurrGame[stepChange.X][stepChange.Y].Num = stepChange.OldVal
	//sudoku.CurrGame = sudoku.Step.Info[sudoku.Step.Num-1]
	delete(sudoku.Step.Info, sudoku.Step.Num) // 清除对应的键值
	sudoku.Step.Num--                         //步数减一
	sudoku.NumFillArr[stepChange.OldVal]++    //对应数字进行增减
	sudoku.NumFillArr[stepChange.NewVal]--    //对应数字进行增减

	// 对局刷新
	eventBus.Publish(event.Event{Type: event.GameUndoStep + strconv.Itoa(stepChange.X) + strconv.Itoa(stepChange.Y)})
}

// GameRestart 游戏重新开始
func GameRestart() {
	CreateGameByLevel(sudoku.LevelInfo)
	// 发送事件
	eventBus.Publish(event.Event{Type: event.GameRefresh})
}

// GameSave 存盘
func GameSave() {
	if saveSudoku == nil {
		saveSudoku = &_types.SudokuInfo{}
	}
	*saveSudoku = *sudoku

}

// GameReStore 恢复存盘至当前对局
func GameReStore() {
	if saveSudoku == nil {
		return
	}
	*sudoku = *saveSudoku
	eventBus.Publish(event.Event{Type: event.GameRefresh})
}

// 获取数字填充的数量
func NumberFillQuantity(num int) int {
	return sudoku.NumFillArr[num]
}
