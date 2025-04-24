package globel

import "com.cyy/sudoku/utils"

const (
	SelectedNum        = "SelectedNum"
	SelectedNumberCell = "SelectedNumberCell"
)

// 等级枚举设置
var LevelList = [...]LevelEnum{{LevelName: "简单", InitSudokuNum: 32, LevelColor: utils.HTML2FyneRGB(50, 173, 94)},
	{LevelName: "中间", InitSudokuNum: 30, LevelColor: utils.HTML2FyneRGB(50, 94, 214)},
	{LevelName: "高级", InitSudokuNum: 27, LevelColor: utils.HTML2FyneRGB(173, 94, 173)},
	{LevelName: "困难", InitSudokuNum: 17, LevelColor: utils.HTML2FyneRGB(253, 94, 94)}}
