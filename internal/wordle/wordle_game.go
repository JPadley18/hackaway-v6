package wordle

import (
	"errors"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

// The maximum number of guesses allowed in a game of Wordle
const MaxGuesses int = 6

// The length of a guess
const GuessLength int = 5

// Possible states a character in a guess could be
type GuessState int

const (
	IncorrectCharacter GuessState = iota
	IncorrectPosition  GuessState = iota
	CorrectCharacter   GuessState = iota
)

// Possible errors that could be encountered when guessing
var (
	ErrOutOfGuesses  = errors.New("You have run out of guesses!")
	ErrInvalidLength = errors.New("Invalid guess length!")
	ErrInvalidFormat = errors.New("Invalid characters in guess!")
	ErrInvalidWord   = errors.New("That is not a valid word!")
)

// Variables to store potential answers and guesses
var (
	answers []string
	guesses []string
)

// Stores information about a game of Wordle (this has no information linking to Discord)
type WordleGame struct {
	Won      bool
	Guesses  []*Guess
	GoalWord string
}

// Guess stores the information about a guess of wordle
type Guess struct {
	GuessWord   string
	GuessResult [5]GuessState
}

// Runs when the wordle module is imported
func init() {
	// Load the possible guesses and answers
	answers, guesses = LoadWords()
}

// NewRandomGame creates a new wordle game with a random solution and returns it
func NewRandomGame() *WordleGame {
	return &WordleGame{
		Won:      false,
		Guesses:  []*Guess{},
		GoalWord: answers[rand.Intn(len(answers))],
	}
}

// Guess allows the submission of a guess to a wordle game, this requires a lowercase guess
// An error is returned if the guess string is invalid or the game has no guesses remaining
func (g *WordleGame) Guess(guess string) (*Guess, error) {
	// Validate the guess
	err := validateGuess(guess)
	if err != nil {
		return nil, err
	}
	// Check that the game has remaining guesses
	if g.GuessesRemaining() == 0 {
		return nil, ErrOutOfGuesses
	}

	// Evaluate the guess
	result, correct := evaluateGuess(guess, g.GoalWord)
	gs := &Guess{
		GuessWord:   guess,
		GuessResult: result,
	}

	// Add it to the guess history
	g.Guesses = append(g.Guesses, gs)

	g.Won = correct

	return gs, nil
}

// GuessesRemaining checks how many guesses remain in a given game
func (g *WordleGame) GuessesRemaining() int {
	return MaxGuesses - len(g.Guesses)
}

// validateGuess returns an appropriate error if the given guess is invalid
func validateGuess(guess string) error {
	// Check the length of the guess
	if len(guess) != GuessLength {
		return ErrInvalidLength
	}
	// Check that the guess contains only lowercase alphabetical letters
	if !regexp.MustCompile(`^[a-z]+$`).MatchString(guess) {
		return ErrInvalidFormat
	}
	// Check that the guess is one of the valid words
	if !isValidWord(guess) {
		return ErrInvalidWord
	}

	return nil
}

// isValidWord determines if the given word is a valid guess
func isValidWord(w string) bool {
	// Use a binary search to check if this word is present
	_, found := sort.Find(len(guesses), func(i int) int {
		return strings.Compare(w, guesses[i])
	})
	if !found {
		// Try the answers list instead
		_, found = sort.Find(len(answers), func(i int) int {
			return strings.Compare(w, answers[i])
		})
	}
	return found
}

// evaluateGuess analyses a wordle guess against a target and returns an array of character statuses
func evaluateGuess(guess string, goal string) ([5]GuessState, bool) {
	// Retrieve the rune counts
	counts := countRunes(goal)

	// Default value = 0 (IncorrectCharacter)
	var result = [5]GuessState{}
	correct := true

	// First pass: check all correct runes
	for i, v := range guess {
		r, _ := utf8.DecodeRuneInString(goal[i:])
		if v == r {
			counts[v]--
			result[i] = CorrectCharacter
		}
	}

	// Second pass: mark incorrect positions
	for i, v := range guess {
		// Only check characters that are not already marked as correct
		if result[i] != CorrectCharacter {
			correct = false
			c, _ := counts[v]
			if c > 0 {
				counts[v]--
				result[i] = IncorrectPosition
			}
		}
	}

	return result, correct
}

// countRunes returns a map with every rune present in the string along with its number of occurrences
func countRunes(s string) map[rune]int {
	var counts = make(map[rune]int)
	// Count the occurrences of each character in the goal string
	for _, c := range s {
		_, exists := counts[c]
		if exists {
			counts[c]++
		} else {
			counts[c] = 1
		}
	}
	return counts
}
