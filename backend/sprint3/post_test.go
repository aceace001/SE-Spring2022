package main_test

import (
	"testing"

 m "SE-Spring2022/backend/sprint3"	
)

func TestPost(t *testing.T) {
	testData := []struct {
	Title   string `json:"title"  binding:"required"`
	Author  string `json:"author" binding:"required"`
	}{
	#predefined parameters
	{"Abc", a},
	{"Aaa", bbbb},
	{"Aaa", bbbb},
	{"123", xyz},
	{"123", xyz},
	{"This is a sample post", xyz},
	}

	for _, dat := range testData {
		t.Run(dat.title, func(t *testing.T) {
			post, err := m.NewPostWithTitle(dat.title)
			if dat.error {
				if err == nil {
					t.Errorf("Expected error Got nil for post: %s", post.Title)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %s for post: %s", err, post.Title)
				}
			}
		})
	}
}
