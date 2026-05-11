package api

import (
	"context"
)

type Converter interface {
	buildVandA(context.Context, StreamWithCodec, StreamWithCodec) (string, error)
	buildOnlyV(context.Context, StreamWithCodec) (string, error)
	buildOnlyA(context.Context, StreamWithCodec) (string, error)
}
