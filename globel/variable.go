package globel

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/server"
)

type sudokuPanel struct {
	game      [9][9]int // 对局信息
	levelInfo LevelEnum // 等级信息
	step      int       // 当前步数
}

var sudokuValue sudokuPanel

// 全局维护一个事件总线
var eventBus *event.Bus

// 初始化处理
func init() {
	CreateGameByLevel(LevelList[0])
	eventBus = event.NewEventBus()
}

// CreateGameByLevel 通过等级进行对应数独游戏的创建
func CreateGameByLevel(level LevelEnum) {
	sudokuValue.levelInfo = level
	sudokuValue.game = server.GenerateSudokuPuzzle(sudokuValue.levelInfo.InitSudokuNum)
	sudokuValue.step = 0
}

func ChangeGameDataVal(i int, j int, newVal int) bool {
	// 进行判断
	isValid := server.IsValid((*server.Sudoku)(&sudokuValue.game), i, j, newVal)
	if isValid {
		sudokuValue.game[i][j] = newVal //设值
		sudokuValue.step++              // 步数加一
	}
	if 1 == sudokuValue.step { // 如果是走了第一步，则进行计时
		eventBus.Publish(event.Event{Type: event.TimeStart})
	}
	// 如果步数和当前等级数之和为数独总个数，则游戏胜利，发布胜利事件
	if 81 == sudokuValue.step+sudokuValue.levelInfo.InitSudokuNum {
		eventBus.Publish(event.Event{Type: event.GameVictory})
	}
	return isValid
}

// GetGameData 获取当前游戏数据
func GetGameData() [9][9]int {
	return sudokuValue.game
}

// GetGameLevel 获取当前等级信息
func GetGameLevel() LevelEnum {
	return sudokuValue.levelInfo
}

func GetGameDataVal(i int, j int) int {
	return sudokuValue.game[i][j]
}

// EventBus 获取事件总线
func EventBus() *event.Bus {
	return eventBus
}
