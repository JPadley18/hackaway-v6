package commands

import (
	"cordle/internal/game"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// surrender allows a player to quit an ongoing game
func surrender(s *discordgo.Session, i *discordgo.InteractionCreate) {
	p := i.Interaction.Member.User
	// Check that there is a game in this channel
	g, exists := game.FindGame(i.ChannelID)
	if exists {
		// Check that the player is part of the current game
		if g.PlayerInGame(p) {
			// Surrender the game
			g.PlayerSurrender(p)
			m := fmt.Sprintf("%s has surrendered!", p.Mention())
			respond(s, i, m, false)
			// Check if the game should now end
			won, id := g.GameWon()
			if won {
				m = fmt.Sprintf("<@%s> wins! The word was `%s`", id, g.GoalWord(p))
				s.ChannelMessageSend(i.ChannelID, m)
				closeGame(s, i.ChannelID)
			}
		} else {
			// The user is not part of this game
			respond(s, i, "You are not part of this game!", true)
		}
	} else {
		// There is no game in this channel
		respond(s, i, "There are no active games in this channel!", true)
	}
}
