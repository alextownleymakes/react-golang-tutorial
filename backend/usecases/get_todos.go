package usecases

import "github.com/alextownleymakes/react-golang-tutorial/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	todos, err := repo.GetAllTodos()

	if err != nil {
		return nil, ErrInternal
	}

	return todos, nil
}