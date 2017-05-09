package currency

import "testing"

func BenchmarkNew(t *testing.B) {
	for i := 0; i < t.N; i++ {
		New(10, 5, "INR", "₹", "paise", 100)
	}
}

func BenchmarkAdd(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Add(*cur2)
	}
}

func BenchmarkSub(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Subtract(*cur2)
	}
}

func BenchmarkMult(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Multiply(2)
	}
}

func BenchmarkMultFloat64(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.MultiplyFloat64(1.01)
	}
}

func BenchmarkFracTotal(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.FractionalTotal()
	}
}

func BenchmarkUpdateWithFrac(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.UpdateWithFractional(2513)
	}
}

func BenchmarkPercent(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Percent(12.18)
	}
}

func BenchmarkFloat64(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.Float64()
	}
}

func BenchmarkString(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.String(true)
	}
}

func BenchmarkStringNoPrefix(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.String(false)
	}
}

func BenchmarkParseString(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ParseString("10.05", "INR", "₹", "paise", 100)
	}
}

func BenchmarkParseFloat64(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ParseFloat64(10.05, "INR", "₹", "paise", 100)
	}
}
