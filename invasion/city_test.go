package invasion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		CityName string
		Roads    []*Road
		want     *City
		wantStr  string
	}{
		{
			name:     "City without roads",
			CityName: "Giethoorn",
			want: &City{
				Name: "Giethoorn",
			},
			wantStr: "Giethoorn",
		},
		{
			name:     "Foo",
			CityName: "Foo",
			Roads: []*Road{
				{
					Direction: DirectionNorth,
					City:      "Bar",
				},
				{
					Direction: DirectionWest,
					City:      "Baz",
				},
				{
					Direction: DirectionSouth,
					City:      "Qu-ux",
				},
			},
			want: &City{
				Name: "Foo",
				Roads: []*Road{
					{
						Direction: DirectionNorth,
						City:      "Bar",
					},
					{
						Direction: DirectionWest,
						City:      "Baz",
					},
					{
						Direction: DirectionSouth,
						City:      "Qu-ux",
					},
				},
			},
			wantStr: "Foo north=Bar west=Baz south=Qu-ux",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewCity(tt.CityName, tt.Roads)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantStr, got.String())
		})
	}
}

func TestCity_Visit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		city *City
		want bool
	}{
		{
			name: "Successfully visit",
			city: &City{
				Name: "City A",
			},
			want: true,
		},
		{
			name: "Encounters another alien",
			city: &City{
				Name: "City A",
				Visitor: &Alien{
					Name: "Alien1",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			alien := NewAlien("👽")

			gotVisitor, got := tt.city.Visit(alien)

			if !tt.want {
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.city.Visitor, gotVisitor)

				assert.Equal(t, HealthKilled, alien.Health)
				assert.Equal(t, HealthKilled, tt.city.Visitor.Health)
				return
			}

			assert.Equal(t, tt.want, got)
			assert.Nil(t, gotVisitor)
		})
	}
}
