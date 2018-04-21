package mob

import (
  "Deadication/util"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  characterImagePath string = "assets/images/zombie.png"
  charScale float64 = 0.75
  maxHealth int = 100
  maxHunger int = 100
  // How much HP is lost while starving
  starvingHP int = 2
)

type CharacterMob struct {
  Mob
  Health int
  Hunger int
}

func Inv(v pixel.Vec) pixel.Vec {
  return pixel.Vec{
    -v.X,
    -v.Y,
  }
}

func rectCollide(r1, r2 pixel.Rect) bool {
  xCol1 := (r1.Max.X > r2.Min.X && r1.Max.X < r2.Max.X) || (r1.Min.X > r2.Min.X && r1.Min.X < r2.Max.X)
  yCol1 := (r1.Max.Y > r2.Min.Y && r1.Max.Y < r2.Max.Y) || (r1.Min.Y > r2.Min.Y && r1.Min.Y < r2.Max.Y)

  xCol2 := (r2.Max.X > r1.Min.X && r2.Max.X < r1.Max.X) || (r2.Min.X > r1.Min.X && r2.Min.X < r1.Max.X)
  yCol2 := (r2.Max.Y > r1.Min.Y && r2.Max.Y < r1.Max.Y) || (r2.Min.Y > r1.Min.Y && r2.Min.Y < r1.Max.Y)

  return (xCol1 && yCol1) || (xCol2 && yCol2)
}

func (c *CharacterMob) CollidesRects(rects []pixel.Rect, camPos pixel.Vec) bool {
  playRect := util.GetSpriteRect(c.Sprites[c.State], camPos)
  for _, r := range rects {
    if rectCollide(playRect, r) {
      return true
    }
  }

  return false
}

func (c *CharacterMob) Collides(statics []*util.StaticObject, camPos pixel.Vec) bool {
  playRect := util.GetSpriteRect(c.Sprites[c.State], camPos)
  for _, obj := range statics {
    objRect := util.GetSpriteRect(obj.Sprite, obj.PosV)
    if !obj.Collision {
      continue
    }
    if rectCollide(playRect, objRect) {
      return true
    }
  }

  return false
}

func (c CharacterMob) Update(win *pixelgl.Window, camPos pixel.Vec) {
  c.Sprites[c.State].Draw(win, pixel.IM.Moved(camPos).Scaled(camPos, charScale))
}

// Increase hunger
func (c *CharacterMob) Eat(points int) {
  c.Hunger += points
  if c.Hunger > maxHunger {
    c.Hunger = maxHunger
  }
}

func (c *CharacterMob) GetHungry(points int) {
  c.Hunger -= points
  if c.Hunger <= 0 {
    // Player is starving
    c.Hurt(starvingHP)
  }
}

// Injures the player
// Returns if the player is dead
func (c *CharacterMob) Hurt(hp int) bool {
  c.Health -= hp
  return c.Health <= 0
}

func (c *CharacterMob) Heal(hp int) {
  c.Health += hp
  if c.Health > maxHealth {
    c.Health = maxHealth
  }
}

func GetChar(stateSprites map[int]*pixel.Sprite) (*CharacterMob, error) {
  return &CharacterMob{
    Mob{
      stateSprites,
      0,
      pixel.IM,
    },
    maxHealth,
    maxHunger,
  }, nil
}
