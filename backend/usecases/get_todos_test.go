package usecases_test

import (
	"fmt"
	"testing"

	"github.com/alextownleymakes/react-golang-tutorial/backend/entities"
	"github.com/alextownleymakes/react-golang-tutorial/backend/usecases"
)

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("Something went wrong!")
}

func TestGetTodos(t *testing.T) {
	repo := new(MockTodosRepo)

	_, err := usecases.GetTodos(repo)

	if err != usecases.ErrInternal {
		t.Fatalf("expected ErrInternal; Got %v", err)
	}
}