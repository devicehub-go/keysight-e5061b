package keysighte5061b_test

import (
	"testing"
	"time"

	keysighte5061b "github.com/devicehub-go/keysight-e5061b"
	"github.com/devicehub-go/keysight-e5061b/protocol"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestGetMeasuremnt(t *testing.T) {
	vna := keysighte5061b.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.4.138",
			Port:         5025,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		},
	})
	if err := vna.Connect(); err != nil {
		panic(err)
	}
	defer vna.Disconnect()

	vna.SetWindowLayout("D12_34")
	for i := range(4) {
		vna.SetSweepPoints(i+1, 1001)
		vna.SetContinuousState(i+1, true)
	}

	vna.SetTraceToWindow(protocol.TraceOptions{
		Channel: 1,
		Trace: 1,
		Parameter: "S11",
		Format: "MLOG",
		CenterFrequency: 476e6,
		SpanFrequency: 3e6,
	});
	vna.SetTraceToWindow(protocol.TraceOptions{
		Channel: 2,
		Trace: 1,
		Parameter: "S11",
		Format: "PHAS",
		CenterFrequency: 476e6,
		SpanFrequency: 3e6,
	});

	vna.SetTraceToWindow(protocol.TraceOptions{
		Channel: 3,
		Trace: 1,
		Parameter: "S11",
		Format: "MLOG",
		CenterFrequency: 476e6,
		SpanFrequency: 3e6,
	});
	vna.SetTraceToWindow(protocol.TraceOptions{
		Channel: 4,
		Trace: 1,
		Parameter: "S11",
		Format: "PHAS",
		CenterFrequency: 476e6,
		SpanFrequency: 3e6,
	});


	/*start := time.Now()
	_, err := vna.GetComplexData(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))

	sweepTime, e := vna.Query(":SENS1:SWE:TIME?")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(sweepTime))*/
}
