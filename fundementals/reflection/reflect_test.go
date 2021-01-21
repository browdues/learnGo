package reflectionplay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"SirNicholasB"},
			[]string{"SirNicholasB"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"SirNicholasB", "NapityTown"},
			[]string{"SirNicholasB", "NapityTown"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"SirNicholasB", 33},
			[]string{"SirNicholasB"},
		},
		{
			"Nested Fields",
			Person{
				"SirNicholasB",
				Profile{33, "NapityTown"},
			},
			[]string{"SirNicholasB", "NapityTown"},
		},
		{
			"Pointers to things",
			&Person{
				"SirNicholasB",
				Profile{33, "NapityTown"},
			},
			[]string{"SirNicholasB", "NapityTown"}
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			assert.Equal(t, test.ExpectedCalls, got)
		})
	}
}
