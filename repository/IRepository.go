package repository

type IDatabase[T any] interface {
	Save(data *T) error
	FindById(id int) (any, error)
	FindAll() ([]T, error)
	Delete(id int) error
}
