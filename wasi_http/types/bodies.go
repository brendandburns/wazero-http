package types

import (
	"bytes"
	"context"
	"encoding/binary"

	"github.com/tetratelabs/wazero/api"
)

type Bodies struct {
	Requests  *Requests
	Responses *Responses
}

func (b *Bodies) incomingBodyStreamFn(ctx context.Context, mod api.Module, res, ptr uint32) {
	// For now just copy the stream forward *hack*
	data := []byte{}
	data = binary.LittleEndian.AppendUint32(data, 0)
	data = binary.LittleEndian.AppendUint32(data, res)

	if !mod.Memory().Write(ptr, data) {
		panic("Failed to write data!")
	}
}

func (b *Bodies) outgoingBodyWriteFn(ctx context.Context, mod api.Module, res, ptr uint32) {
	// For now the body is just the request or response. Eventually we may need an actual body struct.
	response, responseFound := b.Responses.GetResponse(res)
	request, requestFound := b.Requests.GetRequest(res)
	data := []byte{}
	if !responseFound && !requestFound {
		// Error
		data = binary.LittleEndian.AppendUint32(data, 1)
		data = binary.LittleEndian.AppendUint32(data, 0)

		if !mod.Memory().Write(ptr, data) {
			panic("Failed to write data!")
		}
		return
	}
	writer := &bytes.Buffer{}
	stream := b.Responses.streams.NewOutputStream(writer)

	if responseFound {
		response.streamHandle = stream
		response.Buffer = writer
	}
	if requestFound {
		// request.streamHandle = stream
		request.BodyBuffer = writer
	}
	// 0 == no error
	data = binary.LittleEndian.AppendUint32(data, 0)
	data = binary.LittleEndian.AppendUint32(data, stream)

	if !mod.Memory().Write(ptr, data) {
		panic("Failed to write data!")
	}
}

func (b *Bodies) outgoingBodyFinishFn(ctx context.Context, mod api.Module, body, res, opt1, ptr uint32) {
	// TODO: lock buffer here.
	data := []byte{}
	data = binary.LittleEndian.AppendUint32(data, 0)

	if !mod.Memory().Write(ptr, data) {
		panic("Failed to write data!")
	}
}
