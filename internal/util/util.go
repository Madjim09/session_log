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

// Save считывает данные о пациенте из *bufio.Scanner, валидирует их и сохраняет в базу.
// Ожидает три строки подряд: ФИО пациента, специальность врача и дату посещения.
// В случае успешного чтения и валидации добавляет запись в глобальную мапу base,
// где ключ — ФИО пациента, а значение — срез его посещений (UserCard).
//
// Выводит приглашения к вводу на стандартный поток вывода.
// Если запись с такими же данными уже существует, она будет добавлена повторно —
// дедупликация не выполняется.
func Save(scanner *bufio.Scanner) (map[string][]UserCard, error) {
	fmt.Print("ФИО: ")
	if !scanner.Scan() {
		return base, scanner.Err()
	}
	name, err = ValidFIO(scanner.Text())
	if err != nil {
		return base, err
	}

	fmt.Print("Специальность врача: ")
	if !scanner.Scan() {
		return base, scanner.Err()
	}
	spec, err = ValidSpec(scanner.Text())
	if err != nil {
		return base, err
	}

	fmt.Print("Введите дату посещения в формате \"YYYY-MM-DD\": ")
	if !scanner.Scan() {
		return base, scanner.Err()
	}
	date, err = ValidDate(scanner.Text())
	if err != nil {
		return base, err
	}

	base[name] = append(base[name], UserCard{Specialization: spec, Date: date})
	fmt.Print("Карта пациента добавлена.\n\n")

	return base, nil
}

// GetHistory считывает ФИО пациента из *bufio.Scanner, проверяет его корректность
// и возвращает историю посещений этого пациента из глобальной базы данных base.
//
// Функция ожидает одну строку с ФИО, выводит приглашение к вводу на стандартный поток.
// После валидации имени с помощью ValidFIO проверяет наличие записей в базе.
// Если пациент не найден (нет ни одной записи), возвращается ошибка ErrPatientNotFound.
// В случае успеха выводит отформатированный список всех посещений на экран.
func GetHistory(scanner *bufio.Scanner) ([]UserCard, error) {

	fmt.Print("Введите ФИО: ")
	if !scanner.Scan() {
		return nil, scanner.Err()
	}
	name, err = ValidFIO(scanner.Text())
	if err != nil {
		return []UserCard{}, err
	}

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

// GetLastVisit считывает из *bufio.Scanner ФИО пациента и специальность врача,
// затем находит последнюю по дате запись о посещении данного пациента к врачу указанной специальности.
//
// Функция выполняет следующие шаги:
// 1. Считывает строку с ФИО и валидирует её с помощью ValidFIO.
// 2. Считывает строку с названием специальности и валидирует её с помощью ValidSpec.
// 3. Ищет в глобальной базе base все посещения указанного пациента.
// 4. Если посещений нет — возвращает ошибку ErrPatientNotFound.
// 5. Если посещение одно — возвращает его, независимо от специальности.
// 6. Если посещений несколько — находит самое позднее к врачу с указанной специальностью.
func GetLastVisit(scanner *bufio.Scanner) (UserCard, error) {
	fmt.Print("Введите ФИО: ")
	if !scanner.Scan() {
		return UserCard{}, scanner.Err()
	}
	name, err = ValidFIO(scanner.Text())
	if err != nil {
		return UserCard{}, err
	}

	fmt.Print("Специальность врача: ")
	if !scanner.Scan() {
		return UserCard{}, scanner.Err()
	}
	spec, err = ValidSpec(scanner.Text())
	if err != nil {
		return UserCard{}, err
	}

	visits := base[name]
	if len(visits) == 0 {
		return UserCard{}, ErrPatientNotFound
	} else if len(visits) == 1 {
		return base[name][0], nil
	}
	max := base[name][0].Date
	imin := 0
	for i, v := range base[name] {
		if v.Date.After(max) && v.Specialization == spec {
			max = v.Date
			imin = i
		}
	}

	fmt.Printf("Последнее посещение: %s\n\n", base[name][imin].Date.Format("2006-01-02"))

	return base[name][imin], nil
}

func ErrorHandling(e error, scanner *bufio.Scanner) bool {
	if err != nil {
		err = fmt.Errorf("ошибка ввода: %w", err)
		fmt.Print(err)
		fmt.Print("\n\n")
		return true
	}
	return true
}
