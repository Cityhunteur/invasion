package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlien_Land(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		world *World
		want  bool
	}{
		{
			name: "Successfully lands in single city",
			world: &World{
				Cities: map[string]*City{
					"OneCity": {
						Name: "OneCity",
					},
				},
				Aliens: map[string]*Alien{},
			},
			want: true,
		},
		{
			name: "No city found",
			world: &World{
				Cities: map[string]*City{},
				Aliens: map[string]*Alien{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			alien := &Alien{}

			got := alien.Land(tt.world)
			assert.Equal(t, tt.want, got)

			if !tt.want {
				return
			}

			city, exists := tt.world.FindCity(alien.VisitingCity)
			assert.True(t, exists)

			assert.Equal(t, alien.VisitingCity, city.Name)
			assert.Equal(t, alien, city.Visitor)
		})
	}
}

func TestAlien_Travel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		world *World
		want  bool
	}{
		{
			name: "Successfully travels to other city",
			world: &World{
				Cities: map[string]*City{
					"CityA": {
						Name: "CityA",
						Roads: []*Road{
							{
								Direction: "south",
								City:      "CityB",
							},
						},
					},
					"CityB": {
						Name: "CityB",
						Roads: []*Road{
							{
								Direction: "north",
								City:      "CityA",
							},
						},
					},
				},
				Aliens: map[string]*Alien{},
			},
			want: true,
		},
		{
			name: "No roads",
			world: &World{
				Cities: map[string]*City{
					"CityA": {
						Name:  "CityA",
						Roads: []*Road{},
					},
				},
				Aliens: map[string]*Alien{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alien := &Alien{}

			got := alien.Land(tt.world)
			require.True(t, got)

			var visiting *City
			for name, city := range tt.world.Cities {
				if name == alien.VisitingCity {
					visiting = city
				}
			}
			require.NotNil(t, visiting)

			got = alien.Travel(tt.world)
			assert.Equal(t, tt.want, got)

			if !tt.want {
				return
			}

			assert.Equal(t, visiting.Roads[0].City, alien.VisitingCity)
		})
	}
}
