package intro

import "testing"

// Модульный тест.
func TestService_Name(t *testing.T) {
	name := "Test Service"
	s := New(name)
	//s:= New(123)
	got := s.Name()
	if got != name {
		t.Errorf("got %s, want %s", got, name)
	}
}
