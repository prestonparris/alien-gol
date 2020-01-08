package maps

import (
	"fmt"
	"os"
)

func iterateSimulation(mc Frame) Frame {
	shuffleAliens(mc)

	alienCount := 0
	for range mc.Background.Aliens {
		alienCount++
	}

	if alienCount == 0 {
		fmt.Print("All aliens destroyed!\n")
		os.Exit(1)
	}

	return Frame{Foreground: mc.Background, Background: CopyMap(mc.Background)}
}

func StartSimulation(mc Frame, iterations int) {
	maps := mc
	for i := 0; i < iterations; i++ {
		maps = iterateSimulation(maps)
	}

	fmt.Printf("-- FINISHED all %d iterations ---\n", iterations)

	os.Exit(0)
}
