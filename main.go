package main

import (
	"log"
	"time"

	"Deadication/hud"
	"Deadication/player"
	"Deadication/util"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

const (
	winWidth  float64 = 1280
	winHeight float64 = 720
)

var (
	backgroundColour = colornames.Forestgreen

	x1 = 0
	x2 = 0
	y1 = 0
	y2 = 0
)

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

	sprites, pic := util.GetSprites()
	playerObj := player.NewPlayer(sprites)
	gamehud := hud.NewHUD()
	batch, collisions := util.CreateBatch(sprites, pic)
	interactives, zones := util.AllInteractives()

	last := time.Now()
	inZone := ""
	for !win.Closed() {
		win.Clear(backgroundColour)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			log.Println(win.MousePosition())
		}

		dt := time.Since(last).Seconds()
		last = time.Now()

		batch.Draw(win)

		newZone := playerObj.Update(win, dt, collisions, zones)
		if newZone != "" {
			// Player is in a named interactive zone
			// Only activate if changed zones
			if newZone != inZone {
				interactives[newZone].Activate(playerObj.Carrying())
				// Set inZone
				inZone = newZone
			}
		} else {
			// Reset inZone
			inZone = ""
			// Deactivate all zones
			for _, zone := range interactives {
				zone.Deactivate()
			}
		}

		// Update the HUD
		gamehud.Update(win, playerObj)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
