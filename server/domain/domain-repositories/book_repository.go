package repositories

import (
	"context"
	"obserbooks/domain/model"
	"obserbooks/infrastructure/dto"
)

type BookRepository interface {
	CreateBook(context.Context, model.NewBook) (*dto.BookDTO, error)
}
