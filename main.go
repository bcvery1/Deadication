package main

import (
  "log"
  "time"

  "Deadication/hud"
  "Deadication/mob"
  "Deadication/scenes"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"

  "golang.org/x/image/colornames"
)

const (
  winWidth float64  = 1024
  winHeight float64 = 768
  camSpeed float64  = 500.0
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

  // Get the scenes
  menuScene := scenes.GetScene("menu", &sceneChange)
  homeScene := scenes.GetScene("home", &sceneChange)
  farmScene := scenes.GetScene("farm", &sceneChange)
  inventory := scenes.GetScene("inventory", &sceneChange)
  // Set active scene to the menu
  activeScene := "menu"

  // Set up camera
  camPos := pixel.ZV

  last := time.Now()
  for !win.Closed() {
    // Clear previously drawn images
    win.Clear(backgroundColour)
    select {
    case activeScene = <- sceneChange:
      log.Printf("New scene %s", activeScene)
    default:
    }
    // Update the active scene
    switch activeScene {
    case "menu":
      menuScene.Update(win)
    case "home":
      homeScene.Update(win)
    case "farm":
      farmScene.Update(win)
    case "inventory":
      inventory.Update(win)
    default:
      log.Fatalf("Unknown scene %s", activeScene)
    }

    // Draw HUD to screen
    hud.Draw(win, camPos)

    // Draw the character centre screen, move it with the camera position
    character.Update(win, camPos)

    dt := time.Since(last).Seconds()
    last = time.Now()

    // Move the camera
    if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
      camPos.X -= camSpeed * dt
    }
    if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
      camPos.X += camSpeed * dt
    }
    if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
      camPos.Y -= camSpeed * dt
    }
    if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
      camPos.Y += camSpeed * dt
    }

    // Set the cam as the viewpoint
    cam := pixel.IM.Moved(win.Bounds().Center().Sub(camPos))
    win.SetMatrix(cam)

    win.Update()
  }
}

func main() {
  pixelgl.Run(run)
}
