package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name	string
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
			}{"Batman"},
			[]string{"Batman"},
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"Batman", "Gotham"},
			[]string{"Batman", "Gotham"},
		},
		{
			"struct with non string field",
			struct{
				Name string
				Age int
			}{"Batman", 35},
			[]string{"Batman"},
		},
		{
			"struct with nested fields",
			Person{
				"Batman",
				Profile{35, "Gotham"},
			},
			[]string{"Batman", "Gotham"},
		},
		{
			"pointer to struct",
			&Person{
				"Batman",
				Profile{35, "Gotham"},
			},
			[]string{"Batman", "Gotham"},
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
