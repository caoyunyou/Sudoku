package globel

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/server"
	"com.cyy/sudoku/types"
)

var sudoku _types.SudokuInfo
var saveSudoku _types.SudokuInfo

// 全局维护一个事件总线
var eventBus *event.Bus

// 初始化处理
func init() {
	CreateGameByLevel(LevelList[0])
	eventBus = event.NewEventBus()
}

// CreateGameByLevel 通过等级进行对应数独游戏的创建
func CreateGameByLevel(level _types.LevelEnum) {
	sudoku.LevelInfo = level
	sudoku.CurrGame = server.GenerateSudokuPuzzle(sudoku.LevelInfo.InitSudokuNum)
	// 步数信息留存
	sudoku.Step = _types.StepInfo{Num: 0, Info: make(map[int]_types.Game)}
	sudoku.Step.Info[0] = sudoku.CurrGame

	// 对局遍历，并获取其中每个数字的个数
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.NumFillArr[sudoku.CurrGame[i][j].Num]++
		}
	}

}

// ChangeGameDataVal 数独单元格点击
func ChangeGameDataVal(i int, j int, newVal int) bool {
	// 判断对应单元格是否为初始化对局的单元格
	if !sudoku.CurrGame[i][j].IsHole {
		return false
	}
	// 如果对应的数字已满足9个，则不允许再次填入
	if sudoku.NumFillArr[newVal] == 9 {
		return false
	}
	// 进行判断
	isValid := server.IsValid(&sudoku.CurrGame, i, j, newVal)
	if isValid {
		sudoku.CurrGame[i][j].Num = newVal //设值
		sudoku.Step.Num++
		sudoku.Step.Info[sudoku.Step.Num] = sudoku.CurrGame // 步数加一,历史记录留存
		sudoku.NumFillArr[newVal]++                         //对应数字个数增加
		if sudoku.NumFillArr[newVal] == 9 {                 //如果有9个数，则发布事件
			eventBus.Publish(event.Event{Type: event.NumberFillCompleted, Data: newVal})
		}
	}
	if 1 == sudoku.Step.Num { // 如果是走了第一步，则进行计时 TODO 考虑回滚操作
		eventBus.Publish(event.Event{Type: event.TimeStart})
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
	sudoku.CurrGame = sudoku.Step.Info[sudoku.Step.Num-1]
	delete(sudoku.Step.Info, sudoku.Step.Num) // 清除对应的键值
	sudoku.Step.Num--                         //步数减一

	// 对局刷新
	eventBus.Publish(event.Event{Type: event.GameRefresh})
}

// GameRestart 游戏重新开始
func GameRestart() {
	CreateGameByLevel(sudoku.LevelInfo)
	// 发送事件
	eventBus.Publish(event.Event{Type: event.GameRefresh})
}

// GameSave 存盘
func GameSave() {
	saveSudoku = sudoku
}

// GameReStore 恢复存盘至当前对局
func GameReStore() {
	sudoku = saveSudoku
	eventBus.Publish(event.Event{Type: event.GameRefresh})
}
