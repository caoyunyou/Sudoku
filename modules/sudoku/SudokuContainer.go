package sudoku

import (
	"com.cyy/sudoku/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.WidgetRenderer = (*sudokuContainerRenderer)(nil)

// SudokuContainer 创建sudoku实体，用于重写对应的事件
type SudokuContainer struct {
	widget.BaseWidget
	content     *fyne.Container //组合容器
	hoverCircle *ui.HoverEffectCircle
}

func NewSudokuContainer() *SudokuContainer {
	s := &SudokuContainer{}
	sudokuTip := ui.NewSudokuTip(fyne.NewSize(540, 50))
	sudokuPanel := NewSudokuPanel()
	// 等级提示以及时间提示
	//levelContainer := container.NewHBox()
	//hardLevel := widget.NewLabel("难易度: ")
	//levelText := canvas.NewText("高级", utils.HTML2FyneRGB(253, 94, 94))
	//levelContainer.Add(hardLevel)
	//levelContainer.Add(levelText)
	//
	//timeContainer := container.NewHBox()
	//useTime := widget.NewLabel("时间: ")
	//timeText := canvas.NewText("00:00", utils.HTML2FyneRGB(253, 94, 94))
	//timeContainer.Add(useTime)
	//timeContainer.Add(timeText)
	//infoContainer := container.NewCenter(container.NewHBox(levelContainer, timeContainer))
	// 创建数独网格
	//sudokuGrid = utils.SetBackGroundColor(sudokuGrid, utils.HTML2FyneRGB(30, 31, 34))

	// 加一个默认的悬浮
	s.content = container.NewVBox(sudokuTip,
		sudokuPanel)

	s.ExtendBaseWidget(s)

	return s
}

// 自定义渲染器
type sudokuContainerRenderer struct {
	fyne.WidgetRenderer
	s *SudokuContainer
}

func (s *SudokuContainer) Layout(size fyne.Size) {

}

// CreateRenderer 创建渲染器
func (s *SudokuContainer) CreateRenderer() fyne.WidgetRenderer {
	return &sudokuContainerRenderer{widget.NewSimpleRenderer(s.content), s}
}

func (s *SudokuContainer) Refresh() {
	s.BaseWidget.Refresh()
}
