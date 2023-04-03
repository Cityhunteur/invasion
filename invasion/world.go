package invasion

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

var ErrInvalidMap = errors.New("invalid map data")

// World represents a non-existent world containing cities.
type World struct {
	Cities map[string]*City
	Aliens map[string]*Alien
}

// NewWorld constructs a new World using map data read from r.
// The data is a string encoded format of the world. Example:
//
//	Foo north=Bar west=Baz south=Qu-ux
//	Bar south=Foo west=Bee
func NewWorld(r io.Reader) (*World, error) {
	w := &World{
		Cities: make(map[string]*City),
		Aliens: make(map[string]*Alien),
	}
	if err := w.loadMap(r); err != nil {
		return nil, fmt.Errorf("failed to create new world: %w", err)
	}
	return w, nil
}

// Visit indicates the specified Alien is visiting a random city in the world.
// visited is set to true if successful, false otherwise.
func (w *World) Visit(city *City, a *Alien) (visited bool) {
	w.Aliens[a.Name] = a
	visited = true
	if otherAlien, ok := city.Visit(a); !ok {
		// there's been a fight
		w.DestroyCity(city)
		w.DestroyAlien(otherAlien)
		fmt.Printf("%s has been destroyed by %s and %s!\n", city.Name, a.Name, otherAlien.Name)
		visited = false
	}
	return
}

// DestroyCity destroys the specified city in the world.
func (w *World) DestroyCity(c *City) {
	for _, r := range c.Roads {
		otherCity, exists := w.Cities[r.City]
		if !exists {
			continue
		}
		otherCity.RemoveRoad(c.Name)
	}
	delete(w.Cities, c.Name)
}

// DestroyAlien destroys the specified alien from the world.
func (w *World) DestroyAlien(a *Alien) {
	delete(w.Aliens, a.Name)
}

// FindCity returns the City with matching name, if found.
func (w *World) FindCity(name string) (city *City, found bool) {
	city, found = w.Cities[name]
	return
}

func (w *World) IsDestroyed() bool {
	return len(w.Cities) == 0
}

// RandomCity returns a randomly selected city in the world, if any.
// It sets found to true if one is found, otherwise false.
func (w *World) RandomCity() (city *City, found bool) {
	if len(w.Cities) == 0 {
		return nil, false
	}
	cityNames := make([]string, 0)
	for k := range w.Cities {
		cityNames = append(cityNames, k)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(len(cityNames))
	city = w.Cities[cityNames[n]]
	found = true

	return
}

// CityCount returns the number of cities.
func (w *World) CityCount() int {
	return len(w.Cities)
}

// String returns a string representation of the world map.
// Note the original ordering of cities in the original map is not preserved.
func (w *World) String() string {
	var b strings.Builder
	ncity := 0
	for _, c := range w.Cities {
		ncity++
		b.WriteString(c.String())
		if ncity < len(w.Cities) {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (w *World) loadMap(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		subs := strings.Split(scanner.Text(), " ")
		if len(subs) < 1 || subs[0] == "" {
			return ErrInvalidMap
		}
		cityName := subs[0]
		city := &City{Name: cityName}
		for _, road := range subs[1:] { // omit city name
			parts := strings.Split(road, "=")
			if len(parts) != 2 {
				return ErrInvalidMap
			}
			if !IsValidDirection(parts[0]) {
				return ErrInvalidMap
			}
			city.Roads = append(city.Roads, &Road{
				Direction: Direction(parts[0]),
				City:      parts[1],
			})
		}
		w.Cities[cityName] = city
	}
	return nil
}
