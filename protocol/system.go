package protocol

import (
	"fmt"
	"strings"
)

// Resets the instrument to its default state
func (e *E5061B) Reset() error {
	return e.Write("*RST")
}

// sets the instrument mode (NA or ZA)
func (e *E5061B) SetAnalysisMode(mode string) error {
	if err := e.validateAnalysisMode(mode); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SYST:MODE %s", strings.ToUpper(mode))
	return e.Write(cmd)
}

// Sets the screen split configuration
func (e *E5061B) SetWindowLayout(layout string) error {
	if err := e.validateWindowLayout(layout); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":DISP:SPL %s", strings.ToUpper(layout))
	return e.Write(cmd)
}

// Fits the trace data to the display window.
func (e *E5061B) AutoScale(channel, trace int) error {
	if err := e.SelectTrace(channel, trace); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":CALC%d:SEL:RES", channel)
	return e.Write(cmd)
}

// Sets the trigger source (INT, EXT, MAN, BUS)
func (e *E5061B) SetTriggerSource(source string) error {
	if err := e.validateTriggerSource(source); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":TRIG:SOUR %s", strings.ToUpper(source))
	return e.Write(cmd)
}

// Enables or disables continuous sweep for the channel
func (e *E5061B) SetContinuousState(channel int, enabled bool) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	state := "0"
	if enabled {
		state = "1"
	}

	cmd := fmt.Sprintf(":INIT%d:CONT %s", channel, state)
	return e.Write(cmd)
}

// Performs a single sweep and waits for completion
func (e *E5061B) TriggerSweep(channel int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}
	cmd := fmt.Sprintf(":INIT%d:IMM", channel)
	if err := e.Write(cmd); err != nil {
		return err
	}
	_, err := e.Query("*OPC?")
	return err
}
