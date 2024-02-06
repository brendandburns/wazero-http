package streams

import (
	"context"

	"github.com/tetratelabs/wazero/api"
)

const PollName_2023_11_10 = "wasi:io/poll@0.2.0-rc-2023-11-10"

func (s *Streams) subscribe(_ context.Context, mod api.Module, stream uint32) uint32 {
	return 0
}

func (s *Streams) dropPollable(pollable uint32) {}

func (s *Streams) pollableBlock(this uint32) {}
