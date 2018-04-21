package mob

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  StateIdle int = iota
  StateWalk
  StateAttack
  StateDead
)

type IMob interface {
  Update(*pixelgl.Window, pixel.Vec)
}

type Mob struct {
  Sprites []pixel.Sprite
  State int
}
