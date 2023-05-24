package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func Hello(c *gin.Context) {

	response := HelloResponse{
		Message: "Hello! Everything works fine.",
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
