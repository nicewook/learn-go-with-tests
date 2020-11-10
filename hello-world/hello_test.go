package main

import "testing"

func TestHello(t *testing.T) {
	testCases := []struct {
		desc string
		lang string
		want string
		name string
	}{
		{
			desc: "Hello world testing",
			lang: "English",
			want: "Hello, world",
			name: "",
		},
		{
			desc: "Hello Chris testing",
			lang: "English",
			want: "Hello, Chris",
			name: "Chris",
		},
		{
			desc: "Hello world testing - in Spanish",
			lang: "Spanish",
			want: "Hola, world",
			name: "",
		},
		{
			desc: "Hello Chris testing - in Spanish",
			lang: "Spanish",
			want: "Hola, Elodie",
			name: "Elodie",
		},
	}

	assertCorrectMessage := func(got, want string) {
		// https://github.com/godoctor/godoctor/issues/50
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Hello(tC.lang, tC.name)
			want := tC.want
			assertCorrectMessage(got, want)
		})
	}
}
