package common

type Repository interface {
	FindAll() []interface{}
	GetTable() string
}
