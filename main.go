package main

import (
  "log"
  // "time"

  "Deadication/mob"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"

  "golang.org/x/image/colornames"
)

const (
  winWidth float64  = 1024
  winHeight float64 = 768
)

// Main function pixel will run
func run() {
  cfg := pixelgl.WindowConfig{
    Title:  "DEADication",
    Bounds: pixel.R(0, 0, winWidth, winHeight),
    VSync:  true,
  }
  win, err := pixelgl.NewWindow(cfg)
  if err != nil {
    log.Fatal(err)
  }
  win.SetSmooth(true)
  win.Clear(colornames.White)

  characterSprite, err := mob.GetCharacterSprite()
  if err != nil {
    log.Fatal(err)
  }

  // last := time.Now()
  for !win.Closed() {
    // dt := time.Since(last).Seconds()
    // last = time.Now()

    characterSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

    win.Update()
  }
}

func main() {
  pixelgl.Run(run)
}
