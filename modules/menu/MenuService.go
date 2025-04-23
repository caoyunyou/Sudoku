package menu

import (
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
)

// borderMenuGroupGenerate 带边框的按钮组生成
func borderMenuGroupGenerate(title string, btns []*widget.Button, containerSize fyne.Size, innerBorderPadding float32) *fyne.Container {

	// 整体容器
	menuGroupContainer := container.NewWithoutLayout()
	// 横向边距
	innerBorderHPadding := innerBorderPadding
	// 纵向边距
	innerBorderVPadding := 2 * innerBorderPadding
	// 内部边框大小
	innerBorderSize := fyne.NewSize(containerSize.Width-2*innerBorderHPadding, containerSize.Height-2*innerBorderVPadding)
	// 内部边框定位
	innerBorderPos := fyne.NewPos(innerBorderHPadding, innerBorderVPadding)
	// 边框容器定义
	innerBorderContainer := canvas.NewRectangle(color.Transparent)
	innerBorderContainer.StrokeColor = utils.HTML2FyneRGB(0, 0, 255)
	innerBorderContainer.StrokeWidth = 1
	innerBorderContainer.Resize(innerBorderSize)
	innerBorderContainer.Move(innerBorderPos)

	// 按钮列表
	buttons := container.New(layout.NewVBoxLayout()) // 纵向排列
	// 按钮
	for _, btn := range btns {
		buttons.Add(btn)
	}
	// 按钮列表容器：居中展示
	btnContainer := container.NewPadded(
		container.NewVBox(
			layout.NewSpacer(),
			buttons,
			layout.NewSpacer(),
		),
	)

	//按钮列表布局：距离内边框一个内边距padding，起始位置为一个padding位置,容器内部进行居中和
	btnContainer.Resize(fyne.NewSize(innerBorderSize.Width-2*innerBorderHPadding, innerBorderSize.Height-2*innerBorderVPadding))
	btnContainer.Move(innerBorderPos.Add(fyne.NewSize(innerBorderHPadding, innerBorderVPadding)))

	menuGroupContainer.Add(innerBorderContainer)
	menuGroupContainer.Add(btnContainer)

	// 不为空才进行标题设置
	if strings.TrimSpace(title) != "" {
		// 使用变量计算
		textLen := theme.TextSize() * float32(utils.StrLen(title))                             //文本长度 = 字号【有定义则使用定义的，没有就是用系统常量】*数量
		textHeight := theme.TextSize()                                                         //文本高度，按照字号来
		textSize := fyne.NewSize(textLen, textHeight)                                          //文本大小
		textPos := fyne.NewPos((containerSize.Width-textLen)/2, innerBorderPadding+textHeight) // 文本定位：中间位置
		// 文字提示容器定义
		text := canvas.NewText(title, utils.HTML2FyneRGB(0, 0, 85))
		text.Alignment = fyne.TextAlignCenter // 居中对其
		text.Resize(textSize)
		text.Move(textPos)
		// 文字背景容器定义
		textBg := canvas.NewRectangle(utils.HTML2FyneRGB(255, 255, 255))
		textBg.Resize(textSize) //暂时使用文本大小
		textBg.Move(textPos)    // 暂时使用文本定位位置
		menuGroupContainer.Add(textBg)
		menuGroupContainer.Add(text)
	}

	return menuGroupContainer
}
