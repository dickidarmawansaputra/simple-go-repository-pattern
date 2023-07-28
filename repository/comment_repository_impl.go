package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dickidarmawansaputra/simple-go-repository-pattern/entity"
	"strconv"
)

// comment repository implementation
type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	resutl, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		panic(err)
	}

	id, err := resutl.LastInsertId()
	if err != nil {
		panic(err)
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, nil
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}

	return comment, errors.New("id " + strconv.Itoa(int(id)) + " not found")
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
