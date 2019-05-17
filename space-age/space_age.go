package space

type Planet string

const (
	mercury            Planet  = "Mercury"
	venus              Planet  = "Venus"
	earth              Planet  = "Earth"
	mars               Planet  = "Mars"
	jupiter            Planet  = "Jupiter"
	saturn             Planet  = "Saturn"
	uranus             Planet  = "Uranus"
	neptune            Planet  = "Neptune"
	secondsInEarthYear float64 = 31557600
)

var secondsInYearByPlanet = map[Planet]float64{
	mercury: secondsInEarthYear * 0.2408467,
	venus:   secondsInEarthYear * 0.61519726,
	earth:   secondsInEarthYear,
	mars:    secondsInEarthYear * 1.8808158,
	jupiter: secondsInEarthYear * 11.862615,
	saturn:  secondsInEarthYear * 29.447498,
	uranus:  secondsInEarthYear * 84.016846,
	neptune: secondsInEarthYear * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsInYearByPlanet[planet];
}
