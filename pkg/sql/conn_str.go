package sql

import (
	"cordle/pkg/util"
	"encoding/json"
	"fmt"
	"os"
)

type connData struct {
	Username string
	Password string
	Address  string
	Port     string
	Database string
}

func connStr(path string) string {
	content, err := os.ReadFile(path)
	util.CheckErrMsg(err, "Error opening file: ")

	cd := connData{}
	err = json.Unmarshal(content, &cd)
	util.CheckErrMsg(err, "Error duing unmarshall(): ")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cd.Username,
		cd.Password,
		cd.Address,
		cd.Port,
		cd.Database)
}
