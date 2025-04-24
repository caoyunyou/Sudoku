package _types

import "image/color"

// 全局类型定义

type LevelEnum struct {
	LevelName     string
	InitSudokuNum int
	LevelColor    color.Color
}

// GameCell 数字项
type GameCell struct {
	Num    int  // 1-9的数字
	IsHole bool //是否为挖孔数字
}

// Game 游戏对局类型
type Game [9][9]GameCell

// SudokuInfo 数独信息
type SudokuInfo struct {
	CurrGame   Game      // 当前对局
	LevelInfo  LevelEnum // 等级信息
	Step       StepInfo  // 当前步数
	NumFillArr [10]int   // 数字填充数组
}

// GameDump 游戏快照信息，最后的游戏
type GameDump struct {
	X      int // 横坐标
	Y      int // 纵坐标
	OldVal int // 原来的值
	NewVal int //现在的值
}

// StepInfo 步数信息
type StepInfo struct {
	Num  int
	Info map[int]*GameDump
}
