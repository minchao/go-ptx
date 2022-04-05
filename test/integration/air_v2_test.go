package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/air/v2/client"
	"github.com/minchao/go-ptx/air/v2/client/air"
)

var (
	airV2Client *client.MOTCTransportAPIV2
)

func init() {
	airV2Client = client.New(getTransport(), nil)
}

func TestAirV2_Airport(t *testing.T) {
	params := air.NewAirAPIAirport2010Params().
		WithDollarFormat("JSON")
	result, _, err := airV2Client.Air.AirAPIAirport2010(params, nil)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Air.AirAPIAirport return no data")
	}
}
