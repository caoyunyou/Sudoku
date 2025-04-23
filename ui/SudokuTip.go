package ui

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

type SudokuTip struct {
	widget.BaseWidget
	content   *fyne.Container
	levelText *canvas.Text
	timeText  *TimerState
}

func NewSudokuTip(size fyne.Size) *SudokuTip {
	s := &SudokuTip{}

	// 获取当前游戏等级
	currentLevel := globel.GetGameLevel()
	log.Println("level:", currentLevel)

	levelContainer := container.NewHBox()
	hardLevel := widget.NewLabel("难易度: ")
	levelText := canvas.NewText(currentLevel.LevelName, currentLevel.LevelColor)
	s.levelText = levelText
	levelContainer.Add(hardLevel)
	levelContainer.Add(levelText)

	timeContainer := container.NewHBox()

	useTime := widget.NewLabel("时间: ")
	//timeText := canvas.NewText("00:00", utils.HTML2FyneRGB(253, 94, 94))
	timeText := NewTimer()
	s.timeText = timeText
	timeContainer.Add(useTime)
	timeContainer.Add(timeText)
	infoContainer := container.NewCenter(container.NewHBox(levelContainer, timeContainer))
	s.content = container.NewGridWrap(size, infoContainer)
	s.ExtendBaseWidget(s)
	// 订阅对应的等级变更事件
	globel.EventBus().Subscribe(event.GameLevelChange, func(event event.Event) {
		go func() {
			fyne.DoAndWait(func() {
				s.Refresh()
			})
		}()
	})

	// 订阅时间开始事件
	globel.EventBus().Subscribe(event.TimeStart, func(event event.Event) {
		go func() {
			fyne.DoAndWait(func() {
				timeText.TimeStart()
			})
		}()
	})

	return s
}

func (s *SudokuTip) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s.content)
}

func (s *SudokuTip) Refresh() {
	s.BaseWidget.Refresh()
	// 自身刷新代码
	currentLevel := globel.GetGameLevel()
	// 更新对应的文本的颜色和文字信息
	s.levelText.Text = currentLevel.LevelName
	s.levelText.Color = currentLevel.LevelColor
}
