package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/rail/v3/client"
	"github.com/minchao/go-ptx/rail/v3/client/t_r_a"
)

var (
	railV3Client *client.MOTCTransportAPIV3
)

func init() {
	railV3Client = client.New(getTransport(), nil)
}

func TestRailV3_Network(t *testing.T) {
	params := t_r_a.NewNetworkAPIControllerGetParams().
		WithDollarFormat("JSON")
	result, _, err := railV3Client.Tra.NetworkAPIControllerGet(params)
	require.NoError(t, err)
	if len(result.Payload.Networks) == 0 {
		t.Fatal("Tra.NetworkAPIControllerGet return no data")
	}
}

func TestRailV3_TrainType(t *testing.T) {
	params := t_r_a.NewTrainTypeAPIControllerGetParams().
		WithDollarFormat("JSON")
	result, _, err := railV3Client.Tra.TrainTypeAPIControllerGet(params)
	require.NoError(t, err)
	if len(result.Payload.TrainTypes) == 0 {
		t.Fatal("Tra.TrainTypeAPIControllerGet return no data")
	}
}

func TestRailV3_Shape(t *testing.T) {
	params := t_r_a.NewShapeAPIControllerGetParams().
		WithDollarFormat("JSON")
	result, _, err := railV3Client.Tra.ShapeAPIControllerGet(params)
	require.NoError(t, err)
	if len(result.Payload.Shapes) == 0 {
		t.Fatal("Tra.ShapeAPIControllerGet return no data")
	}
}
