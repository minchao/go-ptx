package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/bus/v3/client"
	"github.com/minchao/go-ptx/bus/v3/client/city_bus"
)

var (
	busV23lient *client.MOTCTransportAPIV3
)

func init() {
	busV23lient = client.New(getTransport(), nil)
}

func TestBusV3_RouteNetwork(t *testing.T) {
	params := city_bus.NewCityBusAPIRouteNetworkParams().
		WithDollarFormat("JSON").
		WithCity("Tainan")
	result, _, err := busV23lient.CityBus.CityBusAPIRouteNetwork(params)
	require.NoError(t, err)
	if len(result.Payload.RouteNetworks) == 0 {
		t.Fatal("CityBus.CityBusAPIRouteNetwork return no data")
	}
}
