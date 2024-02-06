package streams

import (
	"context"
	"encoding/binary"
	"log"

	"github.com/brendandburns/wazero-http/wasi_http/common"
	"github.com/tetratelabs/wazero/api"
)

func (s *Streams) streamReadFn(ctx context.Context, mod api.Module, stream_handle uint32, length uint64, out_ptr uint32) {
	rawData := make([]byte, length)
	n, done, err := s.Read(stream_handle, rawData)

	//	data, err := types.ResponseBody()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// fmt.Printf("%d %v %s", n, done, string(rawData))

	if n == 0 && done {
		data := []byte{}
		// 0 == is_ok, 1 == is_err
		le := binary.LittleEndian
		// done is an errorr
		data = le.AppendUint32(data, 1)
		// error tag 1 indicates stream done.
		data = le.AppendUint32(data, 1)
		mod.Memory().Write(out_ptr, data)
		return
	}

	data := rawData[0:n]
	ptr_len := uint32(len(data))
	ptr, err := common.Malloc(ctx, mod, ptr_len)
	if err != nil {
		log.Fatalf(err.Error())
	}
	mod.Memory().Write(ptr, data)

	data = []byte{}
	// 0 == is_ok, 1 == is_err
	le := binary.LittleEndian
	data = le.AppendUint32(data, 0)
	data = le.AppendUint32(data, ptr)
	data = le.AppendUint32(data, ptr_len)
	if done {
		// No more data to read.
		data = le.AppendUint32(data, 0)
	} else {
		data = le.AppendUint32(data, 1)
	}
	mod.Memory().Write(out_ptr, data)
}

func (s *Streams) dropInputStreamFn(_ context.Context, mod api.Module, stream uint32) {
	s.DeleteStream(stream)
}
