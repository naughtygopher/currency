## Currency

Currency is a simple library to do simple currency computations. A currency can be configured by setting the following values in the `Currency` struct. In most cases, a currency would have 2 units, 1 main unit and another fractional/sub unit.

[Ref 1](https://en.wikipedia.org/wiki/Denomination_(currency)), [Ref 2](https://en.wikipedia.org/wiki/Currency) about currencies.
[Non-decimal sub unit in currencies are only used by 2 countries today](https://en.wikipedia.org/wiki/Non-decimal_currency). Even these are getting phased out.

*This library does not support sub units which are not a power of 10. Nor does it support currencies with more than 1 sub unit*

A currency is represented as/configured using the following struct

```
type Currency struct {
	//Code represents the international currency code
	Code string
	//Symbol is the respective currency symbol
	Symbol string
	//Main represents the main value of the currency
	Main int
	//Fractional represents the fractional/sub unit of the currency
	Fractional uint
	//FUName is the name of the fractional/sub unit of the currency. e.g. paise
	FUName string
	//FUShare represents the no.of fractional/sub units that make up 1 main unit. e.g. ₹1 = 100 paise
	//Number of fractional units that make up 1 unit of the main value
	FUShare uint
}
```

### NewCurrency(main int, fractional int, code, symbol string, funame string, fushare uint)

NewCurrency returns a pointer of currency instance created based on the values provided

```
main - Main value of the currency
fractional - Subunit/fractional value of the currency
code - is the currency code according to [ISO 4217 specification](https://en.wikipedia.org/wiki/ISO_4217)
symbol - Unicode symbol of the currency
funame - Name of the fractional/sub unit
fushare - Number of fractional/sub units that make up 1 unit of the main value
```

Fractional unit can be negative only when the main value is 0. If the main value is not 0, fractional unit's negative sign is ignored.

### Parsers

1. `ParseStr(value string, code, symbol string,  fulabel string, fushare uint)` returns a currency struct instance, given a currency value represented as string
2. `ParseFloat64(value float64, code, symbol string, funame string, fushare uint)` returns a 
currency struct instance, given a currency value represented in float64

### Following computations are supported

Computation is supported only between same type of currencies (i.e. currency codes are same)

1. `c1.Add(c2 currency)` add c2 to c1, and update c1
2. `c1.Sub(c2 currency)` subtract c2 from c1, and update c1
3. `c1.Mult(n int)` multiply c1 by n, where n is an integer
4. `c1.MultFloat64(n float64)` multiply c1 by n, where n is a float64 value
5. `c1.UpdateWithFrac(ftotal int)` would update the curreny's respective value, where ftotal is the total value of the currency in its fractional unit. e.g. INR, `UpdateWithFrac(100)` would set the main value as `1` and fractional unit as `0`
6. `c1.FracTotal() int` returns the total value of the currency in its fractional unit. e.g. INR, if the Main value is `1` and fractional unit is `0`, it would return `100`, i.e. 100 paise
7. `c1.Percent(n float64) currency` returns a new currency instance which is n % of c1

### Multiple currency representations

1. `c1.String(prefixSymbol bool)`, returns a string representation of the currency value. Returns string prefixed by its respective symbol if `prefixSymbol` is true
2. `c1.Float64()`, returns a float64 representation of the currency value

## Benchmarks

How to run?

`$ go test -bench=.`

Results when run on a MacBook Pro (13-inch, Early 2015), CPU: 2.7 GHz Intel Core i5, RAM: 8 GB 1867 MHz DDR3, Graphics: Intel Iris Graphics 6100 1536 MB

```
BenchmarkNewCurrency-4              	10000000	       113 ns/op
BenchmarkCurrencyAdd-4              	100000000	        22.8 ns/op
BenchmarkCurrencySub-4              	100000000	        23.3 ns/op
BenchmarkCurrencyMult-4             	100000000	        16.9 ns/op
BenchmarkCurrencyMultFloat64-4      	50000000	        31.5 ns/op
BenchmarkCurrencyFracTotal-4        	2000000000	         0.36 ns/op
BenchmarkCurrencyUpdateWithFrac-4   	100000000	        11.0 ns/op
BenchmarkCurrencyPercent-4          	20000000	        95.3 ns/op
BenchmarkCurrencyFloat64-4          	2000000000	         0.40 ns/op
BenchmarkCurrencyString-4           	 5000000	       248 ns/op
```