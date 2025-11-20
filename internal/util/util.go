package util

import (
	"bufio"
	"errors"
	"fmt"
	"time"
)

var (
	base               = map[string][]UserCard{}
	name               string
	spec               string
	date               time.Time
	err                error
	ErrPatientNotFound = errors.New("patient not found")
)

// Save сохраняет данные о визите в карту
func Save(scanner *bufio.Scanner) (map[string][]UserCard, error) {
	if !scanner.Scan() {
		return base, scanner.Err()
	}

	fmt.Print("ФИО: ")
	name, err = ValidFIO(scanner.Text())
	if err != nil {
		return base, err
	}

	if !scanner.Scan() {
		return base, scanner.Err()
	}

	fmt.Print("Специальность врача: ")
	spec, err = ValidSpec(scanner.Text())
	if err != nil {
		return base, err
	}

	if !scanner.Scan() {
		return base, scanner.Err()
	}

	fmt.Print("Введите дату посещения в формате \"YYYY-MM-DD\": ")
	date, err = ValidDate(scanner.Text())
	if err != nil {
		return base, err
	}

	base[name] = append(base[name], UserCard{Specialization: spec, Date: date})
	fmt.Print("Карта пациента добавлена.\n\n")

	return base, nil
}

// GetHistory позволяет получить историю посещений пациента
func GetHistory(scanner *bufio.Scanner) ([]UserCard, error) {
	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	fmt.Print("Введите ФИО: ")
	name, err = ValidFIO(scanner.Text())

	card := base[name]
	if len(card) == 0 {
		return []UserCard{}, ErrPatientNotFound
	}
	fmt.Printf("Список всех посещений пациента:\n")
	for i, v := range card {
		fmt.Printf("%d. %s %v\n", i+1, v.Specialization, v.Date.Format("2006-01-02"))
	}
	fmt.Println()

	return card, nil
}

// TODO: Операция GetLastVisit позволяет получить последнее посещение пациентом определенного специалиста в больнице
