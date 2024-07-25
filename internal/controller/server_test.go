package controller

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	BaseURL = "http://localhost:8080"
)

func TestStartServer(t *testing.T) {
	r := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := NewMockHandlerInterface(ctrl)

	handler.EXPECT().Connection(gomock.Any()).Times(1)

	go StartServer(r, handler)

	req, _ := http.NewRequest("GET", BaseURL+"/connection", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "could not send request")
	defer res.Body.Close()
}
