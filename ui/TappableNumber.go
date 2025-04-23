package ui

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// TappableNumberCell 实现可点击的圆形数字容器
type TappableNumberCell struct {
	widget.BaseWidget
	content  *fyne.Container
	circle   *canvas.Circle
	text     *canvas.Text
	onTapped func()
}

func NewTappableNumber(c color.Color, num int, size fyne.Size, tapped func()) *TappableNumberCell {
	t := &TappableNumberCell{
		circle:   canvas.NewCircle(c),
		text:     canvas.NewText(strconv.Itoa(num), color.Black),
		onTapped: tapped,
	}
	t.text.Alignment = fyne.TextAlignCenter
	t.text.TextStyle.Bold = true
	t.text.TextSize = 20
	t.text.Color = utils.HTML2FyneRGB(255, 255, 255)

	t.content = container.NewStack(container.NewGridWrap(size, t.circle), t.text)
	t.ExtendBaseWidget(t)
	//t.asyncListen()
	// 事件发布，可抽出
	globel.EventBus().Publish(event.Event{Type: event.SelectNumTap, Data: t})
	return t
}

func (t *TappableNumberCell) Circle() *canvas.Circle {
	return t.circle
}

func (t *TappableNumberCell) Text() *canvas.Text {
	return t.text
}

func (t *TappableNumberCell) asyncListen() {
	// 启动监听Goroutine
	go func() {
		for {
			ob := globel.GetDataObservable(globel.SelectedNum)
			current := ob.Get()
			ob.Lock()
			for ob.Value() == current {
				ob.Wait()
			}
			go func() {
				fyne.DoAndWait(func() {
					num, _ := strconv.Atoi(t.text.Text)
					if ob.Value() == num {
						t.SetFillColor(utils.HTML2FyneRGB(225, 137, 92))
					} else {
						t.SetFillColor(utils.HTML2FyneRGB(26, 56, 226))
					}
					t.circle.Refresh()
				})
			}()
			ob.UnLock()
		}
	}()
}

// CreateRenderer 创建图形渲染
func (t *TappableNumberCell) CreateRenderer() fyne.WidgetRenderer {
	num, _ := strconv.Atoi(t.text.Text)
	if globel.GetDataStorage(globel.SelectedNum) == num {
		t.SetFillColor(utils.HTML2FyneRGB(225, 137, 92))
	} else {
		t.SetFillColor(utils.HTML2FyneRGB(26, 56, 226))
	}
	return widget.NewSimpleRenderer(t.content)
}

func (t *TappableNumberCell) Tapped(*fyne.PointEvent) {
	if t.onTapped != nil {
		t.onTapped()
	}
}

func (t *TappableNumberCell) SetFillColor(newColor color.Color) {
	t.circle.FillColor = newColor
}
