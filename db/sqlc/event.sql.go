// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: event.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
    title,
    start_time,
    end_time,
    is_emegency,
    owner,
    note,
    type,
    visit_type,
    meeting
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, title, start_time, end_time, is_emegency, owner, note, type, visit_type, meeting, created_at
`

type CreateEventParams struct {
	Title      sql.NullString `json:"title"`
	StartTime  time.Time      `json:"start_time"`
	EndTime    time.Time      `json:"end_time"`
	IsEmegency bool           `json:"is_emegency"`
	Owner      int64          `json:"owner"`
	Note       sql.NullString `json:"note"`
	Type       EventTypes     `json:"type"`
	VisitType  VisitTypes     `json:"visit_type"`
	Meeting    sql.NullString `json:"meeting"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Title,
		arg.StartTime,
		arg.EndTime,
		arg.IsEmegency,
		arg.Owner,
		arg.Note,
		arg.Type,
		arg.VisitType,
		arg.Meeting,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.StartTime,
		&i.EndTime,
		&i.IsEmegency,
		&i.Owner,
		&i.Note,
		&i.Type,
		&i.VisitType,
		&i.Meeting,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT id, title, start_time, end_time, is_emegency, owner, note, type, visit_type, meeting, created_at 
FROM events
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id int64) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.StartTime,
		&i.EndTime,
		&i.IsEmegency,
		&i.Owner,
		&i.Note,
		&i.Type,
		&i.VisitType,
		&i.Meeting,
		&i.CreatedAt,
	)
	return i, err
}

const listEvent = `-- name: ListEvent :many
SELECT id, title, start_time, end_time, is_emegency, owner, note, type, visit_type, meeting, created_at FROM events
WHERE start_time >= $1
AND end_time < $2
ORDER BY start_time
LIMIT $3
OFFSET $4
`

type ListEventParams struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

func (q *Queries) ListEvent(ctx context.Context, arg ListEventParams) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, listEvent,
		arg.StartTime,
		arg.EndTime,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.StartTime,
			&i.EndTime,
			&i.IsEmegency,
			&i.Owner,
			&i.Note,
			&i.Type,
			&i.VisitType,
			&i.Meeting,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEvent = `-- name: UpdateEvent :one
UPDATE events
SET  title = $2,
     start_time = $3,
     end_time = $4,
     is_emegency = $5,
     owner = $6,
     note = $7,
     type = $8,
     visit_type = $9,
     meeting = $10
WHERE id = $1
RETURNING id, title, start_time, end_time, is_emegency, owner, note, type, visit_type, meeting, created_at
`

type UpdateEventParams struct {
	ID         int64          `json:"id"`
	Title      sql.NullString `json:"title"`
	StartTime  time.Time      `json:"start_time"`
	EndTime    time.Time      `json:"end_time"`
	IsEmegency bool           `json:"is_emegency"`
	Owner      int64          `json:"owner"`
	Note       sql.NullString `json:"note"`
	Type       EventTypes     `json:"type"`
	VisitType  VisitTypes     `json:"visit_type"`
	Meeting    sql.NullString `json:"meeting"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEvent,
		arg.ID,
		arg.Title,
		arg.StartTime,
		arg.EndTime,
		arg.IsEmegency,
		arg.Owner,
		arg.Note,
		arg.Type,
		arg.VisitType,
		arg.Meeting,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.StartTime,
		&i.EndTime,
		&i.IsEmegency,
		&i.Owner,
		&i.Note,
		&i.Type,
		&i.VisitType,
		&i.Meeting,
		&i.CreatedAt,
	)
	return i, err
}
