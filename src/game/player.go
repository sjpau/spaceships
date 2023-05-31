package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/graphics"
	"github.com/sjpau/vector"
)

type Player struct {
	object       *Object
	acceleration float64
	bullets      []*Bullet
	ammo         int
}

func (p *Player) UpdateBullets() {
	for i := range p.bullets {
		if p.bullets[i] != nil {
			EbitenObjectUpdate(p.bullets[i])
			if p.bullets[i].object.OutsideWindow() &&
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
	p.DrawBullets(screen)
}

func NewPlayer(img *ebiten.Image) *Player {
	//TODO: add cases for different ships
	p := &Player{
		object: &Object{
			image: img,
		},
		ammo: 50,
	}
	p.object.position.X, p.object.position.Y = p.object.Center()
	bullets := make([]*Bullet, p.ammo)
	for i := range bullets {
		bullets[i] = NewBullet(graphics.SpritesBullets[graphics.BLUE], p.object, 10, 10, 1)
	}
	p.bullets = bullets
	return p
}
