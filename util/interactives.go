package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/faiface/pixel"
)

const (
	interactivePlacementPath string = "assets/interactiveplacement.csv"
)

var (
	field2 = Interactive{"Mid field"}
	field3 = Interactive{"Bottom field"}
	pen1   = Interactive{"Top pen"}
	pen2   = Interactive{"Mid pen"}
	pen3   = Interactive{"Bottom pen"}
	river  = Interactive{"River"}
)

var field1 = field{
	Interactive: Interactive{"Top field"},
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
	title string
}

// InteractiveI interface for any Interactive element
type InteractiveI interface {
	opts(string) []optionI
	Activate(string)
	Deactivate()
}

func (i *Interactive) opts(c string) []optionI {
	log.Println("In opts")
	return []optionI{}
}

// Activate the structs function
// Takes what the player is currently carrying
func (i *Interactive) Activate(carrying string) {
	log.Println(i.title)
	i.opts(carrying)
}

// Deactivate stops the interactives behaivour
func (i *Interactive) Deactivate() {
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
