package game

import (
	. "github.com/smuthusa/fungame/src/model"
	"github.com/smuthusa/fungame/src/util"
)

func ShiftCellsOnReachingBoundary(plane [][]Health, boundary Coordinate) [][]Health {
	rowWithLiveCell := map[int]bool{}
	columnWithLiveCell := map[int]bool{}

	for row, rowStatus := range plane {
		for column, health := range rowStatus {
			columnWithLiveCell[column] = updateNewStatus(columnWithLiveCell, column, health)
			rowWithLiveCell[row] = updateNewStatus(rowWithLiveCell, column, health)
		}
	}

	liveCellRows := indicesByLiveCellFn(rowWithLiveCell)
	liveCellColumns := indicesByLiveCellFn(columnWithLiveCell)

	lastRowWithLiveCell, rowErr := util.Max(liveCellRows)
	lastColumnWithLiveCell, colErr := util.Max(liveCellColumns)

	if rowErr == util.EmptyArrayError && colErr == util.EmptyArrayError {
		return plane
	}
	if int(boundary.Row)-1 == lastRowWithLiveCell || int(boundary.Column)-1 == lastColumnWithLiveCell {
		firstRowWithLiveCell, _ := util.Min(liveCellRows)
		firstColumnWithLiveCell, _ := util.Min(liveCellColumns)
		return shiftCellPosition(plane, firstRowWithLiveCell-1, firstColumnWithLiveCell-1)
	}
	return plane
}

func updateNewStatus(statusMap map[int]bool, column int, health Health) bool {
	if val, ok := statusMap[column]; ok {
		return val || health == Live
	} else {
		return health == Live
	}
}

func shiftCellPosition(plane [][]Health, rowsToShift int, colsToShift int) [][]Health {
	for row, rowValues := range plane {
		for column, health := range rowValues {
			if health == Live {
				plane[row-rowsToShift][column-colsToShift] = health
				plane[row][column] = Dead
			}
		}
	}
	return plane
}

func indicesByLiveCellFn(liveStatus map[int]bool) []int {
	var indices []int
	for key, value := range liveStatus {
		if value == true {
			indices = append(indices, key)
		}
	}
	return indices
}
