package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{
				Name string
			}{"John"},
			[]string{"John"},
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"John", "Syktyvkar"},
			[]string{"John", "Syktyvkar"},
		},
		{
			"struct with not string field",
			struct{
				Name string
				Age int
			}{"John", 44},
			[]string{"John"},
		},
		{
			"nested fields",
			Person{"John",
				Profile{44, "Syktyvkar"}},
			[]string{"John", "Syktyvkar"},
		},
		{
			"pointer to struct",
			&Person{"John",
				Profile{44, "Syktyvar"}},
			[]string{"John", "Syktyvar"},
		},
		{
			"array of structs",
			[]Profile{Profile{31, "Viktoria"}, Profile{44, "John"}},
			[]string{"Viktoria", "John"},
		},

	}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string

				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
