package User

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Expense struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Amount     float32   `json:"amount"`
	Date       time.Time `json:"date"`
	UserId     int       `json:"user_id"`
	CategoryId int       `json:"category_id"`
	Continuous bool      `json:"continuous"`
	Fixed      bool      `json:"fixed"`
	Period     int       `json:"period"`
}

func (e *Expense) Create(ctx context.Context, db *pgxpool.Pool) error {
	var id int
	err := db.QueryRow(ctx, `
	INSERT INTO 
    "Expenses" ("Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING "Id";
	`,
		e.Name, e.Amount, e.Date, e.UserId, e.CategoryId, e.Continuous, e.Fixed, e.Period,
	).Scan(id)
	if err != nil {
		return err
	}
	e.Id = id
	return nil
}

func (e *Expense) Save(ctx context.Context, db *pgxpool.Pool) error {
	cmdTag, err := db.Exec(ctx,
		`UPDATE "Expenses"
			SET "Name" = $1,
			"Amount" = $2,
			"Date" = $3,
			"UserId" = $4,
			"CategoryId" = $5,
			"Continuous" = $6,
			"Fixed" = $7,
			"Period" = $8
			WHERE "Id" = $9;`,
		e.Name, e.Amount, e.Date, e.UserId, e.CategoryId, e.Continuous, e.Fixed, e.Period,
	)
	if err != nil {
		return e.Reload(ctx, db)
	}
	if cmdTag.RowsAffected() == 0 {
		return e.Reload(ctx, db)
	}
	return nil
}

func (e *Expense) Delete(ctx context.Context, db *pgxpool.Pool) error {
	cmd, err := db.Exec(ctx,
		`DELETE FROM "Expenses" WHERE "Id" = $1`,
		e.Id,
	)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("expense with id %d not found", e.Id)
	}
	*e = Expense{}
	return nil
}

func (e *Expense) Reload(ctx context.Context, db *pgxpool.Pool) error {
	return db.QueryRow(ctx,
		`SELECT "Id", "Name", "Amount", "Date", "UserId", "CategoryId","Continuous","Fixed","Period" FROM "Expenses" WHERE "Id" = $1`,
		e.Id,
	).Scan(&e.Id, &e.Name, &e.Amount, &e.Date, &e.UserId, &e.CategoryId, &e.Continuous, &e.Fixed, &e.Period)
}
