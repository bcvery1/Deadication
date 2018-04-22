package player

import (
	"time"

	"github.com/bcvery1/Deadication/util"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	playerScale float64 = 1.0
	playerSpeed float64 = 140.0
	hungerAfter         = 4
	hungerBy            = 1
	// How much damage starving does
	starveHurt int = 5
)

var (
	playercycle = 0.0
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
	return (*p).health
}

// Hunger returns the players current health
func (p *Player) Hunger() int {
	return (*p).hunger
}

// Carry causes the player to begin carrying the item
// If already carrying something, it is 'deleted'
func (p *Player) Carry(item string) {
	p.carrying = item
}

// Carrying returns what the player is currently carrying
func (p *Player) Carrying() string {
	return (*p).carrying
}

func (p *Player) setHealth(points int) {
	if points > 100 {
		(*p).health = 100
	} else {
		(*p).health = points
	}
}

func (p *Player) setHunger(points int) {
	if points > 100 {
		(*p).hunger = 100
	} else {
		(*p).hunger = points
	}
}

func (p *Player) statUpdate() {
	for {
		<-time.After(time.Second * hungerAfter)

		p.setHunger((*p).hunger - hungerBy)

		if (*p).hunger < 1 {
			(*p).hunger = 1
			(*p).health -= starveHurt
		}
	}
}

func (p *Player) listenChans() {
	for {
		select {
		case newItem := <-util.PickupChan:
			p.Carry(newItem)
		case hunger := <-util.EatChan:
			p.setHunger((*p).hunger + hunger)
		}
	}
}

// NewPlayer creates a new player object
func NewPlayer(all map[string]*pixel.Sprite) *Player {
	sprites := make(map[string]*pixel.Sprite)
	sprites["idle"] = all["player"]
	sprites["walk1"] = all["playerwalk1"]
	sprites["walk2"] = all["playerwalk2"]
	p := Player{
		sprites,
		"idle",
		pixel.V(120, 120),
		100,
		100,
		"",
	}

	go (&p).statUpdate()
	go (&p).listenChans()

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
	var walking = false
	// Try move right
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		walking = true
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
		walking = true
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
		walking = true
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
		walking = true
		p.pos.Y += playerSpeed * dt
		if p.pos.Y > win.Bounds().Max.Y {
			p.pos.Y = win.Bounds().Max.Y
		}
		if p.Collides(dt, collisions) {
			p.pos.Y -= playerSpeed * dt
		}
	}

	if !walking {
		p.currentAction = "idle"
	} else {
		if playercycle > 0.5 {
			p.currentAction = "walk1"
			if playercycle > 1.0 {
				playercycle = 0.0
			}
		} else {
			p.currentAction = "walk2"
		}
		playercycle += dt
	}

	p.CurrentSprite(dt).Draw(win, pixel.IM.Scaled(p.pos, playerScale).Moved(p.pos))

	// Check if player is within a zone
	for r, z := range zones {
		if p.CollidesWith(dt, r) {
			return z
		}
	}
	return ""
}
