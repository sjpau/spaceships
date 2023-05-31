package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/vector"
)

type Object struct {
	image     *ebiten.Image
	width     int
	height    int
	angle     float64
	position  vector.Vector2D
	direction vector.Vector2D
	velocity  vector.Vector2D
}

type EbitenObject interface {
	Update()
	Draw(screen *ebiten.Image)
}

func (o *Object) OutsideWindow() bool {
	w, h := ebiten.WindowSize()
	if o.position.X < 0 ||
		o.position.X >= float64(w) ||
		o.position.Y < 0 ||
		o.position.Y >= float64(h) {
		return true
	}
	return false
}
func (o *Object) Center() (float64, float64) {
	rect := o.image.Bounds()
	X := rect.Min.X + (rect.Max.X-rect.Min.X)/2
	Y := rect.Min.Y + (rect.Max.Y-rect.Min.Y)/2
	return float64(X), float64(Y)
}

func (o *Object) Collide(n *Object) bool {
	if o.position.X+float64(o.width) >= n.position.X &&
		o.position.X <= n.position.X+float64(n.width) &&
		o.position.Y+float64(o.height) >= n.position.Y &&
		o.position.Y <= n.position.Y+float64(n.height) {
		return true
	}
	return false
}

func (o *Object) SetDirection(v *vector.Vector2D) {
	o.direction = *v.Sub(&o.position)
	o.angle = math.Atan2(o.direction.Y, o.direction.X)
}

func EbitenObjectUpdate(eo EbitenObject) {
	eo.Update()
}

func EbitenObjectDraw(eo EbitenObject, screen *ebiten.Image) {
	eo.Draw(screen)
}
