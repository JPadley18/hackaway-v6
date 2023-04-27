package config

import (
	"cordle/pkg/util"
	"encoding/json"
	"os"
)

// Config structs store configuration information after they are read
type Config struct {
	Token  string
	Status string
}

// LoadConfig returns a Config struct after reading configuration variables from a JSON file
func LoadConfig(path string) Config {
	// Open the configuration file
	file, err := os.ReadFile(path)
	util.CheckErrMsg(err, "Failed to read config file")

	// Create a new Config struct to return
	ret := Config{}
	// Decode JSON into the struct
	err = json.Unmarshal(file, &ret)
	util.CheckErrMsg(err, "Failed to decode JSON from config file")

	return ret
}
