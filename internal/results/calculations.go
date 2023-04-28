package results

import (
	"cordle/internal/users"
)

func winCalc(attempts int, player *users.User) {
	player.Wins += 1
	player.Elo += 4 * (7 - attempts)
}

func loseCalc(attempts int, player *users.User) {
	player.Losses += 1
	player.Elo -= 4 * attempts

	if player.Elo < 0 {
		player.Elo = 0
	}
}

func drawCalc(player *users.User) {
	player.Draws += 1
	player.Elo += 7
}
