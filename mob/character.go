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
)

type CharacterMob struct {
  Mob
}

func (c CharacterMob) Update(win *pixelgl.Window, camPos pixel.Vec) {
  c.Sprites[c.State].Draw(win, pixel.IM.Moved(camPos).Scaled(camPos, charScale))
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
  }, nil
}
