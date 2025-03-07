package dto

import "github.com/google/uuid"

type UpdateUserDTO struct {
	ID    uuid.UUID
	Name  *string
	Email *string
}
