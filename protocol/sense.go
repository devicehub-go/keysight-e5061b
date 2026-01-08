package protocol

import (
	"fmt"
	"strconv"
)

type FrequencyParams struct {
	Start float64
	Stop float64
	Center float64
	Span float64
	Sweep float64
}

// Returns the current frequency parameters
func (e *E5061B) GetFrequencyParameters(channel int) (FrequencyParams, error) {
	var empty FrequencyParams

	start, err := e.Query(fmt.Sprintf(":SENS%d:FREQ:STAR?", channel))
	if err != nil {
		return empty, err
	}
	stop, err := e.Query(fmt.Sprintf(":SENS%d:FREQ:STOP?", channel))
	if err != nil {
		return empty, err
	}
	center, err := e.Query(fmt.Sprintf(":SENS%d:FREQ:CENT?", channel))
	if err != nil {
		return empty, err
	}
	span, err := e.Query(fmt.Sprintf(":SENS%d:FREQ:SPAN?", channel))
	if err != nil {
		return empty, err
	}
	sweep, err := e.Query(fmt.Sprintf(":SENS%d:SWE:POIN?", channel))
	if err != nil {
		return empty, err
	}

	startValue, _ := strconv.ParseFloat(string(start), 64)
	stopValue, _ := strconv.ParseFloat(string(stop), 64)
	centerValue, _ := strconv.ParseFloat(string(center), 64)
	spanValue, _ := strconv.ParseFloat(string(span), 64)
	sweepValue, _ := strconv.ParseFloat(string(sweep), 64)

	return FrequencyParams{
		Start: startValue,
		Stop: stopValue,
		Center: centerValue,
		Span: spanValue,
		Sweep: sweepValue,
	}, nil
}

// Sets the start value of sweep range
func (e *E5061B) SetFrequencyStart(channel int, start float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:FREQ:STAR %g", channel, start)
	return e.Write(cmd)
}

// Sets the stop value of sweep rangel
func (e *E5061B) SetFrequencyStop(channel int, start float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:FREQ:STOP %g", channel, start)
	return e.Write(cmd)
}

// Sets the center value of the sweep range
func (e *E5061B) SetFrequencyCenter(channel int, center float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:FREQ:CENT %g", channel, center)
	return e.Write(cmd)
}

// Sets the span value of the sweep range
func (e *E5061B) SetFrequencySpan(channel int, span float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:FREQ:SPAN %g", channel, span)
	return e.Write(cmd)
}

// Sets the averaging function state
func (e *E5061B) SetAverageState(channel int, state bool) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	stateStr := "OFF"
	if state {
		stateStr = "ON"
	}

	cmd := fmt.Sprintf(":SENS%d:AVER:STAT %s", channel, stateStr)
	return e.Write(cmd)
}

// Sets the averaging factor 
func (e *E5061B) SetAverageFactor(channel, factor int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:AVER:COUN %d", channel, factor)
	return e.Write(cmd)
}

// Resets the average data count to zero
func (e *E5061B) AverageClear(channel int) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:AVER:CLE", channel)
	return e.Write(cmd)
}

// Sets the IF bandwidth 
func (e *E5061B) SetIFBandwidth(channel int, bandwidth float64) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":SENS%d:BAND:RES %f", channel, bandwidth)
	return e.Write(cmd)
}

// Sets the state of IF bandwidth auto function
func (e *E5061B) SetAutoIFBandwidth(channel int, state bool) error {
	if err := e.validateChannel(channel); err != nil {
		return err
	}

	stateStr := "OFF"
	if state {
		stateStr = "ON"
	}

	cmd := fmt.Sprintf(":SENS%d:BWA:STAT %s", channel, stateStr)
	return e.Write(cmd)
}