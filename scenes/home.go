package scenes

func CreateHome(changeScene *chan string) *Scene {
  return &Scene{
    changeScene,
  }
}
