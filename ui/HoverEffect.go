package ui

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

var _ desktop.Hoverable = (*HoverEffectCircle)(nil)

// HoverEffectCircle 自定义悬浮效果组件
type HoverEffectCircle struct {
	widget.BaseWidget
	circle      *canvas.Circle  // 圆形组件
	size        fyne.Size       // 大小
	text        *canvas.Text    //显示文本信息
	content     *fyne.Container // 展示容器，用作渲染
	pos         fyne.Position   // 相对于父元素的定位
	visible     bool            // 是否可见
	cirCleColor color.Color     // 圆形颜色
	textColor   color.Color     // 文字颜色
}

func NewHoverEffect(
	size fyne.Size, // 大小
	circleColor color.Color, // 圆形背景颜色
	text string, //显示文本信息
) *HoverEffectCircle {
	h := &HoverEffectCircle{
		circle:      canvas.NewCircle(color.Transparent),
		text:        canvas.NewText(text, color.Transparent),
		size:        size,
		cirCleColor: circleColor,
		textColor:   utils.HTML2FyneRGB(0, 0, 255),
	}
	h.circle.Resize(size)
	h.text.Alignment = fyne.TextAlignCenter
	h.text.TextSize = 14
	// 初始设置处理：将文字和圆形都设置为透明
	h.visible = false

	h.content = container.NewGridWrap(h.size, container.NewStack(container.NewGridWrap(h.size, h.circle), container.NewCenter(h.text)))
	h.ExtendBaseWidget(h)
	h.eventSubscribe()
	return h
}

// CreateRenderer 创建渲染器
func (h *HoverEffectCircle) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(h.content)
}

// MouseMoved 实现鼠标事件处理
func (h *HoverEffectCircle) MouseMoved(e *desktop.MouseEvent) {
	// 坐标转换为相对组件的位置
	h.pos = e.Position.Subtract(fyne.NewPos(
		h.circle.Size().Width/2,
		h.circle.Size().Height/2,
	))
	h.Move(h.pos)
	h.Refresh()
}

func (h *HoverEffectCircle) MouseIn(e *desktop.MouseEvent) {
	h.visible = true
	// 颜色恢复
	h.circle.FillColor = h.cirCleColor
	h.text.Color = h.textColor

	h.Show()
	// 坐标转换为相对组件的位置
	h.pos = e.Position.Subtract(fyne.NewPos(
		h.circle.Size().Width/2,
		h.circle.Size().Height/2,
	))
	h.Move(h.pos)
	h.Refresh()
}

func (h *HoverEffectCircle) MouseOut() {
	h.visible = false
	h.Hide()
	h.Refresh()
}

func (h *HoverEffectCircle) IsVisible() bool {
	return h.visible
}

func (h *HoverEffectCircle) eventSubscribe() {
	globel.EventBus().Subscribe(event.SelectedNumChange, func(event event.Event) {
		selectedNum := event.Data.(int)
		go func() {
			fyne.DoAndWait(func() {
				h.text.Text = strconv.Itoa(selectedNum)
			})
		}()
	})
}
