package protocol

import (
	"fmt"
	"strings"
)

// Sets the number of active traces on a channel
func (e *E5061B) SetTraceCount(channel, count int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}
	if count < 1 || count > 4 {
		return fmt.Errorf("trace count must be between 1 and 4, got %d", count)
	}

	cmd := fmt.Sprintf(":CALC%d:PAR:COUN %d", channel, count)
	return e.Write(cmd)
}

// Defines the measurement parameter for a specific trace
func (e *E5061B) SetTraceParameter(channel, trace int, param string) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateTrace(trace); err != nil {
		return err
	} else if err := e.validateTraceParameter(param); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":CALC%d:PAR%d:DEF %s", channel, trace, strings.ToUpper(param))
	return e.Write(cmd)
}

// Activates a specific trace for subsequent formatting or scaling
func (e *E5061B) SelectTrace(channel, trace int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateTrace(trace); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":CALC%d:PAR%d:SEL", channel, trace)
	return e.Write(cmd)
}

// Sets the display format for the specified trace
func (e *E5061B) SetTraceFormat(channel, trace int, format string) error {
	if err := e.SelectTrace(channel, trace); err != nil {
		return err
	}
	if err := e.validateDisplayFormat(format); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":CALC%d:SEL:FORM %s", channel, strings.ToUpper(format))
	return e.Write(cmd)
}
