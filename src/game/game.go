package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/vector"
)

const (
	WIDTH  = 640
	HEIGHT = 360
)

type Game struct {
	dt       float64
	fsWidth  int
	fsHeight int
	cursor   vector.Vector2D
	player   *Player
}

func NewGame() *Game {
	fsw, fsh := ebiten.ScreenSizeInFullscreen()
	p := NewPlayer()
	g := &Game{
		fsWidth:  fsw,
		fsHeight: fsh,
		cursor: vector.Vector2D{
			X: 0,
			Y: 0,
		},
		player: p,
	}
	return g
}

func (self *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if self.player != nil {
			if len(self.player.bullets) > 0 {
				for i := range self.player.bullets {
					if self.player.bullets[i] != nil {
						if !self.player.bullets[i].released {
							self.player.bullets[i].ReleaseTo(&self.cursor)
							break
						}
					}
				}
			}
		}
	}
	crx, cry := ebiten.CursorPosition()
	self.cursor.X = float64(crx)
	self.cursor.Y = float64(cry)
	self.player.object.SetDirection(&self.cursor)

	EbitenObjectUpdate(self.player)
	for i := range self.player.bullets {
		if self.player.bullets[i] != nil {
			EbitenObjectUpdate(self.player.bullets[i])
			if self.player.bullets[i].object.OutsideWindow() &&
				self.player.bullets[i].released {
				self.player.bullets[i] = nil
			}
		}
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	for i := range self.player.bullets {
		if self.player.bullets[i] != nil {
			EbitenObjectDraw(self.player.bullets[i], screen)
		}
	}
	EbitenObjectDraw(self.player, screen)
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
