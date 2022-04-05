package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/minchao/go-ptx/basic/v2/client"
	"github.com/minchao/go-ptx/basic/v2/client/basic"
)

var (
	basicV2Client *client.MOTCTransportAPIV2
)

func init() {
	basicV2Client = client.New(getTransport(), nil)
}

func TestBikeV2_Authority(t *testing.T) {
	params := basic.NewBasicAPIAuthority2160Params().
		WithDollarFormat("JSON")
	result, err := basicV2Client.Basic.BasicAPIAuthority2160(params, nil)
	require.NoError(t, err)
	if len(result.Payload) == 0 {
		t.Fatal("Basic.BasicAPIAuthority return no data")
	}
}
