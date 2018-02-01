package domain

import (
	"context"
	"gopkg.in/unrolled/render.v1"
)

type IResource interface {
	Context() context.Context
	Routes() *Routes
	Render() *render.Render
}