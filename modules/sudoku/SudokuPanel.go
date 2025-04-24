package sudoku

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/ui"
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var _ desktop.Hoverable = (*SudokuPanel)(nil)
var _ desktop.Cursorable = (*SudokuPanel)(nil)
var _ fyne.WidgetRenderer = (*panelRenderer)(nil)

// SudokuPanel 创建sudoku实体，用于重写对应的事件
type SudokuPanel struct {
	widget.BaseWidget
	content       *fyne.Container //组合容器
	hoverCircle   *ui.HoverEffectCircle
	fireWorkGroup *ui.FireWorkGroup // 烟花组
}

func NewSudokuPanel() *SudokuPanel {
	s := &SudokuPanel{}
	// 创建数独网格
	sudokuGrid := container.New(layout.NewCustomPaddedVBoxLayout(0))
	for i := 0; i < 3; i++ {
		lineGrid := container.New(layout.NewCustomPaddedHBoxLayout(0))
		for j := 0; j < 3; j++ {
			borderBoolArr := []bool{i == 0, true, true, j == 0}
			sudokuGroup := ui.NewSudokuGroup(i, j, borderBoolArr)
			lineGrid.Add(sudokuGroup)
		}
		sudokuGrid.Add(lineGrid)
	}
	//sudokuGrid = utils.SetBackGroundColor(sudokuGrid, utils.HTML2FyneRGB(30, 31, 34))

	// 加一个默认的悬浮
	s.hoverCircle = ui.NewHoverEffect(fyne.NewSize(30, 30),
		utils.HTML2FyneRGB(176, 157, 121),
		strconv.Itoa(globel.GetDataStorage(globel.SelectedNum).(int)))

	sudokuGrid.Resize(fyne.NewSize(540, 540))
	s.content = container.NewWithoutLayout(sudokuGrid, //边框占掉了一些距离 TODO 后面整一个弹性布局
		s.hoverCircle)

	fireWorkGroup := ui.NewFireWorkGroup(3, 300)
	s.fireWorkGroup = fireWorkGroup
	s.content = container.NewStack(s.content, fireWorkGroup)

	s.ExtendBaseWidget(s)
	s.eventSubscribe()

	return s
}

// 自定义渲染器
type panelRenderer struct {
	fyne.WidgetRenderer
	s *SudokuPanel
}

func (s *SudokuPanel) Layout(size fyne.Size) {

}

// CreateRenderer 创建渲染器
func (s *SudokuPanel) CreateRenderer() fyne.WidgetRenderer {
	return &panelRenderer{widget.NewSimpleRenderer(s.content), s}
}

// Scrolled 滚动事件
func (s *SudokuPanel) Scrolled(e *fyne.ScrollEvent) {
	if e.Scrolled.DY > 0 {
		if globel.GetDataStorage(globel.SelectedNum).(int) == 1 {
			globel.SetDataStorage(globel.SelectedNum, 9)
		} else {
			globel.SetDataStorage(globel.SelectedNum, globel.GetDataStorage(globel.SelectedNum).(int)-1)
		}
		//强制刷新一次
		globel.EventBus().Publish(event.Event{Type: event.SelectedNumChange, Data: globel.GetDataStorage(globel.SelectedNum).(int)})
	} else if e.Scrolled.DY < 0 {
		if globel.GetDataStorage(globel.SelectedNum).(int) == 9 {
			globel.SetDataStorage(globel.SelectedNum, 1)
		} else {
			globel.SetDataStorage(globel.SelectedNum, globel.GetDataStorage(globel.SelectedNum).(int)+1)
		}
		globel.EventBus().Publish(event.Event{Type: event.SelectedNumChange, Data: globel.GetDataStorage(globel.SelectedNum).(int)})
	}
}

func (s *SudokuPanel) MouseIn(e *desktop.MouseEvent) {
	// 悬浮圆形容器进入事件处理
	s.hoverCircle.MouseIn(e)

}

func (s *SudokuPanel) MouseMoved(e *desktop.MouseEvent) {
	s.hoverCircle.MouseMoved(e)
}
func (s *SudokuPanel) MouseOut() {
	s.hoverCircle.MouseOut()
}

// Cursor 鼠标展示判断【这个是实时计算的】
func (s *SudokuPanel) Cursor() desktop.Cursor {
	if s.hoverCircle.Visible() {
		return desktop.HiddenCursor
	}
	return desktop.DefaultCursor
}

func (s *SudokuPanel) Refresh() {
	s.BaseWidget.Refresh()
}

// 事件订阅
func (s *SudokuPanel) eventSubscribe() {
	// 事件订阅 游戏等级变更
	globel.EventBus().Subscribe(event.GameRefresh, func(event event.Event) {
		go func() {
			fyne.DoAndWait(func() {
				s.Refresh()
			})
		}()
	})

	//事件订阅：游戏胜利展示小特效
	// TODO 优化展示，
	globel.EventBus().Subscribe(event.GameVictory, func(e event.Event) {
		// 发布事件：终止时间
		globel.EventBus().Publish(event.Event{Type: event.TimeStop})
		go func() {
			fyne.DoAndWait(func() {
				// 中心点展示烟花特效
				s.fireWorkGroup.Start(fyne.NewPos(s.content.Size().Width/2, s.content.Size().Height/2))
			})
		}()
	})
}
