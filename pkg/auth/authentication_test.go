package auth

import (
	"testing"

	"github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/require"
)

func TestAuthentication_AuthenticateRequest(t *testing.T) {
	serverURL, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", authHandlerFunc(t))

	apiclient := client.New(serverURL.Host, serverURL.Path, []string{"http"})
	apiclient.DefaultAuthentication = NewAuthentication(appID, appKey)
	_, err := apiclient.Submit(newTestClientOperation())
	require.NoError(t, err)
}
