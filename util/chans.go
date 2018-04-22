package util

// PickupChan - All pickups/drops for the player will run through here
var PickupChan = make(chan string, 1)

// EatChan - When the player needs to eat, send to here
var EatChan = make(chan int, 1)

// PopupChan - Displays messages to screen
var PopupChan = make(chan *Popup, 1)
