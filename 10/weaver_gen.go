// go:build !ignoreWeaverGen

package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)
var _ codegen.LatestVersion = codegen.Version[[0][11]struct{}]("You used 'weaver generate' codegen version 0.11.0, but you built your code with an incompatible weaver module version. Try upgrading 'weaver generate' and re-running it.")

func init() {
	codegen.Register(codegen.Registration{
		Name:        "emojis/Cache",
		Iface:       reflect.TypeOf((*Cache)(nil)).Elem(),
		Impl:        reflect.TypeOf(cache{}),
		Routed:      true,
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return cache_local_stub{impl: impl.(Cache), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return cache_client_stub{stub: stub, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "emojis/Cache", Method: "Get"}), putMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "emojis/Cache", Method: "Put"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return cache_server_stub{impl: impl.(Cache), addLoad: addLoad}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "emojis/ChatGPT",
		Iface: reflect.TypeOf((*ChatGPT)(nil)).Elem(),
		Impl:  reflect.TypeOf(chatgpt{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return chatGPT_local_stub{impl: impl.(ChatGPT), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return chatGPT_client_stub{stub: stub, completeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "emojis/ChatGPT", Method: "Complete"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return chatGPT_server_stub{impl: impl.(ChatGPT), addLoad: addLoad}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/Main",
		Iface: reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:  reflect.TypeOf(app{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return main_client_stub{stub: stub, mainMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/Main", Method: "Main"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		RefData: "⟦f3377ce6:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→emojis/Searcher⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "emojis/Searcher",
		Iface: reflect.TypeOf((*Searcher)(nil)).Elem(),
		Impl:  reflect.TypeOf(searcher{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return searcher_local_stub{impl: impl.(Searcher), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return searcher_client_stub{stub: stub, searchMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "emojis/Searcher", Method: "Search"}), searchChatGPTMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "emojis/Searcher", Method: "SearchChatGPT"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return searcher_server_stub{impl: impl.(Searcher), addLoad: addLoad}
		},
		RefData: "⟦e0ea07ba:wEaVeReDgE:emojis/Searcher→emojis/Cache⟧\n⟦92032889:wEaVeReDgE:emojis/Searcher→emojis/ChatGPT⟧\n",
	})
}

// weaver.Instance checks.
var _ weaver.InstanceOf[Cache] = (*cache)(nil)
var _ weaver.InstanceOf[ChatGPT] = (*chatgpt)(nil)
var _ weaver.InstanceOf[weaver.Main] = (*app)(nil)
var _ weaver.InstanceOf[Searcher] = (*searcher)(nil)

// weaver.Router checks.
var _ weaver.RoutedBy[router] = (*cache)(nil)
var _ weaver.Unrouted = (*chatgpt)(nil)
var _ weaver.Unrouted = (*app)(nil)
var _ weaver.Unrouted = (*searcher)(nil)

// Component "cache", router "router" checks.
var _ func(_ context.Context, query string) string = (&router{}).Get             // routed
var _ func(_ context.Context, query string, _ []string) string = (&router{}).Put // routed

// Local stub implementations.

type cache_local_stub struct {
	impl   Cache
	tracer trace.Tracer
}

// Check that cache_local_stub implements the Cache interface.
var _ Cache = (*cache_local_stub)(nil)

func (s cache_local_stub) Get(ctx context.Context, a0 string) (r0 []string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Cache.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx, a0)
}

func (s cache_local_stub) Put(ctx context.Context, a0 string, a1 []string) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Cache.Put", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Put(ctx, a0, a1)
}

type chatGPT_local_stub struct {
	impl   ChatGPT
	tracer trace.Tracer
}

// Check that chatGPT_local_stub implements the ChatGPT interface.
var _ ChatGPT = (*chatGPT_local_stub)(nil)

func (s chatGPT_local_stub) Complete(ctx context.Context, a0 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.ChatGPT.Complete", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Complete(ctx, a0)
}

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

func (s main_local_stub) Main(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Main.Main", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Main(ctx)
}

type searcher_local_stub struct {
	impl   Searcher
	tracer trace.Tracer
}

// Check that searcher_local_stub implements the Searcher interface.
var _ Searcher = (*searcher_local_stub)(nil)

func (s searcher_local_stub) Search(ctx context.Context, a0 string) (r0 []string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Searcher.Search", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Search(ctx, a0)
}

func (s searcher_local_stub) SearchChatGPT(ctx context.Context, a0 string) (r0 []string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Searcher.SearchChatGPT", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SearchChatGPT(ctx, a0)
}

// Client stub implementations.

type cache_client_stub struct {
	stub       codegen.Stub
	getMetrics *codegen.MethodMetrics
	putMetrics *codegen.MethodMetrics
}

// Check that cache_client_stub implements the Cache interface.
var _ Cache = (*cache_client_stub)(nil)

func (s cache_client_stub) Get(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	start := time.Now()
	s.getMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Cache.Get", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)

	// Set the shardKey.
	var r router
	shardKey := _hashCache(r.Get(ctx, a0))

	// Call the remote method.
	s.getMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

func (s cache_client_stub) Put(ctx context.Context, a0 string, a1 []string) (err error) {
	// Update metrics.
	start := time.Now()
	s.putMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Cache.Put", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.putMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.putMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.String(a0)
	serviceweaver_enc_slice_string_4af10117(enc, a1)

	// Set the shardKey.
	var r router
	shardKey := _hashCache(r.Put(ctx, a0, a1))

	// Call the remote method.
	s.putMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.putMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type chatGPT_client_stub struct {
	stub            codegen.Stub
	completeMetrics *codegen.MethodMetrics
}

// Check that chatGPT_client_stub implements the ChatGPT interface.
var _ ChatGPT = (*chatGPT_client_stub)(nil)

func (s chatGPT_client_stub) Complete(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.completeMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.ChatGPT.Complete", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.completeMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.completeMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.completeMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.completeMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

type main_client_stub struct {
	stub        codegen.Stub
	mainMetrics *codegen.MethodMetrics
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

func (s main_client_stub) Main(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.mainMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Main.Main", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.mainMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.mainMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.mainMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.mainMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type searcher_client_stub struct {
	stub                 codegen.Stub
	searchMetrics        *codegen.MethodMetrics
	searchChatGPTMetrics *codegen.MethodMetrics
}

// Check that searcher_client_stub implements the Searcher interface.
var _ Searcher = (*searcher_client_stub)(nil)

func (s searcher_client_stub) Search(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	start := time.Now()
	s.searchMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Searcher.Search", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.searchMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.searchMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.searchMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.searchMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

func (s searcher_client_stub) SearchChatGPT(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	start := time.Now()
	s.searchChatGPTMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Searcher.SearchChatGPT", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.searchChatGPTMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.searchChatGPTMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.searchChatGPTMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.searchChatGPTMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type cache_server_stub struct {
	impl    Cache
	addLoad func(key uint64, load float64)
}

// Check that cache_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*cache_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s cache_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Get":
		return s.get
	case "Put":
		return s.put
	default:
		return nil
	}
}

func (s cache_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var r router
	s.addLoad(_hashCache(r.Get(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s cache_server_stub) put(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 []string
	a1 = serviceweaver_dec_slice_string_4af10117(dec)
	var r router
	s.addLoad(_hashCache(r.Put(ctx, a0, a1)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Put(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type chatGPT_server_stub struct {
	impl    ChatGPT
	addLoad func(key uint64, load float64)
}

// Check that chatGPT_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*chatGPT_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s chatGPT_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Complete":
		return s.complete
	default:
		return nil
	}
}

func (s chatGPT_server_stub) complete(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Complete(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Main":
		return s.main
	default:
		return nil
	}
}

func (s main_server_stub) main(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Main(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type searcher_server_stub struct {
	impl    Searcher
	addLoad func(key uint64, load float64)
}

// Check that searcher_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*searcher_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s searcher_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Search":
		return s.search
	case "SearchChatGPT":
		return s.searchChatGPT
	default:
		return nil
	}
}

func (s searcher_server_stub) search(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Search(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s searcher_server_stub) searchChatGPT(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.SearchChatGPT(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Router methods.

// _hashCache returns a 64 bit hash of the provided value.
func _hashCache(r string) uint64 {
	var h codegen.Hasher
	h.WriteString(string(r))
	return h.Sum64()
}

// _orderedCodeCache returns an order-preserving serialization of the provided value.
func _orderedCodeCache(r string) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteString(string(r))
	return enc.Encode()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}