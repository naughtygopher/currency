package currency

import (
	"strconv"
	"testing"
)

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

	str := cur1.String(true)
	if str != "₹21.98" {
		t.Log("Expected ₹21.98, got:" + str)
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

	str := cur1.String(true)
	if str != "₹0.00" {
		t.Log("Expected ₹0.00, got:" + str)
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

	str := cur1.String(true)
	if str != "-₹20.50" {
		t.Log("Expected -₹20.50, got:" + str)
		t.Fail()
	}

	if cur1.Float64() != -20.50 {
		t.Log("Expected -20.50, got:", cur1.Float64())
		t.Fail()
	}
}

func TestSub(t *testing.T) {
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

	str := cur1.String(true)
	if str != "-₹1.00" {
		t.Log("Expected -₹1.00, got:" + str)
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

	str := cur1.String(true)
	if str != "₹4.70" {
		t.Log("Expected ₹4.70, got:" + str)
		t.Fail()
	}

	if cur1.Float64() != 4.70 {
		t.Log("Expected 4.70, got:", cur1.Float64())
		t.Fail()
	}
}

func TestMult(t *testing.T) {
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

	str := cur1.String(true)
	if str != "₹52.50" {
		t.Log("Expected ₹52.50, got:" + str)
		t.Fail()
	}

	if cur1.Float64() != 52.50 {
		t.Log("Expected 52.50, got:", cur1.Float64())
		t.Fail()
	}
}

func TestMultFloat64(t *testing.T) {
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

	str := cur1.String(true)
	if str != "₹11.03" {
		t.Log("Expected ₹11.03, got:" + str)
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

	str := cur2.String(true)
	if str != "₹0.55" {
		t.Log("Expected ₹0.55, got:" + str)
		t.Fail()
	}

	if cur2.Float64() != 0.55 {
		t.Log("Expected 0.55, got:", cur2.Float64())
		t.Fail()
	}
}

func TestParseStr(t *testing.T) {
	cur, err := ParseString("10.5", "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if cur.Main != 10 {
		t.Log("Expected 10, got:" + strconv.Itoa(cur.Main))
		t.Fail()
	}

	if cur.Fractional != 50 {
		t.Log("Expected 50, got:" + strconv.Itoa(int(cur.Fractional)))
		t.Fail()
	}

	str := cur.String(true)
	if str != "₹10.50" {
		t.Log("Expected ₹10.50, got:" + str)
		t.Fail()
	}

	if cur.Float64() != 10.50 {
		t.Log("Expected 10.50, got:", cur.Float64())
		t.Fail()
	}

}

func TestParseStr2(t *testing.T) {
	cur, err := ParseString("-10.5", "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if cur.Main != -10 {
		t.Log("Expected -10, got:" + strconv.Itoa(cur.Main))
		t.Fail()
	}

	if cur.Fractional != 50 {
		t.Log("Expected 50, got:" + strconv.Itoa(int(cur.Fractional)))
		t.Fail()
	}

	str := cur.String(true)

	if str != "-₹10.50" {
		t.Log("Expected -₹10.50, got:" + str)
		t.Fail()
	}

	//parsing with fractional unit alone
	cur1, err := ParseString("-0.5", "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if cur1.Main != 0 {
		t.Log("Expected 0, got:" + strconv.Itoa(cur1.Main))
		t.Fail()
	}

	if cur1.Fractional != -50 {
		t.Log("Expected -50, got:" + strconv.Itoa(int(cur1.Fractional)))
		t.Fail()
	}

	str = cur1.String(true)

	if str != "-₹0.50" {
		t.Log("Expected -₹0.50, got:" + str)
		t.Fail()
	}

	if cur1.Float64() != -0.50 {
		t.Log("Expected -0.50, got:", cur1.Float64())
		t.Fail()
	}

}

func TestParseFloat64(t *testing.T) {
	cur, err := ParseFloat64(10.05, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if cur.Main != 10 {
		t.Log("Expected 10, got:", cur.Main)
		t.Fail()
	}

	if cur.Fractional != 5 {
		t.Log("Expected 5, got:", cur.Fractional)
		t.Fail()
	}

	s := cur.String(true)
	if s != "₹10.05" {
		t.Log("Expected ₹10.05, got:", s)
		t.Fail()
	}

	if cur.Float64() != 10.05 {
		t.Log("Expected 10.05, got:", cur.Float64())
		t.Fail()
	}
}

func TestString(t *testing.T) {
	cur, err := New(10, 5, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	result := cur.String(true)
	if result != "₹10.05" {
		t.Log("Expected ₹10.05, got:" + result)
		t.Fail()
	}
}

func TestFloat64(t *testing.T) {
	cur, err := New(10, 1, "INR", "₹", "paise", 100)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	f := cur.Float64()
	if f != 10.01 {
		t.Log("Expected 10.50, got:", f)
		t.Fail()
	}
}
