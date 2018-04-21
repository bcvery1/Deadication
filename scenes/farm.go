package scenes

import (
  "log"

  "Deadication/mob"
  "Deadication/util"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
)

const (
  houseImg string = "assets/images/house.png"
  treeImg string = "assets/images/tree.png"
)

type FarmScene struct {
  Scene
  houseImg    *pixel.Sprite
  treesbatch  *pixel.Batch
  treeObjs    []util.StaticObjects
}

func (f *FarmScene) Update(win *pixelgl.Window, camPos pixel.Vec) {
  f.houseImg.Draw(win, pixel.IM.Moved(pixel.V(0, 384)))

}

func (f *FarmScene) Init() {

}

func CreateFarm(changeScene *chan string) *FarmScene {
  housePic, err := util.LoadPic(houseImg)
  if err != nil {
    log.Fatal(err)
  }
  houseSprite := pixel.NewSprite(housePic, housePic.Bounds())

  treePic, err := util.LoadPic(treeImg)
  if err != nil {
    log.Fatal(err)
  }
  batch := pixel.NewBatch(&pixel.TrianglesData{}, treePic)

  t1 := util.NewStaticObject(f.treePics[0], pixel.V(100, 100), true)

  return &FarmScene{
    Scene{changeScene},
    houseSprite,
    batch,
    []util.StaticObjects{t1},
  }
}
