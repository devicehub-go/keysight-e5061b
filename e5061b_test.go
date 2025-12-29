package keysighte5061b_test

import (
	"fmt"
	"testing"
	"time"

	keysighte5061b "github.com/devicehub-go/keysight-e5061b"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestGetMeasuremnt(t *testing.T) {
	vna := keysighte5061b.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.28.39",
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
	vna.Write("*CLS")
	vna.SetSweepPoints(1, 101)
	vna.SetContinuousState(1, true)

	start := time.Now()
	_, err := vna.GetComplexData(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))
}
