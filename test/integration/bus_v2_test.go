package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/bus/v2/client"
	"github.com/minchao/go-ptx/bus/v2/client/city_bus"
)

var (
	busV2Client *client.MOTCTransportAPIV2
)

func init() {
	busV2Client = client.New(getTransport(), nil)
}

func TestBusV2_DataVersion(t *testing.T) {
	params := city_bus.NewCityBusAPIDataVersionParams().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIDataVersion(params)
	require.NoError(t, err)
	assert.NotNil(t, result.Payload.VersionID)
}

func TestBusV2_Route(t *testing.T) {
	params := city_bus.NewCityBusAPIRouteParams().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIRoute(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("CityBus.CityBusAPIRoute return no data")
	}
}

func TestBusV2_Station(t *testing.T) {
	params := city_bus.NewCityBusAPIStationParams().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIStation(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("CityBus.CityBusAPIStation return no data")
	}
}

func TestBusV2_Vehicle(t *testing.T) {
	params := city_bus.NewCityBusAPIVehicleParams().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIVehicle(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("CityBus.CityBusAPIVehicle return no data")
	}
}
