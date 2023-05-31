package game

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/graphics"
)

type Player struct {
	object       *Object
	acceleration float64
	bullets      []*Bullet
	ammo         int
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

func NewPlayer() *Player {
	//TODO: add cases for different ships
	img := graphics.SpritesPlayerShips[graphics.BLUE]
	bimg := graphics.SpritesBullets[graphics.BLUE]
	p := &Player{
		object: &Object{
			image: img,
		},
		ammo: 50,
	}
	p.object.position.X, p.object.position.Y = p.object.Center()
	bullets := make([]*Bullet, p.ammo*2)
	for i := range bullets {
		bullets[i] = &Bullet{
			object: &Object{
				image: bimg,
			},
			owner:        p.object,
			damage:       10 + rand.Intn(10),
			acceleration: 5,
			released:     false,
		}
	}
	p.bullets = bullets
	return p
}
