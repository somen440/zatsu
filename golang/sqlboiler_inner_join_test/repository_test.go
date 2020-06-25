package main

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	return db, mock, func() {
		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
		db.Close()
	}
}

func TestSelectAuthorPosts(t *testing.T) {
	db, mock, def := setUp(t)
	defer def()

	authorID := 1

	query := "SELECT authors.id as author_id, "
	query += "authors.first_name as author_name, "
	query += "p.title as title, "
	query += "p.description as description, "
	query += "p.content as content, "
	query += "p.date as date "
	query += "FROM `authors` "
	query += "INNER JOIN posts p "
	query += "on authors.id = p.author_id "
	query += "WHERE \\(`authors`.`id` IN \\(\\?\\)\\);"

	rows := sqlmock.NewRows([]string{"author_id", "author_name", "title", "description", "content", "date"}).
		AddRow(authorID, "hoge name", "hoge title", "hoge description", "hoge content", time.Date(2019, 10, 10, 11, 11, 11, 0, time.Local))

	mock.ExpectQuery(query).
		WithArgs(authorID).
		WillReturnRows(rows)

	expected := []*AuthorPost{
		{
			int64(authorID),
			"hoge name",
			"hoge title",
			"hoge description",
			"hoge content",
			time.Date(2019, 10, 10, 11, 11, 11, 0, time.Local),
		},
	}
	actual, err := SelectAuthorPosts(db, context.Background(), authorID)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
