package decimal

import "github.com/shopspring/decimal"

func (d1 Decimal) Add(d2 Decimal) Decimal {
	return Decimal(decimal.Decimal(d1).Add(decimal.Decimal(d2)))
}

func (d1 Decimal) Sub(d2 Decimal) Decimal {
	return Decimal(decimal.Decimal(d1).Sub(decimal.Decimal(d2)))
}

func (d1 Decimal) Mul(d2 Decimal) Decimal {
	return Decimal(decimal.Decimal(d1).Mul(decimal.Decimal(d2)))
}

func (d1 Decimal) Div(d2 Decimal) Decimal {
	return Decimal(decimal.Decimal(d1).Div(decimal.Decimal(d2)))
}

func (d1 Decimal) Percent() Decimal {
	return Decimal(decimal.Decimal(d1).Mul(_percent))
}

func (d1 Decimal) GreaterThan(d2 Decimal) bool {
	return decimal.Decimal(d1).GreaterThan(decimal.Decimal(d2))
}

func (d1 Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	return decimal.Decimal(d1).GreaterThanOrEqual(decimal.Decimal(d2))
}

func (d1 Decimal) LessThan(d2 Decimal) bool {
	return decimal.Decimal(d1).LessThan(decimal.Decimal(d2))
}

func (d1 Decimal) LessThanOrEqual(d2 Decimal) bool {
	return decimal.Decimal(d1).LessThanOrEqual(decimal.Decimal(d2))
}

func (d1 Decimal) Equal(d2 Decimal) bool {
	return decimal.Decimal(d1).Equal(decimal.Decimal(d2))
}

func (d Decimal) Truncate(precision int32) Decimal {
	return Decimal(decimal.Decimal(d).Truncate(precision))
}

func (d Decimal) Round(precision int32) Decimal {
	return Decimal(decimal.Decimal(d).Round(precision))
}

func (d Decimal) Neg() Decimal {
	return Decimal(decimal.Decimal(d).Neg())
}

func Max(d1 Decimal, d2 Decimal) Decimal {
	return Decimal(decimal.Max(decimal.Decimal(d1), decimal.Decimal(d2)))
}

func Min(d1 Decimal, d2 Decimal) Decimal {
	return Decimal(decimal.Min(decimal.Decimal(d1), decimal.Decimal(d2)))
}
