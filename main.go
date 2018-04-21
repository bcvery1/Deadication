package main

import (
  "log"
  "time"

  "Deadication/gameinfo"
  "Deadication/mob"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"

  "golang.org/x/image/colornames"
)

const (
  winWidth float64  = 1024
  winHeight float64 = 768
)

var (
  backgroundColour = colornames.Forestgreen
  sceneChange chan string = make(chan string, 1)
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
  win.Clear(backgroundColour)

  // Get main characters sprite
  character, err := mob.GetChar()
  if err != nil {
    log.Fatal(err)
  }

  game := gameinfo.NewGame(win, &(pixel.ZV), character, "farm")

  last := time.Now()
  for !win.Closed() {
    // Clear previously drawn images
    win.Clear(backgroundColour)

    dt := time.Since(last).Seconds()
    last = time.Now()

    // Update the active scene
    game.Update(dt)

    // Set the cam as the viewpoint
    cam := pixel.IM.Moved(win.Bounds().Center().Sub(*game.CamPos))
    win.SetMatrix(cam)

    win.Update()
  }
}

func main() {
  pixelgl.Run(run)
}
