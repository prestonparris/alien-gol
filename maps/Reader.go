package maps

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func LoadFromFile(filename string) *Map {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	m := NewMap()

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading in file %s", err)
		}

		items := strings.FieldsFunc(line, unicode.IsSpace)

		cityName := items[0]

		currentCity := City{
			Name: cityName,
		}

		m.AddCity(currentCity)

		for i, item := range items {
			if i > 0 {
				x := strings.Split(item, "=")
				destinationName := x[1]
				m.AddRoad(currentCity, City{Name: destinationName})
			}
		}
	}

	return m
}
