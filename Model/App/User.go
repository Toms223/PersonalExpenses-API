package App

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Limit int    `json:"limit"`
}

func (u *User) Create(ctx context.Context, db *pgxpool.Pool) error {
	var id int
	err := db.QueryRow(ctx,
		`INSERT INTO "Users" ("Name", "Email", "Limit") VALUES ($1, $2, $3) RETURNING "Id";`,
		u.Name,
		u.Email,
		u.Limit,
	).Scan(id)
	if err != nil {
		return err
	}
	u.Id = id
	return nil
}

func (u *User) Save(ctx context.Context, db *pgxpool.Pool) error {
	// Try update
	cmdTag, err := db.Exec(ctx,
		`UPDATE "Users" SET "Name" = $1, "Email" = $2, "Limit" = $3 WHERE "Id" = $4;`,
		u.Name, u.Email, u.Limit, u.Id,
	)
	if err != nil {
		return u.Reload(ctx, db)
	}
	if cmdTag.RowsAffected() == 0 {
		return u.Reload(ctx, db)
	}
	return nil
}

func (u *User) Delete(ctx context.Context, db *pgxpool.Pool) error {
	cmd, err := db.Exec(ctx,
		`DELETE FROM "Users" WHERE "Id" = $1;`,
		u.Id,
	)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", u.Id)
	}
	*u = User{}
	return nil
}

func (u *User) Reload(ctx context.Context, db *pgxpool.Pool) error {
	return db.QueryRow(ctx,
		`SELECT "Id", "Name", "Email", "Limit" FROM "Users" WHERE "Id" = $1;`,
		u.Id,
	).Scan(&u.Id, &u.Name, &u.Email, &u.Limit)
}
