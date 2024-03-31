package jaegergrpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
)

// MDCarrier 自定义 carrier，并实现接口
type MDCarrier struct {
	metadata.MD
}

func (m MDCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, strs := range m.MD {
		for _, v := range strs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m MDCarrier) Set(key, val string) {
	m.MD[key] = append(m.MD[key], val)
}

// TraceSpanClientInterceptor returns a grpc.UnaryServerInterceptor suitable
// for use in a grpc.Dial() call.
//
// For example:
//
//     conn, err := grpc.Dial(
//         address,
//         ...,  // (existing DialOptions)
//         grpc.WithUnaryInterceptor(rpc.TraceSpanClientInterceptor()),
//     )
//
// It writes current trace span to request metadata.
func TraceSpanClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, request, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		tracer := opentracing.GlobalTracer()
		var parentCtx opentracing.SpanContext
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
		}
		span := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC Client"},
			ext.SpanKindRPCClient,
		)

		defer span.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		err := tracer.Inject(
			span.Context(),
			opentracing.TextMap,
			MDCarrier{md}, // 自定义 carrier
		)

		if err != nil {
			log.Printf("inject span error :%v", err.Error())
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, request, reply, cc, opts...)
		if err != nil {
			log.Printf("call error : %v", err.Error())
		}
		return err
	}
}

// TraceSpanServerInterceptor returns a grpc.UnaryServerInterceptor suitable
// for use in a grpc.NewServer call.
//
// For example:
//
//     s := grpc.NewServer(
//         ...,  // (existing ServerOptions)
//         grpc.UnaryInterceptor(rpc.TraceSpanServerInterceptor()),
//     )
//
// It reads current trace span from request metadata.
func TraceSpanServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		tracer := opentracing.GlobalTracer()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(
			opentracing.TextMap,
			MDCarrier{md},
		)

		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC Server"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()
			ctx = opentracing.ContextWithSpan(ctx, span)
		}
		return handler(ctx, req)
	}
}
