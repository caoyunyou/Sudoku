package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

// AddBorder 容器增加边框
func AddBorder(c *fyne.Container, borderColor color.Color, borderWidth float32) *fyne.Container {
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeColor = borderColor
	border.StrokeWidth = borderWidth
	return container.NewStack(border, container.NewPadded(c))
}

// SetBackGroundColor 设置背景颜色
func SetBackGroundColor(c fyne.CanvasObject, bgColor color.Color) *fyne.Container {
	backGround := canvas.NewRectangle(bgColor)
	return container.NewStack(backGround, c)
}

// AddSquareBorderLine 添加对应的正方向边框线
func AddSquareBorderLine(lineColor color.Color, lineLen float32, lineWight float32, model int) *canvas.Line {
	borderLine := canvas.NewLine(lineColor)
	borderLine.StrokeWidth = lineWight
	switch model {
	case 0: // 上
		borderLine.Position1 = fyne.NewPos(0, 0)
		borderLine.Position2 = fyne.NewPos(lineLen, 0)
		break
	case 1: // 右
		borderLine.Position1 = fyne.NewPos(lineLen, 0)
		borderLine.Position2 = fyne.NewPos(lineLen, lineLen)
		break
	case 2: // 下
		borderLine.Position1 = fyne.NewPos(0, lineLen)
		borderLine.Position2 = fyne.NewPos(lineLen, lineLen)
		break
	case 3: // 左
		borderLine.Position1 = fyne.NewPos(0, 0)
		borderLine.Position2 = fyne.NewPos(0, lineLen)
		break
	default:
		break
	}
	return borderLine

}
