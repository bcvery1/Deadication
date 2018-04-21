package scenes

import (
  "log"

  "Deadication/mob"
  "Deadication/util"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

type Scene struct {
  changeScene *chan string
  collidableObjs  []*util.StaticObject
}

func (s *Scene) Update(win *pixelgl.Window, camPos *pixel.Vec, char *mob.CharacterMob, dt float64, allSprites map[string]*pixel.Sprite) {
  util.MoveCamera(win, camPos, dt)
}

func (s *Scene) Init(batch *pixel.Batch) {
  batch.Clear()
  for _, obj := range s.collidableObjs {
    obj.Sprite.Draw(batch, pixel.IM.Moved(obj.PosV))
  }
}

func (s *Scene) GetCollidables() []*util.StaticObject {
  return s.collidableObjs
}

type IScene interface {
  Update(*pixelgl.Window, *pixel.Vec, *mob.CharacterMob, float64, map[string]*pixel.Sprite)
  Init(*pixel.Batch)
  GetCollidables() []*util.StaticObject
}

func GetScene(scenename string, changeScene *chan string, allSprites map[string]*pixel.Sprite, spritePic pixel.Picture) IScene {
  switch scenename {
  case "menu":
    return CreateMenu(changeScene, allSprites, spritePic)
  case "home":
    return CreateHome(changeScene, allSprites, spritePic)
  case "inventory":
    return CreateInventory(changeScene, allSprites, spritePic)
  case "farm":
    return CreateFarm(changeScene, allSprites, spritePic)
  default:
    log.Printf("Unknown scene %s", scenename)
    return nil
  }
}
