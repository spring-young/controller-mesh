// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pkg/apis/ctrlmesh/proto/throttling.proto

package protoconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	proto "github.com/KusionStack/controller-mesh/pkg/apis/ctrlmesh/proto"
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
	// ThrottlingName is the fully-qualified name of the Throttling service.
	ThrottlingName = "proto.Throttling"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ThrottlingSendConfigProcedure is the fully-qualified name of the Throttling's SendConfig RPC.
	ThrottlingSendConfigProcedure = "/proto.Throttling/SendConfig"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	throttlingServiceDescriptor          = proto.File_pkg_apis_ctrlmesh_proto_throttling_proto.Services().ByName("Throttling")
	throttlingSendConfigMethodDescriptor = throttlingServiceDescriptor.Methods().ByName("SendConfig")
)

// ThrottlingClient is a client for the proto.Throttling service.
type ThrottlingClient interface {
	SendConfig(context.Context, *connect.Request[proto.CircuitBreaker]) (*connect.Response[proto.ConfigResp], error)
}

// NewThrottlingClient constructs a client for the proto.Throttling service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewThrottlingClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ThrottlingClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &throttlingClient{
		sendConfig: connect.NewClient[proto.CircuitBreaker, proto.ConfigResp](
			httpClient,
			baseURL+ThrottlingSendConfigProcedure,
			connect.WithSchema(throttlingSendConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// throttlingClient implements ThrottlingClient.
type throttlingClient struct {
	sendConfig *connect.Client[proto.CircuitBreaker, proto.ConfigResp]
}

// SendConfig calls proto.Throttling.SendConfig.
func (c *throttlingClient) SendConfig(ctx context.Context, req *connect.Request[proto.CircuitBreaker]) (*connect.Response[proto.ConfigResp], error) {
	return c.sendConfig.CallUnary(ctx, req)
}

// ThrottlingHandler is an implementation of the proto.Throttling service.
type ThrottlingHandler interface {
	SendConfig(context.Context, *connect.Request[proto.CircuitBreaker]) (*connect.Response[proto.ConfigResp], error)
}

// NewThrottlingHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewThrottlingHandler(svc ThrottlingHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	throttlingSendConfigHandler := connect.NewUnaryHandler(
		ThrottlingSendConfigProcedure,
		svc.SendConfig,
		connect.WithSchema(throttlingSendConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.Throttling/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ThrottlingSendConfigProcedure:
			throttlingSendConfigHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedThrottlingHandler returns CodeUnimplemented from all methods.
type UnimplementedThrottlingHandler struct{}

func (UnimplementedThrottlingHandler) SendConfig(context.Context, *connect.Request[proto.CircuitBreaker]) (*connect.Response[proto.ConfigResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.Throttling.SendConfig is not implemented"))
}
