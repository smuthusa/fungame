package model

type Health int

const (
	Dead Health = iota
	Live Health = iota
)

type Coordinate struct {
	Row    uint
	Column uint
}

type HealthByPosition func(row int, column int) Health
