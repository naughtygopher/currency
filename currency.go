// Package currency helps represent a currency with high precision, and do currency computations.
package currency

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	// ErrMismatchCurrency is the error returned if trying to do computation between unmatched currencies
	ErrMismatchCurrency = errors.New("Currencies do not match")

	// ErrInvalidCurrency is the error returned while trying parse an invalid currency value
	ErrInvalidCurrency = errors.New("Invalid currency value provided")

	// ErrInvalidFUS is the error returned when Functional unit share is equal to 0
	ErrInvalidFUS = errors.New("Invalid functional unit share provided")
)

// replacer is the regex which replaces all invalid characters inside a string representing a currency value
var replacer = regexp.MustCompile(`([^0-9.\-\+])`)

const (
	replaceWith = ""
)

// Currency represents money with all the meta data required.
type Currency struct {
	// Code represents the international currency code
	Code string
	// Symbol is the respective currency symbol
	Symbol string
	// Main represents the main unit value of the currency
	Main int
	// Fractional represents the fractional unit value of the currency
	Fractional int
	// FUName is the name of the fractional unit of the currency. e.g. paise
	FUName string
	// FUShare represents the number of fractional units that make up 1 main unit. e.g. ₹1 = 100 Paise.
	FUShare uint

	// fuDigits is the number of digits in FUShare-1 (i.e. number of digits in the maximum value which the fractional unit can have, e.g. 99 paise, 2 digits)
	fuDigits int
	// magnitude is the fraction which sets the magnitude required for the rounding function
	magnitude float64
}

// New returns a new instance of currency.
func New(main int, fractional int, code, symbol, funame string, fushare uint) (*Currency, error) {
	if fushare == 0 {
		return nil, ErrInvalidFUS
	}

	if main != 0 && fractional < 0 {
		fractional = -fractional
	}

	fus := int(fushare)

	m := main + (fractional / fus)
	f := fractional % fus

	fudigits := digits(fus - 1)

	mag := float64(5.0)

	for i := 0; i < fudigits-1; i++ {
		mag /= 10
	}

	return &Currency{
		Code:       code,
		Symbol:     symbol,
		Main:       m,
		Fractional: f,
		FUName:     funame,
		FUShare:    fushare,
		fuDigits:   fudigits,
		magnitude:  mag,
	}, nil
}

// NewFractional returns a new instance of currency given the total value of currency in fractional unit.
func NewFractional(ftotal int, code, symbol, funame string, fushare uint) (*Currency, error) {
	if fushare == 0 {
		return nil, ErrInvalidFUS
	}

	fus := int(fushare)

	m := ftotal / fus
	f := (ftotal % fus)

	if m < 0 {
		f = -f
	}

	fudigits := digits(fus - 1)

	mag := float64(5.0)

	for i := 0; i < fudigits-1; i++ {
		mag /= 10
	}

	return &Currency{
		Code:       code,
		Symbol:     symbol,
		Main:       m,
		Fractional: f,
		FUName:     funame,
		FUShare:    fushare,
		fuDigits:   fudigits,
		magnitude:  mag,
	}, nil
}

// ParseString will parse a string representation of the currency and return instance of Currency.
func ParseString(value string, code, symbol, funame string, fushare uint) (*Currency, error) {

	splits := strings.Split(replacer.ReplaceAllString(value, replaceWith), ".")
	if len(splits) != 2 {
		return nil, ErrInvalidCurrency
	}

	mstr := strings.Trim(splits[0], " ")
	fstr := strings.Trim(splits[1], " ")

	m, err := strconv.Atoi(mstr)
	if err != nil {
		return nil, err
	}

	f, err := strconv.Atoi(fstr)
	if err != nil {
		return nil, err
	}

	mpr := len(strconv.Itoa(int(fushare-1))) - len(fstr)

	for i := 0; i < mpr; i++ {
		f *= 10
	}

	if m == 0 && string(mstr[0]) == "-" {
		f = -f
	}

	return New(m, f, code, symbol, funame, fushare)
}

