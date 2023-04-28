package commands

import (
	"fmt"
	"log"

	"cordle/internal/game"

	"github.com/bwmarrin/discordgo"
)

// duelAccept attempts to accept a duel challenge
func duelAccept(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Find the challenge
	c := game.FindChallenge(i.Interaction.Member.User)
	if c != nil {
		// Check that neither player is in a game already
		if game.PlayerFree(c.Source) && game.PlayerFree(c.Target) {
			// Accept the challenge
			game.CloseChallenge(c)
			// Notify the players that the challenge was accepted
			m := fmt.Sprintf(
				"%s, %s has accepted your duel! Please proceed to the breakout thread to play.",
				c.Source.Mention(),
				i.Interaction.Member.Mention(),
			)
			respond(s, i, m, false)
			th, err := startNewDuelThread(s, i, c)
			if err != nil {
				// Report the error
				s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
					Content: "Failed to create a thread for your game",
				})
				log.Printf("Failed to create breakout thread for duel: %s", err)
			} else {
				// Create and store the game
				game.NewDuelGame(th.ID, []*discordgo.User{c.Source, c.Target})
			}
		} else {
			// One or both players is in another game
			respond(s, i, "The duel cannot begin as one or more players are already in a game.", true)
		}
	} else {
		// No challenge was found against this user
		respond(s, i, "You currently have no active challenges against you.", true)
	}
}

// startNewDuelThread separates the functionality of initializing a new thread for a duel game
func startNewDuelThread(s *discordgo.Session, i *discordgo.InteractionCreate, c *game.Challenge) (*discordgo.Channel, error) {
	// Fetch the message to spawn thread from
	msg, _ := s.InteractionResponse(i.Interaction)
	// Create a new thread to duel in
	th, err := s.MessageThreadStartComplex(i.ChannelID, msg.ID, &discordgo.ThreadStart{
		Name:                fmt.Sprintf("%s vs. %s | Duel Game", c.Source.Username, c.Target.Username),
		AutoArchiveDuration: 60,
	})
	if err != nil {
		return nil, err
	}
	// Send a message to begin the game with a reset button for the menu
	s.ChannelMessageSendComplex(th.ID, &discordgo.MessageSend{
		Content: "Use `/guess` to begin guessing. If you wish to leave the game, use `/surrender` or the red button below.",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Surrender",
						Style:    discordgo.DangerButton,
						CustomID: "surrender",
					},
				},
			},
		},
	})
	return th, nil
}
