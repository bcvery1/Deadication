package mob

import (
  "log"

  "Deadication/util"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  characterImagePath string = "assets/zombie.png"
  charScale float64 = 0.5
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

// Loads the main characters sprite
// BUG(needs to load frames of animation)
func GetCharacterSprite() (*pixel.Sprite, error) {
  pic, err := util.LoadPic(characterImagePath)
  if err != nil {
    return nil, err
  }
  // Get single frame of animation
  sprite := pixel.NewSprite(pic, pic.Bounds())

  return sprite, nil
}

func GetChar() (*CharacterMob, error) {
  stateSprites := make(map[int]*pixel.Sprite)
  // Idle
  sprite, err := GetCharacterSprite()
  stateSprites[0] = sprite
  if err != nil {
    log.Fatal(err)
  }

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
