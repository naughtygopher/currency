package currency

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

type output struct {
	main            int
	fractional      int
	totalfractional int
	str             string
	float           float64
	err             error
}

type input struct {
	main            int
	fractional      int
	totalfractional int
	str             string
	float           float64
	code            string
	symbol          string
	fulabel         string
	fushare         uint
}

var newTests = []struct {
	inp input
	out output
}{
	{
		input{
			main:            10,
			fractional:      50,
			totalfractional: 1050,
			str:             "10.50",
			float:           10.50,
			code:            "INR",
			symbol:          "₹",
			fulabel:         "paise",
			fushare:         100},
		output{10, 50, 1050, "₹10.50", 10.50, nil},
	},
	{
		input{
			main:            -10,
			fractional:      -50,
			totalfractional: -1050,
			str:             "-10.50",
			float:           -10.50,
			code:            "INR",
			symbol:          "₹",
			fulabel:         "paise",
			fushare:         100},
		output{-10, 50, -1050, "-₹10.50", -10.50, nil},
	},
	{
		input{
			main:            0,
			fractional:      -50,
			totalfractional: -50,
			str:             "-0.50",
			float:           -0.50,
			code:            "INR",
			symbol:          "₹",
			fulabel:         "paise",
			fushare:         100},
		output{0, -50, -50, "-₹0.50", -0.50, nil},
	},
	{
		input{
			main:            1,
			fractional:      100,
			totalfractional: 0,
			code:            "INR",
			str:             "1.00",
			symbol:          "₹",
			fulabel:         "paise",
			fushare:         0,
		},
		output{0, 0, 0, "", 0, ErrInvalidFUS},
	},
}

func TestNew(t *testing.T) {
	for _, nT := range newTests {
		cur, err := New(
			nT.inp.main,
			nT.inp.fractional,
			nT.inp.code,
			nT.inp.symbol,
			nT.inp.fulabel,
			nT.inp.fushare)
		if err != nil {
			if !errors.Is(err, nT.out.err) {
				t.Fatal(err)
			}
			continue
		}

		if cur.Main != nT.out.main {
			t.Log("Expected:", nT.out.main, "got:", cur.Main)
			t.Fail()
		}

		if cur.Fractional != nT.out.fractional {
			t.Log("Expected:", nT.out.fractional, "got:", cur.Fractional)
			t.Fail()
		}

		cur.PrefixSymbol = true
		str := cur.String()
		if str != nT.out.str {
			t.Log("Expected:", nT.out.str, "got:", str)
			t.Fail()
		}

		if cur.Float64() != nT.out.float {
			t.Log("Expected:", nT.out.float, "got:", cur.Float64())
			t.Fail()
		}

		ft := cur.FractionalTotal()
		if ft != nT.out.totalfractional {
			t.Log("Expected:", nT.out.totalfractional, "got:", ft)
			t.Fail()
		}
	}

}

func TestNewFractional(t *testing.T) {
	for _, nT := range newTests {
		cur, err := NewFractional(
			nT.inp.totalfractional,
			nT.inp.code,
			nT.inp.symbol,
			nT.inp.fulabel,
			nT.inp.fushare)

		if err != nil {
			if !errors.Is(err, nT.out.err) {
				t.Fatal(err)
			}
			continue
		}

		if cur.Main != nT.out.main {
			t.Log("Expected:", nT.out.main, "got:", cur.Main)
			t.Fail()
		}

		if cur.Fractional != nT.out.fractional {
			t.Log("Expected:", nT.out.fractional, "got:", cur.Fractional)
			t.Fail()
		}

		cur.PrefixSymbol = true
		str := cur.String()
		if str != nT.out.str {
			t.Log("Expected:", nT.out.str, "got:", str)
			t.Fail()
		}

		if cur.Float64() != nT.out.float {
			t.Log("Expected:", nT.out.float, "got:", cur.Float64())
			t.Fail()
		}

		ft := cur.FractionalTotal()
		if ft != nT.out.totalfractional {
			t.Log("Expected:", nT.out.totalfractional, "got:", ft)
			t.Fail()
		}
	}
}

func TestParseStr(t *testing.T) {
	for _, nT := range newTests {
		cur, err := ParseString(
			nT.inp.str,
			nT.inp.code,
			nT.inp.symbol,
			nT.inp.fulabel,
			nT.inp.fushare)

		if err != nil {
			if !errors.Is(err, nT.out.err) {
				t.Fatal(err)
			}
			continue
		}

		if cur.Main != nT.out.main {
			t.Log("Expected:", nT.out.main, "got:", cur.Main)
			t.Fail()
		}

		if cur.Fractional != nT.out.fractional {
			t.Log("Expected:", nT.out.fractional, "got:", cur.Fractional)
			t.Fail()
		}

		cur.PrefixSymbol = true
		str := cur.String()
		if str != nT.out.str {
			t.Log("Expected:", nT.out.str, "got:", str)
			t.Fail()
		}

		if cur.Float64() != nT.out.float {
			t.Log("Expected:", nT.out.float, "got:", cur.Float64())
			t.Fail()
		}

		ft := cur.FractionalTotal()
		if ft != nT.out.totalfractional {
			t.Log("Expected:", nT.out.totalfractional, "got:", ft)
			t.Fail()
		}
	}

	// code:            "INR",
	// symbol:          "₹",
	// fulabel:         "paise",
	_, err := ParseString("", "INR", "₹", "paise", 0)
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Error(err)
	}
}

