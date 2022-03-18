package repository

type Authorization interface {
}

type Order interface {
}
type Client interface {
}
type Product interface {
}

type Repository struct {
	Authorization
	Order
	Client
	Product
}

func NewRepository() *Repository {
	return &Repository{}
}
