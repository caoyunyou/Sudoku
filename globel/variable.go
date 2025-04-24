package globel

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/server"
)

type sudokuInfo struct {
	origGame   [9][9]int // 初始对局，用作检测，其中填充数字无法进行更改
	game       [9][9]int // 对局信息
	levelInfo  LevelEnum // 等级信息
	step       stepInfo  // 当前步数
	numFillArr [10]int   // 数字填充数组
}

var sudoku sudokuInfo

// 步数信息
type stepInfo struct {
	num  int
	info map[int][9][9]int
}

// 全局维护一个事件总线
var eventBus *event.Bus

// 初始化处理
func init() {
	CreateGameByLevel(LevelList[0])
	eventBus = event.NewEventBus()
}

// CreateGameByLevel 通过等级进行对应数独游戏的创建
func CreateGameByLevel(level LevelEnum) {
	sudoku.levelInfo = level
	sudoku.origGame = server.GenerateSudokuPuzzle(sudoku.levelInfo.InitSudokuNum)
	sudoku.game = sudoku.origGame
	// 步数信息留存
	sudoku.step = stepInfo{num: 0, info: make(map[int][9][9]int)}
	sudoku.step.info[0] = sudoku.game

	// 对局遍历，并获取其中每个数字的个数
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.numFillArr[sudoku.game[i][j]]++
		}
	}

}

// ChangeGameDataVal 数独单元格点击
func ChangeGameDataVal(i int, j int, newVal int) bool {
	// 判断对应单元格是否为初始化对局的单元格
	if sudoku.origGame[i][j] != 0 {
		return false
	}
	// 如果对应的数字已满足9个，则不允许再次填入
	if sudoku.numFillArr[newVal] == 9 {
		return false
	}
	// 进行判断
	isValid := server.IsValid((*server.Sudoku)(&sudoku.game), i, j, newVal)
	if isValid {
		sudoku.game[i][j] = newVal //设值
		sudoku.step.num++
		sudoku.step.info[sudoku.step.num] = sudoku.game // 步数加一,历史记录留存
		sudoku.numFillArr[newVal]++                     //对应数字个数增加
		if sudoku.numFillArr[newVal] == 9 {             //如果有9个数，则发布事件
			eventBus.Publish(event.Event{Type: event.NumberFillCompleted, Data: newVal})
		}
	}
	if 1 == sudoku.step.num { // 如果是走了第一步，则进行计时 TODO 考虑回滚操作
		eventBus.Publish(event.Event{Type: event.TimeStart})
	}
	// 如果步数和当前等级数之和为数独总个数，则游戏胜利，发布胜利事件
	if 81 == sudoku.step.num+sudoku.levelInfo.InitSudokuNum {
		eventBus.Publish(event.Event{Type: event.GameVictory})
	}
	return isValid
}

// GetGameData 获取当前游戏数据
func GetGameData() [9][9]int {
	return sudoku.game
}

// GetGameLevel 获取当前等级信息
func GetGameLevel() LevelEnum {
	return sudoku.levelInfo
}

func GetGameDataVal(i int, j int) int {
	return sudoku.game[i][j]
}

// EventBus 获取事件总线
func EventBus() *event.Bus {
	return eventBus
}

// UndoStep 操作回滚
func UndoStep() {
	sudoku.game = sudoku.step.info[sudoku.step.num-1]
	delete(sudoku.step.info, sudoku.step.num) // 清除对应的键值
	sudoku.step.num--                         //步数减一

	// 对局刷新
	eventBus.Publish(event.Event{Type: event.GameRefresh})
}
