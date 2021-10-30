package service

import (
	"crawl/pkg/model"
	"crawl/pkg/repo"
)

type Service struct {
	Repo repo.IRepo
}

type IService interface {
	CreateCollection(collect *model.Collection) error
	CreateProduct(collect *model.Item) error
}

func NewService(repo repo.IRepo) IService {
	return &Service{
		Repo: repo,
	}
}

func (h *Service) CreateCollection(collect *model.Collection) error {
	err := h.Repo.CreateCollection(collect)
	if err != nil {
		return err
	}
	return nil
}
func (h *Service) CreateProduct(product *model.Item) error {
	err := h.Repo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}
