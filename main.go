package main

import (
	"embed"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/spaceships/src/game"
	"github.com/sjpau/spaceships/src/graphics"
)

//go:embed asset/*
var assetsFS embed.FS

func init() {
	rand.Seed(time.Now().UnixNano())
	graphics.Preload(&assetsFS)
}

func main() {
	ebiten.SetWindowFloating(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetWindowTitle("Spaceships")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
