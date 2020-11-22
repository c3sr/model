package model

import (
	"time"

	"github.com/c3sr/config"
)

// easyjson:json
type JobResponse struct {
	ID        string                 `json:"id" validate:"required"`
	Kind      ResponseKind           `json:"kind" validate:"required"`
	Error     *Error                 `json:"error" `
	Body      []byte                 `json:"body" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at" validate:"required"`
}

func (JobResponse) TableName() string {
	return config.App.Name + "_jobs"
}
