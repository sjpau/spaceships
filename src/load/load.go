package load

import (
	"embed"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/help"
)

func PNG(files *embed.FS, path string) (*ebiten.Image, error) {
	file, err := files.Open(path)
	help.Check(err)
	img, err := png.Decode(file)
	help.Check(err)
	return ebiten.NewImageFromImage(img), file.Close()
}
