package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/bike/v2/client"
	"github.com/minchao/go-ptx/bike/v2/client/bike"
)

var (
	bikeV2Client *client.MOTCTransportAPIV2
)

func init() {
	bikeV2Client = client.New(getTransport(), nil)
}

func TestBikeV2_Station(t *testing.T) {
	params := bike.NewBikeAPIStation2180Params().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := bikeV2Client.Bike.BikeAPIStation2180(params, nil)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Bike.BikeAPIStation return no data")
	}
}
