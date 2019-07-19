package auth

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

func setup() (serverURL *url.URL, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	u, _ := url.Parse(server.URL)
	return u, mux, server.Close
}

func authHandlerFunc(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.Header.Get("x-date")
		require.NotEmpty(t, date)
		auth := r.Header.Get("Authorization")
		require.Equal(t, authorization(appId, signature(appKey, date)), auth)
	}
}

func newTestClientOperation() *runtime.ClientOperation {
	return &runtime.ClientOperation{
		Method:      "GET",
		PathPattern: "/",
		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
			return nil
		}),
		Reader: runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
			return nil, nil
		}),
	}
}

func TestAuthTransport(t *testing.T) {
	serverURL, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", authHandlerFunc(t))

	httpClient := http.DefaultClient
	httpClient.Transport = NewTransport(appId, appKey)
	apiclient := client.NewWithClient(serverURL.Host, serverURL.Path, []string{"http"}, httpClient)
	_, err := apiclient.Submit(newTestClientOperation())
	require.NoError(t, err)
}

func TestAuthTransport_transport(t *testing.T) {
	// default transport
	tp := &Transport{}
	require.Equal(t, http.DefaultTransport, tp.transport())

	// custom transport
	tp = &Transport{
		Transport: &http.Transport{},
	}
	require.NotEqual(t, http.DefaultTransport, tp.transport())
}
