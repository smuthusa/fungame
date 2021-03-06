package game

import (
	. "github.com/smuthusa/fungame/src/model"
	"github.com/smuthusa/fungame/src/printer"
	"github.com/smuthusa/fungame/src/seed"
	"github.com/smuthusa/fungame/src/visitor"
)

type Controller struct {
	plane       [][]Health
	boundary    Coordinate
	cellVisitor   visitor.NeighbourCellVisitor
	displayStatus printer.Printer
}

func NewController(boundary Coordinate, visitor visitor.NeighbourCellVisitor, seed seed.Seed, printer printer.Printer) *Controller {
	plane := makeNewPlane(boundary)
	plane = seed(plane)
	return &Controller{
		plane:         plane,
		boundary:      boundary,
		cellVisitor:   visitor,
		displayStatus: printer,
	}
}

func makeNewPlane(boundary Coordinate) [][]Health {
	plane := make([][]Health, boundary.Row)
	for i := range plane {
		plane[i] = make([]Health, boundary.Column)
		for c := range plane[i] {
			plane[i][c] = Dead
		}
	}
	return plane
}

func (c *Controller) Transition() {
	healthCallback := func(row int, column int) Health {
		return c.plane[row][column]
	}
	newPlane := makeNewPlane(c.boundary)
	for rowIndex, rowValues := range c.plane {
		for colIndex, health := range rowValues {
			newHealth := c.cellVisitor.Visit(health, rowIndex, colIndex, healthCallback)
			newPlane[rowIndex][colIndex] = newHealth
		}
	}
	c.plane = ShiftCellsOnReachingBoundary(newPlane, c.boundary)
	c.displayStatus(c.plane)
}