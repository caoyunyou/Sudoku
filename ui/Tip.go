package ui

import (
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"time"
)

// 接口声明
var (
	_ desktop.Hoverable = (*Tip)(nil)
	_ fyne.Widget       = (*Tip)(nil)
	_ fyne.Tappable     = (*Tip)(nil)
)

// 变量声明
var (
	normalColor       = utils.HTML2FyneRGB(255, 255, 255) // 默认蓝色
	normalBorderColor = utils.HTML2FyneRGB(220, 223, 230)
	normalTextColor   = utils.HTML2FyneRGB(103, 105, 109)
	hoverColor        = utils.HTML2FyneRGB(236, 245, 255) // 悬浮颜色
	hoverBorderColor  = utils.HTML2FyneRGB(198, 226, 255)
	hoverTextColor    = utils.HTML2FyneRGB(71, 161, 255)
)

type Tip struct {
	widget.BaseWidget
	popUp      *widget.PopUp     // 提示框
	tipContent fyne.CanvasObject // 提示展示容器
	circle     *canvas.Circle    // 圆形
	text       *canvas.Text      //文本信息
	lastShow   time.Time         // 最后一次展示时间【用于防抖】
	isShowing  bool              // 状态：是否展示
	size       fyne.Size         // 控件大小
}

func (t *Tip) MouseIn(e *desktop.MouseEvent) {
	t.circle.FillColor = hoverColor
	t.circle.StrokeColor = hoverBorderColor
	t.text.Color = hoverTextColor
	t.circle.Refresh()
}

// 鼠标移动处理:实时计算对应的位置信息
func (t *Tip) MouseMoved(e *desktop.MouseEvent) {}

// MouseOut 鼠标移出：关闭展示
func (t *Tip) MouseOut() {
	t.circle.FillColor = normalColor
	t.circle.StrokeColor = normalBorderColor
	t.text.Color = normalTextColor
	t.circle.Refresh()
}

func NewTip(text string, size fyne.Size, tipContent fyne.CanvasObject) *Tip {
	t := &Tip{
		tipContent: tipContent,
		circle:     canvas.NewCircle(normalColor),
		text:       canvas.NewText(text, normalTextColor),
		size:       size,
	}
	t.text.Alignment = fyne.TextAlignCenter
	t.circle.StrokeWidth = 1
	t.circle.StrokeColor = normalBorderColor

	t.ExtendBaseWidget(t)
	return t
}

func (t *Tip) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewGridWrap(t.size, container.NewStack(
		t.circle,
		container.NewCenter(t.text),
	)))
}

func (t *Tip) Tapped(*fyne.PointEvent) {
	if t.popUp == nil {
		t.popUp = widget.NewPopUp(t.tipContent, fyne.CurrentApp().Driver().CanvasForObject(t))
		minSize := t.popUp.MinSize()
		// 获取当前容器在界面上的绝对定位
		pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(t)
		newX := pos.X - minSize.Width - 5 //留一点间距
		newY := pos.Y + (t.Size().Height-minSize.Height)/2
		// 窗口大小
		winSize := fyne.CurrentApp().Driver().CanvasForObject(t).Size()
		// 碰撞检测+边界调整
		if newX < 0 {
			newX = 5
		}
		if newX+minSize.Width > winSize.Width {
			newX = newX + minSize.Width - winSize.Width
		}
		if newY < 0 {
			newY = 5
		}
		if newY+minSize.Height > winSize.Height {
			newY = newY + minSize.Height - winSize.Height
		}

		t.tipContent.Resize(minSize)
		t.popUp.ShowAtPosition(fyne.NewPos(newX, newY))
	} else {
		t.popUp.Show()
	}

}
