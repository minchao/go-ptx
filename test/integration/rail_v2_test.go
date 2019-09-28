package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/rail/v2/client"
	"github.com/minchao/go-ptx/rail/v2/client/metro"
	"github.com/minchao/go-ptx/rail/v2/client/t_h_s_r"
	"github.com/minchao/go-ptx/rail/v2/client/t_r_a"
)

var (
	railV2Client *client.MOTCTransportAPIV2
)

func init() {
	railV2Client = client.New(getTransport(), nil)
}

func TestRailV2_Metro_Network(t *testing.T) {
	params := metro.NewMetroAPINetworkParams().
		WithDollarFormat("JSON").
		WithOperator("TRTC")
	result, err := railV2Client.Metro.MetroAPINetwork(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Metro.MetroAPINetwork return no data")
	}
}

func TestRailV2_THSR_Station(t *testing.T) {
	params := t_h_s_r.NewTHSRAPIStationParams().
		WithDollarFormat("JSON")
	result, err := railV2Client.Thsr.THSRAPIStation(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Thsr.THSRAPIStation return no data")
	}
}

func TestRailV2_TRA_Network(t *testing.T) {
	params := t_r_a.NewTRAAPINetworkParams().
		WithDollarFormat("JSON")
	result, err := railV2Client.Tra.TRAAPINetwork(params)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Tra.TRAAPINetwork return no data")
	}
}
