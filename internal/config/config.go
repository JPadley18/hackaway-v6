package config

import (
	"encoding/json"
	"os"

	"cordle/pkg/util"
)

// ConfigData structs store configuration information after they are read
type ConfigData struct {
	Token  string
	Status string
	Game   GameConfig
}

// GameConfig stores game specific settings
type GameConfig struct {
	ChallengeDuration int
}

// The path to read the config from
const configPath = "config/config.json"

// Globally available config data
var Config ConfigData

// When the module is first imported, load the config from a JSON file
func init() {
	// Open the configuration file
	file, err := os.ReadFile(configPath)
	util.CheckErrMsg(err, "Failed to read config file")

	// Decode JSON into the struct
	err = json.Unmarshal(file, &Config)
	util.CheckErrMsg(err, "Failed to decode JSON from config file")
}
