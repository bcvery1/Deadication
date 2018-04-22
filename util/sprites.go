package util

import (
	"encoding/csv"
	"image"
	"io"
	"log"
	"os"
	"strconv"

	// Required for png decoding
	_ "image/png"

	"github.com/faiface/pixel"
)

const (
	spriteMapPath  string  = "assets/map.png"
	spriteMapCSV   string  = "assets/spriteLayout.csv"
	spriteMapWidth float64 = 32

	assetPlacementPath string = "assets/assetplacement.csv"
)

func loadPic(path string) pixel.Picture {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return pixel.PictureDataFromImage(img)
}

// GetSprites loads all sprite from sprite sheet
func GetSprites() (map[string]*pixel.Sprite, pixel.Picture) {
	pic := loadPic(spriteMapPath)

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

	return spriteMap, pic
}

// TranslateRect gets the gameworld rect from sprite
func TranslateRect(s *pixel.Sprite, iV pixel.Vec) pixel.Rect {
	f := s.Frame()
	bottomL := iV.Add(f.Size().Scaled(-0.5))
	topR := bottomL.Add(f.Size())
	return pixel.R(bottomL.X, bottomL.Y, topR.X, topR.Y)
}

// CreateBatch creates the background batch with sprites drawn on
func CreateBatch(sprites map[string]*pixel.Sprite, pic pixel.Picture) (*pixel.Batch, []pixel.Rect) {
	batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)
	collisions := []pixel.Rect{}

	DrawRiver(batch)

	assetsF, err := os.Open(assetPlacementPath)
	if err != nil {
		log.Fatal(err)
	}
	defer assetsF.Close()

	csvFile := csv.NewReader(assetsF)
	for {
		asset, err := csvFile.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		name := asset[0]
		vX, _ := strconv.ParseFloat(asset[1], 64)
		vY, _ := strconv.ParseFloat(asset[2], 64)
		collides := asset[3] == "true"

		v := pixel.V(vX, vY)
		sprites[name].Draw(batch, pixel.IM.Moved(v))

		if collides {
			r := TranslateRect(sprites[name], v)
			collisions = append(collisions, r)
		}
	}

	return batch, collisions
}

// PointsRect gets the vectors of the corners of a Rect
func PointsRect(r pixel.Rect) (bottomleft, bottomright, topleft, topright pixel.Vec) {
	bottomleft = pixel.V(r.Min.X, r.Min.Y)
	bottomright = pixel.V(r.Max.X, r.Min.Y)
	topleft = pixel.V(r.Min.X, r.Max.Y)
	topright = pixel.V(r.Max.X, r.Max.Y)
	return
}

// RectCollide returns whether there is any collision between two rects
func RectCollide(r1, r2 pixel.Rect) bool {
	if (r1.Max.X < r2.Min.X || r1.Min.X > r2.Max.X) && (r1.Max.Y < r2.Min.Y || r1.Min.Y > r2.Max.Y) {
		return false
	}
	return r1.Intersect(r2).Area() > 0.0
}
