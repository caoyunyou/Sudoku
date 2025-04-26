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

// 接口实现
var (
	_ fyne.Tappable = (*TappableNumberCell)(nil)
	_ fyne.Widget   = (*TappableNumberCell)(nil)
)

// 变量声明
var (
	dColor = utils.HTML2FyneRGB(26, 56, 226)
	sColor = utils.HTML2FyneRGB(225, 137, 92)
)

// TappableNumberCell 实现可点击的圆形数字容器
type TappableNumberCell struct {
	widget.BaseWidget
	content       *fyne.Container //组合容器
	circle        *canvas.Circle  //圆形轮廓
	text          *canvas.Text    // 显示文本
	onTapped      func()          //点击事件
	defaultColor  color.RGBA      // 默认颜色
	selectedColor color.RGBA      //选中颜色
	isSelected    bool            // 是否选中
}

func NewTappableNumber(num int, size fyne.Size, tapped func()) *TappableNumberCell {
	t := &TappableNumberCell{
		circle:   canvas.NewCircle(dColor),
		text:     canvas.NewText(strconv.Itoa(num), color.Black),
		onTapped: tapped,
	}
	t.text.Alignment = fyne.TextAlignCenter
	t.text.TextStyle.Bold = true
	t.text.TextSize = 20
	t.text.Color = utils.HTML2FyneRGB(255, 255, 255)
	t.defaultColor = dColor
	t.selectedColor = sColor
	t.isSelected = false

	t.content = container.NewStack(container.NewGridWrap(size, t.circle), t.text)
	t.drawHandle() //画图公共部分
	t.ExtendBaseWidget(t)
	// 事件订阅
	t.eventSubscribe()
	return t
}

func (t *TappableNumberCell) Circle() *canvas.Circle {
	return t.circle
}

func (t *TappableNumberCell) Text() *canvas.Text {
	return t.text
}

// CreateRenderer 创建图形渲染
func (t *TappableNumberCell) CreateRenderer() fyne.WidgetRenderer {
	//num, _ := strconv.Atoi(t.text.Text)
	//if globel.GetDataStorage(globel.SelectedNum) == num {
	//	t.ToSelectedStatus()
	//} else {
	//	t.ToDefaultStatus()
	//}
	return widget.NewSimpleRenderer(t.content)
}

func (t *TappableNumberCell) Tapped(*fyne.PointEvent) {
	if t.onTapped != nil {
		t.onTapped()
	}
}

func (t *TappableNumberCell) Refresh() {
	t.BaseWidget.Refresh()
	// 刷新处理
	t.drawHandle()
}

// ToDefaultStatus 转换成默认状态
func (t *TappableNumberCell) ToDefaultStatus() {
	t.circle.FillColor = t.defaultColor
	t.isSelected = false
}

// ToSelectedStatus 转换成选中状态
func (t *TappableNumberCell) ToSelectedStatus() {
	t.circle.FillColor = t.selectedColor
	t.isSelected = true
}

// 事件订阅
func (t *TappableNumberCell) eventSubscribe() {
	// 事件订阅：数字填满
	globel.EventBus().Subscribe(event.NumberFillCompleted+t.text.Text, func(event event.Event) {
		go func() {
			fyne.DoAndWait(func() {
				// 进行外观上的变更:透明度降低
				t.selectedColor = color.RGBA{R: t.selectedColor.R, G: t.selectedColor.G, B: t.selectedColor.B, A: 180}
				t.defaultColor = color.RGBA{R: t.defaultColor.R, G: t.defaultColor.G, B: t.defaultColor.B, A: 120}
				//通过状态再次进行颜色转变
				if t.isSelected {
					t.ToSelectedStatus()
				} else {
					t.ToDefaultStatus()
				}
			})
		}()
	})

	globel.EventBus().Subscribe(event.NumberFillRollback+t.text.Text, func(event event.Event) {
		go func() {
			fyne.DoAndWait(func() {
				// 进行外观上的变更:透明度降低
				t.selectedColor = sColor
				t.defaultColor = dColor
				//通过状态再次进行颜色转变
				if t.isSelected {
					t.ToSelectedStatus()
				} else {
					t.ToDefaultStatus()
				}
			})
		}()
	})

	// 事件订阅：选择数字变更事件
	globel.EventBus().Subscribe(event.SelectedNumChange, func(event event.Event) {
		num, _ := strconv.Atoi(t.text.Text)
		go func() {
			fyne.DoAndWait(func() {
				if event.Data.(int) == num {
					t.ToSelectedStatus()
				} else {
					t.ToDefaultStatus()
				}
			})
		}()
	})

	// 游戏刷新
	globel.EventBus().Subscribe(event.GameRefresh, func(event event.Event) {
		go func() {
			fyne.DoAndWait(func() {
				t.Refresh()
			})
		}()
	})
}

// 抽出公共画图代码
func (t *TappableNumberCell) drawHandle() {
	// 获取对应的数字
	num, _ := strconv.Atoi(t.text.Text)
	fillQuantity := globel.NumberFillQuantity(num)
	if fillQuantity == 9 {
		t.selectedColor = color.RGBA{R: t.selectedColor.R, G: t.selectedColor.G, B: t.selectedColor.B, A: 180}
		t.defaultColor = color.RGBA{R: t.defaultColor.R, G: t.defaultColor.G, B: t.defaultColor.B, A: 120}
	} else {
		t.selectedColor = sColor
		t.defaultColor = dColor
	}
	if globel.GetDataStorage(globel.SelectedNum) == num {
		t.ToSelectedStatus()
	} else {
		t.ToDefaultStatus()
	}
}
