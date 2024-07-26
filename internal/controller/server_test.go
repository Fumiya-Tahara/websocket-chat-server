package controller

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	testBaseURL = "http://localhost:8080"
)

func TestStartServer(t *testing.T) {
	t.Parallel()

	r := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := NewMockHandlerInterface(ctrl)

	handler.EXPECT().Connection(gomock.Any()).Times(1)

	go StartServer(r, handler)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, testBaseURL+"/connection", nil)
	require.NoError(t, err, "could not create request")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "could not send request")

	defer res.Body.Close()
}
