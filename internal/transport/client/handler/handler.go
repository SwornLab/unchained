package handler

import (
	"context"
)

type Handler interface {
	Challenge(message []byte) []byte
	CorrectnessReport(ctx context.Context, message []byte)
	EventLog(ctx context.Context, message []byte)
	PriceReport(ctx context.Context, message []byte)
	RpcRequest(ctx context.Context, message []byte)
	RpcResponse(ctx context.Context, message []byte)
}
