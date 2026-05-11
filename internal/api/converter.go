package api

import (
	"context"
	"io"
)

type Converter interface {
	buildVandA(context.Context, io.ReadCloser, io.ReadCloser)
	buildOnlyV(context.Context, io.ReadCloser)
	buildOnlyA(context.Context, io.ReadCloser)
}
