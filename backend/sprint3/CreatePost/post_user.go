package post

type PostUser struct {
	ID     uint `json:"id"`
	PostID uint `json:"post_id"`
	UserID uint
}
