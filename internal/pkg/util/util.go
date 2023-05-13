package util

import (
	"database/sql"
	"log"
)

// CheckError logs a fatal error if err != nil
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckErrMsg(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func CheckRow(err error) (bool, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatalln(err)
		}
		return false, nil
	}
	return true, nil
}