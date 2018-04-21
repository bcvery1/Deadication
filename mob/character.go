package mob

import (
  "Deadication/util"

  "github.com/faiface/pixel"
)

const (
  characterImagePath string = "assets/zombie.png"
)

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
