package database

import (
	"fmt"
)

type Stats struct {
	Wins   int
	Losses int
	Draws  int
	Games  int
	Elo    int
	Level  int
}

func (s Stats) ToString() string {
	return fmt.Sprintf("Wins: %d\nLosses: %d\nDraws: %d\nGames Played: %d\nElo: %d\nLevel: %d",
		s.Wins,
		s.Losses,
		s.Draws,
		s.Games,
		s.Elo,
		s.Level)
}

func (s Stats) WinPercentage() float64 {
	return float64(s.Wins) / float64(s.Games)
}
