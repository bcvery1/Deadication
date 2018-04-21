// Package is for general pixel utility functions
package util

import (
  "image"
  "os"

  _ "image/png"

  "github.com/faiface/pixel"
)

var (
  shiftDrawingBy pixel.Vec = pixel.Vec{16.0, 0.0}
)

type StaticObject struct {
  Sprite    *pixel.Sprite
  PosV      pixel.Vec
  Collision bool
}

func (o *StaticObject) Draw(t pixel.Target) {
  o.Sprite.Draw(t, pixel.IM.Moved(o.PosV.Add(shiftDrawingBy)).Scaled(o.PosV, 2.0))
}

func NewStaticObject(spr *pixel.Sprite, pos pixel.Vec, canCollide bool) *StaticObject {
  return &StaticObject{
    spr,
    pos,
    canCollide,
  }
}

// Attempts to load pixel.Picture from image file
// Image file should be PNG
// BUG(Does not check image is png)
func LoadPic(path string) (pixel.Picture, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  img, _, err := image.Decode(f)
  if err != nil {
    return nil, err
  }

  return pixel.PictureDataFromImage(img), nil
}

func GetSpriteRect(s *pixel.Sprite, camPos pixel.Vec) pixel.Rect {
  frame := s.Frame()
  size := frame.Size()
  return frame.Moved(camPos).Moved(size.Scaled(-0.5))
}
