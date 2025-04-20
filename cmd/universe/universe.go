package sphere_platform_universe

type Universe struct {
	Planets *Planets
}

// MakeUniverse - make a new universe.
func MakeUniverse() *Universe {
	return &Universe{
		Planets: MakePlanets(),
	}
}
