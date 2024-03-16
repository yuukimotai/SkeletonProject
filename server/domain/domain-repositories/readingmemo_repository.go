package repositories

import (
	"context"
	"obserbooks/domain/model"
)

type ReadingMemoRepository interface {
	CreateReadingMemo(context.Context, model.NewReadingMemo) (*model.ReadingMemo, error)
	UpdateReadingMemo(context.Context, model.UpdateReadingMemo) (*model.ReadingMemo, error)
	DeleteReadingMemo(context.Context, model.DeleteReadingMemo) (int, error)
}
