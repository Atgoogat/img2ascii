package cli

import (
	"fmt"
	"strconv"
)

type float64Flag struct {
	set   bool
	value *float64
}

func NewFloat64Flag() *float64Flag {
	return &float64Flag{
		false,
		nil,
	}
}

func (ff *float64Flag) Set(value string) error {
	ff.set = true
	v, err := strconv.ParseFloat(value, 64)
	if err == nil {
		ff.value = &v
	}
	return err
}

func (ff *float64Flag) String() string {
	return fmt.Sprint(ff.value)
}
