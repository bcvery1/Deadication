package scenes

import (
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

type MenuScene struct {
  Scene
}

func (s *MenuScene) Update(win *pixelgl.Window, camPos pixel.Vec) {
  *s.changeScene <- "home"
}

func (s *MenuScene) Init() {}

func CreateMenu(changeScene *chan string) *MenuScene {
  return &MenuScene{
    Scene{changeScene},
  }
}
