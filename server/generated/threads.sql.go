// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: threads.sql

package queries

import (
	"context"
)

const getThreadByID = `-- name: GetThreadByID :one
SELECT id, firstname, lastname, title, content, category, upvotes, created_at
FROM threads
WHERE id = $1
`

func (q *Queries) GetThreadByID(ctx context.Context, id int32) (Thread, error) {
	row := q.db.QueryRow(ctx, getThreadByID, id)
	var i Thread
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Title,
		&i.Content,
		&i.Category,
		&i.Upvotes,
		&i.CreatedAt,
	)
	return i, err
}

const insertThread = `-- name: InsertThread :exec
INSERT INTO threads (firstname, lastname, title, content, category, upvotes, created_at)
VALUES ($1, $2, $3, $4, $5, 0, CURRENT_TIMESTAMP)
`

type InsertThreadParams struct {
	Firstname string
	Lastname  string
	Title     string
	Content   string
	Category  string
}

func (q *Queries) InsertThread(ctx context.Context, arg InsertThreadParams) error {
	_, err := q.db.Exec(ctx, insertThread,
		arg.Firstname,
		arg.Lastname,
		arg.Title,
		arg.Content,
		arg.Category,
	)
	return err
}
