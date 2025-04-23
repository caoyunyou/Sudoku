package ui

import (
	"com.cyy/sudoku/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

// 计时器状态结构体
type TimerState struct {
	widget.BaseWidget
	timeLabel *canvas.Text
	content   *fyne.Container
	running   bool      // 是否正在运行
	seconds   int       // 累计秒数
	startTime time.Time // 开始时间
	ticker    *time.Ticker
}

func NewTimer() *TimerState {
	t := &TimerState{}
	timeLabel := canvas.NewText("00:00", utils.HTML2FyneRGB(242, 136, 80))
	timeLabel.Alignment = fyne.TextAlignCenter
	timeLabel.TextStyle = fyne.TextStyle{Bold: true}

	t.timeLabel = timeLabel
	t.content = container.NewCenter(t.timeLabel)
	t.ExtendBaseWidget(t)
	// TODO 事件订阅：第一次填入有效数字
	return t
}

func (t *TimerState) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.content)
}

// 计时开始
func (t *TimerState) TimeStart() {
	if !t.running {
		t.running = true
		t.startTime = time.Now()
		// 创建定时器（每秒触发）
		t.ticker = time.NewTicker(time.Second)
		go func() {
			for range t.ticker.C {
				// TODO 加一些操作：事件变动的事件发布
				// 在主线程更新界面
				fyne.DoAndWait(func() {
					elapsed := int(time.Since(t.startTime).Seconds())
					totalSeconds := t.seconds + elapsed
					// 格式化时间显示
					mins := totalSeconds / 60
					secs := totalSeconds % 60
					t.timeLabel.Text = fmt.Sprintf("%02d:%02d", mins, secs)
					t.timeLabel.Refresh()
				})
			}
		}()
	}
}

// 计时停止
func (t *TimerState) TimeStop() {
	if t.running {
		t.running = false
		t.seconds += int(time.Since(t.startTime).Seconds())
		t.ticker.Stop()
	}
}

// 计时重置
func (t *TimerState) TimeRestart() {
	t.running = false
	t.seconds = 0
	if t.ticker != nil {
		t.ticker.Stop()
	}
	t.timeLabel.Text = "00:00"
}
