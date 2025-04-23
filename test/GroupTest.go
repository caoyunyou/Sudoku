package test

import (
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func CreateSudoku() *fyne.Container {

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			cellGroupV := container.New(layout.NewCustomPaddedVBoxLayout(0))
			// 创建9个
			for i := 0; i < 3; i++ {
				cellGroupH := container.New(layout.NewCustomPaddedHBoxLayout(0))
				for j := 0; j < 3; j++ {
					// 圆形组件设置
					circle := canvas.NewCircle(color.Transparent)               // 圆形
					circle.Resize(fyne.NewSize(60, 60))                         //大小确定
					text := canvas.NewText("1", utils.HTML2FyneRGB(30, 31, 34)) // 文字
					cellS := canvas.NewRectangle(color.Transparent)
					cellS.Resize(fyne.NewSize(60, 60))
					cell := container.NewStack(cellS, circle, text)
					// 添加到横向容器中
					cellGroupH.Add(cell)
				}
				cellGroupV.Add(cellGroupH)
			}
		}
	}

	return nil

}
