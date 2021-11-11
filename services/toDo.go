package services

import (
	"TubesBigData/model"
	"TubesBigData/repository"
)

func AddTodo(toDo *model.ToDo) (interface{}, error) {
	return repository.AddTodo(toDo)
}

func GetOne(id string) (*model.ToDo, error) {
	return repository.GetOne(id)
}

func DeleteOne(id string) (int, error) {
	return repository.DeleteOne(id)
}

func UpdateOne(id string, newToDo model.ToDo) (int, error) {
	return repository.UpdateOne(id, newToDo)
}

func DeleteMultiple(listToDelete []string) (int, error) {
	return repository.DeleteMultiple(listToDelete)
}

func GetAll() ([]model.ToDo, error) {
	return repository.GetAll()
}
