package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/hooklift/gowsdl/soap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zoxpx/gosoapTest/myservice"
)

var calc myservice.CalculatorSoap

func TestAdd(t *testing.T) {
	data := []struct {
		a, b   int32
		expect int32
	}{
		{123, 456, 579},
		{456, -123, 333},
		{123, -456, -333},
	}

	for _, td := range data {
		lab := fmt.Sprintf("%d+(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc.Add(&myservice.Add{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.AddResult)
		})
	}
}

func TestSubtract(t *testing.T) {
	data := []struct {
		a, b   int32
		expect int32
	}{
		{123, 456, -333},
		{456, -123, 579},
		{123, -456, 579},
	}

	for _, td := range data {
		lab := fmt.Sprintf("%d-(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc.Subtract(&myservice.Subtract{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.SubtractResult)
		})
	}
}

func TestMultiply(t *testing.T) {
	data := []struct {
		a, b   int32
		expect int32
	}{
		{123, 456, 56088},
		{456, -123, -56088},
		{123, -456, -56088},
	}

	for _, td := range data {
		lab := fmt.Sprintf("%d*(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc.Multiply(&myservice.Multiply{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.MultiplyResult)
		})
	}
}

func TestDivide(t *testing.T) {
	data := []struct {
		a, b   int32
		expect int32
	}{
		{123, 456, 0},
		{456, 123, 4},
		{4566, -123, -37},
	}

	for _, td := range data {
		lab := fmt.Sprintf(" %d/(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc.Divide(&myservice.Divide{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.DivideResult)
		})
	}
}

func TestMain(m *testing.M) {
	calc = myservice.NewCalculatorSoap(soap.NewClient("http://www.dneonline.com/calculator.asmx"))
	os.Exit(m.Run())
}
