package util

import "time"

// Crop crop
type Crop struct {
	name      string
	spriteFmt string
	// How many months to move to next stage
	harvestRate   int
	stage         int
	harvestAmount int
	// Some crops revert after harvest
	revertTo int
}

// NewCrop creates and returns a crop pointer
func NewCrop(name string) *Crop {
	var c Crop

	switch {
	case name == "corn" || name == cornseed:
		c = Crop{
			name:          "Corn",
			spriteFmt:     "corn%d",
			harvestRate:   1,
			stage:         0,
			harvestAmount: 4,
			revertTo:      -1,
		}
	case name == "apple" || name == appleseed:
		c = Crop{
			name:          "Apple Trees",
			spriteFmt:     "appletree%d",
			harvestRate:   3,
			stage:         0,
			harvestAmount: 2,
			revertTo:      3,
		}
	case name == "cotton" || name == cottonseed:
		c = Crop{
			name:          "Cotton",
			spriteFmt:     "cotton%d",
			harvestRate:   1,
			stage:         0,
			harvestAmount: 4,
			revertTo:      -1,
		}
	default:
		return &Crop{"", "", 0, 0, 0, -1}
	}

	go c.RunUpdateLoop()

	return &c
}

// Revert rolls the crop back if that is the crops' setting
// Otherwise it just returns false
func (c *Crop) Revert() bool {
	c.stage = c.revertTo
	if c.revertTo == 0 {
		go c.RunUpdateLoop()
		return true
	}
	return false
}

// RunUpdateLoop updates the crop growth
func (c *Crop) RunUpdateLoop() {
	ticker := time.NewTicker(time.Duration(c.harvestRate) * time.Second * 30)
	for c.stage != 3 {
		// Wait for the rate to pass
		<-ticker.C
		c.stage++
	}

	// Kill the ticker
	ticker.Stop()
}

// IsReady returns whether the crop is ready for harvest
func (c *Crop) IsReady() bool {
	return c.stage == 3
}

// AmountToHarvest returns how much there is to harvest
func (c *Crop) AmountToHarvest() int {
	return c.harvestAmount
}
