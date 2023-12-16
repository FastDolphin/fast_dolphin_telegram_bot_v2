package utils

import (
	"slices"
)

var constants *Config = New()

func IsAdmin(userChatID string) bool {
	return userChatID == constants.AdminChatIDEvg || userChatID == constants.AdminChatIDNat
}

func IsCoach(userChatID string) bool {
	return slices.Contains(constants.CoachChatIDs, userChatID)
}
