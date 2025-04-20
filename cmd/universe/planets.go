package sphere_platform_universe

type Planets struct {
	// Planets - the list of planets.
	Planets []*Planet

	// HasMaster - whether the universe has a master planet.
	HasMaster bool
}

// MakePlanets - make a new planets.
func MakePlanets() *Planets {
	return &Planets{
		Planets:   []*Planet{},
		HasMaster: false,
	}
}

// AddPlanet - add a planet to the planets.
func (p *Planets) AddPlanet(planet *Planet) {
	p.Planets = append(p.Planets, planet)
}
