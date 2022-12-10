package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"solo string",
			"Nuriddin",
			[]string{"Nuriddin"},
		},
		{
			"struct with one string field",
			struct {
				Name string
			}{"Nuriddin"},
			[]string{"Nuriddin"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Nuriddin", "Tashkent"},
			[]string{"Nuriddin", "Tashkent"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Nuriddin", 22},
			[]string{"Nuriddin"},
		},
		{
			"struct with nested fields",
			Person{
				"Nuriddin",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Nuriddin", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Dilshod",
				Profile{33, "London"},
			},
			[]string{"Dilshod", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Tashkent"},
			},
			[]string{"London", "Tashkent"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Tashkent"},
			},
			[]string{"London", "Tashkent"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with tests", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "London"}
		}

		var got []string
		want := []string{"Berlin", "London"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	containse := false
	for _, x := range haystack {
		if x == needle {
			containse = true
		}
	}
	if !containse {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
