// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: sgroups/api.proto

package sgroupsconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	sgroups "github.com/wildberries-tech/sgroups/v2/pkg/api/sgroups"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SecGroupServiceName is the fully-qualified name of the SecGroupService service.
	SecGroupServiceName = "wb.v2.sgroups.SecGroupService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SecGroupServiceSyncProcedure is the fully-qualified name of the SecGroupService's Sync RPC.
	SecGroupServiceSyncProcedure = "/wb.v2.sgroups.SecGroupService/Sync"
	// SecGroupServiceSyncStatusProcedure is the fully-qualified name of the SecGroupService's
	// SyncStatus RPC.
	SecGroupServiceSyncStatusProcedure = "/wb.v2.sgroups.SecGroupService/SyncStatus"
	// SecGroupServiceSyncStatusesProcedure is the fully-qualified name of the SecGroupService's
	// SyncStatuses RPC.
	SecGroupServiceSyncStatusesProcedure = "/wb.v2.sgroups.SecGroupService/SyncStatuses"
	// SecGroupServiceListNetworksProcedure is the fully-qualified name of the SecGroupService's
	// ListNetworks RPC.
	SecGroupServiceListNetworksProcedure = "/wb.v2.sgroups.SecGroupService/ListNetworks"
	// SecGroupServiceListSecurityGroupsProcedure is the fully-qualified name of the SecGroupService's
	// ListSecurityGroups RPC.
	SecGroupServiceListSecurityGroupsProcedure = "/wb.v2.sgroups.SecGroupService/ListSecurityGroups"
	// SecGroupServiceGetSgSubnetsProcedure is the fully-qualified name of the SecGroupService's
	// GetSgSubnets RPC.
	SecGroupServiceGetSgSubnetsProcedure = "/wb.v2.sgroups.SecGroupService/GetSgSubnets"
	// SecGroupServiceFindSgSgRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindSgSgRules RPC.
	SecGroupServiceFindSgSgRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindSgSgRules"
	// SecGroupServiceFindFqdnRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindFqdnRules RPC.
	SecGroupServiceFindFqdnRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindFqdnRules"
	// SecGroupServiceFindSgIcmpRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindSgIcmpRules RPC.
	SecGroupServiceFindSgIcmpRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindSgIcmpRules"
	// SecGroupServiceFindSgSgIcmpRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindSgSgIcmpRules RPC.
	SecGroupServiceFindSgSgIcmpRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindSgSgIcmpRules"
	// SecGroupServiceFindIECidrSgRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindIECidrSgRules RPC.
	SecGroupServiceFindIECidrSgRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindIECidrSgRules"
	// SecGroupServiceFindIESgSgRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindIESgSgRules RPC.
	SecGroupServiceFindIESgSgRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindIESgSgRules"
	// SecGroupServiceFindIESgSgIcmpRulesProcedure is the fully-qualified name of the SecGroupService's
	// FindIESgSgIcmpRules RPC.
	SecGroupServiceFindIESgSgIcmpRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindIESgSgIcmpRules"
	// SecGroupServiceFindIECidrSgIcmpRulesProcedure is the fully-qualified name of the
	// SecGroupService's FindIECidrSgIcmpRules RPC.
	SecGroupServiceFindIECidrSgIcmpRulesProcedure = "/wb.v2.sgroups.SecGroupService/FindIECidrSgIcmpRules"
	// SecGroupServiceGetSecGroupForAddressProcedure is the fully-qualified name of the
	// SecGroupService's GetSecGroupForAddress RPC.
	SecGroupServiceGetSecGroupForAddressProcedure = "/wb.v2.sgroups.SecGroupService/GetSecGroupForAddress"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	secGroupServiceServiceDescriptor                     = sgroups.File_sgroups_api_proto.Services().ByName("SecGroupService")
	secGroupServiceSyncMethodDescriptor                  = secGroupServiceServiceDescriptor.Methods().ByName("Sync")
	secGroupServiceSyncStatusMethodDescriptor            = secGroupServiceServiceDescriptor.Methods().ByName("SyncStatus")
	secGroupServiceSyncStatusesMethodDescriptor          = secGroupServiceServiceDescriptor.Methods().ByName("SyncStatuses")
	secGroupServiceListNetworksMethodDescriptor          = secGroupServiceServiceDescriptor.Methods().ByName("ListNetworks")
	secGroupServiceListSecurityGroupsMethodDescriptor    = secGroupServiceServiceDescriptor.Methods().ByName("ListSecurityGroups")
	secGroupServiceGetSgSubnetsMethodDescriptor          = secGroupServiceServiceDescriptor.Methods().ByName("GetSgSubnets")
	secGroupServiceFindSgSgRulesMethodDescriptor         = secGroupServiceServiceDescriptor.Methods().ByName("FindSgSgRules")
	secGroupServiceFindFqdnRulesMethodDescriptor         = secGroupServiceServiceDescriptor.Methods().ByName("FindFqdnRules")
	secGroupServiceFindSgIcmpRulesMethodDescriptor       = secGroupServiceServiceDescriptor.Methods().ByName("FindSgIcmpRules")
	secGroupServiceFindSgSgIcmpRulesMethodDescriptor     = secGroupServiceServiceDescriptor.Methods().ByName("FindSgSgIcmpRules")
	secGroupServiceFindIECidrSgRulesMethodDescriptor     = secGroupServiceServiceDescriptor.Methods().ByName("FindIECidrSgRules")
	secGroupServiceFindIESgSgRulesMethodDescriptor       = secGroupServiceServiceDescriptor.Methods().ByName("FindIESgSgRules")
	secGroupServiceFindIESgSgIcmpRulesMethodDescriptor   = secGroupServiceServiceDescriptor.Methods().ByName("FindIESgSgIcmpRules")
	secGroupServiceFindIECidrSgIcmpRulesMethodDescriptor = secGroupServiceServiceDescriptor.Methods().ByName("FindIECidrSgIcmpRules")
	secGroupServiceGetSecGroupForAddressMethodDescriptor = secGroupServiceServiceDescriptor.Methods().ByName("GetSecGroupForAddress")
)

