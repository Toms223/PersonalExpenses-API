package App

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Category struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
}

func (c *Category) Create(ctx context.Context, db *pgxpool.Pool) error {
	var id int
	err := db.QueryRow(ctx,
		`INSERT INTO "Categories" ("UserId", "Name", "Color") VALUES ($1, $2, $3) RETURNING "Id";`,
		c.UserId, c.Name, c.Color,
	).Scan(id)
	if err != nil {
		return err
	}
	c.Id = id
	return nil
}

func (c *Category) Save(ctx context.Context, db *pgxpool.Pool) error {
	cmdTag, err := db.Exec(ctx,
		`UPDATE "Categories" SET "UserId" = $1, "Name" = $2, "Color" = $3 WHERE "Id" = $4;`,
		c.UserId, c.Name, c.Color, c.Id,
	)
	if err != nil {
		return c.Reload(ctx, db)
	}
	if cmdTag.RowsAffected() == 0 {
		return c.Reload(ctx, db)
	}
	return nil
}

func (c *Category) Delete(ctx context.Context, db *pgxpool.Pool) error {
	cmd, err := db.Exec(ctx,
		`DELETE FROM "Categories" WHERE "Id" = $1;`,
		c.Id,
	)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("category with id %d not found", c.Id)
	}
	*c = Category{}
	return nil
}

func (c *Category) Reload(ctx context.Context, db *pgxpool.Pool) error {
	return db.QueryRow(ctx,
		`SELECT "Id", "UserId", "Name", "Color" FROM "Categories" WHERE "Id" = $1;`,
		c.Id,
	).Scan(&c.Id, &c.UserId, &c.Name, &c.Color)
}
