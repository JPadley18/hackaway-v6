package database

import (
	"cordle/pkg/sql"
	"cordle/pkg/util"
	"sync"
)

type Adder interface {
	AddUser(user User)
	AddUsers(users []User)
}

type Updater interface {
	UpdateUser(user *User)
	UpdateUsers(users *[]User)
}

type Reader interface {
	ReadUser(id int) User
	ReadUsers() []User
	ReadTop() []User
	ReadStats(id int) Stats
}

type Checker interface {
	CheckUser(id int) bool
}

type Deleter interface {
	DeleteUser(id int)
	DeleteUsers(ids []int)
}

type Manager interface {
	Adder
	Updater
	Reader
	Checker
	Deleter
}

type Db struct {
	ClientMu sync.Mutex
	Client   sql.Client
}

func NewDb(path string) *Db {
	return &Db{
		Client: *sql.NewClient(path),
	}
}

func (d *Db) Close() {
	err := d.Client.Close()
	util.CheckErr(err)
}
