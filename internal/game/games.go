package game

import (
	"cordle/internal/wordle"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// GameInterface is implemented by all GameModes, allows games to be interacted with
type GameManager interface {
	PlayerInGame(p *discordgo.User) bool
	Menus() map[string]*discordgo.InteractionCreate
	SubmitGuess(guess string, p *discordgo.User) (*wordle.Guess, error)
	PlayerGameBoard(p *discordgo.User) *discordgo.MessageEmbed
	PlayerHasGuesses(p *discordgo.User) bool
	GetPlayerInteractionMenu(p *discordgo.User) (*discordgo.InteractionCreate, bool)
	SetPlayerInteractionMenu(p *discordgo.User, m *discordgo.InteractionCreate)
	GoalWord(p *discordgo.User) string
	PlayerSurrender(p *discordgo.User)
	GameWon() (bool, string)
	ShouldEndInDraw() bool
}

// Thread safe map of channel IDs to the games in them
var games struct {
	g  map[string]GameManager
	mu sync.Mutex
}

// Initialise the games map
func init() {
	games.g = make(map[string]GameManager)
}

// PlayerFree checks whether a player is already in a game, returns true if not
func PlayerFree(p *discordgo.User) bool {
	games.mu.Lock()
	defer games.mu.Unlock()
	for _, g := range games.g {
		if g.PlayerInGame(p) {
			return false
		}
	}
	return true
}

// FindGame returns a game given a channel ID that the game is taking place in
// Returns a reference to the game struct and a boolean confirming if the game exists or not
func FindGame(channelID string) (GameManager, bool) {
	games.mu.Lock()
	defer games.mu.Unlock()
	// For some reason this has to be two lines
	ret, exists := games.g[channelID]
	return ret, exists
}

// CloseGame removes a completed game from the list
func CloseGame(channelID string) {
	games.mu.Lock()
	// Delete the game
	delete(games.g, channelID)
	games.mu.Unlock()
}
