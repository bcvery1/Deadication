package scenes

import (
  "encoding/csv"
  "log"
  "io"
  "os"
  "strconv"

  "Deadication/mob"
  "Deadication/util"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  houseImg string = "assets/images/house.png"
  treeImg string = "assets/images/tree.png"
  treeplacementFile string = "assets/treeplacement.csv"
)

var (
  treePic pixel.Picture
)

type FarmScene struct {
  Scene
  houseImg        *pixel.Sprite
  collidableBatch *pixel.Batch
  collidableObjs  []*util.StaticObject
  init            bool
}

func (f *FarmScene) Update(win *pixelgl.Window, camPos *pixel.Vec, char *mob.CharacterMob, dt float64) {
  f.houseImg.Draw(win, pixel.IM.Moved(pixel.V(0, 384)))
  f.collidableBatch.Draw(win)

  char.Update(win, *camPos)

  newCamPos := util.MoveCamera(win, camPos, dt)
  if !char.Collides(f.collidableObjs, newCamPos) {
    (*camPos) = newCamPos
  }
}

func (f *FarmScene) Init() {
  if f.init {
    return
  }

  f.collidableBatch.Clear()
  for _, obj := range f.collidableObjs {
    obj.Sprite.Draw(f.collidableBatch, pixel.IM.Moved(obj.PosV))
  }
  f.init = true
}

func CreateFarm(changeScene *chan string) *FarmScene {
  housePic, err := util.LoadPic(houseImg)
  if err != nil {
    log.Fatal(err)
  }
  houseSprite := pixel.NewSprite(housePic, housePic.Bounds())

  treePic, err = util.LoadPic(treeImg)
  if err != nil {
    log.Fatal(err)
  }
  batch := pixel.NewBatch(&pixel.TrianglesData{}, treePic)

  // Load tree positions
  trees := []*util.StaticObject{}
  tpF, err := os.Open(treeplacementFile)
  if err != nil {
    log.Fatal(err)
  }
  defer tpF.Close()

  csvFile := csv.NewReader(tpF)
  for {
    pos, err := csvFile.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    x, err := strconv.ParseFloat(pos[0], 64)
    if err != nil {continue}
    y, err := strconv.ParseFloat(pos[1], 64)
    if err != nil {continue}
    t := util.NewStaticObject(treePic, pixel.V(x, y), true)
    trees = append(trees, t)
  }

  return &FarmScene{
    Scene{changeScene},
    houseSprite,
    batch,
    trees,
    false,
  }
}
