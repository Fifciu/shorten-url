package links

import "time"

type CreateLinkDto struct {
	Name        string    `json:"name" validate:"required"`
	OriginalUrl string    `json:"original_url" validate:"required,url"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Alias       string    `json:"alias,omitempty"` // TODO: add better validation
	UserId      uint      `json:"user_id,omitempty"`
}

type UpdateLinkDto struct {
	Name        *string `json:"name"`
	OriginalUrl *string `json:"original_url,omitempty" validate:"omitempty,url""`
	Alias       *string `json:"alias"` // TODO: add better validation
}
