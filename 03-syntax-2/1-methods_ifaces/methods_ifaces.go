package main

import "fmt"

// Интерфейс.
type WalkerTalker interface {
	Walk()        // контракт интерфейса
	Talk() string // контракт интерфейса
	Eater
}

type MyIface interface {
}

type Eater interface {
	Eat() string
}

// Пользовательский тип данных.
type Guy struct {
	Eater
}

// Метод (передача по значению).
func (g *Guy) Walk() {
	fmt.Println("I walk!")
}

// Метод (передача по ссылке).
func (g *Guy) Talk() string {
	return "I talk!"
}

func main() {
	// Переменная интерфейсного типа инициализирована
	// литералом конкретного типа.
	var wt WalkerTalker = &Guy{}
	wt.Walk()
	println(wt.Talk())
	println(wt.Eat())
	var mi MyIface
	i := 10
	s := "str"
	mi = i
	mi = s
	_ = mi
}
