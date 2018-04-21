package scenes

func CreateInventory(changeScene *chan string) *Scene {
  return &Scene{
    changeScene,
  }
}
