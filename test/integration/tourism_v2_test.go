package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/tourism/v2/client"
	"github.com/minchao/go-ptx/tourism/v2/client/taiwan_trip_bus"
	"github.com/minchao/go-ptx/tourism/v2/client/tourism"
)

var (
	tourismV2Client *client.MOTCTransportAPIV2
)

func init() {
	tourismV2Client = client.New(getTransport(), nil)
}

func TestTourismV2_ScenicSpot(t *testing.T) {
	params := tourism.NewTourismAPIScenicSpotParams().
		WithDollarFormat("JSON")
	result, err := tourismV2Client.Tourism.TourismAPIScenicSpot(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Tourism.TourismAPIScenicSpot return no data")
	}
}

func TestTourismV2_TaiwanTripBus_Route(t *testing.T) {
	params := taiwan_trip_bus.NewTaiwanTripBusAPIRouteParams().
		WithDollarFormat("JSON")
	result, err := tourismV2Client.TaiwanTripBus.TaiwanTripBusAPIRoute(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("TaiwanTripBus.TaiwanTripBusAPIRoute return no data")
	}
}
