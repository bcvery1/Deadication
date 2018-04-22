package player

import (
	"Deadication/util"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	playerScale float64 = 1.0
	playerSpeed float64 = 140.0
)

// Player holds data on the player
type Player struct {
	sprites       map[string]*pixel.Sprite
	currentAction string
	pos           pixel.Vec
}

func (p *Player) rect(dt float64) pixel.Rect {
	s := p.CurrentSprite(dt)
	return util.TranslateRect(s, p.pos)
}

// NewPlayer creates a new player object
func NewPlayer(all map[string]*pixel.Sprite) *Player {
	sprites := make(map[string]*pixel.Sprite)
	sprites["idle"] = all["player"]
	return &Player{
		sprites,
		"idle",
		pixel.V(120, 120),
	}
}

// CurrentSprite returns the sprite to display
func (p *Player) CurrentSprite(dt float64) *pixel.Sprite {
	return p.sprites[p.currentAction]
}

// Collides returns whether the player collids with any rect in the slice provided
func (p *Player) Collides(dt float64, collisions []pixel.Rect) bool {
	for _, r := range collisions {
		if util.RectCollide(util.TranslateRect(p.CurrentSprite(dt), p.pos), r) {
			return true
		}
	}
	return false
}

// Update draws the player in a new position on the page
func (p *Player) Update(win *pixelgl.Window, dt float64, collisions []pixel.Rect) {
	p.CurrentSprite(dt).Draw(win, pixel.IM.Scaled(p.pos, playerScale).Moved(p.pos))

	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		p.pos.X -= playerSpeed * dt
		if p.Collides(dt, collisions) {
			p.pos.X += playerSpeed * dt
		}
		if p.pos.X < 0.0 {
			p.pos.X = 0.0
		}
	}
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		p.pos.X += playerSpeed * dt
		if p.pos.X > win.Bounds().Max.X {
			p.pos.X = win.Bounds().Max.X
		}
		if p.Collides(dt, collisions) {
			p.pos.X -= playerSpeed * dt
		}
	}
	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		p.pos.Y -= playerSpeed * dt
		if p.pos.Y < 0.0 {
			p.pos.Y = 0.0
		}
		if p.Collides(dt, collisions) {
			p.pos.Y += playerSpeed * dt
		}
	}
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		p.pos.Y += playerSpeed * dt
		if p.pos.Y > win.Bounds().Max.Y {
			p.pos.Y = win.Bounds().Max.Y
		}
		if p.Collides(dt, collisions) {
			p.pos.Y -= playerSpeed * dt
		}
	}
}
