package gameinfo

import (
  "encoding/csv"
  "log"
  "io"
  "os"
  "strconv"

  "Deadication/hud"
  "Deadication/mob"
  "Deadication/scenes"
  "Deadication/util"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  spriteMapCSV string = "assets/images/spriteLayout.csv"
  spriteMapPath string = "assets/images/map.png"
  spriteMapWidth float64 = 32
)

var (
  allScenes map[string]scenes.IScene = make(map[string]scenes.IScene)
  sceneChange chan string = make(chan string, 1)
)

type GameInfo struct {
  Win         *pixelgl.Window
  CamPos      *pixel.Vec
  Player      *mob.CharacterMob
  ActiveScene scenes.IScene
  SpriteMap   map[string]*pixel.Sprite
  Batch       *pixel.Batch
}

func (g *GameInfo) Update(dt float64) {
  g.ActiveScene.Update(g.Win, g.CamPos, g.Player, dt)
  hud.Draw(g.Win, *g.CamPos)
}

func NewGame(win *pixelgl.Window, camPos *pixel.Vec, initialscene string) *GameInfo {
  allScenes["menu"] = scenes.GetScene("menu", &sceneChange)
  allScenes["home"] = scenes.GetScene("home", &sceneChange)
  allScenes["farm"] = scenes.GetScene("farm", &sceneChange)
  allScenes["inventory"] = scenes.GetScene("inventory", &sceneChange)

  g := GameInfo{
    win,
    camPos,
    &(mob.CharacterMob{}),
    allScenes[initialscene],
    make(map[string]*pixel.Sprite),
    &(pixel.Batch{}),
  }

  pic, err := util.LoadPic(spriteMapPath)
  if err != nil {
    log.Fatal(err)
  }

  spriteF, err := os.Open(spriteMapCSV)
  if err != nil {
    log.Fatal(err)
  }
  defer spriteF.Close()
  csvFile := csv.NewReader(spriteF)
  for {
    spr, err := csvFile.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    name := spr[0]
    x, _ := strconv.ParseFloat(spr[1], 64)
    y, _ := strconv.ParseFloat(spr[2], 64)
    w, _ := strconv.ParseFloat(spr[3], 64)
    h, _ := strconv.ParseFloat(spr[4], 64)
    r := pixel.R(x*spriteMapWidth, y*spriteMapWidth, w*spriteMapWidth+x*spriteMapWidth, h*spriteMapWidth+y*spriteMapWidth)
    g.SpriteMap[name] = pixel.NewSprite(pic, r)
  }

  charSprites := make(map[int]*pixel.Sprite)
  charSprites[0] = g.SpriteMap["player"]
  g.Player, err = mob.GetChar(charSprites)
  if err != nil {
    log.Fatal(err)
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
