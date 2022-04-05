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
	params := t_r_a.NewNetworkAPIControllerGet3200Params().
		WithDollarFormat("JSON")
	result, _, err := railV3Client.Tra.NetworkAPIControllerGet3200(params, nil)
	require.NoError(t, err)
	if len(result.Payload.Networks) == 0 {
		t.Fatal("Tra.NetworkAPIControllerGet return no data")
	}
}

func TestRailV3_TrainType(t *testing.T) {
	params := t_r_a.NewTrainTypeAPIControllerGet3206Params().
		WithDollarFormat("JSON")
	result, _, err := railV3Client.Tra.TrainTypeAPIControllerGet3206(params, nil)
	require.NoError(t, err)
	if len(result.Payload.TrainTypes) == 0 {
		t.Fatal("Tra.TrainTypeAPIControllerGet return no data")
	}
}

func TestRailV3_Shape(t *testing.T) {
	params := t_r_a.NewShapeAPIControllerGet3219Params().
		WithDollarFormat("JSON")
	result, _, err := railV3Client.Tra.ShapeAPIControllerGet3219(params, nil)
	require.NoError(t, err)
	if len(result.Payload.Shapes) == 0 {
		t.Fatal("Tra.ShapeAPIControllerGet return no data")
	}
}
