package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Settings stores the launch params in a json file
type Settings struct {
	GoogleSheetID string
}

// Setup looks for a settings file. If it doesn't exist it creates it and returns false.
func (s *Settings) Setup() error {

	const settingsFilename = "jp_settings.json"

	b, err := ioutil.ReadFile(settingsFilename)
	if err != nil {

		// Assume file doesn't exist, and create it
		s.GoogleSheetID = "Enter Google Sheets ID here"

		settingsJSON, marshalErr := json.Marshal(s)
		if marshalErr != nil {
			return fmt.Errorf("error creating json from settings struct")
		}

		ioutil.WriteFile(settingsFilename, settingsJSON, 0644)
		return fmt.Errorf("settings file %s created - please restart", settingsFilename)
	}

	jsonErr := json.Unmarshal(b, s)
	if jsonErr != nil {
		return fmt.Errorf("json parsing error - please check settings file")
	}

	return nil
}
