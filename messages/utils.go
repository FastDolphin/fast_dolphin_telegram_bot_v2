package messages

import (
	"encoding/json"
	"os"
)

func LoadMessages(path string) (*Messages, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var messages Messages
	err = json.Unmarshal(file, &messages)
	if err != nil {
		return nil, err
	}

	return &messages, nil
}
