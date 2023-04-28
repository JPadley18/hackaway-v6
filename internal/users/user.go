package users

import "fmt"

type User struct {
	Id     int
	Wins   int
	Losses int
	Draws  int
	Elo    int
}

func (u User) ToSqlAdd() string {
	return fmt.Sprintf(
		"%d, %d, %d, %d, %d",
		u.Id,
		u.Wins,
		u.Losses,
		u.Draws,
		u.Elo,
	)
}

func (u User) ToSqlUpdate() string {
	return fmt.Sprintf(
		"wins=%d, losses=%d, draws=%d, elo=%d",
		u.Wins,
		u.Losses,
		u.Draws,
		u.Elo,
	)
}
