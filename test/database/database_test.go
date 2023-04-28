package database

import (
	"cordle/internal/database"
	"cordle/internal/users"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const conf = "../conf/db-key.json"

var d *database.Db

func TestConfigExists(t *testing.T) {
	_, err := os.Stat(conf)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestDb(t *testing.T) {
	d = database.NewDb(conf)
	defer d.Close()

	assert.NotNil(t, d)
}

func TestAddUser(t *testing.T) {
	u := users.User{
		Id:     7567,
		Wins:   20,
		Losses: 53,
		Draws:  151,
		Elo:    341,
	}

	d = database.NewDb(conf)
	defer d.Close()

	d.AddUser(&u)

	e := d.CheckUser(u.Id)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u.Id)
}

func TestAddUsers(t *testing.T) {
	u1 := users.User{
		Id:     7567,
		Wins:   20,
		Losses: 53,
		Draws:  151,
		Elo:    341,
	}
	u2 := users.User{
		Id:     1577,
		Wins:   20,
		Losses: 13,
		Draws:  51,
		Elo:    541,
	}

	d = database.NewDb(conf)
	defer d.Close()

	u := make([]users.User, 2)
	u[0] = u1
	u[1] = u2
	d.AddUsers(&u)

	e := d.CheckUser(u1.Id)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u1.Id)

	e = d.CheckUser(u2.Id)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u2.Id)
}

func TestUpdateUser(t *testing.T) {
	d = database.NewDb(conf)
	defer d.Close()

	u := d.ReadUser(7123)
	draws := u.Draws
	u.Draws += 1
	d.UpdateUser(&u)
	u = d.ReadUser(7123)
	if u.Draws != draws+1 {
		log.Fatalln(errors.New("error updating draw count"))
	}
}

func TestUpdateUsers(t *testing.T) {

}

func TestReadUser(t *testing.T) {
	d = database.NewDb(conf)
	defer d.Close()

	u := d.ReadUser(7123)
	if u.Id != 7123 {
		log.Fatalln(errors.New("read nil user error"))
	}
	assert.NotNil(t, u)
}

func TestReadAllUsers(t *testing.T) {
	d = database.NewDb(conf)
	defer d.Close()

	u := d.ReadAllUsers()
	if len(u) != 3 {
		log.Fatalln(errors.New(fmt.Sprintf(
			"incorrect array length for all users %d", len(u)),
		))
	}
}

func TestReadTop(t *testing.T) {
	d = database.NewDb(conf)
	defer d.Close()

	tt := d.ReadTop()
	for i := 0; i < len(tt) - 2; i++ {
		if tt[i].Elo > tt[i+1].Elo {
			log.Fatalln(errors.New(fmt.Sprintf(
				"top ten wrong order %d", len(tt)),
			))
		} 
	} 
}

func TestCheckUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {
	u := users.User{
		Id:     61567,
		Wins:   22,
		Losses: 51,
		Draws:  101,
		Elo:    371,
	}

	d = database.NewDb(conf)
	defer d.Close()

	d.AddUser(&u)

	d.DeleteUser(u.Id)
	e := d.CheckUser(u.Id)
	if e {
		log.Fatalln(errors.New("failed to delete user"))
	}
}

func TestDeleteUsers(t *testing.T) {

}
