package Auth

import "github.com/google/uuid"

type Client struct {
	Id           uuid.UUID
	Secret       uuid.UUID
	RedirectUris []string
}

func New() {

}
