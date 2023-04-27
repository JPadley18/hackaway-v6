package results

import (
	"cordle/internal/database"
)

func winCalc(attempts int, player *database.User) {
	player.Wins += 1
	player.Games += 1
	player.Elo += 4 * (7 - attempts)
	player.Level += 5 * (7 - attempts)
}

func loseCalc(attempts int, player *database.User) {
	player.Losses += 1
	player.Games += 1
	player.Elo -= 4 * attempts
	player.Level += 3

	if player.Elo < 0 {
		player.Elo = 0
	}
}

func drawCalc(player *database.User) {
	player.Draws += 1
	player.Games += 1
	player.Elo += 7
	player.Level += 10
}
