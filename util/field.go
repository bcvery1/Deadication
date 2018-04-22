package util

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	cottonseed = "cottonseed"
	appleseed  = "appleseed"
	cornseed   = "cornseed"
)

type field struct {
	Interactive
	crop       *Crop
	planted    bool
	amountLeft int
	rect       pixel.Rect
}

// PlantAction action to plant a crop in a field
type PlantAction struct {
	field string
	crop  *Crop
}

// InitFields loops and listens for harvests
func InitFields() {
	go func() {
		for {
			select {
			case fieldName := <-HarvestChan:
				Fields[fieldName].harvest()
			case plantAction := <-PlantChan:
				Fields[plantAction.field].Plant(plantAction.crop)
			}
		}
	}()
}

func (f *field) harvest() {
	// Decrement amount of crops left
	f.amountLeft--
	log.Println(f.amountLeft)

	if f.amountLeft == 0 {
		log.Println(f.planted)
		// If the plant reverts, field is still planted
		f.planted = f.crop.Revert()
		log.Println(f.planted)
	}
}

// UpdateCrop draws the crop to the screen
func (f *field) UpdateCrop(win *pixelgl.Window, allSprites map[string]*pixel.Sprite) {
	if !f.planted {
		return
	}

	// Get the sprite from allSprites
	spriteName := fmt.Sprintf(f.crop.spriteFmt, f.crop.stage)
	cropSprite, ok := allSprites[spriteName]
	if !ok {
		return
	}

	// Draw to each square
	for x := f.rect.Min.X; x < f.rect.Max.X; x += spriteMapWidth {
		for y := f.rect.Min.Y; y < f.rect.Max.Y; y += spriteMapWidth {
			posV := pixel.V(x+(spriteMapWidth/2.0), y+(spriteMapWidth/2.0))
			cropSprite.Draw(win, pixel.IM.Moved(posV))
		}
	}
}

// Plant adds a crop to this field
func (f *field) Plant(c *Crop) {
	f.planted = true
	f.crop = c
	f.amountLeft = c.harvestAmount
}

func (f *field) Update(win *pixelgl.Window, carrying string) {
	if !f.IsActive() {
		return
	}

	// Draw box
	imd := getBox()
	imd.Draw(win)

	// Draw title
	title, scale := getText(-1, f.Title(), 1.4, titleV)
	title.Draw(win, scale)

	shiftV := pixel.V(0, 30)
	fieldoptions := f.opts(carrying)
	for j, opt := range fieldoptions {
		v := menuV.Sub(shiftV.Scaled(float64(j + 1)))
		optTxt, scale := getText(j+1, opt.Text(), 1.1, v)
		optTxt.Draw(win, scale)
	}

	// Check if the user presses a number key to select an option
	doOptions(win, fieldoptions, carrying, f)
}

func (f *field) opts(c string) []optionI {
	observe := observeField{option{"Observe field"}}
	opts := []optionI{&observe}

	if c == "water" {
		o := waterField{option{"Water field"}}
		opts = append(opts, &o)
	}

	if (c == cottonseed || c == appleseed || c == cornseed) && !f.planted {
		o := plantSeeds{option{"Plant seeds"}, c}
		opts = append(opts, &o)
	}

	if f.crop.IsReady() && f.planted {
		s := fmt.Sprintf("Havest (%d left)", f.amountLeft)
		o := havest{option{s}}
		opts = append(opts, &o)
	}

	return opts
}

type observeField struct {
	option
}

func (o *observeField) Action(f InteractiveI, carrying string) {
	PopupChan <- &Popup{"You can grow plants for humans to eat here"}
}

type waterField struct {
	option
}

func (w *waterField) Action(f InteractiveI, carrying string) {
	if carrying == "water" {
		s := fmt.Sprintf("You watered %s.  This will cause the\nplants to grow a bit faster this month", f.Title())
		PopupChan <- &Popup{s}
		PickupChan <- ""
	}
}

type plantSeeds struct {
	option
	seed string
}

func (p *plantSeeds) Action(f InteractiveI, carrying string) {
	s := fmt.Sprintf("You planted %s in %s.\nMake sure you water it each month", p.seed, f.Title())
	PopupChan <- &Popup{s}
	PickupChan <- ""
	PlantChan <- PlantAction{
		f.Title(),
		NewCrop(p.seed),
	}
}

type havest struct {
	option
}

func (h *havest) Action(f InteractiveI, carrying string) {
	PopupChan <- &Popup{"You picked up revolting human food.\nThere isn't even any mold on this!"}
	PickupChan <- "food"
	HarvestChan <- f.Title()
}
