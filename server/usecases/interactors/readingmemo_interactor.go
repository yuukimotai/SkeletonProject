package interactors

import (
	database "obserbooks/infrastructure/mysql"
	repositories "obserbooks/infrastructure/repositories/memo"
	"obserbooks/usecases/boundaries"
)

var _ boundaries.ReadingMemoInputBoundary = &ReadingMemoInteractor{}

type ReadingMemoInteractor struct {
	ReadingMemoRepository *repositories.ReadingMemoRepository
	db                    database.DB
}

func NewReadingMemoInteractor(readingMemoRepository *repositories.ReadingMemoRepository) *ReadingMemoInteractor {
	return &ReadingMemoInteractor{
		ReadingMemoRepository: readingMemoRepository,
	}
}
