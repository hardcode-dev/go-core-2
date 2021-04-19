package main

import (
	"errors"
	"fmt"
	"os"
)

// Пользовательский тип данных для ошибки. Совместим со встроенным.
type customErr struct {
	msg string
}

// Выполнение контракта интерфейса.
func (ce customErr) Error() string {
	return fmt.Sprintf("пользовательская ошибка: %s", ce.msg)
}

func newCustomErr(msg string) customErr {
	return customErr{msg: msg}
}

func main() {
	err := errors.New("пример создания ошибки с помощью пакета errors") // с маленькой буквы!
	fmt.Println(err.Error())

	err = fmt.Errorf("пример создания ошибки с помощью пакета fmt")
	fmt.Println(err) // почему без вызова метода Error()?

	var cErr customErr
	val, err := envVar("1234")
	if err == cErr {
		fmt.Println("customErr")
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	err = cErr // корректно
}

// envVar возвращает переменную окружения, заданную по имени.
// Если переменная не найдена - возвращается ошибка.
func envVar(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "не найдено", newCustomErr("не найдено")
	}
	return val, nil
}
