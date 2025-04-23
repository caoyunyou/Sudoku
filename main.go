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
	//test.FireWork()
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
