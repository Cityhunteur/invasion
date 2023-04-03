package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulation_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		world  *World
		aliens []*Alien
	}{
		{
			name: "Two aliens invade world with two cities",
			world: &World{
				Cities: map[string]*City{
					"Foo": {
						Name: "Foo",
						Roads: []*Road{
							{DirectionNorth, "Bar"},
						},
					},
					"Bar": {
						Name: "Bar",
						Roads: []*Road{
							{DirectionSouth, "Foo"},
							{DirectionWest, "Bee"},
						},
					},
				},
				Aliens: map[string]*Alien{},
			},
			aliens: []*Alien{
				{
					Name: "Alien1",
				},
				{
					Name: "Alien 2",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &Simulation{
				World:  tt.world,
				Aliens: tt.aliens,
			}
			s.Start()

			var destroyedCity string
			// aliens (likely) killed?
			for _, a := range tt.aliens {
				assert.Equal(t, HealthKilled, a.Health)
				destroyedCity = a.VisitingCity
			}

			// was city destroyed?
			_, found := tt.world.FindCity(destroyedCity)
			assert.False(t, found)

			// other cities intact?
			for _, c := range tt.world.Cities {
				if c.Name != destroyedCity {
					_, found := tt.world.FindCity(c.Name)
					assert.True(t, found)
				}
			}
		})
	}
}
