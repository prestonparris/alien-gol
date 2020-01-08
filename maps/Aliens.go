package maps

import (
	"fmt"
	"math/rand"
)

func SeedAliens(m *Map, N int) *Map {

	numberOfCities := len(m.Cities)

	alienId := 0

	for i := 0; i < N; i++ {
		j := 0

		r := rand.Intn(numberOfCities)

		for k := range m.Cities {
			if j == r {
				alien := Alien{Id: alienId}
				m.AddAlien(k, alien)
				alienId++
				break
			}
			j++
		}

	}

	return m
}

func shuffleAliens(mc Frame) Frame {

	// iterate over each city containing aliens
	// for each alien choose a random city that it can hop to
	for c := range mc.Foreground.Aliens {
		for a := range mc.Foreground.Aliens[c] {
			size := len(mc.Foreground.Roads[c])
			if size == 0 {
				break
			}

			rCity := rand.Intn(size)

			i := 0

			for r := range mc.Foreground.Roads[c] {
				// we found the random city destination move the alien
				if i == rCity {
					moveAlien(mc.Background, c, r, a)
					break
				}

				i++
			}
		}
	}

	return mc
}

func moveAlien(m *Map, src City, dst City, alien Alien) {

	m.RemoveAlien(src, alien)

	if _, ok := m.Aliens[dst]; ok {
		if len(m.Aliens[dst]) > 0 {
			for t := range m.Aliens[dst] {
				alien1 := t
				alien2 := alien

				if _, ok := m.Cities[dst]; ok {
					fmt.Printf("%s has been destroyed by alien %d and alien %d\n",
						dst.Name, alien1, alien2)

					m.RemoveCity(dst)
				}

				break
			}
		} else {
			m.AddAlien(dst, alien)
		}
	} else {
		m.AddAlien(dst, alien)
	}

}
