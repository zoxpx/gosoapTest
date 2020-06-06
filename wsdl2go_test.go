package main

import (
	"fmt"
	"github.com/fiorix/wsdl2go/soap"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zoxpx/gosoapTest/calculatorsoap12"
)

var calc2 calculatorsoap12.CalculatorSoap

func init() {
	calc2 = calculatorsoap12.NewCalculatorSoap(&soap.Client{
		URL:       "http://www.dneonline.com/calculator.asmx",
		Namespace: calculatorsoap12.Namespace,
	})
}

/*
 * sends following:
```
Host:             www.dneonline.com
User-Agent:       Go-http-client/1.1
Content-Length:   300
Content-Type:     application/soap+xml; charset=utf-8; action="http://tempuri.org/Subtract"
Accept-Encoding:  gzip

<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://tempuri.org/"
xmlns:tns="http://tempuri.org/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <SOAP-ENV:Body>
    <Subtract>
      <intA>123</intA>
      <intB>-456</intB>
    </Subtract>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>
```
*/
func TestWsdl2goAdd(t *testing.T) {
	data := []struct {
		a, b   int
		expect int
	}{
		{123, 456, 579},
		{456, -123, 333},
		{123, -456, -333},
	}

	for _, td := range data {
		lab := fmt.Sprintf("%d+(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc2.Add(&calculatorsoap12.Add{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.AddResult)
		})
	}
}

func TestWsdl2goSubtract(t *testing.T) {
	data := []struct {
		a, b   int
		expect int
	}{
		{123, 456, -333},
		{456, -123, 579},
		{123, -456, 579},
	}

	for _, td := range data {
		lab := fmt.Sprintf("%d-(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc2.Subtract(&calculatorsoap12.Subtract{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.SubtractResult)
		})
	}
}

func TestWsdl2goMultiply(t *testing.T) {
	data := []struct {
		a, b   int
		expect int
	}{
		{123, 456, 56088},
		{456, -123, -56088},
		{123, -456, -56088},
	}

	for _, td := range data {
		lab := fmt.Sprintf("%d*(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc2.Multiply(&calculatorsoap12.Multiply{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.MultiplyResult)
		})
	}
}

func TestWsdl2goDivide(t *testing.T) {
	data := []struct {
		a, b   int
		expect int
	}{
		{123, 456, 0},
		{456, 123, 4},
		{4566, -123, -37},
	}

	for _, td := range data {
		lab := fmt.Sprintf(" %d/(%d)", td.a, td.b)
		t.Run(lab, func(t *testing.T) {
			t.Parallel()
			resp, err := calc2.Divide(&calculatorsoap12.Divide{
				IntA: td.a,
				IntB: td.b,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, td.expect, resp.DivideResult)
		})
	}
}
