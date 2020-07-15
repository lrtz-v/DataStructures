package replaceSpace

import "strings"

func replaceSpace(s string) string {
	var builder strings.Builder
	
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}