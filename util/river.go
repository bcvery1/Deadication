package util

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

// DrawRiver draws the river to screen
func DrawRiver(batch *pixel.Batch) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Aqua
	imd.EndShape = imdraw.RoundEndShape

	imd.Push(pixel.V(0, 404), pixel.V(120, 434))
	imd.Push(pixel.V(191, 495), pixel.V(298, 469))
	imd.Push(pixel.V(400, 551), pixel.V(621, 554))
	imd.Push(pixel.V(751, 604), pixel.V(867, 605))
	imd.Push(pixel.V(934, 651), pixel.V(978, 701))
	imd.Push(pixel.V(986, 720))

	imd.Line(50)
	imd.Draw(batch)
}

type riverInter struct {
	Interactive
}

func (r *riverInter) Update(win *pixelgl.Window, carrying string) {
	if !r.IsActive() {
		return
	}

	// Draw box
	imd := getBox()
	imd.Draw(win)

	// Draw title
	title, scale := getText(-1, r.Title(), 1.4, titleV)
	title.Draw(win, scale)

	shiftV := pixel.V(0, 30)
	riveroptions := r.opts(carrying)
	for j, opt := range riveroptions {
		v := menuV.Sub(shiftV.Scaled(float64(j + 1)))
		optTxt, scale := getText(j+1, opt.Text(), 1.1, v)
		optTxt.Draw(win, scale)
	}

	// Check if the user presses a number key to select an option
	doOptions(win, riveroptions, carrying, r)
}

func (r *riverInter) opts(c string) []optionI {
	o := observeRiver{option{"Observe river"}}
	opts := []optionI{&o}

	if c == "" {
		o := collectWater{option{"Collect water"}}
		opts = append(opts, &o)
	}

	return opts
}

type observeRiver struct {
	option
}

func (o *observeRiver) Action(f InteractiveI, carrying string) {
	PopupChan <- &Popup{"Observing river"}
}

type collectWater struct {
	option
}

func (c *collectWater) Action(f InteractiveI, carrying string) {
	if carrying == "" {
		PopupChan <- &Popup{"You collected water"}
		PickupChan <- "water"
	}
}
