package keysighte5061b

import (
	"github.com/devicehub-go/keysight-e5061b/protocol"
	"github.com/devicehub-go/unicomm"
)

/*
Creates a new instance of Keysight E5061B Network
Analyzer to communicate and control the device
*/
func New(options unicomm.Options) *protocol.E5061B {
	options.Delimiter = "\n"
	return &protocol.E5061B{
		Communication: unicomm.New(options),
	}
}
