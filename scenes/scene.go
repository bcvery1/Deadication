package scenes

import (
  "log"

  "github.com/faiface/pixel/pixelgl"
)

type Scene struct {
  changeScene *chan string
}

func (s *Scene) Update(win *pixelgl.Window) {}
func (s *Scene) Init() {}

type IScene interface {
  Update(*pixelgl.Window)
  Init()
}

func GetScene(scenename string, changeScene *chan string) IScene {
  switch scenename {
  case "menu":
    return CreateMenu(changeScene)
  case "home":
    return CreateHome(changeScene)
  case "inventory":
    return CreateInventory(changeScene)
  case "farm":
    return CreateFarm(changeScene)
  default:
    log.Printf("Unknown scene %s", scenename)
    return nil
  }
}
