package repository

import (
	"context"

	"github.com/dickidarmawansaputra/simple-go-repository-pattern/entity"
)

type CommentRepository interface {
	// jika butuh DB Transaction masukan parameter tx
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
