package repository

import (
	// koneksi = alias
	"context"
	"fmt"

	"testing"

	"github.com/dickidarmawansaputra/simple-go-repository-pattern/entity"

	koneksi "github.com/dickidarmawansaputra/simple-go-repository-pattern/database"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(koneksi.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repo@mail.com",
		Comment: "Repo comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(koneksi.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 33)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(koneksi.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}

}
