package state

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Draw(screen *ebiten.Image)
	Update()
	Init() //TODO: pass slice of options to init. ex: ship choice from menu
	Change() bool
}
