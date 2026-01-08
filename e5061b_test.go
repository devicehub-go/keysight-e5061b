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

	vna.Reset()
	vna.SetTraceToWindow(protocol.TraceOptions{
		Channel: 1,
		Trace: 1,
		Parameter: "S11",
		Format: "MLOG",
		CenterFrequency: 476e6,
		SpanFrequency: 3e6,
		SweepPoints: 1001,
		Continuos: true,
		AutoIFBW: false,
		IFBandwidth: 10e3,
		AverageState: true,
		AverageFactor: 5,
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
