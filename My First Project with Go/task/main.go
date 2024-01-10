package main

import (
	"fmt"
	"strconv"
)

func priceListPrinter(prices [6]float64, items [6]string) {
	for i := 0; i < len(prices); i++ {
		fmt.Println(items[i] + ": $" + strconv.FormatFloat(prices[i], 'f', 2, 64))
	}
}

func totalIncome(incomeEarned [6]float64) float64 {
	total := 0.0
	for i := 0; i < len(incomeEarned); i++ {
		total += incomeEarned[i]
	}
	return total

}

func incomePrinter(items [6]string, incomeEarned [6]float64, total float64) {
	fmt.Println("Earned amount:")
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i] + ": $" + strconv.FormatFloat(incomeEarned[i], 'f', 2, 64))
	}
	fmt.Println("\nIncome: $" + strconv.FormatFloat(total, 'f', 1, 64))
}
func expensesCalculator(staffExpenses float64, otherExpenses float64) float64 {
	return staffExpenses + otherExpenses
}

func netIncomeCaluclator(totalIncome float64, totalExpenses float64) float64 {
	return totalIncome - totalExpenses
}
func main() {
	//fmt.Println("Prices:")
	//prices := [6]float64{2, 0.2, 5, 4, 2.5, 3.2}
	items := [6]string{"Bubblegum", "Toffee", "Ice cream", "Milk chocolate", "Doughnut", "Pancake"}
	incomeEarned := [6]float64{202, 118, 2250, 1680, 1075, 80}
	//priceListPrinter(prices, items)
	incomePrinter(items, incomeEarned, totalIncome(incomeEarned))
	staffExpenses := 0.0
	otherExpenses := 0.0
	fmt.Println("Staff expenses:")
	fmt.Scan(&staffExpenses)
	fmt.Println("Other expenses:")
	fmt.Scan(&otherExpenses)
	fmt.Println("Net income: $" + strconv.FormatFloat(netIncomeCaluclator(totalIncome(incomeEarned), expensesCalculator(staffExpenses, otherExpenses)), 'f', 1, 64))

}
