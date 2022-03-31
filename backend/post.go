package sprint3

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

var(
	//ErrInvalidContent is used for invalid content.
	ErrInvalidContent = errors.New("invalid content")
	//ErrInvalidSpoiler used for invalid spoiler title.
	ErrInvalidSpoiler = errors.New("invalid spoiler")
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

//CreatePost publishes a post the user timeline and fan-outs it to his follwers
func (s *Service) CreatePost(
	ctx context.Context,
	content string,
	spoilerOf *string,
	nsfw bool,
)(TimelineItem,error){
	var ti TimelineItem
	uid,ok := ctx.Value(KeyAuthUserID).(int64)
	if !ok{
		return ti,ErrUnauthenticated
	}

	content = strings.TrimSpace(content)
	if content == "" || len([]rune(content))>480{
		return ti,ErrInvalidContent
	}

	if spoilerOf !=nil{
		*spoilerOf =strings.TrimSpace(*spoilerOf)
		if *spoilerOf == "" || len([]rune(*spoilerOf)) > 64 {
			return ti,ErrInvalidSpoiler
		}
	}

	tx, err := s.db.BeginTx(ctx,nil)
	if err !=nil{
		return ti,fmt.Errorf("could not begin tx: %v", err)
	}

	defer tx.Rollback()

	query :="INSERT INTO posts (user_id, content, spoiler_of, nsfw) VALUES ($1, $2, $3, $4) "+
		"RETURNIG id, created_at"
	if err = tx.QueryRowContext(ctx, query, uid, content, spoilerOf, nsfw).
		Scan(&ti.Post.ID, &ti.Post.CreatedAt); err != nil{
			return ti, fmt.Errorf("could not insert post:%v",err)
		}

		ti.Post.UerID = uid
		ti.post.Content = content
		ti.Post.SpoilerOf = spoilerOf
		ti.Post.NSFW = nsfw
		ti.Post.Mine = true

	return ti,nil
}
