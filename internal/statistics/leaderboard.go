package statistics

import (
	"cordle/internal/database"
	"fmt"
)

func GetLeaderboard(d *database.Db) string {
	topTen := d.ReadTop()
	output := "``\n"
	for index, user := range topTen {
		output += fmt.Sprintf(
			"%d : %s\n",
			index+1,
			user.ToLeaderboard())
	}
	return output + "``"
}
