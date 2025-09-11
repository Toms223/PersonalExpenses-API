package Client

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Refresh struct {
	Value    uuid.UUID
	Expiry   time.Time
	UserId   int
	ClientId uuid.UUID
}

type Access struct {
	Value    *jwt.Token
	ClientId uuid.UUID
}

func NewRefresh(UserId int, ClientId uuid.UUID, ctx context.Context, db *pgxpool.Pool) *Refresh {
	expiry := time.Now().UTC().Add(24 * 30 * time.Hour)
	value := uuid.New()
	_, err := db.Exec(ctx, `INSERT INTO refresh_tokens (value, expiry, user_id, client_id) VALUES ($1, $2, $3, $4);`,
		value, expiry.UnixMilli(), UserId, ClientId,
	)
	if err != nil {
		fmt.Errorf("could not create new refresh token for %s client", ClientId)
	}
	return &Refresh{
		Value:    value,
		Expiry:   expiry,
		UserId:   UserId,
		ClientId: ClientId,
	}
}

func NewAccess(UserId int, ClientId uuid.UUID, SigningSecret string, ctx context.Context, db *pgxpool.Pool) *Access {
	expiry := time.Now().UTC().Add(time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": UserId,
		"exp":     expiry.UnixMilli(),
	})
	signedString, err := token.SignedString([]byte(SigningSecret))
	if err != nil {
		fmt.Errorf("could not sign new access token for %s client", ClientId)
	}
	_, err = db.Exec(ctx, `INSERT INTO access_tokens (value, client_id) VALUES ($1, $2);`,
		signedString, ClientId,
	)
	if err != nil {
		fmt.Errorf("could not create new acess token for %s client", ClientId)
	}
	return &Access{
		Value:    token,
		ClientId: ClientId,
	}
}
