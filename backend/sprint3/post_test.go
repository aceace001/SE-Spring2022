package models_test

import (
	"testing"

	
)

func TestPost(t *testing.T) {
	testData := []struct {
		title string
		error bool
	}{
		#predefined parameters
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