// SecGroupServiceClient is a client for the wb.v2.sgroups.SecGroupService service.
type SecGroupServiceClient interface {
	Sync(context.Context, *connect.Request[sgroups.SyncReq]) (*connect.Response[emptypb.Empty], error)
	SyncStatus(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[sgroups.SyncStatusResp], error)
	SyncStatuses(context.Context, *connect.Request[emptypb.Empty]) (*connect.ServerStreamForClient[sgroups.SyncStatusResp], error)
	ListNetworks(context.Context, *connect.Request[sgroups.ListNetworksReq]) (*connect.Response[sgroups.ListNetworksResp], error)
	ListSecurityGroups(context.Context, *connect.Request[sgroups.ListSecurityGroupsReq]) (*connect.Response[sgroups.ListSecurityGroupsResp], error)
	GetSgSubnets(context.Context, *connect.Request[sgroups.GetSgSubnetsReq]) (*connect.Response[sgroups.GetSgSubnetsResp], error)
	FindSgSgRules(context.Context, *connect.Request[sgroups.FindSgSgRulesReq]) (*connect.Response[sgroups.SgSgRulesResp], error)
	FindFqdnRules(context.Context, *connect.Request[sgroups.FindFqdnRulesReq]) (*connect.Response[sgroups.FqdnRulesResp], error)
	FindSgIcmpRules(context.Context, *connect.Request[sgroups.FindSgIcmpRulesReq]) (*connect.Response[sgroups.SgIcmpRulesResp], error)
	FindSgSgIcmpRules(context.Context, *connect.Request[sgroups.FindSgSgIcmpRulesReq]) (*connect.Response[sgroups.SgSgIcmpRulesResp], error)
	FindIECidrSgRules(context.Context, *connect.Request[sgroups.FindIECidrSgRulesReq]) (*connect.Response[sgroups.IECidrSgRulesResp], error)
	FindIESgSgRules(context.Context, *connect.Request[sgroups.FindIESgSgRulesReq]) (*connect.Response[sgroups.IESgSgRulesResp], error)
	FindIESgSgIcmpRules(context.Context, *connect.Request[sgroups.FindIESgSgIcmpRulesReq]) (*connect.Response[sgroups.IESgSgIcmpRulesResp], error)
	FindIECidrSgIcmpRules(context.Context, *connect.Request[sgroups.FindIECidrSgIcmpRulesReq]) (*connect.Response[sgroups.IECidrSgIcmpRulesResp], error)
	GetSecGroupForAddress(context.Context, *connect.Request[sgroups.GetSecGroupForAddressReq]) (*connect.Response[sgroups.SecGroup], error)
}

