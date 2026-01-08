package protocol

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devicehub-go/keysight-e5061b/internal/utils"
)

const (
	MinFrequency = 5.0
	MaxFrequency = 3.0e9
	MinPoints    = 2
	MaxPoints    = 1601
	MinChannel   = 1
	MaxChannel   = 4
	MinTrace     = 1
	MaxTrace     = 4
)

func (e *E5061B) validateChannel(channel int) error {
	if channel < MinChannel || channel > MaxChannel {
		return fmt.Errorf(
			"channel must be between %d and %d, got %d",
			MinChannel, MaxChannel, channel,
		)
	}
	return nil
}

func (e *E5061B) validateTrace(trace int) error {
	if trace < MinTrace || trace > MaxTrace {
		return fmt.Errorf(
			"trace must be between %d and %d, got %d",
			MinTrace, MaxTrace, trace,
		)
	}
	return nil
}

func (e *E5061B) validateFrequency(frequency float64) error {
	if frequency < MinFrequency || frequency > MaxFrequency {
		return fmt.Errorf(
			"frequency must be between %.1f and %.1e, got %g",
			MinFrequency, MaxFrequency, frequency,
		)
	}
	return nil
}

func (e *E5061B) validateSweepPoints(points int) error {
	if points < MinPoints || points > MaxPoints {
		return fmt.Errorf("points out of range: %d (must be %d-%d)", points, MinPoints, MaxPoints)
	}
	return nil
}

func (e *E5061B) validateSweepType(sweep string) error {
	valid := []string{"LIN", "LOG", "POW", "SEGM"}
	sweepUpper := strings.ToUpper(sweep)
	if !slices.Contains(valid, sweepUpper) {
		return fmt.Errorf("sweep type must be %s, got %s", sweep, valid)
	}
	return nil
}

func (e *E5061B) validateBandwidth(bandwidth int) error {
	valid := []int{10, 30, 100, 300, 1e3, 3e3, 10e3, 30e3}
	if !slices.Contains(valid, bandwidth) {
		validStr := utils.JoinIntArray(valid, ", ")
		return fmt.Errorf("bandwidth must be %s, got %d", validStr, bandwidth)
	}
	return nil
}

func (e *E5061B) validateAnalysisMode(mode string) error {
	valid := []string{"NA", "ZA"}
	modeUpper := strings.ToUpper(mode)
	if !slices.Contains(valid, modeUpper) {
		return fmt.Errorf("analysis mode must be %s, got %s", strings.Join(valid, ", "), mode)
	}
	return nil
}

func (e *E5061B) validateTraceParameter(param string) error {
	valid := []string{"S11", "S21", "S12", "S22", "R1", "A", "B"}
	paramUpper := strings.ToUpper(param)
	if !slices.Contains(valid, paramUpper) {
		return fmt.Errorf("parameter must be %s, got %s", strings.Join(valid, ", "), param)
	}
	return nil
}

func (e *E5061B) validateDisplayFormat(format string) error {
	valid := []string{"MLOG", "PHAS", "SMIT", "SWR", "REAL", "IMAG", "UPH"}
	formatUpper := strings.ToUpper(format)
	if !slices.Contains(valid, formatUpper) {
		return fmt.Errorf("format must be %s, got %s", strings.Join(valid, ", "), format)
	}
	return nil
}

func (e *E5061B) validateWindowLayout(layout string) error {
	valid := []string{
		"D1", "D12", "D1_2", "D112", "D1_1_2", "D123", 
		"D1_2_3", "D12_33", "D11_23", "D13_23", "D12_13",
		"D1234", "D1_2_3_4", "D12_34",
	}
	layoutUpper := strings.ToUpper(layout)
	if !slices.Contains(valid, layoutUpper) {
		return fmt.Errorf("window layout must be %s, got %s", strings.Join(valid, ", "), layout)
	}
	return nil
}

func (e *E5061B) validateTriggerSource(source string) error {
	valid := []string{"INT", "EXT", "MAN", "BUS"}
	sourceUpper := strings.ToUpper(source)
	if !slices.Contains(valid, sourceUpper) {
		return fmt.Errorf("trigger source must be %s, got %s", strings.Join(valid, ", "), source)
	}
	return nil
}

func (e *E5061B) validateDataFormat(format string) error {
	valid := []string{"ASC", "REAL", "REAL32"}
	formatUpper := strings.ToUpper(format)
	if !slices.Contains(valid, formatUpper) {
		return fmt.Errorf("data format must be %s, got %s", strings.Join(valid, ", "), format)
	}
	return nil
}

func (e *E5061B) validateByteOrder(order string) error {
	valid := []string{"NORM", "SWAP"}
	orderUpper := strings.ToUpper(order)
	if !slices.Contains(valid, orderUpper) {
		return fmt.Errorf("byte order must be %s, got %s", strings.Join(valid, ", "), order)
	}
	return nil
}

func (e *E5061B) validateMarkerIndex(marker int) error {
	if marker < 1 || marker > 9 {
		return fmt.Errorf("marker must be between 1 and 9, got %d", marker)
	}
	return nil
}
