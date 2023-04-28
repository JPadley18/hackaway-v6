package commands

import (
	"cordle/internal/game"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// duelDecline declines a challenge against a given user
func duelDecline(s *discordgo.Session, i *discordgo.InteractionCreate) {
	c := game.FindChallenge(i.Interaction.Member.User)
	if c != nil {
		// Decline the challenge
		game.CloseChallenge(c)
		// Notify players that the challenge has been declined
		m := fmt.Sprintf(
			"%s, %s has declined your duel!",
			c.Source.Mention(),
			i.Interaction.Member.Mention(),
		)
		respond(s, i, m, false)
	} else {
		respond(s, i, "You currently have no active challenges against you.", true)
	}
}
