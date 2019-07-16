package integration

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/client"

	tp "github.com/minchao/go-ptx/transport"
)

var (
	transport *client.Runtime
)

func getTransport() *client.Runtime {
	if transport == nil {
		httpClient := http.DefaultClient
		httpClient.Transport = &tp.AuthTransport{
			AppId:  os.Getenv("APP_ID"),
			AppKey: os.Getenv("APP_KEY"),
		}
		transport = tp.NewWithClient(httpClient)
	}
	return transport
}
