package usecases_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/alextownleymakes/react-golang-tutorial/backend/entities"
	"github.com/alextownleymakes/react-golang-tutorial/backend/usecases"
)

var dummyTodos = []entities.Todo{
	{Title: "Laundry", Description: "take some to your moms", IsCompleted: false},
	{Title: "Dishes", Description: "just wash them", IsCompleted: false},
	{Title: "Organize music room", Description: "gotta get rid of that couch", IsCompleted: false},
	{Title: "Empty boxes", Description: "one at a time", IsCompleted: false},
	{Title: "Hi five sam", Description: "he deserves it", IsCompleted: false},
}

type BadTodosRepo struct{}

func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("Something went wrong!")
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {
	//Test
	t.Run("Returns ErrInternal when TodosRepository returns error", func(t *testing.T) {
		expect := expectate.Expect(t)
	
		repo := new(BadTodosRepo)
	
		todos, err := usecases.GetTodos(repo)
	
		expect(err).ToBe(usecases.ErrInternal)
	
		// expect(todos).ToBe(nil)
		if todos != nil {
			t.Fatalf("Expected todos to be nil; got %v", todos)
		}
	})
	// Test
	t.Run("Returns todos from TodosRepository", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(MockTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})
}