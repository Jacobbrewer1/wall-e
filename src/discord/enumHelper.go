package discord

import "strings"

func parseText(stringValue string) string {
	return strings.ReplaceAll(strings.ToUpper(
		strings.Join(
			strings.Split(
				strings.ReplaceAll(
					strings.Join(
						strings.Fields(stringValue),
						" "),
					" ", "_"),
				""),
			""),
	), "\"", "")
}

func displayFormatter(enum iEnum) (returnText string) {
	split := strings.Split(strings.ToLower(enum.String()), "_")

	for _, val := range split {
		if returnText != "" {
			returnText = returnText + " "
		}

		returnText = returnText + strings.ToUpper(val[:1]) + val[1:]
	}

	return returnText
}
