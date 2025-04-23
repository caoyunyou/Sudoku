package test

//
//
//
//import (
//"fyne.io/fyne/v2"
//"fyne.io/fyne/v2/app"
//"fyne.io/fyne/v2/canvas"
//"fyne.io/fyne/v2/container"
//"fyne.io/fyne/v2/driver/desktop"
//"fyne.io/fyne/v2/widget"
//"image/color"
//)
//
//// 1. 创建自定义光标组件
//type CustomCursor struct {
//	widget.BaseWidget
//	circle  *canvas.Circle
//	text    *canvas.Text
//	visible bool
//}
//
//func NewCustomCursor() *CustomCursor {
//	c := &CustomCursor{
//		circle: canvas.NewCircle(color.RGBA{R: 255, A: 200}),
//		text:   canvas.NewText("", color.Black),
//	}
//	c.circle.Resize(fyne.NewSize(30, 30))
//	c.text.TextSize = 14
//	c.text.Alignment = fyne.TextAlignCenter
//	c.ExtendBaseWidget(c)
//	c.Hide() // 初始隐藏
//	return c
//}
//
//func (c *CustomCursor) UpdatePosition(pos fyne.Position) {
//	// 计算中心点偏移
//	centerPos := pos.Subtract(fyne.NewPos(15, 15))
//	c.Move(centerPos)
//	c.Refresh()
//}
//
//func (c *CustomCursor) CreateRenderer() fyne.WidgetRenderer {
//	return widget.NewSimpleRenderer(
//		container.NewWithoutLayout(
//			c.circle,
//			container.NewCenter(c.text),
//		)
//	)
//}
//
//// 2. 创建可交互父容器
//type CursorContainer struct {
//	widget.BaseWidget
//	content     *fyne.Container
//	cursor      *CustomCursor
//	isHovered   bool
//}
//
//var _ desktop.Hoverable = (*CursorContainer)(nil)
//
//func NewCursorContainer() *CursorContainer {
//	c := &CursorContainer{
//		cursor: NewCustomCursor(),
//	}
//	// 主内容（示例为一个按钮）
//	btn := widget.NewButton("测试按钮", nil)
//
//	c.content = container.NewStack(
//		btn,
//		c.cursor,
//	)
//	c.ExtendBaseWidget(c)
//	return c
//}
//
//// 3. 实现鼠标事件接口
//func (c *CursorContainer) MouseIn(e *desktop.MouseEvent) {
//	c.isHovered = true
//	c.cursor.Show()
//	// 隐藏系统光标（需要自定义驱动实现）
//	if dev, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
//		dev.SetSystemCursor(desktop.HiddenCursor)
//	}
//}
//
//func (c *CursorContainer) MouseMoved(e *desktop.MouseEvent) {
//	if c.isHovered {
//		// 转换为全局坐标
//		globalPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(c.content)
//		c.cursor.UpdatePosition(e.AbsolutePosition.Subtract(globalPos))
//	}
//}
//
//func (c *CursorContainer) MouseOut() {
//	c.isHovered = false
//	c.cursor.Hide()
//	// 恢复系统光标
//	if dev, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
//		dev.SetSystemCursor(desktop.DefaultCursor)
//	}
//}
//
//func (d *CursorContainer) Cursor() desktop.Cursor {
//	if d.split.Horizontal {
//		return desktop.HiddenCursor
//	}
//	return desktop.DefaultCursor
//}
//
//func (c *CursorContainer) CreateRenderer() fyne.WidgetRenderer {
//	return widget.NewSimpleRenderer(c.content)
//}
//
//// 4. 使用示例
//func main() {
//	a := app.New()
//	w := a.NewWindow("自定义光标示例")
//
//	container := NewCursorContainer()
//	w.SetContent(container)
//
//	w.Resize(fyne.NewSize(600, 400))
//	w.ShowAndRun()
//}
