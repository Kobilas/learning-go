package main

import "testing"

func TestFormatAmount(t *testing.T) {
	ans := FormatAmount(2.00)
	if ans != "USD 2.00" {
		// we use t to record the testing error
		t.Errorf("FormatAmount(2.00) = %s; Should be 2.00", ans)
	}
}

func TestFormatAmount2(t *testing.T) {
	ans := FormatAmount(4.00)
	if ans != "USD 4.00" {
		t.Errorf("FormatAmount(4.00) = %s; Should be 4.00", ans)
	}
}

func TestFormatAmount3(t *testing.T) {
	ans := FormatAmount(5.10)
	if ans != "USD 5.10" {
		t.Errorf("FormatAmount(5.10) = %s; Should be 5.10", ans)
	}
}

// test if input value is negative
func TestFormatAmount4(t *testing.T) {
	ans := FormatAmount(-5.10)
	if ans != "Impossible operation" {
		t.Errorf("FormatAmount(-5.10) cannot be performed")
	}
}

func TestSubtractFormatAmount(t *testing.T) {
	ans := SubtractFormatAmount(4.00, 2.00)
	if ans != "USD 2.00" {
		t.Errorf("SubtractFormatAmount(4.00, 2.00) = %s; Should be USD 2.00", ans)
	}
}

func TestSubtractFormatAmount2(t *testing.T) {
	ans := SubtractFormatAmount(3.00, 1.12)
	if ans != "USD 1.88" {
		// continue to use t to record testing error
		t.Errorf("SubtractFormatAmount(3.00, 1.12) = %s; Should be USD 1.88", ans)
	}
}

func TestSubtractFormatAmount3(t *testing.T) {
	ans := SubtractFormatAmount(1.00, 1.12)
	if ans != "Impossible operation" {
		t.Errorf("SubtractFormatAmount(1.00, 1.12) cannot be performed")
	}
}

// test if both input values are negative
func TestSubtractFormat4(t *testing.T) {
	ans := SubtractFormatAmount(-1.00, -1.12)
	if ans != "Impossible operation" {
		t.Errorf("SubtractFormatAmount(-1.00, -1.12) cannot be performed")
	}
}

// test if b is negative
func TestSubtractFormatAmount5(t *testing.T) {
	ans := SubtractFormatAmount(1.00, -1.12)
	if ans != "Impossible operation" {
		t.Errorf("SubtractFormatAmount(1.00, -1.12) cannot be performed")
	}
}

func TestAddFormatAmount(t *testing.T) {
	ans := AddFormatAmount(4.00, 2.00)
	if ans != "USD 6.00" {
		t.Errorf("AddFormatAmount(4.00, 2.00) = %s; Should be USD 6.00", ans)
	}
}

func TestAddFormatAmount2(t *testing.T) {
	ans := AddFormatAmount(3.00, 1.12)
	if ans != "USD 4.12" {
		t.Errorf("AddFormatAmount(3.00, 1.12) = %s; Should be USD 4.12", ans)
	}
}

func TestAddFormatAmount3(t *testing.T) {
	ans := AddFormatAmount(-1.00, -1.12)
	if ans != "Impossible operation" {
		t.Errorf("AddFormatAmount(-1.00, -1.12) cannot be performed")
	}
}
