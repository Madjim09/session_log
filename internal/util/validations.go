package util

import (
	"errors"
	"strings"
	"time"
	"unicode"
)

var (
	ErrEmptyLine = errors.New("empty line")
	text         string
	runes        []rune
	parts        []string
	replaser     = strings.NewReplacer(" ", "-", ".", "-", "\\", "-", "/", "-")
)

// ValidFIO проверяет и форматирует строку с ФИО пациента.
//
// Функция выполняет следующие действия:
//   - Использует ValidSpec для первичной валидации строки (предположительно — проверяет допустимые символы).
//   - Разбивает строку на части по пробельным символам с помощью strings.FieldsFunc.
//   - Приводит первую букву каждого слова к верхнему регистру, остальные — оставляет без изменений.
//   - Объединяет части обратно в строку с одинарными пробелами.
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

// ValidSpec проверяет строку, представляющую специальность врача, на корректность.
//
// Функция выполняет следующие действия:
//   - Проверяет, не является ли входная строка пустой. Если строка пуста — возвращает ошибку ErrEmptyLine.
//   - Удаляет пробельные символы с начала и конца строки с помощью strings.TrimSpace.
func ValidSpec(str string) (string, error) {
	if str == "" {
		return "", ErrEmptyLine
	}

	text = strings.TrimSpace(str)
	return text, nil
}

// ValidDate преобразует строку в объект time.Time, проверяя корректность формата даты.
//
// Функция выполняет следующие шаги:
//   - Очищает входную строку `str` с помощью регулярного выражения `replaser` (например, удаляет лишние пробелы или символы).
//   - Парсит очищенную строку `text` в формате "2006-01-02", который соответствует ISO 8601 (год-месяц-день).
func ValidDate(str string) (time.Time, error) {
	text = replaser.Replace(str)
	date, err := time.Parse("2006-01-02", text)
	return date, err
}
