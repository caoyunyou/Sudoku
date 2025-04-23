package main

import (
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/modules/footer"
	"com.cyy/sudoku/modules/menu"
	"com.cyy/sudoku/modules/sudoku"
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	setupDesktop()
	//test()
}

func test() {
	//bus := event.NewEventBus()
	//
	//// 订阅 "user.login" 事件
	//bus.Subscribe("user.login", func(event event.Event) {
	//	fmt.Printf("Received event: %s, data: %v\n", event.Type, event.Data)
	//})
	//
	//// 发布事件
	//bus.Publish(event.Event{
	//	Type: "user.login",
	//	Data: map[string]string{"username": "john"},
	//})
	//
	//// 保持主 goroutine 运行一段时间（否则程序会立即退出）
	//select {}
}

// setupDesktop 程序启动
func setupDesktop() {
	a := app.New()
	w := a.NewWindow("Sudoku")
	w.Resize(fyne.NewSize(740, 640))
	// 设置值
	globel.SetDataStorage(globel.SelectedNum, 1)

	// 左侧容器:数独网格
	leftContainer := sudoku.NewSudokuContainer()

	// 右侧难度按钮
	rightContainer := menu.ContainerGenerate()

	// 底部数字键
	bottomContainer := footer.ContainerGenerate()

	mainContainer := container.NewBorder(
		nil,
		bottomContainer,
		nil,
		rightContainer,
		leftContainer,
	)
	mainContainer = utils.AddBorder(mainContainer, utils.HTML2FyneRGB(187, 187, 187), 0.5)

	w.CenterOnScreen() // 窗口居中展示
	w.SetContent(mainContainer)
	w.ShowAndRun()
}
