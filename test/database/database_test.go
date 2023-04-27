package database

import (
	"cordle/internal/database"
	"cordle/internal/users"
	"errors"
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

}

func TestUpdateUser(t *testing.T) {

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

func TestReadUsers(t *testing.T) {

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

	e := d.CheckUser(u.Id)
	d.DeleteUser(u.Id)
	if e {
		log.Fatalln(errors.New("failed to add user"))
	}
}

func TestDeleteUsers(t *testing.T) {

}
