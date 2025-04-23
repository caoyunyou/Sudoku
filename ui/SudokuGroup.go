package ui

import (
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"time"
)

// SudokuGroupIdx 数独窗格定位
type SudokuGroupIdx struct {
	x int // 位于整个数独窗体的x坐标
	y int //位于整个数独窗体的y坐标
}

type SudokuGroup struct {
	widget.BaseWidget
	content  *fyne.Container //组合容器
	groupIdx SudokuGroupIdx  // 当前容器为几号窗格
}

// NewSudokuGroup
func NewSudokuGroup(x int, y int, squareBorderArr []bool) *SudokuGroup {

	g := &SudokuGroup{groupIdx: SudokuGroupIdx{x, y}}

	// 纵向
	gridContainer := container.New(layout.NewCustomPaddedVBoxLayout(0))

	for i := 0; i < 3; i++ {
		// 横向
		lineContainer := container.New(layout.NewCustomPaddedHBoxLayout(0))

		for j := 0; j < 3; j++ {
			// 边框
			borderBoolArr := []bool{i > 0, j < 2, false, false}
			var cell *SudokuCell
			cell = NewSudokuCell(
				60, // 直径60像素
				utils.HTML2FyneRGB(255, 255, 255),
				//color.RGBA{152, 251, 251, 255}, // 背景色
				nil, // 文字颜色
				//nil,               // 文字颜色
				34,  // 字号24
				nil, // 点击回调
				borderBoolArr,
				SudokuIdx{x: i, y: j, groupIdx: 3*g.groupIdx.x + g.groupIdx.y},
			)
			lineContainer.Add(cell)
		}
		gridContainer.Add(lineContainer)
	}

	borderContainer := container.NewWithoutLayout()

	go func() {
		time.Sleep(100 * time.Millisecond) // 等待布局完成
		for index, judge := range squareBorderArr {
			if judge {
				borderContainer.Add(utils.AddSquareBorderLine(utils.HTML2FyneRGB(0, 0, 136),
					borderContainer.Size().Height,
					2,
					index))
			}
		}
	}()
	utils.SetBackGroundColor(borderContainer, utils.HTML2FyneRGB(30, 31, 34))
	g.content = borderContainer
	if ((3*x)+y)%2 == 1 { // 奇数个进行设置背景颜色
		// 设置背景颜色
		gridContainer = utils.SetBackGroundColor(gridContainer, utils.HTML2FyneRGB(230, 243, 220))
	} else {
		//g.content = borderContainer
		//g.content = borderContainer
		gridContainer = utils.SetBackGroundColor(gridContainer, utils.HTML2FyneRGB(245, 245, 245))
	}
	g.content = container.NewStack(gridContainer, borderContainer)
	g.ExtendBaseWidget(g)
	return g
}

// CreateRenderer 创建渲染器
func (c *SudokuGroup) CreateRenderer() fyne.WidgetRenderer {

	return widget.NewSimpleRenderer(c.content)
}
