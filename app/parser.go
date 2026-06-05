package main

import (
	"strings"
)

func ParseQuotations(input string) ([]string, error) {
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

func ParseRedirection(parts []string) ([]string, string, error) {
	var args = parts
	var target string

	for i, r := range parts {
		if r == ">" || r == "1>" {
			args = parts[0:i]
			target = parts[i+1]
			break
		}
	}

	return args, target, nil
}
