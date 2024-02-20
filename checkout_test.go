package main

import "testing"

func TestGetTotalPrice(t *testing.T) {
	var totalPrice = 0
	items := [4]string{"A", "B", "C", "D"}
	for _, item := range items {
		totalPrice += GetTotalPrice(item)
	}
	expectedResult := 115
	actualResult := totalPrice

	if expectedResult != actualResult {
		t.Errorf("Expected %+v Got %+v", expectedResult, actualResult)
	}
}

func GetTotalPrice(sku string) int {
	return 0
}
