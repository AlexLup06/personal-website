package ui

import (
	"context"

	"github.com/alexlup06/personal-website/internal/http/kit/httpctx"
	"github.com/alexlup06/personal-website/internal/http/kit/render"
)

func Static(ctx context.Context, name string) string {
	v, ok := httpctx.Assets(ctx)
	if !ok {
		return "/static/" + name
	}
	a, _ := v.(render.Assets)
	return a.Static(name)
}
