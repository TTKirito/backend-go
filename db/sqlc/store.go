package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

type CreateEventTxParams struct {
	Title        string        `json:"title"`
	StartTime    time.Time     `json:"start_time"`
	EndTime      time.Time     `json:"end_time"`
	IsEmegency   bool          `json:"is_emegency"`
	Owner        int64         `json:"owner"`
	Note         string        `json:"note"`
	Type         EventTypes    `json:"type"`
	VisitType    VisitTypes    `json:"visit_type"`
	Meeting      string        `json:"meeting"`
	Location     Location      `json:"location"`
	Participants []Participant `json:"participants"`
}

type CreateEventTxResult struct {
	Event        Event         `json:"event"`
	Location     Location      `json:"location"`
	Participants []Participant `json:"participants"`
}

func (store *Store) CreateEventTx(ctx context.Context, arg CreateEventTxParams) (CreateEventTxResult, error) {
	var result CreateEventTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Event, err = q.CreateEvent(ctx, CreateEventParams{
			Title:      sql.NullString{String: arg.Title, Valid: true},
			StartTime:  arg.StartTime,
			EndTime:    arg.EndTime,
			IsEmegency: arg.IsEmegency,
			Owner:      arg.Owner,
			Note:       sql.NullString{String: arg.Note, Valid: true},
			Type:       EventTypes(arg.Type),
			VisitType:  VisitTypes(arg.VisitType),
			Meeting:    sql.NullString{String: arg.Meeting, Valid: true},
		})

		if err != nil {
			return err
		}

		result.Location, err = q.CreateLocation(ctx, CreateLocationParams{
			Lat:    arg.Location.Lat,
			Long:   arg.Location.Long,
			Street: arg.Location.Street,
			Event:  result.Event.ID,
		})

		if err != nil {
			return err
		}

		for _, participant := range arg.Participants {

			res, err := q.CreateParticipant(ctx, CreateParticipantParams{
				Event:   result.Event.ID,
				Account: participant.ID,
			})

			if err != nil {
				return err
			}

			result.Participants = append(result.Participants, res)

		}

		return nil
	})

	return result, err
}
