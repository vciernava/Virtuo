package environment

import (
	"sync"

	"emperror.dev/errors"

	"github.com/docker/docker/client"
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
