package util

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  CamSpeed float64  = 500.0
)

func MoveCamera(win *pixelgl.Window, camPos *pixel.Vec, dt float64) {
  if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
    (*camPos).X -= CamSpeed * dt
  }
  if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
    (*camPos).X += CamSpeed * dt
  }
  if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
    (*camPos).Y -= CamSpeed * dt
  }
  if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
    (*camPos).Y += CamSpeed * dt
  }
}
