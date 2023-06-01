package state

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/spaceships/src/game/component"
	"github.com/sjpau/spaceships/src/graphics"
	"github.com/sjpau/vector"
)

type Gameplay struct {
	changeState bool
	fsWidth     int
	fsHeight    int
	cursor      vector.Vector2D
	player      *component.Player
}

func (g *Gameplay) Change() bool {
	return g.changeState
}

func (g *Gameplay) Init() {
	fsw, fsh := ebiten.ScreenSizeInFullscreen()
	p := component.NewPlayer(graphics.SpritesPlayerShips[graphics.BLUE])
	g.fsWidth = fsw
	g.fsHeight = fsh
	g.cursor = vector.Vector2D{
		X: 0,
		Y: 0,
	}
	g.player = p
	fmt.Println(g.player)
}

func (g *Gameplay) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.player != nil {
			g.player.Shoot(&g.cursor)
		}
	}
	crx, cry := ebiten.CursorPosition()
	g.cursor.X = float64(crx)
	g.cursor.Y = float64(cry)
	g.player.Object.SetDirection(&g.cursor)

	component.EbitenObjectUpdate(g.player)
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	component.EbitenObjectDraw(g.player, screen)
}
