package main

import (
	"net/http"
	"testing"

	"go-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router = mapUrls()

func Strigify(payload map[string]interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func TestPostHomePage(t *testing.T) {

	payload := gin.H{
		"title": "I am happy now",
		"author":  michael,
	}


  router.POST("/post", users.PostHomePage)
  
	rr, req := utils.HttpTestRequest("POST", "/post", Strigify(payload))

	req.Header.Set("Content-Type", "application/json")//?
	router.ServeHTTP(rr, <-request)

	response := utils.Parse(rr.Body.Bytes())

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "successfuly", response.Status)
	assert.Equal(t, "post request successfuly", response.Message)
}
