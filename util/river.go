package util

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

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
