package maps

type City struct {
	Name string
}

type Alien struct {
	Id int
}

type Map struct {
	Cities map[City]struct{}
	Roads map[City]map[City]struct{}
	Aliens map[City]map[Alien]struct{}
}

type Frame struct {
	Foreground *Map
	Background *Map
}

func NewMap() *Map {
	return &Map{
		Cities: make(map[City]struct{}),
		Roads: make(map[City]map[City]struct{}),
		Aliens: make(map[City]map[Alien]struct{}),
	}
}

func CopyMap(m *Map) *Map {
	nm := NewMap()

	for c := range m.Cities {
		nm.AddCity(c)
	}

	for c := range m.Roads {
		for r := range m.Roads[c] {
			nm.AddRoad(c, r)
		}
	}

	for c := range m.Aliens {
		for a := range m.Aliens[c] {
			nm.AddAlien(c, a)
		}
	}

	return nm
}

func (m *Map) AddCity(city City) bool {
	if _, ok := m.Cities[city]; ok {
		return false
	}

	m.Cities[city] = struct{}{}

	return true
}

func (m *Map) RemoveCity(city City) bool {
	if _, ok := m.Cities[city]; ok {
		delete(m.Cities, city)
	}

	if _, ok := m.Roads[city]; ok {
		delete(m.Roads, city)
	}

	for k := range m.Roads {
		for k2 := range m.Roads[k] {
			if k2 == city {
				delete(m.Roads[k], k2)
			}
		}
	}

	if _, ok := m.Aliens[city]; ok {
		delete(m.Aliens, city)
	}

	return true
}

func (m *Map) AddRoad(src, destination City) {
	if _, ok := m.Cities[src]; !ok {
		m.AddCity(src)
	}

	if _, ok := m.Cities[destination]; !ok {
		m.AddCity(destination)
	}

	if _, ok := m.Roads[src]; !ok {
		m.Roads[src] = make(map[City]struct{})
	}

	m.Roads[src][destination] = struct{}{}
}

func (m *Map) AddAlien(city City, alien Alien) {
	if _, ok := m.Aliens[city]; !ok {
		m.Aliens[city] = make(map[Alien]struct{})
	}

	m.Aliens[city][alien] = struct{}{}
}

func (m *Map) RemoveAlien(city City, alien Alien) {
	if _, ok := m.Aliens[city]; !ok {
		m.Aliens[city] = make(map[Alien]struct{})
	}

	delete(m.Aliens[city], alien)

	if len(m.Aliens[city]) == 0 {
		delete(m.Aliens, city)
	}
}
