package component

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/vector"
)

type Bullet struct {
	Object       *Object
	owner        *Object
	damage       int
	acceleration float64
	released     bool
	releasePoint vector.Vector2D
}

func (b *Bullet) ReleaseTo(releasePoint *vector.Vector2D) {
	b.released = true
	b.releasePoint = *releasePoint
	b.Object.direction = *b.releasePoint.Sub(&b.Object.position)
	b.Object.angle = math.Atan2(b.Object.direction.Y, b.Object.direction.X)
}

func (b *Bullet) Update() {
	if b.released {
		b.Object.velocity.X += b.acceleration * math.Cos(b.Object.angle)
		b.Object.velocity.Y += b.acceleration * math.Sin(b.Object.angle)
		b.Object.position.X += b.Object.velocity.X
		b.Object.position.Y += b.Object.velocity.Y
	} else {
		b.Object.position.X = b.owner.position.X
		b.Object.position.Y = b.owner.position.Y
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	if b.released {
		s := b.Object.image.Bounds().Size()
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(1, 1)
		o.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
		o.GeoM.Rotate(b.Object.angle)
		o.GeoM.Translate(b.Object.position.X, b.Object.position.Y)
		screen.DrawImage(b.Object.image, o)
	}
}

func NewBullet(img *ebiten.Image, owner *Object, damage int, seed int, a float64) *Bullet {
	b := &Bullet{
		Object: &Object{
			image: img,
		},
		owner:        owner,
		damage:       damage + rand.Intn(seed),
		acceleration: a,
		released:     false,
	}
	return b
}
