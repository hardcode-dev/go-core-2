// Пакет - основная программная единица приложения.
package intro

// Пользовательский тип. Структура.
type Service struct {
	name string
}

// Функция. Выполняет роль конструкторы.
func New(name string) *Service {
	var s Service
	s.name = name
	return &s
}

// Метод типа даных Service.
func (s *Service) Name() string {
	return s.name
}
