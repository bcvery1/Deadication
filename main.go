package main

import (
  "log"
  "time"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  winWidth float64 = 1024
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

  last := time.Now()
  for !win.Closed() {
    dt := time.Since(last).Seconds()
    last = time.Now()
    
    win.Update()
  }
}

func main() {
  pixelgl.Run(run)
}
