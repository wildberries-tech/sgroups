package sgroups

import (
	"context"

	"github.com/wildberries-tech/sgroups/v2/internal/domains/sgroups"
	registry "github.com/wildberries-tech/sgroups/v2/internal/registry/sgroups"
	"github.com/wildberries-tech/sgroups/v2/pkg/api/common"
	sg "github.com/wildberries-tech/sgroups/v2/pkg/api/sgroups"
)

// ListNetworks impl 'sgroups' service
func (srv *sgService) ListNetworks(ctx context.Context, req *sg.ListNetworksReq) (resp *sg.ListNetworksResp, err error) {
	defer func() {
		err = correctError(err)
	}()
	var reader registry.Reader
	if reader, err = srv.registryReader(ctx); err != nil {
		return resp, err
	}
	defer reader.Close() //lint:nolint
	var scope registry.Scope = registry.NoScope
	if nws := req.GetNetworkNames(); len(nws) > 0 {
		scope = registry.NetworkNames(nws...)
	}
	resp = new(sg.ListNetworksResp)
	err = reader.ListNetworks(ctx, func(nw sgroups.Network) error {
		resp.Networks = append(resp.Networks, &sg.Network{
			Name: nw.Name,
			Network: &common.Networks_NetIP{
				CIDR: nw.Net.String(),
			},
		})
		return nil
	}, scope)
	return resp, err
}
