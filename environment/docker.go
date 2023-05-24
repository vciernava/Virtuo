package environment

import (
	"context"
	"emperror.dev/errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"sync"
)

var (
	_conce  sync.Once
	_client *client.Client
)

func Docker() (*client.Client, error) {
	var err error
	_conce.Do(func() {
		_client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	})
	return _client, errors.Wrap(err, "environment/docker: could not create client")
}

type ImagePullRequest struct {
	Image string `json:"image"`
}

func PullImage(c *gin.Context) {
	var req ImagePullRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set options for pulling the image
	options := types.ImagePullOptions{}

	// Pull the image
	response, err := _client.ImagePull(context.Background(), req.Image, options)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to pull image: %s", err.Error()))
		return
	}
	defer response.Close()

	// Read and print the pull output
	c.Writer.WriteString("Pulling image...\n")
	_, err = io.Copy(c.Writer, response)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read image pull response: %s", err.Error()))
		return
	}

	c.Writer.WriteString("Image pulled successfully.\n")
}
