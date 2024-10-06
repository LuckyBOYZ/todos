package repository

type IRepository[T any] interface {
	Save(data *T) error
	FindById(id int) (*T, error)
	FindAll() ([]T, error)
	Delete(id int) error
}

type ITodoRepository[T any] interface {
	IRepository[T]
	FindAllNotFinishedTodos() ([]Todo, error)
}
