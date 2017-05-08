package main

import (
	"time"

	"github.com/coderconvoy/money"
)

type Client struct {
	Name               string
	DOB                time.Time
	Gender, PrefGender rune
	Location           string
	Locx, locy         float32
	MinPrice, MaxPrice money.M
	Aims               []Aim
}

type Trainer struct {
	Name     string
	DOB      time.Time
	Gender   rune
	Location string
}
