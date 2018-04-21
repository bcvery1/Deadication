package scenes

func CreateFarm(changeScene *chan string) *Scene {
  return &Scene{
    changeScene,
  }
}
