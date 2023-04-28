package commands

import (
	"cordle/internal/game"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// guess submits a guess to an ongoing game of Cordle
func guess(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, exists := game.FindGame(i.Interaction.ChannelID)
	if exists {
		p := i.Interaction.Member.User
		if g.PlayerInGame(p) {
			// Retrive the guess
			guess := strings.ToLower(i.ApplicationCommandData().Options[0].StringValue())
			_, err := g.SubmitGuess(guess, p)
			if err == nil {
				// Guess was valid, display result
				updateGameBoard(s, i, p, g)
				// Update all game boards
				for id, in := range g.Menus() {
					if id != p.ID {
						// Update the other user's existing game board
						gameBoardEdit(s, in, g.PlayerGameBoard(in.Interaction.Member.User))
					}
				}
				// Check if a player has won the game
				won, id := g.GameWon()
				if won {
					// Notify the players that the game has been won
					s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("<@%s> has won the game! The word was `%s`.", id, g.GoalWord(p)))
					// Close the game
					closeGame(s, i.ChannelID)
				} else if !g.PlayerHasGuesses(p) {
					// Notify the players that one has run out of guesses
					s.ChannelMessageSend(i.ChannelID, p.Mention()+" has run out of guesses!")
					// Check if both players have run out of guesses
					if g.ShouldEndInDraw() {
						// End the game in a draw
						s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("All players are out of guesses! The word was `%s`.", g.GoalWord(p)))
						closeGame(s, i.ChannelID)
					}
				}
			} else {
				// Send the error message back to the user
				respond(s, i, err.Error(), true)
			}
		} else {
			// The player does not belong to this game
			respond(s, i, "You are not part of this game!", true)
		}
	} else {
		// No game in this channel
		respond(s, i, "There are no active games in this channel", true)
	}
}

// updateGameBoard updates the player's game board
func updateGameBoard(s *discordgo.Session, i *discordgo.InteractionCreate, p *discordgo.User, g game.GameManager) {
	// Start by retrieving the correct interaction to respond to
	prev, exists := g.GetPlayerInteractionMenu(p)
	if exists {
		// Delete the existing interaction
		s.InteractionResponseDelete(prev.Interaction)
	}
	// Send the new board and store it
	gameBoardRespond(s, i, g.PlayerGameBoard(p))
	g.SetPlayerInteractionMenu(p, i)
}

// Helper function to send a game board as an embed response to an interaction
func gameBoardRespond(s *discordgo.Session, i *discordgo.InteractionCreate, emb *discordgo.MessageEmbed) {
	// Send the response
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:  discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{emb},
		},
	})
}

// Helper function to edit an existing game board response
func gameBoardEdit(s *discordgo.Session, i *discordgo.InteractionCreate, emb *discordgo.MessageEmbed) {
	// Edit the given game board
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{emb},
	})
}

// Close game closes the current game
func closeGame(s *discordgo.Session, th string) {
	// Remove the game internally
	game.CloseGame(th)
	// Archive and lock the thread from discord
	archived := true
	locked := true
	_, err := s.ChannelEditComplex(th, &discordgo.ChannelEdit{
		Archived: &archived,
		Locked:   &locked,
	})
	if err != nil {
		log.Println(err.Error())
	}
}
