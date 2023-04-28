package wordle

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateGuess(t *testing.T) {
	var tests = []struct {
		guess   string
		outcome error
	}{
		{"abcde", nil},
		{"abCde", ErrInvalidFormat},
		{"ab.de", ErrInvalidFormat},
		{"abcd", ErrInvalidLength},
		{"abcdef", ErrInvalidLength},
	}

	for _, test := range tests {
		t.Run(test.guess, func(t *testing.T) {
			outcome := validateGuess(test.guess)
			if outcome != test.outcome {
				t.Errorf("got '%v', want '%v'", outcome, test.outcome)
			}
		})
	}
}

func TestCountRunes(t *testing.T) {
	var tests = []struct {
		input   string
		outcome map[rune]int
	}{
		{"aaaaa", map[rune]int{'a': 5}},
		{"abcde", map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1}},
		{"abbbb", map[rune]int{'a': 1, 'b': 4}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			outcome := countRunes(test.input)
			if !reflect.DeepEqual(outcome, test.outcome) {
				t.Errorf("got '%v', want '%v'", outcome, test.outcome)
			}
		})
	}
}

func TestEvaluateGuess(t *testing.T) {
	var tests = []struct {
		input   string
		goal    string
		outcome [5]GuessState
	}{
		{"aaaaa", "aaaaa", [5]GuessState{CorrectCharacter, CorrectCharacter, CorrectCharacter, CorrectCharacter, CorrectCharacter}},
		{"bbbbb", "aaaaa", [5]GuessState{IncorrectCharacter, IncorrectCharacter, IncorrectCharacter, IncorrectCharacter, IncorrectCharacter}},
		{"ecdba", "abcde", [5]GuessState{IncorrectPosition, IncorrectPosition, IncorrectPosition, IncorrectPosition, IncorrectPosition}},
		{"aanon", "sunny", [5]GuessState{IncorrectCharacter, IncorrectCharacter, CorrectCharacter, IncorrectCharacter, IncorrectPosition}},
		{"nnoon", "sunny", [5]GuessState{IncorrectPosition, IncorrectPosition, IncorrectCharacter, IncorrectCharacter, IncorrectCharacter}},
		{"nonny", "sunny", [5]GuessState{IncorrectCharacter, IncorrectCharacter, CorrectCharacter, CorrectCharacter, CorrectCharacter}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%s/%s", test.input, test.goal)
		t.Run(name, func(t *testing.T) {
			outcome, _ := evaluateGuess(test.input, test.goal)
			if outcome != test.outcome {
				t.Errorf("got %v, want %v", outcome, test.outcome)
			}
		})
	}
}
