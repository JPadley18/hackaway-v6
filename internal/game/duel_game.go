package game

import (
	"cordle/internal/wordle"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DuelGame holds the information about a DuelGame
type DuelGame struct {
	// games stores a map of user IDs to their game
	games map[string]*wordle.WordleGame
	// menus stores the interaction to edit to display games to each user
	menus map[string]*discordgo.InteractionCreate
}

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func NewDuelGame(th string, p []*discordgo.User) {
	// Create the shared game
	g0 := wordle.NewRandomGame()
	// Manually create a second game with the same goal word
	// This is more efficient than doing a deep copy
	g1 := &wordle.WordleGame{
		Guesses:  []*wordle.Guess{},
		GoalWord: g0.GoalWord,
	}

	// Create the game struct and store it
	games.mu.Lock()
	games.g[th] = &DuelGame{
		games: map[string]*wordle.WordleGame{
			p[0].ID: g0,
			p[1].ID: g1,
		},
		menus: make(map[string]*discordgo.InteractionCreate),
	}
	games.mu.Unlock()
}

// PlayerInGame returns true if the given player is a part of the game
func (g *DuelGame) PlayerInGame(p *discordgo.User) bool {
	_, exists := g.games[p.ID]
	return exists
}

// Menus returns the map of menus for this game
func (g *DuelGame) Menus() map[string]*discordgo.InteractionCreate {
	return g.menus
}

// PlayerHasGuesses returns true if the player has guesses remaining in the game
func (g *DuelGame) PlayerHasGuesses(p *discordgo.User) bool {
	return g.games[p.ID].GuessesRemaining() > 0
}

// GetPlayerInteractionMenu searches for and returns the interaction menu for the given player
// Returns a boolean to indicate whether or not the menu was found
func (g *DuelGame) GetPlayerInteractionMenu(p *discordgo.User) (*discordgo.InteractionCreate, bool) {
	r, exists := g.menus[p.ID]
	return r, exists
}

// SetPlayerInteractionMenu stores an interaction to be used as the user's menu
func (g *DuelGame) SetPlayerInteractionMenu(p *discordgo.User, m *discordgo.InteractionCreate) {
	g.menus[p.ID] = m
}

// SubmitGuess allows a guess to be submitted to the game of a given player
// Returns the result as an array of wordle.GuessState
func (g *DuelGame) SubmitGuess(guess string, p *discordgo.User) (*wordle.Guess, error) {
	pg := g.games[p.ID]
	return pg.Guess(guess)
}

// PlayerGuessHistory returns the formatted game history of the player and their opponent
func (g *DuelGame) PlayerGameBoard(p *discordgo.User) *discordgo.MessageEmbed {
	var ghp, gho []*wordle.Guess
	for id, game := range g.games {
		// If it is the local player, store in ghp
		if id == p.ID {
			ghp = game.Guesses
		} else {
			gho = game.Guesses
		}
	}
	gbp := displayGame(ghp, false)
	gbo := displayGame(gho, true)
	return renderGameBoard(gbp, gbo, p)
}

// GoalWord returns the goal word for this game
func (g *DuelGame) GoalWord(p *discordgo.User) string {
	return g.games[p.ID].GoalWord
}

// PlayerSurrender allows a player to quit an ongoing game
// In a duel game, this should immediately end the game
func (g *DuelGame) PlayerSurrender(p *discordgo.User) {
	for id, g := range g.games {
		if p.ID != id {
			g.Won = true
			return
		}
	}
}

// GameWon returns true if the game has been won, as well as the ID of the winner
func (g *DuelGame) GameWon() (bool, string) {
	for id, g := range g.games {
		if g.Won {
			return true, id
		}
	}
	return false, ""
}

// ShouldEndInDraw returns true if the current game has reached a stalemate and should end in a draw
func (g *DuelGame) ShouldEndInDraw() bool {
	for _, g := range g.games {
		if g.GuessesRemaining() > 0 {
			return false
		}
	}
	return true
}

// displayGame returns a string displaying the given guess history.
// If hide is true, returns the game board without the letters
func displayGame(gh []*wordle.Guess, hide bool) string {
	// Iterate over the slice to build the guess board
	var gb strings.Builder
	for i := 0; i < wordle.MaxGuesses; i++ {
		// Write the player's board
		if i < len(gh) {
			gb.WriteString(displayGuess(gh[i], hide))
		} else {
			// If not all guesses have been filled, add a blank line
			gb.WriteString(blankLine())
		}
		gb.WriteRune('\n')
	}
	// Return the board as a string
	return gb.String()
}

// Given 2 game boards as a string, creates a MessageEmbed containing the board
func renderGameBoard(gbp string, gbo string, p *discordgo.User) *discordgo.MessageEmbed {
	// Create a message embed with the game board inside
	return &discordgo.MessageEmbed{
		Type:  discordgo.EmbedTypeRich,
		Title: "Cordle Game | Duel",
		Color: 0x00b503,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    p.Username,
			IconURL: p.AvatarURL("64"),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "You",
				Value:  gbp,
				Inline: true,
			},
			{
				Name:   "Your opponent",
				Value:  gbo,
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Your opponent is guessing at the same time as you, try to solve the puzzle before they do! Use /guess to guess again.",
		},
	}
}

// displayGuess returns a nicely formatted response from a guess result to send back to the user.
// If hide is true, returns blank boxes instead of boxes with letters
func displayGuess(r *wordle.Guess, hide bool) string {
	var s strings.Builder
	runes := []rune(r.GuessWord)
	for i, gs := range r.GuessResult {
		prefix := ""
		if gs == wordle.CorrectCharacter {
			prefix = "green"
		} else if gs == wordle.IncorrectPosition {
			prefix = "yellow"
		} else {
			prefix = "grey"
		}
		// Calculate the name of the required emoji and write it
		var e string
		if hide {
			e = fmt.Sprintf("%s_blank", prefix)
		} else {
			e = fmt.Sprintf("%s_%c", prefix, runes[i])
		}
		s.WriteString(Emojis[e])
	}
	return s.String()
}

// blankLine generates a line of five blank emojis
func blankLine() string {
	var s strings.Builder
	for i := 0; i < 5; i++ {
		s.WriteString(Emojis["blank"])
	}
	return s.String()
}
