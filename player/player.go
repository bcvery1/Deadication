package player

import (
	"Deadication/util"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	playerScale float64 = 1.0
	playerSpeed float64 = 140.0
	hungerAfter         = 4
	// How much damage starving does
	starveHurt int = 5
)

// Player holds data on the player
type Player struct {
	sprites       map[string]*pixel.Sprite
	currentAction string
	pos           pixel.Vec
	health        int
	hunger        int
	carrying      string
}

func (p *Player) rect(dt float64) pixel.Rect {
	s := p.CurrentSprite(dt)
	return util.TranslateRect(s, p.pos)
}

// Health returns the players current health
func (p *Player) Health() int {
	return p.health
}

// Hunger returns the players current health
func (p *Player) Hunger() int {
	return p.hunger
}

// Carry causes the player to begin carrying the item
// If already carrying something, it is 'deleted'
func (p *Player) Carry(item string) {
	p.carrying = item
}

// Carrying returns what the player is currently carrying
func (p *Player) Carrying() string {
	return p.carrying
}

func (p *Player) eat(points int) {
	p.health += points
	if p.Health() > 100 {
		p.health = 100
	}
}

func (p *Player) statUpdate() {
	for {
		<-time.After(time.Second * hungerAfter)
		p.hunger--
		if p.hunger < 1 {
			p.hunger = 1
			p.health -= starveHurt
		}
	}
}

func (p *Player) listenChans() {
	for {
		select {
		case newItem := <-util.PickupChan:
			p.Carry(newItem)
		case health := <-util.EatChan:
			p.eat(health)
		}
	}
}

// NewPlayer creates a new player object
func NewPlayer(all map[string]*pixel.Sprite) *Player {
	sprites := make(map[string]*pixel.Sprite)
	sprites["idle"] = all["player"]
	p := Player{
		sprites,
		"idle",
		pixel.V(120, 120),
		100,
		100,
		"",
	}

	go p.statUpdate()
	go p.listenChans()

	return &p
}

// CurrentSprite returns the sprite to display
func (p *Player) CurrentSprite(dt float64) *pixel.Sprite {
	return p.sprites[p.currentAction]
}

// Collides returns whether the player collids with any rect in the slice provided
func (p *Player) Collides(dt float64, collisions []pixel.Rect) bool {
	for _, r := range collisions {
		if p.CollidesWith(dt, r) {
			return true
		}
	}
	return false
}

// CollidesWith checks if the player collides with a specific rect
func (p *Player) CollidesWith(dt float64, rect pixel.Rect) bool {
	return util.RectCollide(util.TranslateRect(p.CurrentSprite(dt), p.pos), rect)
}

// Update draws the player in a new position on the page
// Returns if it is within a zone
func (p *Player) Update(win *pixelgl.Window, dt float64, collisions []pixel.Rect, zones map[pixel.Rect]string) string {
	p.CurrentSprite(dt).Draw(win, pixel.IM.Scaled(p.pos, playerScale).Moved(p.pos))

	// Try move right
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		p.pos.X -= playerSpeed * dt
		if p.Collides(dt, collisions) {
			p.pos.X += playerSpeed * dt
		}
		if p.pos.X < 0.0 {
			p.pos.X = 0.0
		}
	}
	// Try move left
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		p.pos.X += playerSpeed * dt
		if p.pos.X > win.Bounds().Max.X {
			p.pos.X = win.Bounds().Max.X
		}
		if p.Collides(dt, collisions) {
			p.pos.X -= playerSpeed * dt
		}
	}
	// Try move down
	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		p.pos.Y -= playerSpeed * dt
		if p.pos.Y < 0.0 {
			p.pos.Y = 0.0
		}
		if p.Collides(dt, collisions) {
			p.pos.Y += playerSpeed * dt
		}
	}
	// Try move up
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		p.pos.Y += playerSpeed * dt
		if p.pos.Y > win.Bounds().Max.Y {
			p.pos.Y = win.Bounds().Max.Y
		}
		if p.Collides(dt, collisions) {
			p.pos.Y -= playerSpeed * dt
		}
	}

	// Check if player is within a zone
	for r, z := range zones {
		if p.CollidesWith(dt, r) {
			return z
		}
	}
	return ""
}
