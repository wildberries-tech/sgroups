package pkg

import (
	"testing"

	"github.com/wildberries-tech/sgroups/v2/pkg/api/sgroups"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestSwaggerUtil(t *testing.T) {
	var u SwaggerUtil[sgroups.SecGroupServiceServer]
	s, e := u.GetSpec()
	require.NoError(t, e)
	require.NotNil(t, s)
}

func TestClosableClient(t *testing.T) {
	var conn *grpc.ClientConn
	var c ClosableClient[sgroups.SecGroupServiceClient]
	err := c.Init(conn)
	require.NoError(t, err)
	require.NotNil(t, c.C)
}
