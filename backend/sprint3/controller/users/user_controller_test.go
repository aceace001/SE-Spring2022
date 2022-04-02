package users

import (
	"testing"
	// "SE-Spring2022/backend/sprint3"
)

// func TestPostHomePage(t *testing.T) {
// 	testData := []struct {
// 		Title  string `json:"title"  binding:"required"`
// 		Author string `json:"author" binding:"required"`
// 	}{
// 		// #predefined parameters
// 		{"Abc", "a"},
// 		{"Aaa", "bbbb"},
// 		{"Aaa", "bbbb"},
// 		{"123", "xyz"},
// 		{"123", "xyz"},
// 		{"This is a sample post", xyz},
// 	}

// 	for _, dat := range testData {
// 		t.Run(dat.title, func(t *testing.T) {
// 			post, err := m.NewPostWithTitle(dat.title)
// 			if dat.error {
// 				if err == nil {
// 					t.Errorf("Expected error Got nil for post: %s", post.Title)
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("Unexpected error: %s for post: %s", err, post.Title)
// 				}
// 			}
// 		})
// 	}
// }

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}