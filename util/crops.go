package util

import "time"

// Crop crop
type Crop struct {
	name string
	// How many months to move to next stage
	harvestRate   int
	stage         int
	harvestAmount int
	ticker        *time.Ticker
	// Some crops revert after harvest
	revertTo int
}

// NewCrop creates and returns a crop pointer
func NewCrop(name string) *Crop {
	var c Crop

	switch name {
	case "corn":
		ticker := time.NewTicker(1 * time.Second * 30)
		c = Crop{"Corn", 1, 0, 4, ticker, -1}
	case "apple":
		ticker := time.NewTicker(3 * time.Second * 30)
		c = Crop{"Apple Trees", 3, 0, 2, ticker, 3}
	case "cotton":
		ticker := time.NewTicker(1 * time.Second * 30)
		c = Crop{"Cotton", 1, 0, 4, ticker, -1}
	default:
		return &Crop{"", 0, 0, 0, nil, -1}
	}

	go c.RunUpdateLoop()

	return &c
}

// RunUpdateLoop updates the crop growth
func (c *Crop) RunUpdateLoop() {
	for c.stage != 4 {
		// Wait for the rate to pass
		<-c.ticker.C
		c.stage++
	}

	// Kill the ticker if the crop dies
	if c.revertTo < 0 {
		c.ticker.Stop()
	}
}

// IsReady returns whether the crop is ready for harvest
func (c *Crop) IsReady() bool {
	return c.stage == 4
}

// AmountToHarvest returns how much there is to harvest
func (c *Crop) AmountToHarvest() int {
	return c.harvestAmount
}
