package ui

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	myTheme "com.cyy/sudoku/theme"
	"com.cyy/sudoku/utils"
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
	sudokuX  int // 对应数独数组的X坐标
	sudokuY  int // 对应数独数组的X坐标
}

// 接口声明
var (
	_ fyne.Tappable = (*SudokuCell)(nil)
	_ fyne.Widget   = (*SudokuCell)(nil)
)

var (
	dTextColor = myTheme.SimpleTextColor()
	sTextColor = utils.HTML2FyneRGB(0, 187, 0)
)

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
	defaultTextColor  color.Color       // 默认的文字颜色
	tappedTextColor   color.Color       //点击之后的文字颜色
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
	c.text.Color = textColor
	// 颜色设置
	c.defaultTextColor = textColor
	c.tappedTextColor = sTextColor

	// 对应数组坐标转换
	c.data.sudokuX = 3*(data.groupIdx/3) + data.x
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
	c.eventSubscribe()
	return c
}

// 画图处理
func (c *SudokuCell) drawHandle() {
	// 获取数据信息

	cell := globel.GetGameDataVal(c.data.sudokuX, c.data.sudokuY)
	//log.Printf("drawHandle-->cell:%v,x:%d,y:%d", cell, c.data.sudokuX, c.data.sudokuY)
	if cell.Num != 0 {
		c.text.Text = strconv.Itoa(cell.Num)
		// 是否是挖出来的孔
		if cell.IsHole { //选中的颜色
			c.text.Color = c.tappedTextColor
			c.text.TextStyle.Italic = true // 设置斜体
		} else {
			c.text.Color = c.defaultTextColor
			c.text.TextStyle.Italic = false // 取消斜体
		}
	} else {
		c.text.Text = ""
	}

	// 如果选中了对应数字，才进行边框设定
	if globel.GetDataStorage(globel.SelectedNum) == cell.Num {
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

func (c *SudokuCell) eventSubscribe() {
	// 事件订阅:选择数字滚动
	globel.EventBus().Subscribe(event.SelectedNumChange, func(event event.Event) {
		go func() { // 异步执行渲染进程，防止出现问题
			fyne.DoAndWait(func() {
				c.Refresh()
			})
		}()
	})
	//事件订阅::指定坐标回滚事件
	globel.EventBus().Subscribe(event.GameUndoStep+strconv.Itoa(c.data.sudokuX)+strconv.Itoa(c.data.sudokuY), func(event event.Event) {
		go func() { // 异步执行渲染进程，防止出现问题
			fyne.DoAndWait(func() {
				c.Refresh()
			})
		}()
	})
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
		canChange := globel.ChangeGameDataVal(c.data.sudokuX, c.data.sudokuY, globel.GetDataStorage(globel.SelectedNum).(int))
		if canChange {
			c.Refresh() //复用刷新
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
