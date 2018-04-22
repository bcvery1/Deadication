package util

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	humanSpeed = 10.0
)

var (
	src    = rand.NewSource(time.Now().Unix())
	myRand = rand.New(src)
)

// Human holds human in pen object
type Human struct {
	pos     pixel.Vec
	nextpos pixel.Vec
	sprite  *pixel.Sprite
}

func randomVec(r pixel.Rect) pixel.Vec {
	xDiff := myRand.Intn(int(r.W()))
	yDiff := myRand.Intn(int(r.H()))
	diffV := pixel.V(float64(xDiff), float64(yDiff))

	return pixel.V(r.Min.X, r.Min.Y).Add(diffV)
}

// this is 'false' bias
func randomBool() bool {
	r := myRand.Intn(200)
	return r == 100
}

// Update sets the humans movements
func (h *Human) Update(p *pen, win *pixelgl.Window, dt float64) {
	h.sprite.Draw(win, pixel.IM.Moved(h.pos))

	// Try move human
	if randomBool() {
		h.nextpos = randomVec(p.rect)
	}

	// Move towards the next pos
	// Definitely better way to do this with vectors...
	if h.pos.X > h.nextpos.X {
		h.pos.X -= humanSpeed * dt
	} else {
		h.pos.X += humanSpeed * dt
	}
	if h.pos.Y > h.nextpos.Y {
		h.pos.Y -= humanSpeed * dt
	} else {
		h.pos.Y += humanSpeed * dt
	}
}

// NewHuman creates a new human in the pen
func NewHuman(p *pen, sprites map[string]*pixel.Sprite) *Human {
	return &Human{
		p.rect.Center(),
		p.rect.Center(),
		sprites["human"],
	}
}
