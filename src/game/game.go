package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/graphics"
	"github.com/sjpau/vector"
)

const (
	WIDTH  = 640
	HEIGHT = 360
)

type Game struct {
	fsWidth  int
	fsHeight int
	cursor   vector.Vector2D
	canvas   *ebiten.Image
	player   *Player
}

func NewGame() *Game {
	fsw, fsh := ebiten.ScreenSizeInFullscreen()
	img := graphics.PlayerShips[graphics.SHIP_BLUE]
	p := &Player{
		object: &Object{
			image: img,
		},
	}
	p.object.position.X, p.object.position.Y = p.object.Center()
	g := &Game{
		fsWidth:  fsw,
		fsHeight: fsh,
		cursor: vector.Vector2D{
			X: 0,
			Y: 0,
		},
		canvas: ebiten.NewImage(WIDTH, HEIGHT),
		player: p,
	}
	return g
}

func (self *Game) Update() error {
	crx, cry := ebiten.CursorPosition()
	self.cursor.X = float64(crx)
	self.cursor.Y = float64(cry)
	self.player.object.SetDirection(&self.cursor)

	EbitenObjectUpdate(self.player)
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	EbitenObjectDraw(self.player, screen)
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
