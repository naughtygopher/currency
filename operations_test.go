package currency

import (
	"testing"
)

func TestUpdateWithFractional(t *testing.T) {
	cur, err := New(1, 0, "INR", "₹", "paise", 100)
	if err != nil {
		t.Fatal(err)
	}
	cur.UpdateWithFractional(1005)
	if cur.Main != 10 {
		t.Log("Expected 10, got:", cur.Main)
		t.Fail()
	}

	if cur.Fractional != 5 {
		t.Log("Expected 5, got:", cur.Fractional)
		t.Fail()
	}

	cur.PrefixSymbol = true
	s := cur.String()
	if s != "₹10.05" {
		t.Log("Expected ₹10.05, got:", s)
		t.Fail()
	}

	if cur.Float64() != 10.05 {
		t.Log("Expected 10.05, got:", cur.Float64())
		t.Fail()
	}

}

func TestAdd(t *testing.T) {
	cur1, err := New(10, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur2, err := New(10, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.Add(*cur2)
	if cur1.Main != 21 {
		t.Log("Expected 21, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 98 {
		t.Log("Expected 98, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "₹21.98" {
		t.Log("Expected ₹21.98, got:", str)
		t.Fail()
	}

	if cur1.Float64() != 21.98 {
		t.Log("Expected 21.98, got:", cur1.Float64())
		t.Fail()
	}
}

func TestAdd2(t *testing.T) {
	cur1, err := New(10, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur2, err := New(-10, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.Add(*cur2)

	if cur1.Main != 0 {
		t.Log("Expected 0, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 0 {
		t.Log("Expected 0, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "₹0.00" {
		t.Log("Expected ₹0.00, got:", str)
		t.Fail()
	}

	if cur1.Float64() != 0.00 {
		t.Log("Expected 0.00, got:", cur1.Float64())
		t.Fail()
	}
}

func TestAdd3(t *testing.T) {
	cur1, err := New(-10, 25, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur2, err := New(-10, 25, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.Add(*cur2)
	if cur1.Main != -20 {
		t.Log("Expected -20, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 50 {
		t.Log("Expected 50, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "-₹20.50" {
		t.Log("Expected -₹20.50, got:", str)
		t.Fail()
	}

	if cur1.Float64() != -20.50 {
		t.Log("Expected -20.50, got:", cur1.Float64())
		t.Fail()
	}
}

func TestAddInt(t *testing.T) {
	cur, err := New(10, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur.AddInt(10, 10)

	if cur.Main != 21 {
		t.Log("Expected:", 21, "got:", cur.Main)
		t.Fail()
	}

	if cur.Fractional != 9 {
		t.Log("Expected:", 9, "got:", cur.Fractional)
		t.Fail()
	}

	cur.PrefixSymbol = true
	str := cur.String()
	if str != "₹21.09" {
		t.Log("Expected:", "₹21.09", "got:", str)
		t.Fail()
	}

	if cur.Float64() != 21.09 {
		t.Log("Expected:", 21.09, "got:", cur.Float64())
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	cur1, err := New(10, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur2, err := New(11, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.Subtract(*cur2)

	if cur1.Main != -1 {
		t.Log("Expected -1, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 0 {
		t.Log("Expected 0, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "-₹1.00" {
		t.Log("Expected -₹1.00, got:", str)
		t.Fail()
	}

	if cur1.Float64() != -1.00 {
		t.Log("Expected -1.00, got:", cur1.Float64())
		t.Fail()
	}
}

func TestSub2(t *testing.T) {
	cur1, err := New(10, 69, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur2, err := New(5, 99, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.Subtract(*cur2)

	if cur1.Main != 4 {
		t.Log("Expected 4, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 70 {
		t.Log("Expected 70, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "₹4.70" {
		t.Log("Expected ₹4.70, got:", str)
		t.Fail()
	}

	if cur1.Float64() != 4.70 {
		t.Log("Expected 4.70, got:", cur1.Float64())
		t.Fail()
	}
}

func TestSubtractInt(t *testing.T) {
	cur, err := New(10, 90, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur.SubtractInt(10, 99)

	if cur.Main != 0 {
		t.Log("Expected:", 0, "got:", cur.Main)
		t.Fail()
	}

	if cur.Fractional != -9 {
		t.Log("Expected:", -9, "got:", cur.Fractional)
		t.Fail()
	}

	cur.PrefixSymbol = true
	str := cur.String()
	if str != "-₹0.09" {
		t.Log("Expected:", "-₹0.09", "got:", str)
		t.Fail()
	}

	if cur.Float64() != -0.09 {
		t.Log("Expected:", -0.09, "got:", cur.Float64())
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	cur1, err := New(10, 50, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.Multiply(5)

	if cur1.Main != 52 {
		t.Log("Expected 52, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 50 {
		t.Log("Expected 50, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "₹52.50" {
		t.Log("Expected ₹52.50, got:", str)
		t.Fail()
	}

	if cur1.Float64() != 52.50 {
		t.Log("Expected 52.50, got:", cur1.Float64())
		t.Fail()
	}
}

func TestMultiplyFloat64(t *testing.T) {
	cur1, err := New(10, 50, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur1.MultiplyFloat64(1.05)

	if cur1.Main != 11 {
		t.Log("Expected 11, got:", cur1.Main)
		t.Fail()
	}

	if cur1.Fractional != 3 {
		t.Log("Expected 3, got:", cur1.Fractional)
		t.Fail()
	}

	cur1.PrefixSymbol = true
	str := cur1.String()
	if str != "₹11.03" {
		t.Log("Expected ₹11.03, got:", str)
		t.Fail()
	}

	if cur1.Float64() != 11.03 {
		t.Log("Expected 11.03, got:", cur1.Float64())
		t.Fail()
	}
}

func TestPercent(t *testing.T) {
	cur1, err := New(10, 50, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	cur2 := cur1.Percent(5.25)

	if cur2.Main != 0 {
		t.Log("Expected 0, got:", cur2.Main)
		t.Fail()
	}

	if cur2.Fractional != 55 {
		t.Log("Expected 55, got:", cur2.Fractional)
		t.Fail()
	}

	cur2.PrefixSymbol = true
	str := cur2.String()
	if str != "₹0.55" {
		t.Log("Expected ₹0.55, got:", str)
		t.Fail()
	}

	if cur2.Float64() != 0.55 {
		t.Log("Expected 0.55, got:", cur2.Float64())
		t.Fail()
	}
}

func TestDivideRetain(t *testing.T) {
	cur, err := New(1, 0, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	splits, _ := cur.Divide(3, true)
	if cur.Fractional != 1 {
		t.Log("Expected", 1, "got:", cur.Fractional)
		t.Fail()
	}

	for idx := range splits {
		split := splits[idx]
		if split.Main != 0 {
			t.Log("Expected", 0, "got:", split.Main)
			t.Fail()
		}

		if split.Fractional != 33 {
			t.Log("Expected", 33, "got:", split.Fractional)
			t.Fail()
		}

		split.PrefixSymbol = true
		str := split.String()
		if str != "₹0.33" {
			t.Log("Expected:", "₹0.33", "got:", str)
			t.Fail()
		}

		if split.Float64() != 0.33 {
			t.Log("Expected:", 0.33, "got:", split.Float64())
			t.Fail()
		}

		if split.FractionalTotal() != 33 {
			t.Log("Expected:", 33, "got:", split.FractionalTotal())
			t.Fail()
		}
	}
}

func TestDivideNoRetain(t *testing.T) {
	cur, err := New(1, 0, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	splits, _ := cur.Divide(3, false)
	if cur.Fractional != 0 {
		t.Log("Expected", 0, "got:", cur.Fractional)
		t.Fail()
	}

	for idx := range splits {
		split := splits[idx]
		if split.Main != 0 {
			t.Log("Expected", 0, "got:", split.Main)
			t.Fail()
		}

		if split.Fractional != 33 && split.Fractional != 34 {
			t.Log("Expected", "33 or 34", "got:", split.Fractional)
			t.Fail()
		}

		split.PrefixSymbol = true
		str := split.String()
		if str != "₹0.33" && str != "₹0.34" {
			t.Log("Expected:", "₹0.33 or ₹0.34", "got:", str)
			t.Fail()
		}

		if split.Float64() != 0.33 && split.Float64() != 0.34 {
			t.Log("Expected:", "0.33 or 0.34", "got:", split.Float64())
			t.Fail()
		}

		if split.FractionalTotal() != 33 && split.FractionalTotal() != 34 {
			t.Log("Expected:", "33 or 34", "got:", split.FractionalTotal())
			t.Fail()
		}
	}
}

func BenchmarkUpdateWithFractional(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.UpdateWithFractional(2513)
	}
}

func BenchmarkAdd(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Add(*cur2)
	}
}

func BenchmarkAddInt(t *testing.B) {
	cur, _ := New(10, 99, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.AddInt(10, 10)
	}
}

func BenchmarkSubtract(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Subtract(*cur2)
	}
}

func BenchmarkSubtractInt(t *testing.B) {
	cur, _ := New(10, 99, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.SubtractInt(10, 10)
	}
}

func BenchmarkMultiply(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Multiply(2)
	}
}

func BenchmarkMultiplyFloat64(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.MultiplyFloat64(1.01)
	}
}

func BenchmarkPercent(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Percent(12.18)
	}
}

func BenchmarkDivide(t *testing.B) {
	cur1, _ := New(9999, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.Divide(2, true)
	}
}
