package main

import "testing"

func TestGetTotalPrice(t *testing.T) {
	var totalPrice = 0
	items := []string{"A", "B", "C", "D"}

	scanned := Scan(items) // get occurences
	for item, quantity := range scanned {
		totalPrice += GetTotalPrice(item, quantity)
	}

	expectedResult := 115
	actualResult := totalPrice

	if expectedResult != actualResult {
		t.Errorf("Expected %+v Got %+v", expectedResult, actualResult)
	}
}

func TestGetTotalPriceWithDiscountForA(t *testing.T) {
	var totalPrice = 0
	items := []string{"A", "A", "A"}

	scanned := Scan(items) // get occurences
	for item, quantity := range scanned {
		totalPrice += GetTotalPrice(item, quantity)
	}

	expectedResult := 130
	actualResult := totalPrice

	if expectedResult != actualResult {
		t.Errorf("Expected %+v Got %+v", expectedResult, actualResult)
	}
}

func GetTotalPrice(sku string, quantity int) int {
	if sku == "A" {
		if quantity >= 3 {
			return 130
		}
		return 50 * quantity
	}
	if sku == "B" {
		if quantity >= 2 {
			return 45
		}
		return 30 * quantity
	}
	if sku == "C" {
		return 20 * quantity
	}
	if sku == "D" {
		return 15 * quantity
	}
	return 0
}

func Scan(items []string) map[string]int {
	dict := make(map[string]int)
	for _, num := range items {
		dict[num] = dict[num] + 1
	}
	return dict
}
