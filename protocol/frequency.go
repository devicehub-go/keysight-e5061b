package protocol

import (
	"fmt"
)

// Sets the start frequency in Hz
func (e *E5061B) SetStartFrequency(channel int, frequency float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateFrequency(frequency); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:FREQ:STAR %g", channel, frequency)
	return e.Write(cmd)
}

// Sets the stop frequency in Hz
func (e *E5061B) SetStopFrequency(channel int, frequency float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateFrequency(frequency); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:FREQ:STOP %g", channel, frequency)
	return e.Write(cmd)
}

// Sets the number of points for the measurement
func (e *E5061B) SetSweepPoints(channel, points int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateSweepPoints(points); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:SWE:POIN %d", channel, points)
	return e.Write(cmd)
}

// Sets the sweep mode
func (e *E5061B) SetSweepType(channel int, sweep string) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateSweepType(sweep); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:SWE:TYPE %s", channel, sweep)
	return e.Write(cmd)
}

// Sets the IF bandwidth in Hz
func (e *E5061B) SetBandwidth(channel, bandwidth int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	} else if err := e.validateBandwidth(bandwidth); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:BWID %d", channel, bandwidth)
	return e.Write(cmd)
}
