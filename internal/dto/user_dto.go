package dto

import "github.com/google/uuid"

type WorkspaceDto struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Domain  string    `json:"domain"`
	OwnerID uuid.UUID `json:"owner_id"`
	Plan    any       `json:"plan"`
}

type CreateWorkspaceDTO struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type UpdateWorkspaceDTO struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}
