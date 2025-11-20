package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Madjim09/session_log/internal/util"
)

var err error

func main() {
	flag := true
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Команды:\nSave\nGetHistory\nExit\n\n")
	for flag {
		fmt.Println("Введите команду")
		if !scanner.Scan() {
			err := fmt.Errorf("ошибка ввода: %w", scanner.Err())
			fmt.Print(err)
			flag = false
			continue
		}

		command := strings.TrimSpace(strings.ToLower(scanner.Text()))
		switch command {
		case "save", "s":
			_, err = util.Save(scanner)
			flag = errorHandling(err, scanner)
		case "gethistory", "gh":
			_, err = util.GetHistory(scanner)
			flag = errorHandling(err, scanner)
		case "getlastvisit", "glv":
			_, err = util.GetHistory(scanner)
			flag = errorHandling(err, scanner)
		case "exit", "e":
			flag = false
		default:
			fmt.Print("Неизвестная команда\n\n")
		}

	}
}

func errorHandling(e error, scanner *bufio.Scanner) bool {
	if err == scanner.Err() {
		err = fmt.Errorf("ошибка ввода: %w", err)
		fmt.Print(err)
		return false
	}
	if err != nil {
		err = fmt.Errorf("ошибка ввода: %w", err)
		fmt.Print(err)
		fmt.Print("\n\n")
		return true
	}
	return true
}
