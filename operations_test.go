package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateWithFractional(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(1, 0, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur.UpdateWithFractional(1005)
	asserter.Equal(10, cur.Main)
	asserter.Equal(5, cur.Fractional)

	cur.PrefixSymbol = true
	asserter.Equal("₹10.05", cur.String())
	asserter.Equal(10.05, cur.Float64())
}

func TestAdd(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur1, err := New(10, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur2, err := New(10, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	asserter.NoError(cur1.Add(*cur2))

	asserter.Equal(21, cur1.Main)
	asserter.Equal(98, cur1.Fractional)
	asserter.Equal(21.98, cur1.Float64())

	cur1.PrefixSymbol = true
	asserter.Equal("₹21.98", cur1.String())
}

func TestAdd2(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur1, err := New(10, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur2, err := New(-10, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	asserter.NoError(cur1.Add(*cur2))

	asserter.Equal(0, cur1.Main)
	asserter.Equal(0, cur1.Fractional)
	asserter.Equal(0.00, cur1.Float64())

	cur1.PrefixSymbol = true
	asserter.Equal("₹0.00", cur1.String())
}

func TestAdd3(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur1, err := New(-10, 25, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur2, err := New(-10, 25, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	asserter.NoError(cur1.Add(*cur2))

	asserter.Equal(-20, cur1.Main)
	asserter.Equal(50, cur1.Fractional)
	asserter.Equal(-20.50, cur1.Float64())

	cur1.PrefixSymbol = true
	asserter.Equal("-₹20.50", cur1.String())
}

func TestAddInt(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(10, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur.AddInt(10, 10)
	asserter.Equal(21, cur.Main)
	asserter.Equal(9, cur.Fractional)
	asserter.Equal(21.09, cur.Float64())

	cur.PrefixSymbol = true
	asserter.Equal("₹21.09", cur.String())
}

func TestSubtract(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur1, err := New(10, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur2, err := New(11, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	asserter.NoError(cur1.Subtract(*cur2))

	asserter.Equal(-1, cur1.Main)
	asserter.Equal(0, cur1.Fractional)
	asserter.Equal(-1.00, cur1.Float64())

	cur1.PrefixSymbol = true
	asserter.Equal("-₹1.00", cur1.String())
}

func TestSub2(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur1, err := New(10, 69, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur2, err := New(5, 99, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	asserter.NoError(cur1.Subtract(*cur2))

	asserter.Equal(4, cur1.Main)
	asserter.Equal(70, cur1.Fractional)
	asserter.Equal(4.70, cur1.Float64())

	cur1.PrefixSymbol = true
	asserter.Equal("₹4.70", cur1.String())
}

func TestSubtractInt(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(10, 90, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur.SubtractInt(10, 99)

	asserter.Equal(0, cur.Main)
	asserter.Equal(-9, cur.Fractional)
	asserter.Equal(-0.09, cur.Float64())

	cur.PrefixSymbol = true
	asserter.Equal("-₹0.09", cur.String())
}

func TestMultiply(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(10, 50, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur.Multiply(5)

	asserter.Equal(52, cur.Main)
	asserter.Equal(50, cur.Fractional)
	asserter.Equal(52.50, cur.Float64())

	cur.PrefixSymbol = true
	asserter.Equal("₹52.50", cur.String())
}

func TestMultiplyFloat64(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(10, 50, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur.MultiplyFloat64(1.05)

	asserter.Equal(11, cur.Main)
	asserter.Equal(3, cur.Fractional)
	asserter.Equal(11.03, cur.Float64())

	cur.PrefixSymbol = true
	asserter.Equal("₹11.03", cur.String())
}

func TestPercent(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur1, err := New(10, 50, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	cur2 := cur1.Percent(5.25)

	asserter.Equal(0, cur2.Main)
	asserter.Equal(55, cur2.Fractional)
	asserter.Equal(0.55, cur2.Float64())

	cur2.PrefixSymbol = true
	asserter.Equal("₹0.55", cur2.String())
}

func TestDivideRetain(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(1, 0, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	splits, _ := cur.Divide(3, true)
	asserter.Equal(1, cur.Fractional)

	for idx := range splits {
		split := splits[idx]
		asserter.Equal(0, split.Main)
		asserter.Equal(33, split.Fractional)
		asserter.Equal(0.33, split.Float64())
		asserter.Equal(33, split.FractionalTotal())

		split.PrefixSymbol = true
		asserter.Equal("₹0.33", split.String())
	}
}

func TestDivideNoRetain(t *testing.T) {
	requirer := require.New(t)
	asserter := assert.New(t)

	cur, err := New(1, 0, "INR", "₹", "paise", 100)
	requirer.NoError(err)

	splits, _ := cur.Divide(3, false)
	asserter.Equal(0, cur.Fractional)

	for idx := range splits {
		split := splits[idx]
		asserter.Equal(0, split.Main)

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
	cur, _ := New(1, 0, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.UpdateWithFractional(2513)
	}
}

func BenchmarkAdd(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur2, _ := New(10, 5, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		_ = cur1.Add(*cur2)
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
		_ = cur1.Subtract(*cur2)
	}
}

func BenchmarkSubtractInt(t *testing.B) {
	cur, _ := New(10, 99, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.SubtractInt(10, 10)
	}
}

func BenchmarkMultiply(t *testing.B) {
	cur, _ := New(1, 0, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.Multiply(2)
	}
}

func BenchmarkMultiplyFloat64(t *testing.B) {
	cur, _ := New(1, 0, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.MultiplyFloat64(1.01)
	}
}

func BenchmarkPercent(t *testing.B) {
	cur, _ := New(1, 0, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur.Percent(12.18)
	}
}

func BenchmarkAllocate(t *testing.B) {
	cur, _ := New(9999, 0, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		_, _ = cur.Allocate(2, true)
	}
}
