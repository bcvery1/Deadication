package hud

import (
	"fmt"
	"strings"
	"time"

	"github.com/bcvery1/Deadication/player"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	// How long is a game day in reallife seconds
	// Update this in util/crops.go if changed
	daySeconds = 1
)

var (
	bottomLeft = pixel.V(0, 690)
	topRight   = pixel.V(1280, 720)
	lineLeft   = pixel.V(0, 689)
	lineRight  = pixel.V(1280, 689)

	months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	textColor = colornames.Black
)

// HUD is the hud struct
type HUD struct {
	day   int
	month int
	atlas *text.Atlas
}

// Update updates the hud and draws it the window
func (h *HUD) Update(win *pixelgl.Window, player *player.Player) {
	imd := imdraw.New(nil)

	// Draw the box
	imd.Color = colornames.Whitesmoke
	imd.Push(bottomLeft, topRight)
	imd.Rectangle(0)

	// Draw the line shadow
	imd.Color = colornames.Darkslategrey
	imd.Push(lineLeft, lineRight)
	imd.Line(1)

	imd.Draw(win)

	date, dateMat := h.getDate()
	date.Draw(win, dateMat)

	health, healthMat := h.getHealth(player)
	health.Draw(win, healthMat)

	hunger, hungerMat := h.getHunger(player)
	hunger.Draw(win, hungerMat)

	carry, carryMat := h.getCarry(player)
	carry.Draw(win, carryMat)
}

func scaleHUDText(txt *text.Text) pixel.Matrix {
	return pixel.IM.Scaled(txt.Orig, 1.4)
}

func (h *HUD) getCarry(p *player.Player) (*text.Text, pixel.Matrix) {

	txt := text.New(pixel.V(550, 700), h.atlas)
	txt.Color = textColor

	if p.Carrying() != "" {
		fmt.Fprintf(txt, "Carrying: %s", strings.Title(p.Carrying()))
	}
	return txt, scaleHUDText(txt)
}

func (h *HUD) getDate() (*text.Text, pixel.Matrix) {
	txt := text.New(pixel.V(1155, 700), h.atlas)
	txt.Color = textColor

	fmt.Fprintf(txt, "%d %s", h.day, months[h.month])

	return txt, scaleHUDText(txt)
}

func (h *HUD) getHealth(p *player.Player) (*text.Text, pixel.Matrix) {
	txt := text.New(pixel.V(10, 700), h.atlas)
	txt.Color = textColor

	fmt.Fprintf(txt, "Health %d/100", p.Health())

	return txt, scaleHUDText(txt)
}

func (h *HUD) getHunger(p *player.Player) (*text.Text, pixel.Matrix) {
	txt := text.New(pixel.V(250, 700), h.atlas)
	txt.Color = textColor

	fmt.Fprintf(txt, "Hunger %d/100", p.Hunger())

	return txt, scaleHUDText(txt)
}

func (h *HUD) updateDate() {
	for {
		<-time.After(time.Second * daySeconds)
		if h.day == 30 {
			h.day = 1
			if h.month == 11 {
				h.month = 0
			} else {
				h.month++
			}
		} else {
			h.day++
		}
	}
}

// NewHUD creates a hud object to display
func NewHUD() *HUD {
	h := HUD{
		// State 1st March
		1,
		2,
		// For writing text
		text.NewAtlas(basicfont.Face7x13, text.ASCII),
	}

	// Update date in background
	go h.updateDate()

	return &h
}
