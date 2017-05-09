package currency

import "testing"

func BenchmarkNewCurrency(t *testing.B) {
	for i := 0; i < t.N; i++ {
		New(10, 5, "INR", "₹", "paise", 100)
	}
}

func BenchmarkCurrencyAdd(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Add(*cur2)
	}
}

func BenchmarkCurrencySub(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Sub(*cur2)
	}
}

func BenchmarkCurrencyMult(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Mult(2)
	}
}

func BenchmarkCurrencyMultFloat64(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.MultFloat64(1.01)
	}
}

func BenchmarkCurrencyFracTotal(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.FracTotal()
	}
}

func BenchmarkCurrencyUpdateWithFrac(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.UpdateWithFrac(2513)
	}
}

func BenchmarkCurrencyPercent(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Percent(12.18)
	}
}

func BenchmarkCurrencyFloat64(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.Float64()
	}
}

func BenchmarkCurrencyString(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.String(true)
	}
}
