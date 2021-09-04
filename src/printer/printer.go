package printer

import (
	"fmt"
	. "github.com/smuthusa/fungame/src/model"
)

type Printer func(plane [][]Health)

func ConsolePrinter(plane [][]Health) {
	fmt.Println("\nCurrent Health")
	for _, rowValues := range plane {
		for _, health := range rowValues {
			if health == Dead {
				fmt.Print(fmt.Sprintf(".  "))
			} else {
				fmt.Print(fmt.Sprintf("*  "))
			}
		}
		fmt.Println()
	}
}
