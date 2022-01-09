package usecases

import "github.com/alextownleymakes/react-golang-tutorial/backend/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}