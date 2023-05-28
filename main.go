package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/game"
)

/*All `Game` code is suggested to be separeted from main.go*/
func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
