package commands

import (
	"fmt"
	"log"

	"cordle/pkg/util"

	"github.com/bwmarrin/discordgo"
)

// Big list of available commands
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "duel",
		Description: "Send a duel challenge to another player",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "opponent",
				Description: "The player you wish to challenge",
				Required:    true,
			},
		},
	},
	{
		Name:        "accept",
		Description: "Accept a duel challenge against you",
	},
	{
		Name:        "decline",
		Description: "Decline a duel challenge against you",
	},
	{
		Name:        "surrender",
		Description: "Surrender your current game",
	},
	{
		Name:        "guess",
		Description: "Submit a guess to a game of wordle",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "guess",
				Description: "Your five-letter guess",
				Required:    true,
			},
		},
	},
}

var regCommands []*discordgo.ApplicationCommand

// Big map linking command names to their handling functions
// Handler functions are stored in separate go files
var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"duel":      duel,
	"accept":    duelAccept,
	"decline":   duelDecline,
	"surrender": surrender,
	"guess":     guess,
}

// Map linking button CustomIDs to their handler functions
var buttonHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"surrender": surrender,
}

// RegisterCommands registers all command with Discord, this is necessary to allow users to run them
func RegisterCommands(s *discordgo.Session) {
	// Initialise the registered command array
	regCommands = make([]*discordgo.ApplicationCommand, len(commands))
	// Iterate over each command and register it
	for i, cmd := range commands {
		c, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		util.CheckErrMsg(err, "Failed to create command: /"+cmd.Name)
		// Log the command for later deletion
		regCommands[i] = c
		log.Printf("	- /%s", cmd.Name)
	}

	// Create a handler to map commands to their handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			// Handlers for application commands
			h, exists := commandHandlers[i.ApplicationCommandData().Name]
			if exists {
				h(s, i)
			}
		} else if i.Type == discordgo.InteractionMessageComponent {
			// Handlers for buttons
			h, exists := buttonHandlers[i.MessageComponentData().CustomID]
			if exists {
				h(s, i)
			}
		}
	})
}

// ClearCommands deletes all commands created during this session
func ClearCommands(s *discordgo.Session) {
	// Iterate over every registered command and delete it
	for _, cmd := range regCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		util.CheckErrMsg(err, fmt.Sprintf("Failed to delete command /%s", cmd.Name))
		log.Printf("	- /%s", cmd.Name)
	}
}

// respond sends a response to an interaction
func respond(s *discordgo.Session, i *discordgo.InteractionCreate, message string, ephemeral bool) {
	// Configure whether the message is ephemeral or not
	var f discordgo.MessageFlags
	if ephemeral {
		f = discordgo.MessageFlagsEphemeral
	}

	// Send the response
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   f,
		},
	})
}
