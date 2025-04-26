package menu

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ContainerGenerate 菜单容器生成
func ContainerGenerate() *fyne.Container {

	content := container.New(layout.NewCustomPaddedVBoxLayout(0))

	// 1.提示窗格
	tipContent := widget.NewLabel("1.滚动滑轮实现数字切换\r\n2.点击下方数字按钮区域实现数字切换\r\n3.点击数独内容区域实现数字填入和取消")
	tips := ui.NewTip("提示", fyne.NewSize(50, 50), tipContent)

	content.Add(container.NewCenter(tips))

	// 2:操作窗格
	// 整体容器
	menuContainer := container.New(layout.NewVBoxLayout())

	levelBtns := make([]*widget.Button, 0, len(globel.LevelList))

	for _, enum := range globel.LevelList {
		levelBtns = append(levelBtns, widget.NewButton(enum.LevelName, func() {
			//点击后触发
			globel.CreateGameByLevel(enum)
			// 发送事件
			globel.EventBus().Publish(event.Event{Type: event.GameRefresh})
		}))

	}

	menuContainer.Add(container.NewGridWrap(fyne.NewSize(200, 270), borderMenuGroupGenerate("新游戏", levelBtns, fyne.NewSize(200, 300), 20)))

	// 创建新容器

	handleBtns := []*widget.Button{
		{Text: "撤销", OnTapped: func() {
			globel.UndoStep()
		}},
		{Text: "重新开始", OnTapped: func() {
			globel.GameRestart()
		}},
		{Text: "打印", OnTapped: func() {
			//TODO 这个不太好整
		}},
		{Text: "保存", OnTapped: func() {
			globel.GameSave()
		}},
		{Text: "恢复", OnTapped: func() {
			globel.GameReStore()
		}},
	}

	menuContainer.Add(container.NewGridWrap(fyne.NewSize(200, 270), borderMenuGroupGenerate("", handleBtns, fyne.NewSize(200, 300), 20)))

	bottomContainer := container.NewGridWrap(fyne.NewSize(200, 540), menuContainer)
	content.Add(bottomContainer)

	return content
}
