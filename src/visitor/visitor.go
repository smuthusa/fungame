package visitor

import (
	. "github.com/smuthusa/fungame/src/model"
	"github.com/smuthusa/fungame/src/rule"
)

type NeighbourCellVisitor struct {
	transitionRule rule.TransitionRule
	boundary       Coordinate
}

func NewNeighbourCellVisitor(transitionRule rule.TransitionRule, boundary Coordinate) NeighbourCellVisitor {
	return NeighbourCellVisitor{
		transitionRule: transitionRule,
		boundary:       boundary,
	}
}

var neighboursRelativePath = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1} /*{0, 0},*/, {0, 1}, //{0, 0} is current position, ignore traversing it.
	{1, -1}, {1, 0}, {1, 1},
}

func (v *NeighbourCellVisitor) Visit(currentHealth Health, row int, column int, healthByPosition HealthByPosition) Health {
	liveNeighbours := 0
	for _, path := range neighboursRelativePath {
		neighbourRow := row + path[0]
		neighbourCol := column + path[1]
		if isWithinBoundary(neighbourRow, v.boundary.Row) && isWithinBoundary(neighbourCol, v.boundary.Column) {
			health := healthByPosition(neighbourRow, neighbourCol)
			if health == Live {
				liveNeighbours++
			}
		}
	}
	return v.transitionRule.Transition(currentHealth, liveNeighbours)
}

func isWithinBoundary(pos int, maxPos uint) bool {
	return pos >= 0 && pos < int(maxPos)
}
