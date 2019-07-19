package integration

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/client"
	"github.com/minchao/go-ptx/pkg/auth"
	tp "github.com/minchao/go-ptx/pkg/transport"
)

var (
	transport *client.Runtime
)

func getTransport() *client.Runtime {
	if transport == nil {
		httpClient := http.DefaultClient
		httpClient.Transport = auth.NewTransport(os.Getenv("APP_ID"), os.Getenv("APP_KEY"))
		transport = tp.NewWithClient(httpClient)
	}
	return transport
}
