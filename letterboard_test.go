package main

import (
	"reflect"
	"testing"
)

func assertMoves(t *testing.T, expected Moves, actual Moves) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected\n  %v,\nbut received\n  %v", expected, actual)
	}
}

func TestSolveLetterboard_With_Cat(t *testing.T) {
	actual := SolveLetterboard([]rune{'a', 'z', 'c', 't', 'v', 'a'}, "cat")
	expected := Moves{
		/* c */ Move{Left, -1}, Move{Left, -1}, Move{Left, 'c'},
		/* a */ Move{Right, -1}, Move{Right, 'a'},
		/* t */ Move{Left, -1}, Move{Left, 't'},
	}

	assertMoves(t, expected, actual)
}

func TestSolveLetterboard_With_TV(t *testing.T) {
	actual := SolveLetterboard([]rune{'a', 'z', 'c', 't', 'v', 'a'}, "tv")
	expected := Moves{
		/* t */ Move{Right, -1}, Move{Right, -1}, Move{Right, 't'},
		/* v */ Move{Left, 'v'},
	}

	assertMoves(t, expected, actual)
}
