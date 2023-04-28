package game

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Challenge stores information about an active challenge between two players
type Challenge struct {
	Source *discordgo.User
	Target *discordgo.User
}

// Stores all currently active challenges
// Uses a mutex to stop expireChallenge while FindChallenge is in operation
var challenges struct {
	c  []*Challenge
	mu sync.Mutex
}

// NewChallenge creates a new challenge between two players
func NewChallenge(s *discordgo.User, t *discordgo.User) *Challenge {
	c := &Challenge{
		Source: s,
		Target: t,
	}
	challenges.mu.Lock()
	challenges.c = append(challenges.c, c)
	challenges.mu.Unlock()

	return c
}

// Locates and returns a challenge between two players given the target user object
// Returns nil if one is not found
func FindChallenge(t *discordgo.User) *Challenge {
	challenges.mu.Lock()
	defer challenges.mu.Unlock()

	for _, c := range challenges.c {
		if c.Target.ID == t.ID {
			return c
		}
	}

	return nil
}

// CloseChallenge removes a challenge from the active list
func CloseChallenge(c *Challenge) {
	challenges.mu.Lock()
	defer challenges.mu.Unlock()

	// Find the index of the challenge
	idx := -1
	for i, ch := range challenges.c {
		if ch == c {
			idx = i
		}
	}

	// If idx == -1, the challenge has already been removed
	if idx != -1 {
		challenges.c[idx] = challenges.c[len(challenges.c)-1]
		challenges.c = challenges.c[:len(challenges.c)-1]
	}
}
