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
  houseX float64 = 0
  houseY float64 = 384
)

var (
  allScenes map[string]scenes.IScene  = make(map[string]scenes.IScene)
  sceneChange chan string             = make(chan string, 1)
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
  g.ActiveScene.Update(g.Win, g.CamPos, g.Player, dt, g.SpriteMap)

  houseV := pixel.V(houseX, houseY)
  g.SpriteMap["house"].Draw(g.Win, pixel.IM.Moved(houseV).Scaled(houseV, 2.5))

  g.Batch.Draw(g.Win)

  g.Player.Update(g.Win, *(g.CamPos))

  houseRec := []pixel.Rect{pixel.R(-200, 790, -37, 1080), pixel.R(40, 790, 184, 1080), pixel.R(-195, 935, 184, 1080)}

  newCamPos := util.MoveCamera(g.Win, g.CamPos, dt)
  if !g.Player.Collides(g.ActiveScene.GetCollidables(), newCamPos) && !g.Player.CollidesRects(houseRec, newCamPos) {
    g.CamPos = &newCamPos
  }
  hud.Draw(g.Win, *g.CamPos)
}

func NewGame(win *pixelgl.Window, camPos *pixel.Vec, initialscene string) *GameInfo {
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
  spriteMap := make(map[string]*pixel.Sprite)
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
    spriteMap[name] = pixel.NewSprite(pic, r)
  }

  charSprites := make(map[int]*pixel.Sprite)
  charSprites[0] = spriteMap["player"]
  player, err := mob.GetChar(charSprites)
  if err != nil {
    log.Fatal(err)
  }

  allScenes["menu"] = scenes.GetScene("menu", &sceneChange, spriteMap, pic)
  allScenes["home"] = scenes.GetScene("home", &sceneChange, spriteMap, pic)
  allScenes["farm"] = scenes.GetScene("farm", &sceneChange, spriteMap, pic)
  allScenes["inventory"] = scenes.GetScene("inventory", &sceneChange, spriteMap, pic)

  batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)

  g := GameInfo{
    win,
    camPos,
    player,
    allScenes[initialscene],
    spriteMap,
    batch,
  }

  g.ActiveScene.Init(g.Batch)

  go func(g *GameInfo) {
    for !win.Closed() {
      newScene := <- sceneChange
      g.ActiveScene = allScenes[newScene]
      g.ActiveScene.Init(g.Batch)
    }
  }(&g)

  return &g
}
