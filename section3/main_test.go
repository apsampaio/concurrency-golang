package main

import (
	"sync"
	"testing"
)

func Test_CalcBalance(t *testing.T) {
	var balance sync.Mutex
	var wg sync.WaitGroup

	// define weekly revenue
	incomes := []Income{
		{Source: "Main jobs", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	bankBalance := CalcBalance(incomes, &balance, &wg)

	if bankBalance != 34320 {
		t.Error("bank balance different from expected")
	}
}
