// Handles the heads-up-display
package hud

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  "github.com/faiface/pixel/imdraw"

  "golang.org/x/image/colornames"
)

// Draws the hud to the screen
func Draw(win *pixelgl.Window, camPos pixel.Vec) {
  // Get the bottom left corner of the camera
  bottomLeftV := camPos.Sub(win.Bounds().Max.Scaled(0.5))

  imd := imdraw.New(nil)

  imd.Color = colornames.Ivory
  // Push from bottom left corner, up-right to form a rectangle
  imd.Push(bottomLeftV, bottomLeftV.Add(pixel.V(160, 120)))
  imd.Rectangle(0)

  imd.Draw(win)
}
