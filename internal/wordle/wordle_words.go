package wordle

import (
	"encoding/json"
	"os"

	"cordle/pkg/util"
)

// LoadWords reads the json files containing allowed words and answers and returns them (answers, guesses)
func LoadWords() ([]string, []string) {
	afile, err := os.ReadFile("wordle/answers.json")
	util.CheckErrMsg(err, "Failed to load answers")
	gfile, err := os.ReadFile("wordle/guesses.json")
	util.CheckErrMsg(err, "Failed to load guesses")

	// Decode JSON
	answers := []string{}
	err = json.Unmarshal(afile, &answers)
	util.CheckErrMsg(err, "Failed to decode answers")

	guesses := []string{}
	err = json.Unmarshal(gfile, &guesses)
	util.CheckErrMsg(err, "Failed to decode guesses")

	return answers, guesses
}
