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
	params := city_bus.NewCityBusAPIDataVersion2033Params().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIDataVersion2033(params, nil)
	require.NoError(t, err)
	assert.NotNil(t, result.Payload.VersionID)
}

func TestBusV2_Route(t *testing.T) {
	params := city_bus.NewCityBusAPIRoute2035Params().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIRoute2035(params, nil)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("CityBus.CityBusAPIRoute return no data")
	}
}

func TestBusV2_Station(t *testing.T) {
	params := city_bus.NewCityBusAPIStation2037Params().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIStation2037(params, nil)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("CityBus.CityBusAPIStation return no data")
	}
}

func TestBusV2_Vehicle(t *testing.T) {
	params := city_bus.NewCityBusAPIVehicle2041Params().
		WithDollarFormat("JSON").
		WithCity("Taipei")
	result, _, err := busV2Client.CityBus.CityBusAPIVehicle2041(params, nil)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("CityBus.CityBusAPIVehicle return no data")
	}
}
