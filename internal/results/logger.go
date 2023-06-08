package results

import (
	"cordle/internal/database"
	"fmt"
)

func LogWin(d *database.Db, attempts int, id string) {
	user := d.ReadUser(id)
	winCalc(attempts, &user)
	fmt.Println(user.ToStr())

	d.UpdateUser(&user)
}

func LogLoss(d *database.Db, attempts int, id string) {
	user := d.ReadUser(id)
	loseCalc(attempts, &user)
	fmt.Println(user.ToStr())

	d.UpdateUser(&user)
}

func LogDraw(d *database.Db, id string) {
	user := d.ReadUser(id)
	drawCalc(&user)
	fmt.Println(user.ToStr())

	d.UpdateUser(&user)
}
