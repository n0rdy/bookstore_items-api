package services

import (
	"github.com/n0rdy/bookstore_items-api/domain/items"
	"github.com/n0rdy/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	return nil, nil
}
