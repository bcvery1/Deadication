package main

import (
	"log"
	"time"

	"github.com/bcvery1/Deadication/hud"
	"github.com/bcvery1/Deadication/player"
	"github.com/bcvery1/Deadication/util"

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

	// Assign to global var
	util.AllSprites = sprites

	// Add two initial humans
	util.Pens["Bottom pen"].AddHuman()
	util.Pens["Bottom pen"].AddHuman()

	// Add initial corn to top field
	util.Fields["Top field"].Plant(util.NewCrop("corn"))
	util.Fields["Mid field"].Plant(util.NewCrop("apple"))

	// Start listening for popups
	util.InitPopups()
	// Start listening for eatings
	util.InitPens()
	// Start listening for harvests
	util.InitFields()

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

		// Update crop growth via fields
		for _, f := range util.Fields {
			f.UpdateCrop(win, sprites)
		}

		newZone := playerObj.Update(win, dt, collisions, zones)
		if newZone != "" {
			// Player is in a named interactive zone
			// Only activate if changed zones
			if newZone != inZone {
				interactives[newZone].Activate(playerObj.Carrying())
				// Set inZone
				inZone = newZone
			}
			interactives[newZone].Update(win, playerObj.Carrying())
		} else {
			// Reset inZone
			inZone = ""
			// Deactivate all zones
			for _, zone := range interactives {
				zone.Deactivate()
			}
		}

		// Update human movements
		for _, p := range util.Pens {
			p.UpdateHumans(win, dt)
		}

		// Update the HUD
		gamehud.Update(win, playerObj)

		// Display any popups
		popup, show := util.GetMessage()
		if show {
			popup.Draw(win)
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
