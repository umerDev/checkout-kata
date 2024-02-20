package main

import "testing"

func AssertResult(expectedResult int, actualResult int, t *testing.T) {
	if expectedResult != actualResult {
		t.Errorf("Expected %+v Got %+v", expectedResult, actualResult)
	}
}

func TestGetTotalPrice(t *testing.T) {
	var totalPrice = 0
	items := []string{"A", "B", "C", "D"}

	scanned := Scan(items) // get occurences
	for item, quantity := range scanned {
		totalPrice += GetTotalPrice(item, quantity)
	}

	expectedResult := 115
	actualResult := totalPrice

	AssertResult(expectedResult, actualResult, t)
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

	AssertResult(expectedResult, actualResult, t)

}

func TestGetTotalPriceWithDiscountAB(t *testing.T) {
	var totalPrice = 0
	items := []string{"A", "A", "A", "B", "B"}

	scanned := Scan(items) // get occurences
	for item, quantity := range scanned {
		totalPrice += GetTotalPrice(item, quantity)
	}

	expectedResult := 175
	actualResult := totalPrice

	AssertResult(expectedResult, actualResult, t)
}

func TestGetTotalPriceWithDiscountA4(t *testing.T) {
	var totalPrice = 0
	items := []string{"A", "A", "A", "A"}

	scanned := Scan(items) // get occurences
	for item, quantity := range scanned {
		totalPrice += GetTotalPrice(item, quantity)
	}

	expectedResult := 180
	actualResult := totalPrice

	AssertResult(expectedResult, actualResult, t)
}

func TestGetTotalPriceWithDiscountA4B3(t *testing.T) {
	var totalPrice = 0
	items := []string{"A", "A", "A", "A", "B", "B", "B"}

	scanned := Scan(items) // get occurences
	for item, quantity := range scanned {
		totalPrice += GetTotalPrice(item, quantity)
	}

	expectedResult := 180 + 45 + 30
	actualResult := totalPrice

	AssertResult(expectedResult, actualResult, t)
}

func GetTotalPrice(sku string, quantity int) int {
	if sku == "A" {
		if quantity >= 3 {
			return CalculateDiscount(quantity, 3, 130, 50)
		}
		return 50 * quantity
	}
	if sku == "B" {
		if quantity >= 2 {
			return CalculateDiscount(quantity, 2, 45, 30)
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

func CalculateDiscount(quantity int, multiplier int, specialPrice int, normalPrice int) int {
	price := 0
	remaining := quantity % multiplier

	if remaining == 0 {
		for i := 0; i < quantity; i += multiplier {
			price += specialPrice
		}
		return price
	}

	if remaining != 0 {
		discounted := quantity - remaining
		for i := 0; i < discounted; i += multiplier {
			price += specialPrice
		}
		return remaining*normalPrice + price
	}

	return price
}
