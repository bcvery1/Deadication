package util

import "log"

type crop struct {
	name        string
	harvestRate int
}

type field struct {
	Interactive
	havestPerc int
	crop       crop
	planted    bool
}

func (f *field) Activate(carrying string) {
	log.Printf("Field activate %s", carrying)
	f.opts(carrying)
}

func (f *field) opts(c string) []optionI {
	observe := observeField{option{"Observe field"}}
	opts := []optionI{&observe}

	if c == "water" {
		o := waterField{option{"Water field"}}
		opts = append(opts, &o)
	}

	if c == "seed" && !f.planted {
		o := plantSeeds{option{"Plant seeds"}}
		opts = append(opts, &o)
	}

	return opts
}

type observeField struct {
	option
}

type waterField struct {
	option
}

func (w *waterField) Action(f InteractiveI, carrying string) {}

type plantSeeds struct {
	option
}