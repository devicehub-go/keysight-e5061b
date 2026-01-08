package protocol

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

type ComplexData struct {
	Real      []float64
	Imaginary []float64
	Magnitude []float64
	Phase     []float64
	Frequency []float64
}

// Sets the data type for transfers
func (e *E5061B) SetDataFormat(format string) error {
	if err := e.validateDataFormat(format); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":FORM:DATA %s", strings.ToUpper(format))
	return e.Write(cmd)
}

// Sets the byte order for binary transfers
func (e *E5061B) SetByteOrder(order string) error {
	if err := e.validateByteOrder(order); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":FORM:BORD %s", strings.ToUpper(order))
	return e.Write(cmd)
}

// Converts a byte array to an array of float64 values (SWAP, 64-bit)
func (e *E5061B) ByteToFloatArray(payload []byte) ([]float64, error) {
	if len(payload)%8 != 0 {
		return nil, fmt.Errorf("payload must be aligned to 8 bytes, got %d", len(payload))
	}
	numValues := len(payload) / 8

	values := make([]float64, numValues)
	for i := range numValues {
		bits := binary.LittleEndian.Uint64(payload[i*8 : i*8+8])
		values[i] = math.Float64frombits(bits)
	}

	return values, nil
}

// Fetches formatted trace data as a float64 slice using 64-bit precision
func (e *E5061B) GetFormattedData(channel int) ([]float64, error) {
	if err := e.validateChannel(channel); err != nil {
		return nil, err
	}
	if err := e.SetDataFormat("REAL"); err != nil {
		return nil, err
	}
	if err := e.SetByteOrder("SWAP"); err != nil {
		return nil, err
	}

	cmd := fmt.Sprintf(":CALC%d:DATA:FDAT?", channel)
	payload, err := e.QueryByteSequence(cmd)
	if err != nil {
		return nil, err
	}

	return e.ByteToFloatArray(payload)
}

// Fetches raw complex data as a float64 slice
func (e *E5061B) GetComplexData(channel int) (ComplexData, error) {
	if err := e.validateChannel(channel); err != nil {
		return ComplexData{}, err
	}
	if err := e.SetDataFormat("REAL"); err != nil {
		return ComplexData{}, err
	}
	if err := e.SetByteOrder("SWAP"); err != nil {
		return ComplexData{}, err
	}

	cmd := fmt.Sprintf(":CALC%d:DATA:SDAT?", channel)
	payload, err := e.QueryByteSequence(cmd)
	if err != nil {
		return ComplexData{}, err
	}
	values, err := e.ByteToFloatArray(payload)
	if err != nil {
		return ComplexData{}, err
	}

	numPoints := len(values) / 2
    output := ComplexData{
        Real:      make([]float64, numPoints),
        Imaginary: make([]float64, numPoints),
        Magnitude: make([]float64, numPoints),
        Phase:     make([]float64, numPoints),
		Frequency: make([]float64, numPoints),
    }

	params, err := e.GetFrequencyParameters(channel)
	if err != nil {
		return ComplexData{}, err
	}

    for i := range numPoints {
        r, im := values[i*2], values[i*2+1]

        output.Real[i] = r
        output.Imaginary[i] = im
        output.Magnitude[i] = 20*math.Log10(math.Sqrt(r*r + im*im))
        output.Phase[i] = math.Atan2(im, r) * (180/math.Pi)
		output.Frequency[i] = params.Start + float64(i)*(params.Stop - params.Start) / (params.Sweep - 1)
    }

	return output, nil
}

// Selects a specific trace and returns its complex data
func (e *E5061B) GetTraceComplexData(channel, trace int) (ComplexData, error) {
	if err := e.SelectTrace(channel, trace); err != nil {
		return ComplexData{}, err
	}
	return e.GetComplexData(channel)
}

// Enables or disables a specific marker on a trace
func (e *E5061B) SetMarkerState(channel, trace, marker int, enabled bool) error {
	if err := e.SelectTrace(channel, trace); err != nil {
		return err
	} else if err := e.validateMarkerIndex(marker); err != nil {
		return err
	}

	state := "0"
	if enabled {
		state = "1"
	}

	cmd := fmt.Sprintf(":CALC%d:MARK%d:STAT %s", channel, marker, state)
	return e.Write(cmd)
}

// Moves the marker to a specific frequency position in Hz
func (e *E5061B) SetMarkerX(channel, trace, marker int, frequency float64) error {
	if err := e.SelectTrace(channel, trace); err != nil {
		return err
	} else if err := e.validateMarkerIndex(marker); err != nil {
		return err
	} else if err := e.validateFrequency(frequency); err != nil {
		return err
	}

	cmd := fmt.Sprintf(":CALC%d:MARK%d:X %g", channel, marker, frequency)
	return e.Write(cmd)
}

// Queries the measurement data at the marker's current position
func (e *E5061B) GetMarkerY(channel, trace, marker int) ([]byte, error) {
	if err := e.SelectTrace(channel, trace); err != nil {
		return nil, err
	} else if err := e.validateMarkerIndex(marker); err != nil {
		return nil, err
	}

	cmd := fmt.Sprintf(":CALC%d:MARK%d:Y?", channel, marker)
	return e.Query(cmd)
}