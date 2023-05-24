package routes

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/vciernava/Virtuo/environment"
	"net/http"
)

var (
	cli *client.Client
)

func NewInstall() string {
	if c, err := environment.Docker(); err != nil {
		return "Client not found."
	} else {
		cli = c
		return "Client Connected."
	}
}

func GetServers(c *gin.Context) {
	resp, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(jsonResponse)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

type ContainerCreateRequest struct {
	Image              string   `json:"image"`
	ContainerName      string   `json:"containername"`
	Env                []string `json:"env"`
	StartAfterCreation bool     `json:"startaftercreation"`
}

type ContainerCreateResponse struct {
	ContainerID string `json:"container_id"`
}

func CreateServer(c *gin.Context) {
	var req ContainerCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	//portBinding := nat.PortBinding{
	//	HostIP:   "0.0.0.0",
	//	HostPort: "25567",
	//}
	//
	//portBindings := nat.PortMap{
	//	nat.Port("25567/tcp"): []nat.PortBinding{portBinding},
	//}

	mountVolumes := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: "C:/minecraft-server",
			Target: "/data",
		},
	}

	config := &container.Config{
		Image:       req.Image,
		Env:         req.Env,
		AttachStdin: true,
		OpenStdin:   true,
	}
	hostConf := &container.HostConfig{
		Mounts: mountVolumes,
	}

	resp, err := cli.ContainerCreate(context.Background(), config, hostConf, nil, nil, req.ContainerName)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	if req.StartAfterCreation == true {
		err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	}

	response := ContainerCreateResponse{
		ContainerID: resp.ID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusCreated)
	_, err = c.Writer.Write(jsonResponse)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

type ContainerStartRequest struct {
	ContainerID string `json:"container_id"`
}
type ContainerStartResponse struct {
	ContainerID string `json:"container_id"`
}

func StartServer(c *gin.Context) {
	var req ContainerStartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err := cli.ContainerStart(context.Background(), req.ContainerID, types.ContainerStartOptions{})
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	response := ContainerStartResponse{
		ContainerID: req.ContainerID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(jsonResponse)
}

type ContainerStopRequest struct {
	ContainerID string `json:"container_id"`
}
type ContainerStopResponse struct {
	ContainerID string `json:"container_id"`
}

func StopServer(c *gin.Context) {
	var req ContainerStopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err := cli.ContainerStop(context.Background(), req.ContainerID, container.StopOptions{})
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	response := ContainerStopResponse{
		ContainerID: req.ContainerID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(jsonResponse)
}

type ContainerDeleteRequest struct {
	ContainerID string `json:"container_id"`
}
type ContainerDeleteResponse struct {
	ContainerID string `json:"container_id"`
}

func DeleteServer(c *gin.Context) {
	var req ContainerDeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err := cli.ContainerRemove(context.Background(), req.ContainerID, types.ContainerRemoveOptions{})
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	response := ContainerStopResponse{
		ContainerID: req.ContainerID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(jsonResponse)
}
