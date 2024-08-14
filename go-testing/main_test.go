package main

import (
	"testing"
)

func TestMonth(t *testing.T) {
	if ConvertMonth(8) != "August" {
		t.Error("The 8th month was supposed to be August.")
	}
}

func TestMonths(t *testing.T) {
	months := []int{
		8, 8, 8,
	}

	for _, month := range months { // Goes through each month
		if ConvertMonth(month) != "August" {
			t.Error("The month isn't August.")
		}
	}
}
