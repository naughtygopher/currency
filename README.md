## Currency

Currency is a package to do simple currency computations. A currency can be configured by setting the following values in the `Currency` struct. In most cases, a currency would have 2 units, 1 main unit and another fractional/sub unit.

[Ref 1](https://en.wikipedia.org/wiki/Denomination_(currency)), [Ref 2](https://en.wikipedia.org/wiki/Currency) about currencies.

[Non-decimal sub unit in currencies are only used by 2 countries today](https://en.wikipedia.org/wiki/Non-decimal_currency). Even these are getting phased out.

*This package does not support sub units which are not a power of 10. Nor does it support currencies with more than 1 sub unit*

A currency is represented as/configured using the following struct

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
main - Main value of the currency
fractional - Subunit/fractional value of the currency
code - is the currency code according to [ISO 4217 specification](https://en.wikipedia.org/wiki/ISO_4217)
symbol - Unicode symbol of the currency
funame - Name of the fractional/sub unit
fushare - Number of fractional/sub units that make up 1 unit of the main value
```

Fractional unit can be negative only when the main value is 0. If the main value is not 0, fractional unit's negative sign is ignored.

### Parsers & convenience methods

1. `NewFractional(fractional int, symbol string,  fulabel string, fushare uint)` returns a currency struct instance, given a currency's total value represented by the fractional unit
3. `ParseString(value string, code, symbol string,  fulabel string, fushare uint)` returns a currency struct instance, given a currency value represented as string
4. `ParseFloat64(value float64, code, symbol string, funame string, fushare uint)` returns a currency struct instance, given a currency value represented in float64

### Following computations are supported

Computation is supported only between same type of currencies (i.e. currency codes are same)

1. `c1.Add(c2 currency)` add c2 to c1, and update c1
2. `c1.Subtract(c2 currency)` subtract c2 from c1, and update c1
3. `c1.Multiply(n int)` multiply c1 by n, where n is an integer
4. `c1.MultiplyFloat64(n float64)` multiply c1 by n, where n is a float64 value
5. `c1.UpdateWithFractional(ftotal int)` would update the curreny's respective value, where ftotal is the total value of the currency in its fractional unit. e.g. INR, `UpdateWithFractional(100)` would set the main value as `1` and fractional unit as `0`
6. `c1.FractionalTotal() int` returns the total value of the currency in its fractional unit. e.g. INR, if the Main value is `1` and fractional unit is `0`, it would return `100`, i.e. 100 paise
7. `c1.Percent(n float64) currency` returns a new currency instance which is n % of c1

### Multiple currency representations

1. `c1.String(prefixSymbol bool)`, returns a string representation of the currency value. Returns string prefixed by its respective symbol if `prefixSymbol` is true
2. `c1.Float64()`, returns a float64 representation of the currency value

## Benchmarks

How to run?

`$ go test -bench=.`

Results when run on a MacBook Pro (13-inch, Early 2015), CPU: 2.7 GHz Intel Core i5, RAM: 8 GB 1867 MHz DDR3, Graphics: Intel Iris Graphics 6100 1536 MB

```
BenchmarkNew-4                    	20000000	        70.3 ns/op
BenchmarkNewFractional-4          	20000000	        73.0 ns/op
BenchmarkParseString-4            	 2000000	       918 ns/op
BenchmarkParseFloat64-4           	10000000	       135 ns/op
BenchmarkFractionalTotal-4        	2000000000	         0.35 ns/op
BenchmarkUpdateWithFractional-4   	100000000	        10.9 ns/op
BenchmarkAdd-4                    	100000000	        22.3 ns/op
BenchmarkSubtract-4               	100000000	        22.7 ns/op
BenchmarkMultiply-4               	100000000	        17.0 ns/op
BenchmarkMultiplyFloat64-4        	50000000	        31.6 ns/op
BenchmarkPercent-4                	20000000	        69.9 ns/op
BenchmarkString-4                 	10000000	       227 ns/op
BenchmarkStringNoPrefix-4         	10000000	       176 ns/op
BenchmarkFloat64-4                	2000000000	         0.36 ns/op
```