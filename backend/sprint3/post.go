package post

type Post struct {
	ID                    uint   `json:"id"`
	Description           string `json:"description"`
	Date                  int64  `json:"date"`
	MarkedAsInappropriate bool   `json:"marked_as_inappropriate"`
	UserEmail             string
	MediaID               uint
}
