package util

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  CamSpeed float64  = 500.0
)

// Returns potential new position for camera
// Caller must move it after checking collisions
func MoveCamera(win *pixelgl.Window, camPos *pixel.Vec, dt float64) pixel.Vec {
  retV := (*camPos)
  if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
     retV.X -= CamSpeed * dt
  }
  if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
    retV.X += CamSpeed * dt
  }
  if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
    retV.Y -= CamSpeed * dt
  }
  if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
    retV.Y += CamSpeed * dt
  }

  return retV
}