// ParseFloat64 will parse a float value into currency.
func ParseFloat64(value float64, code, symbol, funame string, fushare uint) (*Currency, error) {
	if fushare == 0 {
		return nil, ErrInvalidFUS
	}

	fus := int(fushare)
	fudigits := len(strconv.Itoa(fus - 1))

	mag := float64(5.0)
	for i := 0; i < fudigits-1; i++ {
		mag /= 10
	}

	ftotal := round(value*float64(fushare), mag)

	main := ftotal / fus
	fractional := (ftotal % fus)

	if main < 0 {
		fractional = -fractional
	}

	m := main + (fractional / fus)
	f := fractional % fus

	return &Currency{
		Code:       code,
		Symbol:     symbol,
		Main:       m,
		Fractional: f,
		FUName:     funame,
		FUShare:    fushare,
		fuDigits:   fudigits,
		magnitude:  mag,
	}, nil
}

// FractionalTotal returns the total value in fractional int.
func (c *Currency) FractionalTotal() int {
	cFrac := c.Fractional

	if c.Main < 0 {
		cFrac = -cFrac
	}

	return ((c.Main * int(c.FUShare)) + cFrac)
}

// UpdateWithFractional will update all the relevant values of currency based on the fractional unit provided.
func (c *Currency) UpdateWithFractional(frac int) {
	fus := int(c.FUShare)

	c.Main = (frac / fus)
	c.Fractional = (frac % fus)

	if c.Main < 0 {
		c.Fractional = -c.Fractional
	}
}

// Add adds the given currency with the base currency.
func (c *Currency) Add(acur Currency) error {
	if c.Code != acur.Code {
		return ErrMismatchCurrency
	}

	c.UpdateWithFractional(c.FractionalTotal() + acur.FractionalTotal())
	return nil
}

// Subtract subtracts the given currency from the base currency.
func (c *Currency) Subtract(scur Currency) error {
	if c.Code != scur.Code {
		return ErrMismatchCurrency
	}

	c.UpdateWithFractional(c.FractionalTotal() - scur.FractionalTotal())
	return nil
}

// Percent returns a new instance of currency which is n percent of c.
func (c *Currency) Percent(n float64) *Currency {
	totalFrac := round(float64(c.FractionalTotal())*(n/100.00), c.magnitude)
	c1 := *c
	c1.UpdateWithFractional(totalFrac)
	return &c1
}

// Multiply multiplies the currency by an integer.
func (c *Currency) Multiply(by int) {
	c.UpdateWithFractional(c.FractionalTotal() * by)
}

// MultiplyFloat64 multiplies the currency by a float value.
func (c *Currency) MultiplyFloat64(by float64) {
	t := float64(c.FractionalTotal()) * by
	c.UpdateWithFractional(round(t, c.magnitude))
}

// Float64 returns the currency in float64 format.
func (c *Currency) Float64() float64 {
	frac := c.Fractional
	if c.Main < 0 {
		frac = -frac
	}

	return float64(c.Main) + (float64(frac) / float64(c.FUShare))
}

// String returns the currency represented as string.
func (c *Currency) String(prefixSymbol bool) string {
	frc := c.Fractional
	if c.Fractional < 0 {
		frc = -frc
	}

	fstr := strconv.Itoa(frc)

	//all the missing digits are added to the string
	for i := 0; i < c.fuDigits-len(fstr); i++ {
		fstr = "0" + fstr
	}

	str := strconv.Itoa(c.Main) + "." + fstr

	if c.Fractional < 0 {
		str = "-" + str
	}

	if prefixSymbol {
		if c.Fractional < 0 || c.Main < 0 {
			str = strings.Replace(str, "-", "-"+c.Symbol, 1)
		} else {
			str = c.Symbol + str
		}
	}

	return str
}

// round rounds off the float value to the configured precision and returns an integer.
func round(f float64, magnitude float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}

	return int(f + math.Copysign(magnitude, f))
}

// digits returns the number of digits in an integer
func digits(n int) int {
	if n < 0 {
		n = -n
	}

	n /= 10
	d := 1

	for n > 0 {
		d++
		n /= 10
	}

	return d
}
