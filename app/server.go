package app

import (
	"fmt"
	"net/http"

	"github.com/agusheryanto182/go-CRUD/models/config"
)

func NewServer(middleware http.Handler, c *config.Global) *http.Server {
	address := fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
	return &http.Server{
		Addr:    address,
		Handler: middleware,
	}
}
