package main

import (
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		"struct with one string field",
		struct{
			field1 string
		}{"someword"},
		[]string{"someword"},
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
