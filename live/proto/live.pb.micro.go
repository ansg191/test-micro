// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/live.proto

package live

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Live service

func NewLiveEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Live service

type LiveService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Live_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Live_PingPongService, error)
}

type liveService struct {
	c    client.Client
	name string
}

func NewLiveService(name string, c client.Client) LiveService {
	return &liveService{
		c:    c,
		name: name,
	}
}

func (c *liveService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Live.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Live_StreamService, error) {
	req := c.c.NewRequest(c.name, "Live.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &liveServiceStream{stream}, nil
}

type Live_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type liveServiceStream struct {
	stream client.Stream
}

func (x *liveServiceStream) Close() error {
	return x.stream.Close()
}

func (x *liveServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *liveServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *liveServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *liveServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *liveService) PingPong(ctx context.Context, opts ...client.CallOption) (Live_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Live.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &liveServicePingPong{stream}, nil
}

type Live_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type liveServicePingPong struct {
	stream client.Stream
}

func (x *liveServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *liveServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *liveServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *liveServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *liveServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *liveServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Live service

type LiveHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Live_StreamStream) error
	PingPong(context.Context, Live_PingPongStream) error
}

func RegisterLiveHandler(s server.Server, hdlr LiveHandler, opts ...server.HandlerOption) error {
	type live interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Live struct {
		live
	}
	h := &liveHandler{hdlr}
	return s.Handle(s.NewHandler(&Live{h}, opts...))
}

type liveHandler struct {
	LiveHandler
}

func (h *liveHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.LiveHandler.Call(ctx, in, out)
}

func (h *liveHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.LiveHandler.Stream(ctx, m, &liveStreamStream{stream})
}

type Live_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type liveStreamStream struct {
	stream server.Stream
}

func (x *liveStreamStream) Close() error {
	return x.stream.Close()
}

func (x *liveStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *liveStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *liveStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *liveStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *liveHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.LiveHandler.PingPong(ctx, &livePingPongStream{stream})
}

type Live_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type livePingPongStream struct {
	stream server.Stream
}

func (x *livePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *livePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *livePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *livePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *livePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *livePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
