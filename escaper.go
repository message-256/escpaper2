package escpaper

import (
	"strings"
)

func Escape(input string) string {
	var escaped string
	if strings.Contains(input, "\x1b") {
		return input
	}
	if len(input) <= 1 {
		return ""
	}
	//escapin
	for i := range input {
		if input[i] == '\\' {
			if i == 0 {
				escaped += "\x1b"
			} else if escaped[i-1] != '\x1b' {
				escaped += "\x1b"
			} else {
				escaped += "\\"
			}
		} else {
			escaped += string(input[i])
		}
	}
	return escaped
}
func SubString(input string, delim rune) string {
	var next int
	var returned string
	if len(input) <= 1 {
		return ""
	}
	for {
		next = strings.Index(input, string(delim))
		if next == -1 {
			return returned + input
		}
		escapedsection := Escape(string(input[:next+1]))
		returned += escapedsection
		input = input[next+1:]
		if escapedsection[len(escapedsection)-2] != '\x1b' {
			break
		}

	}
	return returned

}
