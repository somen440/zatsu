package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/somen440/zatsu/golang/sqlboiler_inner_join_test/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type AuthorPost struct {
	ID          int64     `boil:"author_id"`
	Name        string    `boil:"author_name"`
	Title       string    `boil:"title"`
	Description string    `boil:"description"`
	Content     string    `boil:"content"`
	Date        time.Time `boil:"date"`
}

func (ap *AuthorPost) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("ID: %d\n", ap.ID))
	out.WriteString(fmt.Sprintf("Name: %s\n", ap.Name))
	out.WriteString(fmt.Sprintf("Title: %s\n", ap.Title))
	out.WriteString(fmt.Sprintf("Description: %s\n", ap.Description))
	out.WriteString(fmt.Sprintf("Content: %s\n", ap.Content))
	out.WriteString(fmt.Sprintf("Date: %+v\n", ap.Date))

	return out.String()
}

func SelectAuthorPosts(db *sql.DB, ctx context.Context, userIDs ...interface{}) ([]*AuthorPost, error) {
	entities := []*AuthorPost{}

	if err := models.Authors(
		qm.Select(
			"authors.id as author_id",
			"authors.first_name as author_name",
			"p.title as title",
			"p.description as description",
			"p.content as content",
			"p.date as date",
		),
		qm.InnerJoin("posts p on authors.id = p.author_id"),
		qm.WhereIn("authors.id in ?", userIDs...),
	).Bind(ctx, db, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}
