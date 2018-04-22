package util

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type crop struct {
	name        string
	harvestRate int
}

type field struct {
	Interactive
	havestPerc int
	crop       crop
	planted    bool
	amountLeft int
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

	if c == "seed" && !f.planted {
		o := plantSeeds{option{"Plant seeds"}}
		opts = append(opts, &o)
	}

	if f.havestPerc == 100 {
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
	log.Println("Observing field")
}

type waterField struct {
	option
}

func (w *waterField) Action(f InteractiveI, carrying string) {
	if carrying == "water" {
		log.Println("Watering field")
		PickupChan <- ""
	}
}

type plantSeeds struct {
	option
}

func (p *plantSeeds) Action(f InteractiveI, carrying string) {
	PickupChan <- ""
}

type havest struct {
	option
}

func (h *havest) Action(f InteractiveI, carrying string) {
	PickupChan <- "food"
}
