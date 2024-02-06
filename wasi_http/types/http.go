package types

import (
	"context"
	"fmt"

	"github.com/brendandburns/wazero-http/wasi_http/common"
	"github.com/brendandburns/wazero-http/wasi_http/streams"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

const (
	ModuleName            = "types"
	ModuleName_2023_10_18 = "wasi:http/types@0.2.0-rc-2023-10-18"
	ModuleName_2023_11_10 = "wasi:http/types@0.2.0-rc-2023-11-10"
)

func logFn(ctx context.Context, mod api.Module, ptr, len uint32) {
	str, _ := common.ReadString(mod, ptr, len)
	fmt.Print(str)
}

func Instantiate_v1(ctx context.Context, rt wazero.Runtime, s *streams.Streams, r *Requests, rs *Responses, f *FieldsCollection, o *OutResponses) error {
	_, err := rt.NewHostModuleBuilder(ModuleName).
		NewFunctionBuilder().WithFunc(r.newOutgoingRequestFn).Export("new-outgoing-request").
		NewFunctionBuilder().WithFunc(f.newFieldsFn).Export("[constructor]fields").
		NewFunctionBuilder().WithFunc(f.dropFieldsFn).Export("drop-fields").
		NewFunctionBuilder().WithFunc(f.fieldsEntriesFn).Export("fields-entries").
		NewFunctionBuilder().WithFunc(r.dropOutgoingRequestFn).Export("drop-outgoing-request").
		NewFunctionBuilder().WithFunc(r.outgoingRequestWriteFn).Export("outgoing-request-write").
		NewFunctionBuilder().WithFunc(rs.dropIncomingResponseFn).Export("drop-incoming-response").
		NewFunctionBuilder().WithFunc(rs.incomingResponseStatusFn).Export("incoming-response-status").
		NewFunctionBuilder().WithFunc(rs.incomingResponseHeadersFn).Export("incoming-response-headers").
		NewFunctionBuilder().WithFunc(rs.incomingResponseConsumeFn).Export("incoming-response-consume").
		NewFunctionBuilder().WithFunc(futureResponseGetFn).Export("future-incoming-response-get").
		NewFunctionBuilder().WithFunc(r.incomingRequestMethodFn).Export("[method]incoming-request.method").
		NewFunctionBuilder().WithFunc(r.incomingRequestPathFn).Export("[method]incoming-request.path-with-query").
		NewFunctionBuilder().WithFunc(r.incomingRequestAuthorityFn).Export("[method]incoming-request.authority").
		NewFunctionBuilder().WithFunc(r.incomingRequestHeadersFn).Export("[method]incoming-request.headers").
		NewFunctionBuilder().WithFunc(incomingRequestConsumeFn).Export("[method]incoming-request.consume").
		NewFunctionBuilder().WithFunc(r.dropIncomingRequestFn).Export("drop-incoming-request").
		NewFunctionBuilder().WithFunc(o.setResponseOutparamFn).Export("[static]response-outparam.set").
		NewFunctionBuilder().WithFunc(rs.newOutgoingResponseFn).Export("[constructor]outgoing-response").
		NewFunctionBuilder().WithFunc(rs.outgoingResponseWriteFn).Export("[method]outgoing-response.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyWriteFn).Export("[method]outgoing-body.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyFinishFn).Export("[static]outgoing-body.finish").
		NewFunctionBuilder().WithFunc(dropOutgoingResponseFn).Export("drop-outgoing-response").
		NewFunctionBuilder().WithFunc(logFn).Export("log-it").
		Instantiate(ctx)
	return err
}

func Instantiate_2023_10_18(ctx context.Context, rt wazero.Runtime, s *streams.Streams, r *Requests, rs *Responses, f *FieldsCollection, o *OutResponses) error {
	_, err := rt.NewHostModuleBuilder(ModuleName_2023_10_18).
		NewFunctionBuilder().WithFunc(r.newOutgoingRequestFn_2023_10_18).Export("[constructor]outgoing-request").
		NewFunctionBuilder().WithFunc(f.newFieldsFn).Export("[constructor]fields").
		NewFunctionBuilder().WithFunc(f.dropFieldsFn).Export("[resource-drop]fields").
		NewFunctionBuilder().WithFunc(f.fieldsEntriesFn).Export("[method]fields.entries").
		NewFunctionBuilder().WithFunc(r.dropOutgoingRequestFn).Export("[resource-drop]outgoing-request").
		NewFunctionBuilder().WithFunc(r.outgoingRequestWriteFn).Export("[method]outgoing-request.write").
		NewFunctionBuilder().WithFunc(rs.dropIncomingResponseFn).Export("[resource-drop]incoming-response").
		NewFunctionBuilder().WithFunc(rs.incomingResponseStatusFn).Export("[method]incoming-response.status").
		NewFunctionBuilder().WithFunc(rs.incomingResponseHeadersFn).Export("[method]incoming-response.headers").
		NewFunctionBuilder().WithFunc(rs.incomingResponseConsumeFn).Export("[method]incoming-response.consume").
		NewFunctionBuilder().WithFunc(rs.incomingResponseConsumeFn).Export("[method]incoming-body.stream").
		NewFunctionBuilder().WithFunc(futureResponseGetFn_2023_10_18).Export("[method]future-incoming-response.get").
		NewFunctionBuilder().WithFunc(r.incomingRequestMethodFn).Export("[method]incoming-request.method").
		NewFunctionBuilder().WithFunc(r.incomingRequestPathFn_2023_10_18).Export("[method]incoming-request.path-with-query").
		NewFunctionBuilder().WithFunc(r.incomingRequestAuthorityFn_2023_10_18).Export("[method]incoming-request.authority").
		NewFunctionBuilder().WithFunc(r.incomingRequestHeadersFn).Export("[method]incoming-request.headers").
		NewFunctionBuilder().WithFunc(incomingRequestConsumeFn).Export("[method]incoming-request.consume").
		NewFunctionBuilder().WithFunc(r.dropIncomingRequestFn).Export("[resource-drop]incoming-request").
		NewFunctionBuilder().WithFunc(o.setResponseOutparamFn_2023_10_18).Export("[static]response-outparam.set").
		NewFunctionBuilder().WithFunc(rs.newOutgoingResponseFn).Export("[constructor]outgoing-response").
		NewFunctionBuilder().WithFunc(rs.outgoingResponseWriteFn_2023_10_18).Export("[method]outgoing-response.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyWriteFn).Export("[method]outgoing-body.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyFinishFn).Export("[static]outgoing-body.finish").
		NewFunctionBuilder().WithFunc(dropOutgoingResponseFn).Export("[resource-drop]outgoing-response").
		NewFunctionBuilder().WithFunc(logFn).Export("log-it").
		Instantiate(ctx)
	return err
}

func Instantiate_2023_11_10(ctx context.Context, rt wazero.Runtime, s *streams.Streams, r *Requests, rs *Responses, f *FieldsCollection, o *OutResponses) error {
	_, err := rt.NewHostModuleBuilder(ModuleName_2023_11_10).
		NewFunctionBuilder().WithFunc(r.newOutgoingRequestFn_2023_11_10).Export("[constructor]outgoing-request").
		NewFunctionBuilder().WithFunc(f.newFieldsFn).Export("[constructor]fields").
		NewFunctionBuilder().WithFunc(f.dropFieldsFn).Export("[resource-drop]fields").
		NewFunctionBuilder().WithFunc(f.fieldsEntriesFn).Export("[method]fields.entries").
		NewFunctionBuilder().WithFunc(r.dropOutgoingRequestFn).Export("[resource-drop]outgoing-request").
		NewFunctionBuilder().WithFunc(r.outgoingRequestWriteFn).Export("[method]outgoing-request.write").
		NewFunctionBuilder().WithFunc(rs.dropIncomingResponseFn).Export("[resource-drop]incoming-response").
		NewFunctionBuilder().WithFunc(rs.incomingResponseStatusFn).Export("[method]incoming-response.status").
		NewFunctionBuilder().WithFunc(rs.incomingResponseHeadersFn).Export("[method]incoming-response.headers").
		NewFunctionBuilder().WithFunc(rs.incomingResponseConsumeFn).Export("[method]incoming-response.consume").
		NewFunctionBuilder().WithFunc(rs.incomingResponseConsumeFn).Export("[method]incoming-body.stream").
		NewFunctionBuilder().WithFunc(rs.incomingResponseSubscribe).Export("[method]future-incoming-response.subscribe").
		NewFunctionBuilder().WithFunc(futureResponseGetFn_2023_11_10).Export("[method]future-incoming-response.get").
		NewFunctionBuilder().WithFunc(r.incomingRequestMethodFn).Export("[method]incoming-request.method").
		NewFunctionBuilder().WithFunc(r.incomingRequestPathFn_2023_10_18).Export("[method]incoming-request.path-with-query").
		NewFunctionBuilder().WithFunc(r.incomingRequestAuthorityFn_2023_10_18).Export("[method]incoming-request.authority").
		NewFunctionBuilder().WithFunc(r.incomingRequestHeadersFn).Export("[method]incoming-request.headers").
		NewFunctionBuilder().WithFunc(incomingRequestConsumeFn).Export("[method]incoming-request.consume").
		NewFunctionBuilder().WithFunc(r.dropIncomingRequestFn).Export("[resource-drop]incoming-request").
		NewFunctionBuilder().WithFunc(o.setResponseOutparamFn_2023_11_10).Export("[static]response-outparam.set").
		NewFunctionBuilder().WithFunc(rs.newOutgoingResponseFn_2023_11_10).Export("[constructor]outgoing-response").
		NewFunctionBuilder().WithFunc(rs.outgoingResponseWriteFn_2023_10_18).Export("[method]outgoing-response.body").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyWriteFn).Export("[method]outgoing-body.write").
		NewFunctionBuilder().WithFunc(rs.outgoingBodyFinishFn).Export("[static]outgoing-body.finish").
		NewFunctionBuilder().WithFunc(dropOutgoingResponseFn).Export("[resource-drop]outgoing-response").
		NewFunctionBuilder().WithFunc(logFn).Export("log-it").
		NewFunctionBuilder().WithFunc(dropFutureIncomingResponse).Export("[resource-drop]future-incoming-response").
		NewFunctionBuilder().WithFunc(r.fields.fieldsFromList).Export("[static]fields.from-list").
		NewFunctionBuilder().WithFunc(r.outgoingRequestBody).Export("[method]outgoing-request.body").
		NewFunctionBuilder().WithFunc(r.outgoingRequestSetMethod).Export("[method]outgoing-request.set-method").
		NewFunctionBuilder().WithFunc(r.outgoingRequestSetPathWithQuery).Export("[method]outgoing-request.set-path-with-query").
		NewFunctionBuilder().WithFunc(r.outgoingRequestSetAuthority).Export("[method]outgoing-request.set-authority").
		NewFunctionBuilder().WithFunc(r.outgoingRequestSetScheme).Export("[method]outgoing-request.set-scheme").
		NewFunctionBuilder().WithFunc(rs.outgoingResponseSetStatus_2023_11_10).Export("[method]outgoing-response.set-status-code").
		Instantiate(ctx)
	return err
}
