package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/game"
)

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
