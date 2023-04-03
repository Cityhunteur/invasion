package invasion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		in      string
		want    *World
		wantErr error
	}{
		{
			name: "Valid example map",
			in: `Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee`,
			want: &World{
				Cities: map[string]*City{
					"Foo": {
						Name: "Foo",
						Roads: []*Road{
							{DirectionNorth, "Bar"},
							{DirectionWest, "Baz"},
							{DirectionSouth, "Qu-ux"},
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
			wantErr: nil,
		},
		{
			name: "Invalid line",
			in: `Foo north=Bar

`,
			wantErr: ErrInvalidMap,
		},
		{
			name:    "Invalid road",
			in:      `Foo north`,
			wantErr: ErrInvalidMap,
		},
		{
			name:    "Invalid direction",
			in:      `Foo invalid=Bar`,
			wantErr: ErrInvalidMap,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewWorld(strings.NewReader(tt.in))

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.ErrorAs(t, err, &tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWorld_String(t *testing.T) {
	t.Parallel()
	// we use a slice here to ignore the order of cities in map for assertion
	want := []string{"A north=B", "B south=A"}
	w, err := NewWorld(strings.NewReader(strings.Join(want, "\n")))
	assert.NoError(t, err)

	got := w.String()

	assert.ElementsMatch(t, want, strings.Split(got, "\n"))
}

func TestWorld_DestroyCity(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		world *World
		city  *City
	}{
		{
			name: "Existing city",
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
			city: &City{
				Name: "Foo",
				Roads: []*Road{
					{DirectionNorth, "Bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.world.DestroyCity(tt.city)

			assert.NotContains(t, tt.world.Cities, tt.city)

			// assert no roads lead to destroyed city
			for _, c := range tt.world.Cities {
				for _, r := range c.Roads {
					assert.NotEqual(t, tt.city.Name, r.City)
				}
			}
		})
	}
}
