package main

import "strings"

func ParseInput(input string) ([]string, error) {
	var parts []string
	var builder strings.Builder
	var inSingleQuote, inDoubleQuote, isEscaped bool

	for _, r := range input {
		if isEscaped {
			if inDoubleQuote && r != '"' && r != '\\' {
				builder.WriteRune('\\')
			}
			builder.WriteRune(r)
			isEscaped = false
			continue
		}

		switch r {
		case '\\':
			if inSingleQuote {
				builder.WriteRune(r)
			} else {
				isEscaped = true
			}

		case '\'':
			if inDoubleQuote {
				builder.WriteRune(r)
			} else {
				inSingleQuote = !inSingleQuote
			}

		case '"':
			if inSingleQuote {
				builder.WriteRune(r)
			} else {
				inDoubleQuote = !inDoubleQuote
			}

		case ' ':
			if inSingleQuote || inDoubleQuote {
				builder.WriteRune(r)
			} else if builder.Len() > 0 {
				parts = append(parts, builder.String())
				builder.Reset()
			}

		default:
			builder.WriteRune(r)
		}
	}

	if isEscaped {
		builder.WriteRune('\\')
	}
	if builder.Len() > 0 {
		parts = append(parts, builder.String())
	}

	return parts, nil
}