// NewSecGroupServiceClient constructs a client for the wb.v2.sgroups.SecGroupService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSecGroupServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SecGroupServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &secGroupServiceClient{
		sync: connect.NewClient[sgroups.SyncReq, emptypb.Empty](
			httpClient,
			baseURL+SecGroupServiceSyncProcedure,
			connect.WithSchema(secGroupServiceSyncMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		syncStatus: connect.NewClient[emptypb.Empty, sgroups.SyncStatusResp](
			httpClient,
			baseURL+SecGroupServiceSyncStatusProcedure,
			connect.WithSchema(secGroupServiceSyncStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		syncStatuses: connect.NewClient[emptypb.Empty, sgroups.SyncStatusResp](
			httpClient,
			baseURL+SecGroupServiceSyncStatusesProcedure,
			connect.WithSchema(secGroupServiceSyncStatusesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listNetworks: connect.NewClient[sgroups.ListNetworksReq, sgroups.ListNetworksResp](
			httpClient,
			baseURL+SecGroupServiceListNetworksProcedure,
			connect.WithSchema(secGroupServiceListNetworksMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSecurityGroups: connect.NewClient[sgroups.ListSecurityGroupsReq, sgroups.ListSecurityGroupsResp](
			httpClient,
			baseURL+SecGroupServiceListSecurityGroupsProcedure,
			connect.WithSchema(secGroupServiceListSecurityGroupsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSgSubnets: connect.NewClient[sgroups.GetSgSubnetsReq, sgroups.GetSgSubnetsResp](
			httpClient,
			baseURL+SecGroupServiceGetSgSubnetsProcedure,
			connect.WithSchema(secGroupServiceGetSgSubnetsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findSgSgRules: connect.NewClient[sgroups.FindSgSgRulesReq, sgroups.SgSgRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindSgSgRulesProcedure,
			connect.WithSchema(secGroupServiceFindSgSgRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findFqdnRules: connect.NewClient[sgroups.FindFqdnRulesReq, sgroups.FqdnRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindFqdnRulesProcedure,
			connect.WithSchema(secGroupServiceFindFqdnRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findSgIcmpRules: connect.NewClient[sgroups.FindSgIcmpRulesReq, sgroups.SgIcmpRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindSgIcmpRulesProcedure,
			connect.WithSchema(secGroupServiceFindSgIcmpRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findSgSgIcmpRules: connect.NewClient[sgroups.FindSgSgIcmpRulesReq, sgroups.SgSgIcmpRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindSgSgIcmpRulesProcedure,
			connect.WithSchema(secGroupServiceFindSgSgIcmpRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findIECidrSgRules: connect.NewClient[sgroups.FindIECidrSgRulesReq, sgroups.IECidrSgRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindIECidrSgRulesProcedure,
			connect.WithSchema(secGroupServiceFindIECidrSgRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findIESgSgRules: connect.NewClient[sgroups.FindIESgSgRulesReq, sgroups.IESgSgRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindIESgSgRulesProcedure,
			connect.WithSchema(secGroupServiceFindIESgSgRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findIESgSgIcmpRules: connect.NewClient[sgroups.FindIESgSgIcmpRulesReq, sgroups.IESgSgIcmpRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindIESgSgIcmpRulesProcedure,
			connect.WithSchema(secGroupServiceFindIESgSgIcmpRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		findIECidrSgIcmpRules: connect.NewClient[sgroups.FindIECidrSgIcmpRulesReq, sgroups.IECidrSgIcmpRulesResp](
			httpClient,
			baseURL+SecGroupServiceFindIECidrSgIcmpRulesProcedure,
			connect.WithSchema(secGroupServiceFindIECidrSgIcmpRulesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSecGroupForAddress: connect.NewClient[sgroups.GetSecGroupForAddressReq, sgroups.SecGroup](
			httpClient,
			baseURL+SecGroupServiceGetSecGroupForAddressProcedure,
			connect.WithSchema(secGroupServiceGetSecGroupForAddressMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// secGroupServiceClient implements SecGroupServiceClient.
type secGroupServiceClient struct {
	sync                  *connect.Client[sgroups.SyncReq, emptypb.Empty]
	syncStatus            *connect.Client[emptypb.Empty, sgroups.SyncStatusResp]
	syncStatuses          *connect.Client[emptypb.Empty, sgroups.SyncStatusResp]
	listNetworks          *connect.Client[sgroups.ListNetworksReq, sgroups.ListNetworksResp]
	listSecurityGroups    *connect.Client[sgroups.ListSecurityGroupsReq, sgroups.ListSecurityGroupsResp]
	getSgSubnets          *connect.Client[sgroups.GetSgSubnetsReq, sgroups.GetSgSubnetsResp]
	findSgSgRules         *connect.Client[sgroups.FindSgSgRulesReq, sgroups.SgSgRulesResp]
	findFqdnRules         *connect.Client[sgroups.FindFqdnRulesReq, sgroups.FqdnRulesResp]
	findSgIcmpRules       *connect.Client[sgroups.FindSgIcmpRulesReq, sgroups.SgIcmpRulesResp]
	findSgSgIcmpRules     *connect.Client[sgroups.FindSgSgIcmpRulesReq, sgroups.SgSgIcmpRulesResp]
	findIECidrSgRules     *connect.Client[sgroups.FindIECidrSgRulesReq, sgroups.IECidrSgRulesResp]
	findIESgSgRules       *connect.Client[sgroups.FindIESgSgRulesReq, sgroups.IESgSgRulesResp]
	findIESgSgIcmpRules   *connect.Client[sgroups.FindIESgSgIcmpRulesReq, sgroups.IESgSgIcmpRulesResp]
	findIECidrSgIcmpRules *connect.Client[sgroups.FindIECidrSgIcmpRulesReq, sgroups.IECidrSgIcmpRulesResp]
	getSecGroupForAddress *connect.Client[sgroups.GetSecGroupForAddressReq, sgroups.SecGroup]
}

// Sync calls wb.v2.sgroups.SecGroupService.Sync.
func (c *secGroupServiceClient) Sync(ctx context.Context, req *connect.Request[sgroups.SyncReq]) (*connect.Response[emptypb.Empty], error) {
	return c.sync.CallUnary(ctx, req)
}

// SyncStatus calls wb.v2.sgroups.SecGroupService.SyncStatus.
func (c *secGroupServiceClient) SyncStatus(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[sgroups.SyncStatusResp], error) {
	return c.syncStatus.CallUnary(ctx, req)
}

// SyncStatuses calls wb.v2.sgroups.SecGroupService.SyncStatuses.
func (c *secGroupServiceClient) SyncStatuses(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.ServerStreamForClient[sgroups.SyncStatusResp], error) {
	return c.syncStatuses.CallServerStream(ctx, req)
}

// ListNetworks calls wb.v2.sgroups.SecGroupService.ListNetworks.
func (c *secGroupServiceClient) ListNetworks(ctx context.Context, req *connect.Request[sgroups.ListNetworksReq]) (*connect.Response[sgroups.ListNetworksResp], error) {
	return c.listNetworks.CallUnary(ctx, req)
}

// ListSecurityGroups calls wb.v2.sgroups.SecGroupService.ListSecurityGroups.
func (c *secGroupServiceClient) ListSecurityGroups(ctx context.Context, req *connect.Request[sgroups.ListSecurityGroupsReq]) (*connect.Response[sgroups.ListSecurityGroupsResp], error) {
	return c.listSecurityGroups.CallUnary(ctx, req)
}

// GetSgSubnets calls wb.v2.sgroups.SecGroupService.GetSgSubnets.
func (c *secGroupServiceClient) GetSgSubnets(ctx context.Context, req *connect.Request[sgroups.GetSgSubnetsReq]) (*connect.Response[sgroups.GetSgSubnetsResp], error) {
	return c.getSgSubnets.CallUnary(ctx, req)
}

// FindSgSgRules calls wb.v2.sgroups.SecGroupService.FindSgSgRules.
func (c *secGroupServiceClient) FindSgSgRules(ctx context.Context, req *connect.Request[sgroups.FindSgSgRulesReq]) (*connect.Response[sgroups.SgSgRulesResp], error) {
	return c.findSgSgRules.CallUnary(ctx, req)
}

// FindFqdnRules calls wb.v2.sgroups.SecGroupService.FindFqdnRules.
func (c *secGroupServiceClient) FindFqdnRules(ctx context.Context, req *connect.Request[sgroups.FindFqdnRulesReq]) (*connect.Response[sgroups.FqdnRulesResp], error) {
	return c.findFqdnRules.CallUnary(ctx, req)
}

// FindSgIcmpRules calls wb.v2.sgroups.SecGroupService.FindSgIcmpRules.
func (c *secGroupServiceClient) FindSgIcmpRules(ctx context.Context, req *connect.Request[sgroups.FindSgIcmpRulesReq]) (*connect.Response[sgroups.SgIcmpRulesResp], error) {
	return c.findSgIcmpRules.CallUnary(ctx, req)
}

// FindSgSgIcmpRules calls wb.v2.sgroups.SecGroupService.FindSgSgIcmpRules.
func (c *secGroupServiceClient) FindSgSgIcmpRules(ctx context.Context, req *connect.Request[sgroups.FindSgSgIcmpRulesReq]) (*connect.Response[sgroups.SgSgIcmpRulesResp], error) {
	return c.findSgSgIcmpRules.CallUnary(ctx, req)
}

// FindIECidrSgRules calls wb.v2.sgroups.SecGroupService.FindIECidrSgRules.
func (c *secGroupServiceClient) FindIECidrSgRules(ctx context.Context, req *connect.Request[sgroups.FindIECidrSgRulesReq]) (*connect.Response[sgroups.IECidrSgRulesResp], error) {
	return c.findIECidrSgRules.CallUnary(ctx, req)
}

// FindIESgSgRules calls wb.v2.sgroups.SecGroupService.FindIESgSgRules.
func (c *secGroupServiceClient) FindIESgSgRules(ctx context.Context, req *connect.Request[sgroups.FindIESgSgRulesReq]) (*connect.Response[sgroups.IESgSgRulesResp], error) {
	return c.findIESgSgRules.CallUnary(ctx, req)
}

// FindIESgSgIcmpRules calls wb.v2.sgroups.SecGroupService.FindIESgSgIcmpRules.
func (c *secGroupServiceClient) FindIESgSgIcmpRules(ctx context.Context, req *connect.Request[sgroups.FindIESgSgIcmpRulesReq]) (*connect.Response[sgroups.IESgSgIcmpRulesResp], error) {
	return c.findIESgSgIcmpRules.CallUnary(ctx, req)
}

// FindIECidrSgIcmpRules calls wb.v2.sgroups.SecGroupService.FindIECidrSgIcmpRules.
func (c *secGroupServiceClient) FindIECidrSgIcmpRules(ctx context.Context, req *connect.Request[sgroups.FindIECidrSgIcmpRulesReq]) (*connect.Response[sgroups.IECidrSgIcmpRulesResp], error) {
	return c.findIECidrSgIcmpRules.CallUnary(ctx, req)
}

// GetSecGroupForAddress calls wb.v2.sgroups.SecGroupService.GetSecGroupForAddress.
func (c *secGroupServiceClient) GetSecGroupForAddress(ctx context.Context, req *connect.Request[sgroups.GetSecGroupForAddressReq]) (*connect.Response[sgroups.SecGroup], error) {
	return c.getSecGroupForAddress.CallUnary(ctx, req)
}

// SecGroupServiceHandler is an implementation of the wb.v2.sgroups.SecGroupService service.
type SecGroupServiceHandler interface {
	Sync(context.Context, *connect.Request[sgroups.SyncReq]) (*connect.Response[emptypb.Empty], error)
	SyncStatus(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[sgroups.SyncStatusResp], error)
	SyncStatuses(context.Context, *connect.Request[emptypb.Empty], *connect.ServerStream[sgroups.SyncStatusResp]) error
	ListNetworks(context.Context, *connect.Request[sgroups.ListNetworksReq]) (*connect.Response[sgroups.ListNetworksResp], error)
	ListSecurityGroups(context.Context, *connect.Request[sgroups.ListSecurityGroupsReq]) (*connect.Response[sgroups.ListSecurityGroupsResp], error)
	GetSgSubnets(context.Context, *connect.Request[sgroups.GetSgSubnetsReq]) (*connect.Response[sgroups.GetSgSubnetsResp], error)
	FindSgSgRules(context.Context, *connect.Request[sgroups.FindSgSgRulesReq]) (*connect.Response[sgroups.SgSgRulesResp], error)
	FindFqdnRules(context.Context, *connect.Request[sgroups.FindFqdnRulesReq]) (*connect.Response[sgroups.FqdnRulesResp], error)
	FindSgIcmpRules(context.Context, *connect.Request[sgroups.FindSgIcmpRulesReq]) (*connect.Response[sgroups.SgIcmpRulesResp], error)
	FindSgSgIcmpRules(context.Context, *connect.Request[sgroups.FindSgSgIcmpRulesReq]) (*connect.Response[sgroups.SgSgIcmpRulesResp], error)
	FindIECidrSgRules(context.Context, *connect.Request[sgroups.FindIECidrSgRulesReq]) (*connect.Response[sgroups.IECidrSgRulesResp], error)
	FindIESgSgRules(context.Context, *connect.Request[sgroups.FindIESgSgRulesReq]) (*connect.Response[sgroups.IESgSgRulesResp], error)
	FindIESgSgIcmpRules(context.Context, *connect.Request[sgroups.FindIESgSgIcmpRulesReq]) (*connect.Response[sgroups.IESgSgIcmpRulesResp], error)
	FindIECidrSgIcmpRules(context.Context, *connect.Request[sgroups.FindIECidrSgIcmpRulesReq]) (*connect.Response[sgroups.IECidrSgIcmpRulesResp], error)
	GetSecGroupForAddress(context.Context, *connect.Request[sgroups.GetSecGroupForAddressReq]) (*connect.Response[sgroups.SecGroup], error)
}

// NewSecGroupServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSecGroupServiceHandler(svc SecGroupServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	secGroupServiceSyncHandler := connect.NewUnaryHandler(
		SecGroupServiceSyncProcedure,
		svc.Sync,
		connect.WithSchema(secGroupServiceSyncMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceSyncStatusHandler := connect.NewUnaryHandler(
		SecGroupServiceSyncStatusProcedure,
		svc.SyncStatus,
		connect.WithSchema(secGroupServiceSyncStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceSyncStatusesHandler := connect.NewServerStreamHandler(
		SecGroupServiceSyncStatusesProcedure,
		svc.SyncStatuses,
		connect.WithSchema(secGroupServiceSyncStatusesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceListNetworksHandler := connect.NewUnaryHandler(
		SecGroupServiceListNetworksProcedure,
		svc.ListNetworks,
		connect.WithSchema(secGroupServiceListNetworksMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceListSecurityGroupsHandler := connect.NewUnaryHandler(
		SecGroupServiceListSecurityGroupsProcedure,
		svc.ListSecurityGroups,
		connect.WithSchema(secGroupServiceListSecurityGroupsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceGetSgSubnetsHandler := connect.NewUnaryHandler(
		SecGroupServiceGetSgSubnetsProcedure,
		svc.GetSgSubnets,
		connect.WithSchema(secGroupServiceGetSgSubnetsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindSgSgRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindSgSgRulesProcedure,
		svc.FindSgSgRules,
		connect.WithSchema(secGroupServiceFindSgSgRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindFqdnRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindFqdnRulesProcedure,
		svc.FindFqdnRules,
		connect.WithSchema(secGroupServiceFindFqdnRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindSgIcmpRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindSgIcmpRulesProcedure,
		svc.FindSgIcmpRules,
		connect.WithSchema(secGroupServiceFindSgIcmpRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindSgSgIcmpRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindSgSgIcmpRulesProcedure,
		svc.FindSgSgIcmpRules,
		connect.WithSchema(secGroupServiceFindSgSgIcmpRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindIECidrSgRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindIECidrSgRulesProcedure,
		svc.FindIECidrSgRules,
		connect.WithSchema(secGroupServiceFindIECidrSgRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindIESgSgRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindIESgSgRulesProcedure,
		svc.FindIESgSgRules,
		connect.WithSchema(secGroupServiceFindIESgSgRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindIESgSgIcmpRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindIESgSgIcmpRulesProcedure,
		svc.FindIESgSgIcmpRules,
		connect.WithSchema(secGroupServiceFindIESgSgIcmpRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceFindIECidrSgIcmpRulesHandler := connect.NewUnaryHandler(
		SecGroupServiceFindIECidrSgIcmpRulesProcedure,
		svc.FindIECidrSgIcmpRules,
		connect.WithSchema(secGroupServiceFindIECidrSgIcmpRulesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secGroupServiceGetSecGroupForAddressHandler := connect.NewUnaryHandler(
		SecGroupServiceGetSecGroupForAddressProcedure,
		svc.GetSecGroupForAddress,
		connect.WithSchema(secGroupServiceGetSecGroupForAddressMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/wb.v2.sgroups.SecGroupService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SecGroupServiceSyncProcedure:
			secGroupServiceSyncHandler.ServeHTTP(w, r)
		case SecGroupServiceSyncStatusProcedure:
			secGroupServiceSyncStatusHandler.ServeHTTP(w, r)
		case SecGroupServiceSyncStatusesProcedure:
			secGroupServiceSyncStatusesHandler.ServeHTTP(w, r)
		case SecGroupServiceListNetworksProcedure:
			secGroupServiceListNetworksHandler.ServeHTTP(w, r)
		case SecGroupServiceListSecurityGroupsProcedure:
			secGroupServiceListSecurityGroupsHandler.ServeHTTP(w, r)
		case SecGroupServiceGetSgSubnetsProcedure:
			secGroupServiceGetSgSubnetsHandler.ServeHTTP(w, r)
		case SecGroupServiceFindSgSgRulesProcedure:
			secGroupServiceFindSgSgRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindFqdnRulesProcedure:
			secGroupServiceFindFqdnRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindSgIcmpRulesProcedure:
			secGroupServiceFindSgIcmpRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindSgSgIcmpRulesProcedure:
			secGroupServiceFindSgSgIcmpRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindIECidrSgRulesProcedure:
			secGroupServiceFindIECidrSgRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindIESgSgRulesProcedure:
			secGroupServiceFindIESgSgRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindIESgSgIcmpRulesProcedure:
			secGroupServiceFindIESgSgIcmpRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceFindIECidrSgIcmpRulesProcedure:
			secGroupServiceFindIECidrSgIcmpRulesHandler.ServeHTTP(w, r)
		case SecGroupServiceGetSecGroupForAddressProcedure:
			secGroupServiceGetSecGroupForAddressHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSecGroupServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSecGroupServiceHandler struct{}

func (UnimplementedSecGroupServiceHandler) Sync(context.Context, *connect.Request[sgroups.SyncReq]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.Sync is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) SyncStatus(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[sgroups.SyncStatusResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.SyncStatus is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) SyncStatuses(context.Context, *connect.Request[emptypb.Empty], *connect.ServerStream[sgroups.SyncStatusResp]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.SyncStatuses is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) ListNetworks(context.Context, *connect.Request[sgroups.ListNetworksReq]) (*connect.Response[sgroups.ListNetworksResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.ListNetworks is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) ListSecurityGroups(context.Context, *connect.Request[sgroups.ListSecurityGroupsReq]) (*connect.Response[sgroups.ListSecurityGroupsResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.ListSecurityGroups is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) GetSgSubnets(context.Context, *connect.Request[sgroups.GetSgSubnetsReq]) (*connect.Response[sgroups.GetSgSubnetsResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.GetSgSubnets is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindSgSgRules(context.Context, *connect.Request[sgroups.FindSgSgRulesReq]) (*connect.Response[sgroups.SgSgRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindSgSgRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindFqdnRules(context.Context, *connect.Request[sgroups.FindFqdnRulesReq]) (*connect.Response[sgroups.FqdnRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindFqdnRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindSgIcmpRules(context.Context, *connect.Request[sgroups.FindSgIcmpRulesReq]) (*connect.Response[sgroups.SgIcmpRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindSgIcmpRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindSgSgIcmpRules(context.Context, *connect.Request[sgroups.FindSgSgIcmpRulesReq]) (*connect.Response[sgroups.SgSgIcmpRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindSgSgIcmpRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindIECidrSgRules(context.Context, *connect.Request[sgroups.FindIECidrSgRulesReq]) (*connect.Response[sgroups.IECidrSgRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindIECidrSgRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindIESgSgRules(context.Context, *connect.Request[sgroups.FindIESgSgRulesReq]) (*connect.Response[sgroups.IESgSgRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindIESgSgRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindIESgSgIcmpRules(context.Context, *connect.Request[sgroups.FindIESgSgIcmpRulesReq]) (*connect.Response[sgroups.IESgSgIcmpRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindIESgSgIcmpRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) FindIECidrSgIcmpRules(context.Context, *connect.Request[sgroups.FindIECidrSgIcmpRulesReq]) (*connect.Response[sgroups.IECidrSgIcmpRulesResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.FindIECidrSgIcmpRules is not implemented"))
}

func (UnimplementedSecGroupServiceHandler) GetSecGroupForAddress(context.Context, *connect.Request[sgroups.GetSecGroupForAddressReq]) (*connect.Response[sgroups.SecGroup], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.v2.sgroups.SecGroupService.GetSecGroupForAddress is not implemented"))
}