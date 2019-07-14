package transport

import (
	"net/http"

	"github.com/go-openapi/runtime/client"
)

const (
	DefaultHost     string = "ptx.transportdata.tw"
	DefaultBasePath string = "/MOTC/"
)

func NewWithClient(c *http.Client) *client.Runtime {
	return client.NewWithClient(DefaultHost, DefaultBasePath, []string{"https"}, c)
}
