package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance

	var bankBalance int
	var balance sync.Mutex
	var wg sync.WaitGroup

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main jobs", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	bankBalance = CalcBalance(incomes, &balance, &wg)

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
}

func CalcBalance(incomes []Income, balance *sync.Mutex, wg *sync.WaitGroup) int {
	var result int

	wg.Add(len(incomes))
	// loop through 52 weeks and print out  how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := result
				temp += income.Amount
				result = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	return result
}
