package service

import (
	"context"
)

type Converter interface {
	BuildVandA(context.Context, StreamWithCodec, StreamWithCodec) (string, error)
	BuildOnlyV(context.Context, StreamWithCodec) (string, error)
	BuildOnlyA(context.Context, StreamWithCodec) (string, error)
}