func TestParseFloat64(t *testing.T) {
	for _, nT := range newTests {
		cur, err := ParseFloat64(
			nT.inp.float,
			nT.inp.code,
			nT.inp.symbol,
			nT.inp.fulabel,
			nT.inp.fushare)

		if err != nil {
			if !errors.Is(err, nT.out.err) {
				t.Fatal(err)
			}
			continue
		}

		if cur.Main != nT.out.main {
			t.Log("Expected:", nT.out.main, "got:", cur.Main)
			t.Fail()
		}

		if cur.Fractional != nT.out.fractional {
			t.Log("Expected:", nT.out.fractional, "got:", cur.Fractional)
			t.Fail()
		}

		cur.PrefixSymbol = true
		str := cur.String()
		if str != nT.out.str {
			t.Log("Expected:", nT.out.str, "got:", str)
			t.Fail()
		}

		if cur.Float64() != nT.out.float {
			t.Log("Expected:", nT.out.float, "got:", cur.Float64())
			t.Fail()
		}

		ft := cur.FractionalTotal()
		if ft != nT.out.totalfractional {
			t.Log("Expected:", nT.out.totalfractional, "got:", ft)
			t.Fail()
		}
	}
}

func TestFormat(t *testing.T) {
	c, _ := New(12, 75, "INR", "₹", "paise", 100)
	list := []struct {
		Verb     string
		Expected string
		Prefix   bool
		Suffix   bool
	}{
		{
			Verb:     "s",
			Expected: "12.75",
		},
		{
			Verb:     "s",
			Prefix:   true,
			Expected: "₹12.75",
		},
		{
			Verb:     "s",
			Suffix:   true,
			Expected: "12.75₹",
		},
		{
			Verb:     "s",
			Suffix:   true,
			Prefix:   true,
			Expected: "₹12.75₹",
		},
		{
			Verb:     "d",
			Expected: "12",
		},
		{
			Verb:     "m",
			Expected: "75",
		},
		{
			Verb:     "f",
			Expected: "12.75",
		},
		{
			Verb:     "y",
			Expected: "₹",
		},
	}
	for _, l := range list {
		c.PrefixSymbol = l.Prefix
		c.SuffixSymbol = l.Suffix

		formatstr := "%" + l.Verb
		str := fmt.Sprintf(formatstr, c)
		if str != l.Expected {
			t.Errorf("Format string: %s, Expected '%s', got '%s'", formatstr, l.Expected, str)
		}
	}
}

func BenchmarkNew(t *testing.B) {
	for i := 0; i < t.N; i++ {
		New(10, 50, "INR", "₹", "paise", 100)
	}
}

func BenchmarkNewFractional(t *testing.B) {
	for i := 0; i < t.N; i++ {
		NewFractional(1005, "INR", "₹", "paise", 100)
	}
}

func BenchmarkParseFloat64(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ParseFloat64(10.05, "INR", "₹", "paise", 100)
	}
}

func BenchmarkParseString(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ParseString("10.05", "INR", "₹", "paise", 100)
	}
}

func BenchmarkString(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	cur1.PrefixSymbol = true
	for i := 0; i < t.N; i++ {
		cur1.String()
	}
}

func BenchmarkStringNoPrefix(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.StringWithoutSymbols()
	}
}

func BenchmarkFloat64(t *testing.B) {
	cur1, _ := New(10, 5, "INR", "₹", "paise", 100)
	for i := 0; i < t.N; i++ {
		cur1.Float64()
	}
}

func BenchmarkFractionalTotal(t *testing.B) {
	cur1, _ := New(1, 0, "INR", "₹", "paise", 100)

	for i := 0; i < t.N; i++ {
		cur1.FractionalTotal()
	}
}

func Test_digits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single positive",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "single negative",
			args: args{
				n: -1,
			},
			want: 1,
		},
		{
			name: "more digits positive",
			args: args{
				n: 100,
			},
			want: 3,
		},
		{
			name: "more digits negative",
			args: args{
				n: -100,
			},
			want: 3,
		},
		{
			name: "zero",
			args: args{
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digits(tt.args.n); got != tt.want {
				t.Errorf("digits() = %v, want %v", got, tt.want)
			}
		})
	}
}
