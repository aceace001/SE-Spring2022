package sprint3

import (
	"time"
)

//post model
type Post struct {
	ID        int64     `json:"id"`
	UerID     int64     `json:""`
	Content   string    `json:"content"`
	SpoilerOf *string   `json:"spoilerOf"`
	NSFW      bool      `json:"nsfw"`
	CreatedAt time.Time `json:"createdAt"`
	User      *User     `json:"user.omitempty"`
	Mine      bool      `json:"mine"`
}
