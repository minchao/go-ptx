package transport

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewWithClient(t *testing.T) {
	tp := NewWithClient(nil)
	require.Equal(t, DefaultHost, tp.Host)
	require.Equal(t, DefaultBasePath, tp.BasePath)
}
