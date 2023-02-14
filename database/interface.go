package database

import (
	"database/sql"
	"fmt"
)

type Interface struct {
	Connection string
}

func getRecord(selection string, table string, query string) *sql.Rows {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf("select %s from %s where %s;", selection, table, query))
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	result.Next()

	return result
}

func getRecords(table string, query string) [][]string {
	conn := Connect()
	record, err := conn.db.Query(fmt.Sprintf("select * from %s where %s;", table, query))
	if err != nil {
		panic(err.Error())
	}
	defer record.Close()

	return nil
}

func getTable(table string) [][]string {
	conn := Connect()
	record, err := conn.db.Query(fmt.Sprintf("select * from %s;", table))
	if err != nil {
		panic(err.Error())
	}
	defer record.Close()

	return nil
}

func insertRecord(query string) {
	conn := Connect()
	insert, err := conn.db.Query(query);
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func insertRecords() {

}

func updateRecord() {

}

func updateRecords() {

}

func deleteRecord() {

}

func deleteRecords() {

}

func deleteTable() {

}

func deleteTables() {

}
