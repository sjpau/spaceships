package component

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/graphics"
	"github.com/sjpau/vector"
)

type Player struct {
	Object       *Object
	acceleration float64
	bullets      []*Bullet
	ammo         int
}

func (p *Player) UpdateBullets() {
	for i := range p.bullets {
		if p.bullets[i] != nil {
			EbitenObjectUpdate(p.bullets[i])
			if p.bullets[i].Object.OutsideWindow() &&
				p.bullets[i].released {
				p.bullets[i] = nil
			}
		}
	}
}

func (p *Player) DrawBullets(screen *ebiten.Image) {
	for i := range p.bullets {
		if p.bullets[i] != nil {
			EbitenObjectDraw(p.bullets[i], screen)
		}
	}
}

func (p *Player) Shoot(cursor *vector.Vector2D) {
	if len(p.bullets) > 0 {
		for i := range p.bullets {
			if p.bullets[i] != nil {
				if !p.bullets[i].released {
					p.bullets[i].ReleaseTo(cursor)
					break
				}
			}
		}
	}
}

func (p *Player) Update() {
	p.UpdateBullets()
	p.AdjustAcceleration()
	p.Object.velocity.X += p.acceleration * math.Cos(p.Object.angle)
	p.Object.velocity.Y += p.acceleration * math.Sin(p.Object.angle)
	p.Object.position.X += p.Object.velocity.X
	p.Object.position.Y += p.Object.velocity.Y
}

func (p *Player) AdjustAcceleration() {
	dir := p.Object.direction.Magnitude()
	p.acceleration = dir / 1500
}

func (p *Player) Draw(screen *ebiten.Image) {
	s := p.Object.image.Bounds().Size()
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	o.GeoM.Rotate(p.Object.angle)
	o.GeoM.Translate(p.Object.position.X, p.Object.position.Y)
	screen.DrawImage(p.Object.image, o)
	p.DrawBullets(screen)
}

func NewPlayer(img *ebiten.Image) *Player {
	//TODO: add cases for different ships
	p := &Player{
		Object: &Object{
			image: img,
		},
		ammo: 50,
	}
	p.Object.position.X, p.Object.position.Y = p.Object.Center()
	bullets := make([]*Bullet, p.ammo)
	for i := range bullets {
		bullets[i] = NewBullet(graphics.SpritesBullets[graphics.BLUE], p.Object, 10, 10, 1)
	}
	p.bullets = bullets
	return p
}
