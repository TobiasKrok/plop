package utils

import "strings"

func GetAndParseSessionID(username string) (id string, exists bool) {
	return strings.CutPrefix(username, "s:")
}
