package util

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// House house
type House struct {
	Interactive
	inventory map[string]int
}

// NewHouse creates new house interactive entity
func NewHouse() *House {
	inv := make(map[string]int)
	inv["cornseed"] = 1
	inv["appleseed"] = 0
	inv["cottonseed"] = 1
	inv["food"] = 4

	h := House{
		Interactive{"House", false},
		inv,
	}

	h.initHouse()

	return &h
}

// Update Updates the options on screen
func (h *House) Update(win *pixelgl.Window, carrying string) {
	if !h.IsActive() {
		return
	}

	// Draw box
	imd := getBox()
	imd.Draw(win)

	// Draw title
	title, scale := getText(-1, h.Title(), 1.4, titleV)
	title.Draw(win, scale)

	shiftV := pixel.V(0, 30)
	houseoptions := h.opts(carrying)
	for j, opt := range houseoptions {
		v := menuV.Sub(shiftV.Scaled(float64(j + 1)))
		optTxt, scale := getText(j+1, opt.Text(), 1.1, v)
		optTxt.Draw(win, scale)
	}

	// Check if the user presses a number key to select an option
	doOptions(win, houseoptions, carrying, h)
}

func (h *House) initHouse() {
	go func() {
		for {
			select {
			case item := <-HouseInvChan:
				h.inventory[item]++
			case item := <-TakeFromHouseChan:
				h.inventory[item]--
			}
		}
	}()
}

func (h *House) opts(c string) []optionI {
	opts := []optionI{}

	if c == cornseed {
		o := storeItem{option{"Store corn seed"}}
		opts = append(opts, &o)
	}

	if c == appleseed {
		o := storeItem{option{"Store apple seed"}}
		opts = append(opts, &o)
	}

	if c == cottonseed {
		o := storeItem{option{"Store cotton seed"}}
		opts = append(opts, &o)
	}

	if c == "food" {
		o := storeItem{option{"Store food"}}
		opts = append(opts, &o)
	}

	if c == "" {
		var s string
		s = cornseed
		if h.inventory[s] == 0 && Money >= cornseedPrice {
			o := buy{option{"Buy corn seed 4gp"}, s, cornseedPrice}
			opts = append(opts, &o)
		} else {
			str := fmt.Sprintf("Pickup corn seed [%d]", h.inventory[s])
			o := collect{option{str}, s}
			opts = append(opts, &o)
		}

		s = appleseed
		if h.inventory[s] == 0 && Money >= appleseedPrice {
			o := buy{option{"Buy apple seed 6gp"}, s, appleseedPrice}
			opts = append(opts, &o)
		} else {
			str := fmt.Sprintf("Pickup apple seed [%d]", h.inventory[s])
			o := collect{option{str}, s}
			opts = append(opts, &o)
		}

		s = cottonseed
		if h.inventory[s] == 0 && Money >= cottonseedPrice {
			o := buy{option{"Buy cotton seed 3gp"}, s, cottonseedPrice}
			opts = append(opts, &o)
		} else {
			str := fmt.Sprintf("Pickup cotton seed [%d]", h.inventory[s])
			o := collect{option{str}, s}
			opts = append(opts, &o)
		}
	}

	return opts
}

type storeItem struct{ option }

func (ss *storeItem) Action(f InteractiveI, carrying string) {
	HouseInvChan <- carrying
	PickupChan <- ""
}

type buy struct {
	option
	item  string
	price int
}

func (b *buy) Action(f InteractiveI, carrying string) {
	HouseInvChan <- b.item
	Money -= b.price
}

type collect struct {
	option
	item string
}

func (c *collect) Action(f InteractiveI, carrying string) {
	TakeFromHouseChan <- c.item
	PickupChan <- c.item
}
