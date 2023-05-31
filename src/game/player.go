package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	object       *Object
	acceleration float64
}

func (p *Player) Update() {
	p.AdjustAcceleration()
	p.object.velocity.X += p.acceleration * math.Cos(p.object.angle)
	p.object.velocity.Y += p.acceleration * math.Sin(p.object.angle)
	p.object.position.X += p.object.velocity.X
	p.object.position.Y += p.object.velocity.Y
}

func (p *Player) AdjustAcceleration() {
	dir := p.object.direction.Magnitude()
	p.acceleration = dir / 1500
}

func (p *Player) Draw(screen *ebiten.Image) {
	s := p.object.image.Bounds().Size()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	o.GeoM.Rotate(p.object.angle)
	o.GeoM.Translate(p.object.position.X, p.object.position.Y)
	screen.DrawImage(p.object.image, o)
}
