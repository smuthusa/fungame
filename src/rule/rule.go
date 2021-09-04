package rule

import (
	. "github.com/smuthusa/fungame/src/model"
)

type TransitionRule interface {
	Transition(currentHealth Health, liveCellsCount int) Health
}

type DefaultTransitionRule struct{}

func (r *DefaultTransitionRule) Transition(currentHealth Health, liveCellsCount int) Health {
	if currentHealth == Live && canCellSurvive(liveCellsCount) {
		return Live
	} else if currentHealth == Dead && canReincarnate(liveCellsCount) {
		return Live
	} else {
		return Dead
	}
}

func canCellSurvive(currentPopulation int) bool {
	return currentPopulation == 2 || currentPopulation == 3
}

func canReincarnate(currentPopulation int) bool {
	return currentPopulation == 3
}
