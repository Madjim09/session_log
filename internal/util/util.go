package util

import (
	"bufio"
	"fmt"
	"time"
)

var (
	base = map[string]UserCard{}
	name string
	spec string
	date time.Time
	err  error
)

func Save(scanner *bufio.Scanner) error {
	if !scanner.Scan() {
		return scanner.Err()
	}

	fmt.Print("ФИО: ")
	name, err = ValidFIO(scanner.Text())
	if err != nil {
		return err
	}

	if !scanner.Scan() {
		return scanner.Err()
	}

	fmt.Print("Специальность врача: ")
	spec, err = ValidSpec(scanner.Text())
	if err != nil {
		return err
	}

	fmt.Print("Введите дату посещения в формате \"YYYY-MM-DD\": ")
	date, err = ValidDate(scanner.Text())
	if err != nil {
		return err
	}

	base[name] = UserCard{Specialization: spec, Date: date}

	return nil
}
