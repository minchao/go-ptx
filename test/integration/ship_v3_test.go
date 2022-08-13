package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/ship/v3/client"
	"github.com/minchao/go-ptx/ship/v3/client/common"
)

var (
	shipV3Client *client.MOTCTransportAPIV3
)

func init() {
	shipV3Client = client.New(getTransport(), nil)
}

func TestShipV3_ShipAPIPort(t *testing.T) {
	params := common.NewShipBasicPort3231Params().
		WithDollarFormat("JSON")
	result, _, err := shipV3Client.Common.ShipBasicPort3231(params, nil)
	require.NoError(t, err)
	if len(result.Payload.Ports) == 0 {
		t.Fatal("Ship.ShipAPIPort return no data")
	}
}
