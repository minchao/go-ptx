package transport

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
)

const (
	appId  = "APP_ID"
	appKey = "APP_KEY"
)

func TestAuthTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		date := r.Header.Get("x-date")
		require.NotEmpty(t, date)
		auth := r.Header.Get("Authorization")
		require.Equal(t, authorization(appId, signature(appKey, date)), auth)
	}))
	defer server.Close()

	httpClient := http.DefaultClient
	httpClient.Transport = &AuthTransport{AppId: appId, AppKey: appKey}
	u, _ := url.Parse(server.URL)
	apiclient := client.NewWithClient(u.Host, "/", []string{"http"}, httpClient)
	_, err := apiclient.Submit(&runtime.ClientOperation{
		Method:      "GET",
		PathPattern: "/",
		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
			return nil
		}),
		Reader: runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
			return nil, nil
		}),
	})
	require.NoError(t, err)
}

func TestAuthTransport_transport(t *testing.T) {
	// default transport
	tp := &AuthTransport{}
	require.Equal(t, http.DefaultTransport, tp.transport())

	// custom transport
	tp = &AuthTransport{
		Transport: &http.Transport{},
	}
	require.NotEqual(t, http.DefaultTransport, tp.transport())
}
