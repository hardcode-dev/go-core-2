package main

// Интерфейс.
type WalkerTalker interface {
	Walk()        // контракт интерфейса
	Talk() string // контракт интерфейса
}

// Пользовательский тип данных.
type Guy struct{}

// Метод (передача по значению).
func (g Guy) Walk() {
	println("I walk!")
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
}
