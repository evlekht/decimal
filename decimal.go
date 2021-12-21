package decimal

import (
	"time"

	"github.com/shopspring/decimal"
)

type Decimal decimal.Decimal

var (
	zero       = Decimal(decimal.Zero)
	one        = FromFloat(1)
	percent    = FromFloat(0.01)
	secondNano = FromInt64(int64(time.Second))

	_percent = decimal.Decimal(percent)
)

func Zero() Decimal       { return zero }
func One() Decimal        { return one }
func Percent() Decimal    { return percent }
func SecondNano() Decimal { return secondNano }

func FromInt64(n int64) Decimal {
	return Decimal(decimal.NewFromFloat(float64(n)))
}

func FromInt(n int) Decimal {
	return Decimal(decimal.NewFromFloat(float64(n)))
}

func FromFloat(n float64) Decimal {
	return Decimal(decimal.NewFromFloat(n))
}

func FromString(str string) (Decimal, error) {
	dec, err := decimal.NewFromString(str)
	return Decimal(dec), err
}

func (d Decimal) String() string {
	return decimal.Decimal(d).String()
}

func (d Decimal) Float64() float64 {
	f, _ := decimal.Decimal(d).Float64()
	return f
}
