package util

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var (
	backgroundcolor = colornames.Whitesmoke
	bordercolor     = colornames.Black
	textcolor       = colornames.Black

	bottomLeft = pixel.V(300, 300)
	topRight   = pixel.V(900, 420)

	currentPopup = &Popup{}
	displayPopup = false
)

// Popup defines a popup message to show center screen
type Popup struct {
	message string
}

// Draw draws to the target
func (p *Popup) Draw(t pixel.Target) {
	imd := imdraw.New(nil)

	imd.Color = backgroundcolor
	imd.Push(bottomLeft, topRight)
	imd.Rectangle(0)

	imd.Color = bordercolor
	imd.Push(bottomLeft, topRight)
	imd.Rectangle(2)

	imd.Draw(t)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(320, 400), atlas)
	txt.Color = textcolor

	fmt.Fprintf(txt, p.message)
	txt.Draw(t, pixel.IM.Scaled(txt.Orig, 1.7))
}

// GetMessage returns whether a message needs to be shown
func GetMessage() (*Popup, bool) {
	return currentPopup, displayPopup
}

// InitPopups begins the loop to listen for popups
func InitPopups() {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 3):
				displayPopup = false
			case p := <-PopupChan:
				currentPopup = p
				displayPopup = true
			}
		}
	}()
}
