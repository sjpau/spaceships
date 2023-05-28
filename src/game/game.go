package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CANVAS_WIDTH  = 640
	CANVAS_HEIGHT = 360
)

type Game struct {
	canvas *ebiten.Image
}

func NewGame() *Game {
	g := &Game{
		canvas: ebiten.NewImage(CANVAS_WIDTH, CANVAS_HEIGHT),
	}
	return g
}

func (self *Game) Update() error {
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {

	w, h := screen.Size()
	xscale, yscale := float64(w)/CANVAS_WIDTH, float64(h)/CANVAS_HEIGHT
	scaling := int(math.Min(xscale, yscale))
	Width, Height := CANVAS_WIDTH*scaling, CANVAS_HEIGHT*scaling
	ysf, xsf := float64(scaling), float64(scaling)
	marginHor := (w - Width) / 2
	marginVer := (h - Height) / 2

	o := &ebiten.DrawImageOptions{}
	if float64(scaling) < 1 {
		// Resolution is too low
		marginHor, marginVer = 0, 0
		xsf, ysf = float64(w)/CANVAS_WIDTH, float64(h)/CANVAS_HEIGHT
		if w >= CANVAS_WIDTH {
			xsf = 1.0
			marginHor = (w - CANVAS_WIDTH) / 2
		}
		if h >= CANVAS_HEIGHT {
			ysf = 1.0
			marginVer = (h - CANVAS_HEIGHT) / 2
		}
	}
	o.GeoM.Scale(xsf, ysf)
	if marginHor != 0 || marginVer != 0 {
		screen.Fill(color.RGBA{24, 24, 24, 255})
		o.GeoM.Translate(float64(marginHor), float64(marginVer))
	}
	self.canvas.Fill(color.RGBA{0, 0, 0, 255})
	screen.DrawImage(self.canvas, o)
	/* Render High RES if needed*/
	if scaling >= 1 {
	}
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
