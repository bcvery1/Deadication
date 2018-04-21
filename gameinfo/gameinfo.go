package gameinfo

import (
  "Deadication/hud"
  "Deadication/mob"
  "Deadication/scenes"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

var (
  allScenes map[string]scenes.IScene = make(map[string]scenes.IScene)
  sceneChange chan string = make(chan string, 1)
)

type GameInfo struct {
  Win *pixelgl.Window
  CamPos *pixel.Vec
  Player *mob.CharacterMob
  ActiveScene scenes.IScene
}

func (g *GameInfo) Update(dt float64) {
  hud.Draw(g.Win, *g.CamPos)
  g.ActiveScene.Update(g.Win, g.CamPos, g.Player, dt)
}

func NewGame(win *pixelgl.Window, camPos *pixel.Vec, player *mob.CharacterMob, initialscene string) *GameInfo {
  allScenes["menu"] = scenes.GetScene("menu", &sceneChange)
  allScenes["home"] = scenes.GetScene("home", &sceneChange)
  allScenes["farm"] = scenes.GetScene("farm", &sceneChange)
  allScenes["inventory"] = scenes.GetScene("inventory", &sceneChange)

  g := GameInfo{
    win,
    camPos,
    player,
    allScenes[initialscene],
  }

  g.ActiveScene.Init()

  go func(g *GameInfo) {
    for !win.Closed() {
      newScene := <- sceneChange
      g.ActiveScene = allScenes[newScene]
      g.ActiveScene.Init()
    }
  }(&g)

  return &g
}
