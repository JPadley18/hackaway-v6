package database

import (
	"cordle/internal/database"
	"cordle/internal/users"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const conf = "../conf/db-key.json"

var d *database.Db

func TestMain(m *testing.M) {

}

func TestDb(t *testing.T) {
	d = database.NewDb(conf)
	defer d.Close()

	assert.NotNil(t, d)
}

func TestAddUser(t *testing.T) {
	u := users.User{
		Id:     123,
		Wins:   10,
		Losses: 3,
		Draws:  5,
		Elo:    567,
	}

	d = database.NewDb(conf)
	defer d.Close()

	d.AddUser(&u)
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

}

func TestDeleteUsers(t *testing.T) {

}
