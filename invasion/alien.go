package invasion

import (
	"math/rand"
	"time"
)

const minCities = 10000

type Health string

var (
	HealthAlive  Health = "Alive"
	HealthKilled Health = "Killed"
)

// Alien represents an alien invading the world.
type Alien struct {
	Name   string
	Health Health

	VisitingCity  string
	CitiesVisited int
}

// NewAlien creates a new Alien with the given name.
func NewAlien(name string) *Alien {
	return &Alien{
		Name:   name,
		Health: HealthAlive,
	}
}

// Land places the alien at a random city in the world.
// It returns true if the landing was successful.
func (a *Alien) Land(w *World) bool {
	city, ok := w.RandomCity()
	if !ok {
		return false
	}
	a.VisitingCity = city.Name

	if ok := w.Visit(city, a); !ok {
		return false
	}

	a.CitiesVisited++
	return true
}

// Travel moves the alien to another random city using available roads.
// It returns true if the alien successfully moved to another city, false otherwise.
func (a *Alien) Travel(w *World) bool {
	if a.CitiesVisited == minCities {
		return false
	}
	// we need to determine the roads still available
	currentCity, found := w.FindCity(a.VisitingCity)
	if !found || len(currentCity.Roads) == 0 {
		// alien is trapped
		return false
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	road := currentCity.Roads[r.Intn(len(currentCity.Roads))]
	nextCity, found := w.FindCity(road.City)
	if !found {
		return false
	}
	a.VisitingCity = nextCity.Name

	if ok := w.Visit(nextCity, a); ok {
		a.CitiesVisited++
	}
	return true
}

func (a *Alien) Kill() {
	a.Health = HealthKilled
}

// CitiesVisitedCount returns the number of cities the alien visited so far.
func (a *Alien) CitiesVisitedCount() int {
	return a.CitiesVisited
}
