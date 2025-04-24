package ui

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	myTheme "com.cyy/sudoku/theme"
	"com.cyy/sudoku/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// SudokuIdx 数独定位信息
type SudokuIdx struct {
	groupIdx int // 窗格定位，哪个窗格里面
	x        int // 窗格中的x坐标
	y        int // 窗格中的y坐标
	sudukuX  int // 对应数独数组的X坐标
	sudokuY  int // 对应数独数组的X坐标
}

// SudokuCell /*数独单元格*/
type SudokuCell struct {
	widget.BaseWidget                   // 继承
	circle            *canvas.Circle    //圆形背景
	squareBorder      *canvas.Rectangle // 正方形边框容器
	text              *canvas.Text      // 文字
	content           *fyne.Container   //组合容器
	onTapped          func()            //点击回调
	size              fyne.Size         //大小
	minSize           fyne.Size         //最小尺寸
	data              SudokuIdx         // 对应数独的定位信息
	uuid              string            //随机序列号
}

func NewSudokuCell(diameter float32, // 直径
	bgColor color.Color, // 背景色
	textColor color.Color, // 文字颜色
	textSize float32, // 字号
	tapFunc func(), // 点击回调
	squareBorderArr []bool, //正方形边框设置数组
	data SudokuIdx, // 对应数独定位信息
) *SudokuCell {
	c := &SudokuCell{
		circle:   canvas.NewCircle(bgColor),
		text:     canvas.NewText("", myTheme.SimpleTextColor()),
		onTapped: tapFunc,
		size:     fyne.NewSize(diameter, diameter),
		minSize:  fyne.NewSize(diameter, diameter),
		data:     data,
	}
	// 设置文本样式
	c.text.Alignment = fyne.TextAlignCenter
	c.text.TextSize = textSize
	//c.text.TextStyle.Bold = true
	c.text.Color = textColor

	// 对应数组坐标转换
	c.data.sudukuX = 3*(data.groupIdx/3) + data.x
	c.data.sudokuY = (data.groupIdx%3)*3 + data.y
	c.circle.FillColor = color.Transparent // 设置透明背景
	// 抽出画图公共部分
	c.drawHandle()
	// 创建一个正方形边框容器
	c.squareBorder = canvas.NewRectangle(color.Transparent)

	squareBorderContainer := container.NewWithoutLayout(
		container.NewGridWrap(fyne.NewSize(diameter, diameter), c.squareBorder))

	for index, judge := range squareBorderArr {
		if judge {
			squareBorderContainer.Add(utils.AddSquareBorderLine(utils.HTML2FyneRGB(0, 0, 136),
				diameter,
				1,
				index))
		}
	}

	// 组合图形元素
	c.content = container.NewStack(
		container.NewCenter(squareBorderContainer),
		container.NewCenter(container.NewGridWrap(fyne.NewSize(0.9*diameter, 0.9*diameter), c.circle)),
		container.NewCenter(c.text),
	)

	c.ExtendBaseWidget(c)

	// 事件订阅:选择数字滚动
	globel.EventBus().Subscribe(event.SelectNumScroll, func(event event.Event) {
		selectedNum := event.Data.(int)
		go func() { // 异步执行渲染进程，防止出现问题
			fyne.DoAndWait(func() {
				num, _ := strconv.Atoi(c.Text().Text)
				if selectedNum == num {
					c.Circle().StrokeColor = utils.HTML2FyneRGB(238, 119, 80)
					c.Circle().StrokeWidth = 2
				} else {
					c.Circle().StrokeWidth = 0
				}
				c.Circle().Refresh()
			})
		}()
	})
	return c
}

// 画图处理
func (c *SudokuCell) drawHandle() {
	// 获取数据信息
	number := globel.GetGameDataVal(c.data.sudukuX, c.data.sudokuY)
	if number != 0 {
		c.text.Text = strconv.Itoa(number)
	} else {
		c.text.Text = ""
	}

	// 如果选中了对应数字，才进行边框设定
	if globel.GetDataStorage(globel.SelectedNum) == number {
		c.circle.StrokeColor = utils.HTML2FyneRGB(238, 119, 80)
		c.circle.StrokeWidth = 2
	} else {
		c.circle.StrokeWidth = 0
	}
}

// Circle 获取圆形组件
func (c *SudokuCell) Circle() *canvas.Circle {
	return c.circle
}

func (c *SudokuCell) Text() *canvas.Text {
	return c.text
}

func busMethod(c *SudokuCell) {
	for {
		ob := globel.GetDataObservable(globel.SelectedNum)
		current := ob.Get()
		ob.Lock()
		for ob.Value() == current {
			ob.Wait()
		}
		go func() {
			fyne.DoAndWait(func() {
				num, _ := strconv.Atoi(c.text.Text)
				if ob.Value() == num {
					c.circle.StrokeColor = utils.HTML2FyneRGB(238, 119, 80)
					c.circle.StrokeWidth = 2
				} else {
					c.circle.StrokeWidth = 0
				}
				c.circle.Refresh()
			})
		}()
		ob.UnLock()
		fmt.Println("busMethod -> 监听到变化，新值:", ob.Value())
	}
}

// CreateRenderer 创建渲染器
func (c *SudokuCell) CreateRenderer() fyne.WidgetRenderer {
	// 如果选中了对应数字，才进行边框设定
	num, _ := strconv.Atoi(c.text.Text)
	if globel.GetDataStorage(globel.SelectedNum) == num {
		c.circle.StrokeColor = utils.HTML2FyneRGB(238, 119, 80)
		c.circle.StrokeWidth = 2
	} else {
		c.circle.StrokeWidth = 0
	}
	return widget.NewSimpleRenderer(c.content)
}

// Tapped 实现点击接口
func (c *SudokuCell) Tapped(*fyne.PointEvent) {
	if c.onTapped != nil {
		c.onTapped()
	} else { // 设置点击处理事件
		canChange := globel.ChangeGameDataVal(c.data.sudukuX, c.data.sudokuY, globel.GetDataStorage(globel.SelectedNum).(int))
		if canChange {
			c.circle.StrokeColor = utils.HTML2FyneRGB(238, 119, 80)
			c.circle.StrokeWidth = 2
			c.text.Text = strconv.Itoa(globel.GetDataStorage(globel.SelectedNum).(int))
			c.text.Color = utils.HTML2FyneRGB(0, 187, 0)
			c.text.TextStyle.Italic = true // 设置斜体
			c.text.Refresh()
			c.circle.Refresh()
		} else {
			//TODO 看看是不是要弹出警告啥的
		}
	}
}

// Refresh 刷新操作
func (c *SudokuCell) Refresh() {
	c.BaseWidget.Refresh()
	c.drawHandle()
}
