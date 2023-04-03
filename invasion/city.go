package invasion

import (
	"strings"
)

// Direction indicates the direction of travel.
type Direction string

var (
	DirectionNorth Direction = "north"
	DirectionSouth Direction = "south"
	DirectionEast  Direction = "east"
	DirectionWest  Direction = "west"
)

// IsValidDirection reports whether v is a valid Direction.
func IsValidDirection(s string) bool {
	switch Direction(s) {
	case DirectionNorth, DirectionSouth, DirectionEast, DirectionWest:
		return true
	default:
		return false
	}
}

// Road represents a road to a city and the direction of travel.
type Road struct {
	Direction Direction
	City      string
}

// City represents a city and roads connecting it to another cities.
type City struct {
	Name  string
	Roads []*Road

	Destroyed bool
	Visitor   *Alien
}

// NewCity construct a new City.
func NewCity(name string, roads []*Road) *City {
	return &City{
		Name:  name,
		Roads: roads,
	}
}

// Visit indicates an alien is visiting the city.
// It sets ok to false and returns the visiting alien encounters another alien.
// Otherwise, sets ok to false and returns the other alien.
func (c *City) Visit(a *Alien) (visitor *Alien, ok bool) {
	if c.Visitor == nil {
		c.Visitor = a
		return nil, true
	}
	// aliens fight
	c.Visitor.Kill()
	a.Kill()
	return c.Visitor, false
}

// RemoveRoad removes roads leading to the given city.
func (c *City) RemoveRoad(toCity string) {
	remainingRoads := make([]*Road, 0)
	for _, r := range c.Roads {
		if r.City == toCity {
			// omit road leading to specified city
			continue
		}
		remainingRoads = append(remainingRoads, r)
	}
	c.Roads = remainingRoads
}

func (c *City) String() string {
	var b strings.Builder
	b.WriteString(c.Name)
	for _, r := range c.Roads {
		b.WriteString(" ")
		b.WriteString(string(r.Direction))
		b.WriteString("=")
		b.WriteString(r.City)
	}
	return b.String()
}
