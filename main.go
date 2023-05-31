package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/game"
	"github.com/sjpau/spaceships/src/graphics"
)

//go:embed asset/*
var assetsFS embed.FS

func init() {
	graphics.Preload(&assetsFS)
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowTitle("Spaceships")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
