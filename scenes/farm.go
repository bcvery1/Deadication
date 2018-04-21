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
)

func CreateFarm(changeScene *chan string, allSprites map[string]*pixel.Sprite, spritePic pixel.Picture) *Scene {
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
    t := util.NewStaticObject(allSprites["tree"], pixel.V(x, y), true)
    trees = append(trees, t)
  }

  return &Scene{changeScene, trees}
}
