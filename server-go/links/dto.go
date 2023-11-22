package links

import "time"

type CreateLinkDto struct {
	Name        string    `json:"name" validate:"required"`
	OriginalUrl string    `json:"original_url" validate:"required,url"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Alias       string    `json:"alias,omitempty"` // TODO: add better validation
	UserId      uint      `json:"user_id,omitempty"`
}
