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
	plane := newPlane(boundary)
	plane = seed(plane)
	return &Controller{
		plane:         plane,
		boundary:      boundary,
		cellVisitor:   visitor,
		displayStatus: printer,
	}
}

func newPlane(boundary Coordinate) [][]Health {
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
	boundary := c.boundary
	newPlane := newPlane(boundary)
	for rowIndex, rowValues := range c.plane {
		for colIndex, health := range rowValues {
			newHealth := c.cellVisitor.Visit(health, rowIndex, colIndex, healthCallback)
			newPlane[rowIndex][colIndex] = newHealth
		}
	}
	c.displayStatus(newPlane)
	c.plane = ShiftCellsOnReachingBoundary(newPlane, boundary)
	c.displayStatus(c.plane)
}

func (c *Controller) shiftPlane(plane [][]Health, shift Coordinate) [][]Health {
	for row, rowValues := range plane {
		for column, health := range rowValues {
			if health == Live {
				plane[uint(row)-shift.Row][uint(column)-shift.Column] = health
			}
		}
	}
	return plane
}
