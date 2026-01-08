package protocol

import (
	"fmt"
	"strings"
)

// Sets the screen split configuration
func (e *E5061B) SetWindowLayout(layout string) error {
	if err := e.validateWindowLayout(layout); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":DISP:SPL %s", strings.ToUpper(layout))
	return e.Write(cmd)
}

// Maximize a window or trace graph
func (e *E5061B) MaximizeWindow(channel int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":DISP:WIND%d MAX", channel)
	return e.Write(cmd)
}

// Adjust the value of the reference division line and scale
// per division to display the trace appropriately
func (e *E5061B) SetAutoScale(channel, trace int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateTrace(trace); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":DISP:WIND%d:TRAC%d:Y:SCALe:AUTO", channel, trace)
	return e.Write(cmd)
}