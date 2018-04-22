package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	interactivePlacementPath string = "assets/interactiveplacement.csv"
)

var (
	field2 = Interactive{"Mid field", false}
	field3 = Interactive{"Bottom field", false}
	pen1   = Interactive{"Top pen", false}
	pen2   = Interactive{"Mid pen", false}
	pen3   = Interactive{"Bottom pen", false}
	river  = Interactive{"River", false}
)

var field1 = field{
	Interactive: Interactive{"Top field", false},
	havestPerc:  0,
	crop:        crop{"", 0},
	planted:     false,
}

type option struct {
	text string
}

func (o *option) Text() string {
	return o.text
}
func (o *option) Action(i InteractiveI, carrying string) {
	log.Printf("Acting with %s", carrying)
}

type optionI interface {
	Text() string
	Action(InteractiveI, string)
}

// Interactive action for in game objects
type Interactive struct {
	title  string
	active bool
}

// InteractiveI interface for any Interactive element
type InteractiveI interface {
	opts(string) []optionI
	Activate(string, *pixelgl.Window)
	Deactivate()
	Title() string
	IsActive() bool
	Update(*pixelgl.Window)
}

// Update updates the interactive in the game world
func (i *Interactive) Update(win *pixelgl.Window) {
	if !i.IsActive() {
		return
	}
	log.Println("updating")

	imd := imdraw.New(nil)
	imd.Color = colornames.Whitesmoke
	imd.Push(pixel.V(0, 0), pixel.V(300, 175))
	imd.Rectangle(0)
	imd.Draw(win)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	title := text.New(pixel.V(30, 155), atlas)
	title.Color = colornames.Black
	fmt.Fprintf(title, i.Title())
	title.Draw(win, pixel.IM.Scaled(title.Orig, 1.4))
}

// IsActive returns whether this interactive is currently active
func (i *Interactive) IsActive() bool {
	return i.active
}

// Title returns the interactives' title
func (i *Interactive) Title() string {
	return i.title
}

func (i *Interactive) opts(c string) []optionI {
	log.Println("In opts")
	return []optionI{}
}

// Activate the structs function
// Takes what the player is currently carrying
func (i *Interactive) Activate(carrying string, win *pixelgl.Window) {
	i.active = true
	i.opts(carrying)
}

// Deactivate stops the interactives behaivour
func (i *Interactive) Deactivate() {
	i.active = false
}

// AllInteractives gets all interactive entities and collision areas
func AllInteractives() (map[string]InteractiveI, map[pixel.Rect]string) {
	r := make(map[pixel.Rect]string)
	interactiveF, err := os.Open(interactivePlacementPath)
	if err != nil {
		log.Fatal(err)
	}
	defer interactiveF.Close()

	csvFile := csv.NewReader(interactiveF)
	for {
		i, err := csvFile.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		xMin, _ := strconv.ParseFloat(i[0], 64)
		yMin, _ := strconv.ParseFloat(i[1], 64)
		xMax, _ := strconv.ParseFloat(i[2], 64)
		yMax, _ := strconv.ParseFloat(i[3], 64)
		interName := i[4]
		rect := pixel.R(xMin, yMin, xMax, yMax)
		r[rect] = interName
	}

	m := make(map[string]InteractiveI)
	m["field1"] = &field1
	m["field2"] = &field2
	m["field3"] = &field3
	m["pen1"] = &pen1
	m["pen2"] = &pen2
	m["pen3"] = &pen3
	m["river"] = &river

	return m, r
}

func drawBox(i InteractiveI, win *pixelgl.Window) {
	go func(i InteractiveI) {
		for i.IsActive() {
			imd := imdraw.New(nil)
			imd.Color = colornames.Whitesmoke
			imd.Push(pixel.V(0, 0), pixel.V(300, 175))
			imd.Rectangle(0)
			imd.Draw(win)

			atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
			title := text.New(pixel.V(30, 155), atlas)
			title.Color = colornames.Black
			fmt.Fprintf(title, i.Title())
			title.Draw(win, pixel.IM.Scaled(title.Orig, 1.4))
		}
	}(i)
}
