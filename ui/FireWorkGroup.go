package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

type FireWorkGroup struct {
	widget.BaseWidget
	content      *fyne.Container
	fireWorkList []*FireworkLauncher // 烟花列表
	pos          fyne.Position       // 指定燃放位置
	fireWorkNum  int                 // 烟花数量
	fireInterval time.Duration       // 燃放间隔时间:单位毫秒
}

// NewFireWorkGroup 创建烟花组，用于在指定区域半径内进行烟花的渲染
func NewFireWorkGroup(fireWorkNum int, fireInterval time.Duration) *FireWorkGroup {
	fwg := &FireWorkGroup{fireWorkNum: fireWorkNum, fireInterval: fireInterval}

	fwg.content = container.NewWithoutLayout()
	fwg.fireWorkList = make([]*FireworkLauncher, fwg.fireWorkNum, fwg.fireWorkNum)
	for i := 0; i < fwg.fireWorkNum; i++ {
		fwg.fireWorkList[i] = NewFireworkLauncher()
		fwg.content.Add(fwg.fireWorkList[i])
	}

	fwg.ExtendBaseWidget(fwg)
	return fwg
}

func (fwg *FireWorkGroup) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(fwg.content)
}

// Start 烟花绽放
func (fwg *FireWorkGroup) Start(pos fyne.Position) {
	go func() {
		for idx, launcher := range fwg.fireWorkList {
			//间隔时间内进行激活
			if idx > 0 {
				time.Sleep((fwg.fireInterval) * time.Millisecond)
				launcher.LaunchFirework(fyne.NewPos(pos.X+(rand.Float32()*100-50), pos.Y+(rand.Float32()*100-50)),
					FireworkConfig{
						ParticleCount: 50,
						SpeedBase:     6.0,
					})
			} else {
				launcher.LaunchFirework(fyne.NewPos(pos.X+(rand.Float32()*100-50), pos.Y+(rand.Float32()*100-50)),
					FireworkConfig{
						ParticleCount: 50,
						SpeedBase:     6.0,
					})
			}
		}
	}()

}
