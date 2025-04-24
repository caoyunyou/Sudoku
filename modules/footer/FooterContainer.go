package footer

import (
	"com.cyy/sudoku/event"
	"com.cyy/sudoku/globel"
	"com.cyy/sudoku/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"strconv"
)

// ContainerGenerate 数字按钮生成
func ContainerGenerate() *fyne.Container {

	eventSubscribe()

	// 底部数字按钮
	numButtons := container.New(layout.NewHBoxLayout())
	for i := 1; i <= 9; i++ {
		num := i
		var tc *ui.TappableNumberCell
		tc = ui.NewTappableNumber(num, fyne.NewSize(55, 55), func() {
			selected := globel.GetDataStorage(globel.SelectedNumberCell)
			if selected != nil { // 清除之前按钮的颜色
				selected.(*ui.TappableNumberCell).ToDefaultStatus()
			}
			selected = tc
			globel.SetDataStorage(globel.SelectedNumberCell, tc)
			selected.(*ui.TappableNumberCell).ToSelectedStatus()
			globel.SetDataStorage(globel.SelectedNum, num)
		})
		numButtons.Add(tc)
	}
	return container.NewGridWrap(fyne.NewSize(540, 100),
		container.NewCenter(numButtons))

}

// 事件订阅
func eventSubscribe() {
	globel.EventBus().Subscribe(event.SelectNumTap, func(event event.Event) {
		t := event.Data.(*ui.TappableNumberCell)
		for {
			ob := globel.GetDataObservable(globel.SelectedNum)
			current := ob.Get()
			ob.Lock()
			for ob.Value() == current {
				ob.Wait()
			}
			go func() {
				fyne.DoAndWait(func() {
					num, _ := strconv.Atoi(t.Text().Text)
					if ob.Value() == num {
						t.ToSelectedStatus()
					} else {
						t.ToDefaultStatus()
					}
					t.Circle().Refresh()
				})
			}()
			ob.UnLock()
		}
	})
}
