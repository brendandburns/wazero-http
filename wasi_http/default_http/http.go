package default_http

import (
	"context"
	"fmt"

	"github.com/brendandburns/wazero-http/wasi_http/types"
	"github.com/tetratelabs/wazero"
)

const (
	ModuleName            = "default-outgoing-HTTP"
	ModuleName_2023_10_18 = "wasi:http/outgoing-handler@0.2.0-rc-2023-10-18"
	ModuleName_2023_11_10 = "wasi:http/outgoing-handler@0.2.0-rc-2023-11-10"
	ModuleName_0_2_0      = "wasi:http/outgoing-handler@0.2.0"
)

func Instantiate(ctx context.Context, r wazero.Runtime, req *types.Requests, res *types.Responses, f *types.FieldsCollection, version string) error {
	handler := &Handler{req, res, f}
	var name string
	switch version {
	case "v1":
		name = ModuleName
	case "2023_10_18":
		name = ModuleName_2023_10_18
	case "2023_11_10":
		name = ModuleName_2023_11_10
	case "v0.2.0":
		name = ModuleName_0_2_0
	default:
		return fmt.Errorf("unknown version: %s", version)
	}
	builder := r.NewHostModuleBuilder(name).
		NewFunctionBuilder().WithFunc(requestFn).Export("request")
	switch version {
	case "v1":
		builder.NewFunctionBuilder().WithFunc(handler.handleFn).Export("handle")
	case "2023_10_18":
		builder.NewFunctionBuilder().WithFunc(handler.handleFn_2023_10_18).Export("handle")
	case "2023_11_10":
		builder.NewFunctionBuilder().WithFunc(handler.handleFn_2023_11_10).Export("handle")
	case "v0.2.0":
		builder.NewFunctionBuilder().WithFunc(handler.handleFn_2023_11_10).Export("handle")
	}
	_, err := builder.Instantiate(ctx)
	return err
}
