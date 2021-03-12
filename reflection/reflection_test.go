package reflection

import (
	"reflect"
	"testing"
)

func TestWakl(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name           string
		Input          interface{}
		ExepectedCalls []string
	}{
		{
			"struct with non string",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},

			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Floripa"},
			},
			[]string{"London", "Floripa"},
		},
		{
			"Slices",
			[2]Profile{
				{33, "London"},
				{34, "Floripa"},
			},
			[]string{"London", "Floripa"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExepectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExepectedCalls)
			}

		})
	}

}
