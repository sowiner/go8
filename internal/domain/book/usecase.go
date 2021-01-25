package book

import (
	"context"

	"github.com/gmhafiz/go8/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, title, description, imageURL, publishedDate string) (*models.Book, error)
	All(ctx context.Context) ([]*models.Book, error)
	Find(ctx context.Context, bookID int64) (*models.Book, error)
	Update(ctx context.Context, book *models.Book) (*models.Book, error)
	Delete(ctx context.Context, bookID int64) error
}
