package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/vector"
)

type Bullet struct {
	object       *Object
	owner        *Object
	damage       int
	acceleration float64
	released     bool
	releasePoint vector.Vector2D
}

func (b *Bullet) ReleaseTo(releasePoint *vector.Vector2D) {
	b.released = true
	b.releasePoint = *releasePoint
	b.object.direction = *b.releasePoint.Sub(&b.object.position)
	b.object.angle = math.Atan2(b.object.direction.Y, b.object.direction.X)
}

func (b *Bullet) Update() {
	if b.released {
		b.object.velocity.X += b.acceleration * math.Cos(b.object.angle)
		b.object.velocity.Y += b.acceleration * math.Sin(b.object.angle)
		b.object.position.X += b.object.velocity.X
		b.object.position.Y += b.object.velocity.Y
	} else {
		b.object.position.X = b.owner.position.X
		b.object.position.Y = b.owner.position.Y
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	if b.released {
		s := b.object.image.Bounds().Size()
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(1, 1)
		o.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
		o.GeoM.Rotate(b.object.angle)
		o.GeoM.Translate(b.object.position.X, b.object.position.Y)
		screen.DrawImage(b.object.image, o)
	}
}
