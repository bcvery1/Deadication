// Package is for general pixel utility functions
package util

import (
  "image"
  "os"

  _ "image/png"

  "github.com/faiface/pixel"
)

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
