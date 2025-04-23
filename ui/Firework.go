package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"math/rand"
	"time"
)

// 烟花特效批次配置
type FireworkConfig struct {
	Color         color.Color //颜色
	ParticleCount int         //粒子总数
	SpeedBase     float64     //速度
}

// 烟花特效管理器
type FireworkLauncher struct {
	widget.BaseWidget
	container       *fyne.Container
	activeParticles []*Particle
	animating       bool
}

type Particle struct {
	circle    *canvas.Circle
	pos       fyne.Position
	velocity  fyne.Position
	startTime time.Time
}

func NewFireworkLauncher() *FireworkLauncher {
	return &FireworkLauncher{
		container: container.NewWithoutLayout(),
	}
}

func (fl *FireworkLauncher) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(fl.container)
}

// LaunchFirework 生成指定配置的烟花特效（线程安全）
func (fl *FireworkLauncher) LaunchFirework(pos fyne.Position, config FireworkConfig) {
	if fl.animating {
		return
	}
	fl.animating = true

	go func() {
		defer func() { fl.animating = false }()

		// 连续触发三次
		for i := 0; i < 3; i++ {
			// 每次生成前清理旧粒子
			fl.cleanParticles()

			// 根据次数修改配置
			currentConfig := config
			currentConfig.Color = generateStageColor(i)
			currentConfig.SpeedBase += float64(i) * 0.5

			// 生成粒子
			particles := fl.createParticles(pos, currentConfig)
			fl.activeParticles = particles

			// 运行动画1秒
			fl.runAnimation(1 * time.Second)

			// 清理本次粒子
			fl.cleanParticles()
		}
	}()
}

// 创建不同阶段的粒子
func (fl *FireworkLauncher) createParticles(pos fyne.Position, config FireworkConfig) []*Particle {
	var particles []*Particle

	for i := 0; i < config.ParticleCount; i++ {
		angle := rand.Float64() * 2 * math.Pi
		speed := config.SpeedBase + rand.Float64()*2

		p := &Particle{
			circle:    canvas.NewCircle(config.Color),
			pos:       pos,
			velocity:  fyne.NewPos(float32(math.Cos(angle))*float32(speed), float32(math.Sin(angle))*float32(-speed)),
			startTime: time.Now(),
		}
		p.circle.Resize(fyne.NewSize(8, 8))
		p.circle.Move(pos.Subtract(fyne.NewPos(4, 4)))

		particles = append(particles, p)
		fl.container.Add(p.circle)
	}
	return particles
}

// 运行动画并自动清理
func (fl *FireworkLauncher) runAnimation(duration time.Duration) {
	start := time.Now()
	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()

	for range ticker.C {
		elapsed := time.Since(start)
		if elapsed > duration {
			return
		}
		fyne.DoAndWait(func() {
			delta := float64(time.Since(start)) / float64(time.Second)
			fl.updateParticles(delta)
			fl.container.Refresh()
		})
	}
}

// 更新粒子状态
func (fl *FireworkLauncher) updateParticles(delta float64) {
	gravity := fyne.NewPos(0, 0.3)
	for _, p := range fl.activeParticles {
		p.velocity = p.velocity.Add(gravity)
		p.pos = p.pos.Add(p.velocity)
		p.circle.Move(p.pos.Subtract(fyne.NewPos(4, 4)))

		// 透明度渐变
		alpha := uint8(255 * (1 - delta))
		if col, ok := p.circle.FillColor.(color.NRGBA); ok {
			p.circle.FillColor = color.NRGBA{R: col.R, G: col.G, B: col.B, A: alpha}
		}
	}
}

// 清理粒子
func (fl *FireworkLauncher) cleanParticles() {
	fyne.DoAndWait(func() {
		for _, p := range fl.activeParticles {
			fl.container.Remove(p.circle)
		}
		fl.activeParticles = nil
		fl.container.Refresh()
	})
}

// 生成不同阶段的颜色
func generateStageColor(stage int) color.Color {
	switch stage {
	case 0:
		return color.NRGBA{R: 255, G: 50, B: 0, A: 255} // 红色
	case 1:
		return color.NRGBA{R: 0, G: 200, B: 255, A: 255} // 蓝色
	case 2:
		return color.NRGBA{R: 255, G: 215, B: 0, A: 255} // 金色
	default:
		return color.White
	}
}
