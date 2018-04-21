package scenes

import (
  "Deadication/util"

  "github.com/faiface/pixel"
)

func CreateMenu(changeScene *chan string, allSprites map[string]*pixel.Sprite, spritePic pixel.Picture) *Scene {
  return &Scene{
    changeScene,
    []*util.StaticObject{},
  }
}
