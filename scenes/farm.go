package scenes

import (
  "encoding/csv"
  "log"
  "io"
  "os"
  "strconv"

  "Deadication/util"

  "github.com/faiface/pixel"
)

const (
  treeplacementFile string = "assets/treeplacement.csv"
  penlayoutFile string = "assets/penLayout.csv"
)

func CreateFarm(changeScene *chan string, allSprites map[string]*pixel.Sprite, spritePic pixel.Picture) *Scene {
  // Load tree positions
  collidables := []*util.StaticObject{}
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
    t := util.NewStaticObject(allSprites["tree"], pixel.V(x, y), true)
    collidables = append(collidables, t)
  }

  penF, err := os.Open(penlayoutFile)
  if err != nil {
    log.Fatal(err)
  }
  defer penF.Close()

  csvFile = csv.NewReader(penF)
  for {
    pos, err := csvFile.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    x, _ := strconv.ParseFloat(pos[0], 64)
    y, _ := strconv.ParseFloat(pos[1], 64)
    p := util.NewStaticObject(allSprites["pen"], pixel.V(x, y), true)
    collidables = append(collidables, p)
  }

  return &Scene{changeScene, collidables}
}
