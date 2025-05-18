package utils

import "strings"

func TrimTelegramCommand(sourceString string) (string) {
	spaceIndex := strings.Index(sourceString, " ")
	if spaceIndex == -1 {
		return ""
	}
	return strings.TrimSpace(sourceString[spaceIndex+1:])
}

func JoinWithParagraphs(lines []string) string {
	return strings.Join(lines, "\n\n")
}