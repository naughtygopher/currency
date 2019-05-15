[![Build Status](https://travis-ci.org/bnkamalesh/currency.svg?branch=master)](https://travis-ci.org/bnkamalesh/currency)
 [![gocover.run](https://gocover.run/github.com/bnkamalesh/currency.svg?style=flat&tag=1.10)](https://gocover.run?tag=1.10&repo=github.com%2Fbnkamalesh%2Fcurrency)
[![](https://goreportcard.com/badge/github.com/bnkamalesh/currency)](https://goreportcard.com/report/github.com/bnkamalesh/currency)
[![](https://api.codeclimate.com/v1/badges/d181840cf59e912b6c31/maintainability)](https://codeclimate.com/github/bnkamalesh/currency/maintainability)
[![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/bnkamalesh/currency)

## Currency

Currency package helps you do currency computations accurately, by avoiding *_peddling_*.
The `Currency` struct holds all the data required to define a currency.

```
type Currency struct {
	// Code represents the international currency code
	Code string
	// Symbol is the respective currency symbol
	Symbol string
	// Main represents the main value of the currency
	Main int
	// Fractional represents the fractional/sub unit of the currency
	Fractional uint
	// FUName is the name of the fractional/sub unit of the currency. e.g. paise
	FUName string
	// FUShare represents the no.of fractional/sub units that make up 1 main unit. e.g. ₹1 = 100 paise
	// Number of fractional units that make up 1 unit of the main value
	FUShare uint
}
```

### New(main int, fractional int, code, symbol string, funame string, fushare uint)

New returns a pointer of currency instance created based on the values provided

```
main - Main/Super unit of the currency
fractional - Subunit/fractional unit of the currency
code - is the currency code according to [ISO 4217 specification](https://en.wikipedia.org/wiki/ISO_4217)
symbol - Unicode symbol of the currency
funame - Name of the fractional/sub unit
fushare - Number of fractional/sub units that make up 1 unit of the main/super unit
```

*IMPORTANT! Fractional unit can be negative only when the main value is 0. If the main value is not 0, fractional unit's negative sign is ignored.*

### Parsers & convenience methods

1. `NewFractional(fractional int, symbol string,  fulabel string, fushare uint)` returns a currency struct instance, given a currency's total value represented by the fractional unit
3. `ParseString(value string, code, symbol string,  fulabel string, fushare uint)` returns a currency struct instance, given a currency value represented as string
4. `ParseFloat64(value float64, code, symbol string, funame string, fushare uint)` returns a currency struct instance, given a currency value represented in float64

### Computational methods

IMPORTANT: Computation is supported only between same type of currencies (i.e. currency codes should be same)

1. `c1.Add(c2 currency)` add c2 to c1, and update c1
2. `c1.AddInt(main int, fractional int)` add the currency equivalent of the main & fractional int to c1
3. `c1.Subtract(c2 currency)` subtract c2 from c1, and update c1
4. `c1.SubtractIn(main int, fractional int)` subtract the currency equivalent of the main & fractional int from c1
5. `c1.Multiply(n int)` multiply c1 by n, where n is an integer
6. `c1.MultiplyFloat64(n float64)` multiply c1 by n, where n is a float64 value
7. `c1.UpdateWithFractional(ftotal int)` would update the the value of c1,  where *ftotal* is the total value of the currency in fractional unit. e.g. INR, `UpdateWithFractional(100)` would set the main value as `1` and fractional unit as `0`
8. `c1.FractionalTotal() int` returns the total value of the currency in its fractional unit. e.g. INR, if the Main value is `1` and fractional unit is `0`, it would return `100`, i.e. 100 paise
9. `c1.Percent(n float64) currency` returns a new currency instance which is n percentage of c1
10. `c1.Divide(n int, retain bool)[]currency, ok ` returns a slice of currency of size n. `ok` if returned as `true` means the currency value was perfectly divisible by n. If `retain` is true,
then `c1` will have the remainder value after dividing otherwise it is distributed among the returned currencies.

#### Why does `Divide(n int, retain bool)` return a slice of currencies?

`Divide` unlike other operations, cannot be rounded off. If it is rounded, it would result in currency peddling.

e.g. ₹1/- (INR 1) is to be divided by 3. There are 2 options of dividing this by 3.

	1. Set 33 paise per split, and retain the remaining 1 paise at source. (`Divide(n, true)`)

	2. Set 1 of the split with an extra value, i.e. 34 + 33 + 33. (`Divide(n, false)`)

### Multiple currency representations

1. `c1.String(prefixSymbol bool)`, returns a string representation of the currency value. Returns string prefixed by its respective symbol if `prefixSymbol` is true
2. `c1.Float64()`, returns a float64 representation of the currency value

## Benchmarks

How to run?

`$ go test -bench=.`

Results when run on a MacBook Pro (13-inch, Early 2015), CPU: 2.7 GHz Intel Core i5, RAM: 8 GB 1867 MHz DDR3, Graphics: Intel Iris Graphics 6100 1536 MB

```
BenchmarkNew-4                    	20000000	        67.3 ns/op
BenchmarkNewFractional-4          	20000000	        65.9 ns/op
BenchmarkParseFloat64-4           	20000000	        87.4 ns/op
BenchmarkParseString-4            	 3000000	       544 ns/op
BenchmarkString-4                 	10000000	       211 ns/op
BenchmarkStringNoPrefix-4         	10000000	       164 ns/op
BenchmarkFloat64-4                	2000000000	         0.34 ns/op
BenchmarkFractionalTotal-4        	2000000000	         0.33 ns/op
BenchmarkUpdateWithFractional-4   	100000000	        10.3 ns/op
BenchmarkAdd-4                    	100000000	        20.8 ns/op
BenchmarkAddInt-4                 	100000000	        18.9 ns/op
BenchmarkSubtract-4               	100000000	        21.2 ns/op
BenchmarkSubtractInt-4            	100000000	        18.3 ns/op
BenchmarkMultiply-4               	100000000	        16.2 ns/op
BenchmarkMultiplyFloat64-4        	50000000	        30.1 ns/op
BenchmarkPercent-4                	20000000	        67.1 ns/op
BenchmarkDivide-4                 	10000000	       155 ns/op
```

## References

1. [Ref - Sub unit or fractional unit](https://en.wikipedia.org/wiki/Denomination_(currency))
2. [Ref - Currencies](https://en.wikipedia.org/wiki/Currency) - about currencies
3. [Non-decimal sub unit in currencies are only used by 2 countries today](https://en.wikipedia.org/wiki/Non-decimal_currency). These are getting phased out.

*IMPORTANT! This package does not support sub units which are not a power of 10. Nor does it support currencies with more than 1 sub unit*
