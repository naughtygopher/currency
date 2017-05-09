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
		cur1.Sub(*cur2)
	}
}

func BenchmarkMult(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Mult(2)
	}
}

func BenchmarkMultFloat64(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.MultFloat64(1.01)
	}
}

func BenchmarkFracTotal(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.FracTotal()
	}
}

func BenchmarkUpdateWithFrac(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.UpdateWithFrac(2513)
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
