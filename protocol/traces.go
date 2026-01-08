package protocol

import (
	"fmt"
	"strings"
)

type TraceOptions struct {
	Channel int
	Trace int
	Parameter string
	Format string
	StartFrequency float64
	StopFrequency float64
	CenterFrequency float64
	SpanFrequency float64
}

// Assign a channel and trace to specific window
func (e *E5061B) SetTraceToWindow(opt TraceOptions) error {
	if err := e.SetTraceParameter(opt.Channel, opt.Trace, opt.Parameter); err != nil {
		return err
	} else if err := e.SelectTrace(opt.Channel, opt.Trace); err != nil {
		return err
	} else if err := e.SetTraceFormat(opt.Channel, opt.Trace, opt.Format); err != nil {
		return err
	}

	if opt.StartFrequency != 0 && opt.StopFrequency != 0 {
		if err := e.SetFrequencyStart(opt.Channel, opt.StartFrequency); err != nil {
			return err;
		} else if err := e.SetFrequencyStop(opt.Channel, opt.StopFrequency); err != nil {
			return err;
		}
	}

	if opt.CenterFrequency != 0 && opt.SpanFrequency != 0 {
		if err := e.SetFrequencyCenter(opt.Channel, opt.CenterFrequency); err != nil {
			return err;
		} else if err := e.SetFrequencySpan(opt.Channel, opt.SpanFrequency); err != nil {
			return err;
		}
	}

	return e.SetAutoScale(opt.Channel, opt.Trace)
}

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

	cmd := fmt.Sprintf(":CALC%d:FORM %s", channel, strings.ToUpper(format))
	return e.Write(cmd)
}

// Sets the smoothing state for the selected trace
func (e *E5061B) SetSmoothingState(channel int, state bool) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	stateStr := "OFF"
	if state {
		stateStr = "ON"
	}

	cmd := fmt.Sprintf(":CALC%d:SEL:SMO:STAT %s", channel, stateStr)
	return e.Write(cmd)
}

// Sets the smoothing aperture (percentage to the sweep span value)
// for the selected channel
func (e *E5061B) SetSmoothingAperture(channel int, aperture float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}
	if aperture < 0.05 || 25 < aperture {
		return fmt.Errorf("aperture must be between 0.05 and 25, got %f", aperture)
	}
	
	cmd := fmt.Sprintf(":CALC%d:SEL:SMO:APER %f", channel, aperture)
	return e.Write(cmd)
}