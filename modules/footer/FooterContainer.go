package footer

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// ContainerGenerate 数字按钮生成
func ContainerGenerate() *fyne.Container {

	// 底部数字按钮
	numButtons := container.New(layout.NewHBoxLayout())
	for i := 1; i <= 9; i++ {
		num := i
		var tc *ui.TappableNumberCell
		tc = ui.NewTappableNumber(num, fyne.NewSize(55, 55), func() {
			// 设置全局变量，并发布事件
			globel.SetDataStorage(globel.SelectedNum, num)
			globel.EventBus().Publish(event.Event{Type: event.SelectedNumChange, Data: num})
		})
		numButtons.Add(tc)
	}
	return container.NewGridWrap(fyne.NewSize(540, 100),
		container.NewCenter(numButtons))

}
