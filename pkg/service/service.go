package service

import "ClientBaseWEBStudio/pkg/repository"

type Authorization interface {
}

type Order interface {
}
type Client interface {
}
type Product interface {
}

type Service struct {
	Authorization
	Order
	Client
	Product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
