package database

import (
	"cordle/internal/users"
	"cordle/pkg/util"
	"fmt"
)

func (d *Db) AddUser(user *users.User) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	insert, err := db.Query(fmt.Sprintf(
		`insert into users(id, wins, losses, draws, elo) 
		values(%s);`,
		user.ToSqlAdd(),
	))

	util.CheckErr(err)
	defer insert.Close()
}

func (d *Db) AddUsers(users *[]users.User) {
	for _, user := range *users {
		d.AddUser(&user)
	}
}

func (d *Db) UpdateUser(user *users.User) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	updates := user.ToSqlUpdate()
	query := fmt.Sprintf(
		"id=%d",
		user.Id,
	)

	update, err := db.Query(fmt.Sprintf(
		`update users
		set %s
		where %s;`,
		updates,
		query,
	))

	util.CheckErr(err)
	defer update.Close()
}

func (d *Db) UpdateUsers(users *[]users.User) {
	for _, user := range *users {
		d.UpdateUser(&user)
	}
}

func (d *Db) ReadUser(id int) users.User {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	result, err := db.Queryx(fmt.Sprintf(
		"select * from users where id=%d;",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var user users.User
	result.Next()
	err = result.StructScan(&user)

	util.CheckErr(err)
	return user
}

func (d *Db) ReadAllUsers() []users.User {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	result, err := db.Queryx("select * from users;")
	util.CheckErr(err)
	defer result.Close()

	us := make([]users.User, 0)
	var u users.User
	for i := 0; result.Next(); i++ {
		err := result.StructScan(&u)
		if err != nil {
			panic(err.Error())
		}
		us = append(us, u)
	}

	return us
}

func (d *Db) ReadTop() []users.User {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	results, err := db.Queryx("select * from users order by elo, id asc limit 0,10;")
	util.CheckErr(err)
	defer results.Close()

	tt := make([]users.User, 0)
	var u users.User
	for i := 0; results.Next(); i++ {
		err := results.StructScan(&u)
		util.CheckErr(err)
		tt = append(tt, u)
	}

	return tt
}

func (d *Db) ReadStats(id int) Stats {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	result, err := db.Queryx(fmt.Sprintf(
		"select wins, losses, draws, games, elo, level from users where id=%d;",
		id))

	util.CheckErr(err)
	defer result.Close()

	var stats Stats
	result.Next()
	err = result.StructScan(&stats)
	util.CheckErr(err)

	return stats
}

func (d *Db) CheckUser(id int) bool {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	err := db.QueryRow(fmt.Sprintf(
		"select id from users where id=%d",
		id)).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}

func (d *Db) DeleteUser(id int) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	query := fmt.Sprintf(
		"id='%d'",
		id,
	)

	delete, err := db.Query(fmt.Sprintf(
		"delete from users where %s;",
		query,
	))

	util.CheckErr(err)
	defer delete.Close()
}

func (d *Db) DeleteUsers(ids []int) {
	for _, id := range ids {
		d.DeleteUser(id)
	}
}
