package graphics

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/help"
	"github.com/sjpau/spaceships/src/load"
)

func Preload(files *embed.FS) error {
	var err error
	PlayerShips = make([]*ebiten.Image, 1)
	path := "asset/graphics/player/"
	PlayerShips[SHIP_BLUE], err = load.PNG(files, path+"ship_blue_1.png")
	help.Check(err)
	return nil
}
