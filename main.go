package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"../alien-gol/maps"
)

func main() {

	rand.Seed(time.Now().Unix())

	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Println("Please provide the number of Aliens to generate and a map file to load!")
		os.Exit(1)
	}

	filename := arguments[1]
	numberOfAliens, err := strconv.Atoi(arguments[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := maps.LoadFromFile(filename)

	ma := maps.SeedAliens(m, numberOfAliens)

	bg := maps.CopyMap(ma)

	iterations := 10000

	maps.StartSimulation(maps.Frame{Foreground: ma, Background: bg}, iterations)
}
