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
			Host:         "10.0.9.37",
			Port:         5025,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		},
	})
	if err := vna.Connect(); err != nil {
		panic(err)
	}
	defer vna.Disconnect()

	start := time.Now()
	vna.SetMarkerSearch(1, 2, 1, "MIN")
	vna.SetMarkerTrackingState(1, 2, 1, true)
	valueBytes, err := vna.GetMarkerY(1, 4, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(valueBytes))
	fmt.Println(time.Since(start))
}
