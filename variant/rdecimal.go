package variant

import (
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/vbdate"
	"golang.org/x/text/collate"
)

// RDecimal creates a decimal variant type.
func RDecimal(v decimal.Decimal) Variant {
	return rDecimal(v)
}

// rDecimal implements a time.
type rDecimal decimal.Decimal

// Type returns the type of the variant.
func (s rDecimal) getType() Type {
	return TypeRDecimal
}

func (s rDecimal) equals(v Variant, _ bool) bool {
	if v == nil {
		return s == 0
	}
	switch val := v.(type) {
	case rBool:
		return s.toBool() == bool(val)
	case rString:
		return s.toString() == string(val)
	case rDate:
		return s.toDate() == vbdate.VBDate(val)
	}
	return decimal.Decimal(s) == v.toDecimal()
}

func (s rDecimal) less(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return s < 0
	}
	switch val := v.(type) {
	case rString:
		return s.toString() < string(val)
	case rDate:
		return s.toDate().Before(vbdate.VBDate(val))
	case rDecimal:
		return s < val
	}
	return decimal.Decimal(s) < v.toDecimal()
}

func (s rDecimal) greater(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return 0 < s
	}
	switch val := v.(type) {
	case rString:
		return string(val) < s.toString()
	case rDate:
		return vbdate.VBDate(val).Before(s.toDate())
	case rDecimal:
		return val < s
	}
	return v.toDecimal() < decimal.Decimal(s)
}

// ToFloat64 converts the time to float64 representing seconds.
func (s rDecimal) toFloat64() float64 {
	return decimal.Decimal(s).ToFloat64()
}

func (s rDecimal) toString() string {
	return decimal.Decimal(s).ToString()
}
func (s rDecimal) toStringWithDateFormat(string) string {
	return s.toString()
}

func (s rDecimal) toDate() vbdate.VBDate {
	return toTime(s.toFloat64())
}

func (s rDecimal) toBool() bool {
	return s != 0
}

func (s rDecimal) toInt() int {
	return decimal.Decimal(s).ToInt()
}

func (s rDecimal) toDecimal() decimal.Decimal {
	return decimal.Decimal(s)
}

func (s rDecimal) abs() Variant {
	if s < 0 {
		return -s
	}
	return s
}
func (s rDecimal) val() Variant {
	return s
}

func (s rDecimal) MarshalJSON() ([]byte, error) {
	return []byte(s.toString()), nil
}
func (s rDecimal) plus(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	switch p.(type) {
	case rFloat:
		return RFloat(s.toFloat64()).plus(p)
	default:
		return rDecimal(s.toDecimal() + p.toDecimal())
	}
}

func (s rDecimal) minus(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	switch p.(type) {
	case rFloat:
		return RFloat(s.toFloat64()).minus(p)
	default:
		return rDecimal(s.toDecimal() - p.toDecimal())
	}
}

func (s rDecimal) mult(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	switch p.(type) {
	case rFloat:
		return RFloat(s.toFloat64()).mult(p)
	default:
		return rDecimal(s.toDecimal() * p.toDecimal() / decimal.Decimals)
	}
}

func (s rDecimal) div(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	x := p.toFloat64()
	if x == 0 {
		return nil
	}
	return RFloat(s.toDecimal().Div(x))
}

func (s rDecimal) divInt(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	x := p.toDecimal()
	if x == 0 {
		return nil
	}
	x = s.toDecimal() / x
	return rDecimal(decimal.Decimals * x)
}

func (s rDecimal) mod(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	x := p.toDecimal()
	if x == 0 {
		return nil
	}
	return rDecimal(s.toDecimal() % x)
}

func (s rDecimal) exp(p Variant) Variant {
	return RFloat(s.toFloat64()).exp(p)
}

func (s rDecimal) isNumeric() bool {
	return true
}

func (s rDecimal) toJSON() []byte {
	return []byte(s.toString())
}
