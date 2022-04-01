package sprint3

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

var(
	//ErrInvalidContent is used for invalid content.
	ErrInvalidContent = errors.New("invalid content")
	//ErrInvalidSpoiler used for invalid spoiler title.
	ErrInvalidSpoiler = errors.New("invalid spoiler")
	ErrUnauthenticated = errors.New("Unauthenticated")
)

//post model
type Post struct {
	ID        int64     `json:"id"`
	UerID     int64     `json:""`
	Content   string    `json:"content"`
	SpoilerOf *string   `json:"spoilerOf"`
	NSFW      bool      `json:"nsfw"`
	LikesCount int 		`json:"likesCount"`
	CreatedAt time.Time `json:"createdAt"`
	User      *User     `json:"user.omitempty"`
	Mine      bool      `json:"mine"`
	Liked	  bool		`json:"liked"`
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
		Scan(&ti.Post.ID, &ti.Post.CreatedAt); err != nil {
		return ti, fmt.Errorf("could not insert post:%v", err)
	}

	ti.Post.UerID = uid
	ti.post.Content = content
	ti.Post.SpoilerOf = spoilerOf
	ti.Post.NSFW = nsfw
	ti.Post.Mine = true

	query = "INSERT INTO timeline (user_id, post_id) VALUES ($1, $2) RETURNING id"
	if err = tx.QueryRowContext(ctx, query, uid, ti.Post.ID).Scan(&ti.ID); err != nil {
		return ti, fmt.Errorf("could not insert timeline item: %v", err)
	}

	ti.UerID = uid
	ti.postID = ti.Post.ID

	if err = tx.Commit(); err !=nil{
		return ti, fmt.Errorf("could not commit to create post: %v", err)
	}

	go func(p Post) {
		u, err := s.userByID(context.Background(), p.UserID)
		if err != nil {
			log.Printf("could not get post user: %v\n", err)
			return
		}

		p.User = &u
		p.Mine = false

		tt, err := s.fanoutPost(p)
		if err != nil {
			log.Printf("could not fanout post: %v\n", err)
			return
		}

		for _, ti = range tt {
			log.Println(litter.Sdump(ti))
			//TODO:broadcast timeline items.
		}

	}(ti.Post)

	return ti, nil
}

func (s *Service) fanoutPost(p Post) ([]TimelineItem, error) {
	query := "INSERT INTO timeline(user_id, post_id)" +
		"SELECT follwer_id, $1 FROM follows WHERE followee_id = $2" +
		"RETURNING id, user_id"
	rows, err := s.db.Query(query, p.ID, p.UserID)
	if err != nil {
		return nil, fmt.Errorf("could not insert timeline: %v", err)
	}

	defer rows.Close()

	tt := []TimelineItem{}
	for rows.Next() {
		var ti TimelineItem
		if err = rows.Scan(&ti.ID, &ti.UserID); err != nil {
			return nil, fmt.Errorf("could not scan timeline item: %v", err)
		}

		ti.PostID = p.ID
		ti.Post = p
		tt = append(tt, ti)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate timeline rows: %v", err)
	}

	return tt, nil
}

// get posts from a user in descending order
func (s *Service) Posts() {
	ctx context.Context,
	username string,
	last int,
	before int64,
} ([]Post, error) {
	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username) {
		return nil, ErrInvalidUsername 
	}
	uid, auth := ctx.Value(KeyAuthUserID).(int64)
	last = normalizePageSize(last)
	query, args, err := buildQuery(`
		SELECT id, content, spoiler_of, nsfw, likes_count, created_at
		{{if .auth}} 
		, posts.user_id = @uid AS mine,
		, likes.user_id IS NOT NULL AS liked 
		{{end}}
		FROM posts
		{{if .auth}}
		LEFT JOIN post_likes AS likes
			ON likes.user_id = @uid AND likes.post_id = posts.id 
		{{end}}
		WHERE posts.user_id = (SELECT id FROM users WHERE username = @username)
		{{if .before}}AND posts.id < @before{{end}}
		ORDER BY created_at DESC 
		LIMIT @last
	`, map[string]interface{}{
		"auth": auth,
		"uid": uid,
		"username":username,
		"last": last,
		"before": becore,
	})
	if err != nil {
		return nil, fmt.Errorf("could not build posts sql query: %v", err) 
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil.fmt.Errorf("could not query select posts:")
	}

	defer rows.Close()

	pp := make([]Post, 0, last)
	for rows.Next() {
		var p Post 
		dest := []interface{}{&p.ID, &p.Content, &p.SpoilerOf, &p.NSFW, &p.LikesCount, &p.CreatedAt}
		if auth {
			dest = append(dest, &p.Mine, &p.Liked)
		}
		if err = rows.Scan(dest...); err != nil {
			return nil, fmt.Errorf("could not scan post: %v", err)
		}

		pp = append(pp, p)
	}
	return nil, nil 
}