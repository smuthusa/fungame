package main

import (
	"time"

	"github.com/smuthusa/fungame/src/game"
	"github.com/smuthusa/fungame/src/model"
	"github.com/smuthusa/fungame/src/printer"
	"github.com/smuthusa/fungame/src/rule"
	"github.com/smuthusa/fungame/src/seed"
	"github.com/smuthusa/fungame/src/visitor"
)

func main() {
	boundary := model.Coordinate{Row: 25, Column: 25}
	transitionInterval := time.Second
	cellVisitor := visitor.NewNeighbourCellVisitor(&rule.DefaultTransitionRule{}, boundary)
	planeController := game.NewController(boundary, cellVisitor, seed.GliderSeed, printer.ConsolePrinter)

	for {
		planeController.Transition()
		time.Sleep(transitionInterval)
	}
}
