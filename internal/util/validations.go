package util

import (
	"errors"
	"strings"
	"time"
	"unicode"
)

var (
	errEmptyLine = errors.New("empty line")
	text         string
	runes        []rune
	parts        []string
	replaser     = strings.NewReplacer(" ", "-", ".", "-", "\\", "-", "/", "-")
)

func ValidFIO(str string) (string, error) {
	text, err = ValidSpec(str)

	if err != nil {
		return text, err
	}

	parts = strings.FieldsFunc(text, unicode.IsSpace)

	for i, p := range parts {
		runes = []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}

	return strings.Join(parts, " "), nil
}

func ValidSpec(str string) (string, error) {
	if str == "" {
		return "", errEmptyLine
	}

	text = strings.TrimSpace(str)
	return text, nil
}

func ValidDate(str string) (time.Time, error) {
	text = replaser.Replace(str)
	date, err := time.Parse("2006-01-02", str)
	return date, err
}
