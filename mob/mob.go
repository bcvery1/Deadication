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
  SetState(int)
}

type Mob struct {
  // Maps state ID to annimation sequence
  // BUG(Currently only mapping to one sprite)
  Sprites map[int]*pixel.Sprite
  State int
  PosMat pixel.Matrix
}

func (m Mob) Update(win *pixelgl.Window, camPos pixel.Vec) {
  m.Sprites[m.State].Draw(win, m.PosMat)
}
