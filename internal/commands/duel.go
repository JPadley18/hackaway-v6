package commands

import (
	"fmt"
	"time"

	"cordle/internal/config"
	"cordle/internal/game"

	"github.com/bwmarrin/discordgo"
)

func duel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check whether the target is valid (not a role and not a bot)
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if !user.Bot && user.ID != i.Interaction.Member.User.ID {
		// Check that the user invoking the challenge is allowed to invoke a challenge
		if game.UserCanChallenge(i.Member.User, i.Interaction.ChannelID) {
			// Check that the target does not already have a challenge against them
			if game.FindChallenge(user, i.Interaction.ChannelID) == nil {
				// Create a new challenge
				c := game.NewChallenge(i.Interaction.Member.User, user, i.Interaction)
				// Respond to the interaction, notifying the other player of the duel
				m := fmt.Sprintf(
					"%s, %s has challenged you to a duel! You have %d seconds to either `/accept` or `/decline` this duel.",
					user.Mention(),
					i.Interaction.Member.Mention(),
					config.Config.Game.ChallengeDuration,
				)
				respond(s, i, m, false)
				// After the delay, delete the challenge notification and expire the challenge if it hasn't been declined
				time.AfterFunc(time.Duration(config.Config.Game.ChallengeDuration)*time.Second, func() {
					s.InteractionResponseDelete(i.Interaction)
					game.CloseChallenge(c)
				})
			} else {
				// Challenge already exists
				respond(s, i, "That player already has an active challenge against them! Please wait for it to expire and try again", true)
			}
		} else {
			// User cannot start a challenge
			respond(s, i, "You already have an active challenge against another player! Please wait for that challenge to expire before starting another.", true)
		}
	} else {
		// Challenge target is invalid
		respond(s, i, "That is not a valid duel target!", true)
	}
}
