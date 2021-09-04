package seed

import (
	. "github.com/smuthusa/fungame/src/model"
)

type Seed func(plane [][]Health) [][]Health

var gliderPositions = [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}

func GliderSeed(plane [][]Health) [][]Health {
	for _, position := range gliderPositions {
		plane[position[0]][position[1]] = Live
	}
	return plane
}