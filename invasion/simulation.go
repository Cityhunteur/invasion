// Package invasion implements a simulator for an alien invasion.
package invasion

import "fmt"

type Simulation struct {
	World  *World
	Aliens []*Alien
}

// NewSimulation creates a new Simulation to simulate the invasion of the given World
// by the specified number of Alien(s).
func NewSimulation(w *World, nAliens int) *Simulation {
	var aliens []*Alien
	for i := 1; i <= nAliens; i++ {
		name := fmt.Sprintf("Alien %d", i)
		a := NewAlien(name)
		aliens = append(aliens, a)
	}
	return &Simulation{
		World:  w,
		Aliens: aliens,
	}
}

func (s *Simulation) Start() {
	// unleash aliens in the world
	for _, a := range s.Aliens {
		if !a.Land(s.World) {
			// all cities were destroyed prior to all aliens landing
			break
		}
	}

	visits := 0
	for {
		for _, a := range s.World.Aliens {
			if a.Travel(s.World) {
				visits++
				continue
			}
		}

		// are aliens still visiting?
		if visits < 1 {
			break
		}

		visits = 0
	}

	if s.World.IsDestroyed() {
		fmt.Println("Entire world has been destroyed 👻")
		return
	}

	fmt.Println(s.World.String())
}
