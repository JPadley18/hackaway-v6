package database

import (
	"fmt"
)

type User struct {
	Id     int
	Name   string
	Wins   int
	Losses int
	Draws  int
	Games  int
	Elo    int
	Level  int
}

func (u User) ToString() string {
	return fmt.Sprintf(
		"ID: %d\nName: %s\nWins: %d\nLosses: %d\nDraws: %d\nGames Played: %d\nElo: %d\nLevel: %d",
		u.Id,
		u.Name,
		u.Wins,
		u.Losses,
		u.Draws,
		u.Games,
		u.Elo,
		u.Level)
}

func (u User) ToSqlAdd() string {
	return fmt.Sprintf(
		"%d, '%s', %d, %d, %d, %d, %d, %d",
		u.Id,
		u.Name,
		u.Wins,
		u.Losses,
		u.Draws,
		u.Games,
		u.Elo,
		u.Level)
}

func (u User) ToSqlUpdate() string {
	return fmt.Sprintf(
		"wins=%d, losses=%d, draws=%d, games=%d, elo=%d, level=%d",
		u.Wins,
		u.Losses,
		u.Draws,
		u.Games,
		u.Elo,
		u.Level)
}

func (u User) ToStat() string {
	return fmt.Sprintf(
		"%s's stats:\nWins: %d\nLosses: %d\nDraws: %d\nGames: %d\nElo: %d\nLevel: %d",
		u.Name,
		u.Wins,
		u.Losses,
		u.Draws,
		u.Games,
		u.Elo,
		u.Level)
}

func (u User) ToLeaderboard() string {
	return fmt.Sprintf(
		"%s : %d",
		u.Name,
		u.Elo)
}

func (u User) WinPercentage() float64 {
	return float64(u.Wins) / float64(u.Games)
}
